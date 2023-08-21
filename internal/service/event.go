package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/cursorutils"
	"exusiai.dev/roguestats-backend/internal/exprutils"
	"exusiai.dev/roguestats-backend/internal/model"
	"exusiai.dev/roguestats-backend/internal/repo"
)

type Event struct {
	fx.In

	EventRepo       repo.Event
	ResearchService Research
	AuthService     Auth
}

func (s Event) CreateEvent(ctx context.Context, event *model.Event) error {
	return s.EventRepo.CreateEvent(ctx, event)
}

func (s Event) GetEvents(ctx context.Context) ([]*model.Event, error) {
	return s.EventRepo.GetEvents(ctx)
}

func (s Event) GetEventsByResearchID(ctx context.Context, researchID string) ([]*model.Event, error) {
	return s.EventRepo.GetEventsByResearchID(ctx, researchID)
}

func (s Event) CreateEventFromInput(ctx context.Context, input *model.NewEvent) (*model.Event, error) {
	event, err := s.convertFromEventInputToEvent(ctx, input)
	if err != nil {
		return nil, err
	}
	err = s.CreateEvent(ctx, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s Event) GetPaginatedEvents(ctx context.Context, researchID string, first int, after string) (*model.EventsConnection, error) {
	events, err := s.EventRepo.GetPaginatedEventsByResearchID(ctx, researchID, first, after)
	if err != nil {
		return nil, err
	}
	eventsConnection := &model.EventsConnection{
		Edges:    make([]*model.EventsEdge, 0),
		PageInfo: &model.PageInfo{},
	}
	for _, event := range events {
		eventsConnection.Edges = append(eventsConnection.Edges, &model.EventsEdge{
			Node:   event,
			Cursor: cursorutils.EncodeCursor(event.ID),
		})
	}

	// decide StartCursor and EndCursor
	endID := events[len(events)-1].ID
	if len(events) > 0 {
		eventsConnection.PageInfo.StartCursor = cursorutils.EncodeCursor(events[0].ID)
		eventsConnection.PageInfo.EndCursor = cursorutils.EncodeCursor(endID)
	}

	// decide HasNextPage
	nextEvent, err := s.EventRepo.GetPaginatedEventsByResearchID(ctx, researchID, 1, endID)
	if err != nil {
		return nil, err
	}
	eventsConnection.PageInfo.HasNextPage = new(bool)
	*eventsConnection.PageInfo.HasNextPage = len(nextEvent) > 0

	return eventsConnection, nil
}

/**
 * CalculateStats filters events by filterInput and maps every event to a result using resultMappingInput, then group by result and count
 * @param {string} filterInput can be jsonLogic or expr expression, depends on the filter implementation
 * @param {string} resultMappingInput must be an expr expression
 */
func (s Event) CalculateStats(ctx context.Context, researchID string, filterInput string, resultMappingInput string) (*model.GroupCountResult, error) {
	// filter events
	filteredEvents, err := s.getEventsWithFilter(ctx, researchID, filterInput)
	if err != nil {
		return nil, err
	}

	categoryCountMap := make(map[interface{}]int)

	totalCount := 0
	for _, event := range filteredEvents {
		// map event to result
		results, err := s.mapEventToResult(event, resultMappingInput)
		if err != nil {
			return nil, err
		}
		if results == nil {
			continue
		}
		totalCount++
		// group by result and count
		for _, result := range results {
			categoryCountMap[result]++
		}
	}

	// convert map into array
	results := make([]*model.CategoryCount, 0)
	for category, count := range categoryCountMap {
		results = append(results, &model.CategoryCount{
			Category: category,
			Count:    count,
		})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Count > results[j].Count
	})

	groupCountResult := &model.GroupCountResult{
		Results: results,
		Total:   totalCount,
	}
	return groupCountResult, nil
}

/**
 * GetEventsWithFilter filters events by filterInput. For current implementation, we query the database for all events and filter them in memory.
 * In the future, we should implement a filter that can be translated into a SQL query.
 * @param {string} filterInput For current implementation, we use expr
 */
func (s Event) getEventsWithFilter(ctx context.Context, researchID string, filterInput string) ([]*model.Event, error) {
	events, err := s.GetEventsByResearchID(ctx, researchID)
	if err != nil {
		return nil, err
	}
	if filterInput == "" {
		return events, nil
	}
	exprRunner := exprutils.GetExprRunner()
	filteredEvents := make([]*model.Event, 0)
	for _, event := range events {
		output, err := exprRunner.RunCode(filterInput, exprRunner.PrepareEnv(event))
		if err != nil {
			return nil, err
		}
		if output == nil {
			continue
		}
		if output.(bool) {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents, nil
}

func (s Event) mapEventToResult(event *model.Event, resultMappingInput string) ([]interface{}, error) {
	exprRunner := exprutils.GetExprRunner()
	output, err := exprRunner.RunCode(resultMappingInput, exprRunner.PrepareEnv(event))
	if err != nil {
		return nil, err
	}
	if output == nil {
		return nil, nil
	}

	mappedResults := make([]interface{}, 0)
	if isArray(output) {
		mappedResults = output.([]interface{})
	} else {
		mappedResults = append(mappedResults, output)
	}
	// if result is not hashable, convert it to string
	for i, result := range mappedResults {
		if !isHashable(result) {
			mappedResults[i] = fmt.Sprintf("%v", result)
		}
	}

	return mappedResults, nil
}

func (s Event) convertFromEventInputToEvent(ctx context.Context, input *model.NewEvent) (*model.Event, error) {
	user, err := s.AuthService.CurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	// get schema from research
	research, err := s.ResearchService.GetResearchByID(ctx, input.ResearchID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("research not found")
		}
		return nil, err
	}

	// validate event json
	schema, err := json.Marshal(research.Schema)
	if err != nil {
		return nil, err
	}
	sch, err := jsonschema.CompileString("schema.json", string(schema))
	if err != nil {
		return nil, err
	}
	if err = sch.Validate(input.Content); err != nil {
		return nil, err
	}

	// FIXME: should use a global snowflake node or something like an ID generator
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	event := &model.Event{
		ID:         node.Generate().String(),
		ResearchID: input.ResearchID,
		Content:    input.Content,
		UserID:     user.ID,
		CreatedAt:  time.Now(),
		UserAgent:  input.UserAgent,
	}
	return event, nil
}

func isHashable(v interface{}) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice, reflect.Map, reflect.Func:
		return false
	default:
		return true
	}
}

func isArray(input interface{}) bool {
	kind := reflect.TypeOf(input).Kind()
	return kind == reflect.Array || kind == reflect.Slice
}
