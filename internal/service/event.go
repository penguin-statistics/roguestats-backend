package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"

	"github.com/pkg/errors"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"go.uber.org/fx"

	"exusiai.dev/roguestats-backend/internal/ent"
	"exusiai.dev/roguestats-backend/internal/ent/event"
	"exusiai.dev/roguestats-backend/internal/ent/research"
	"exusiai.dev/roguestats-backend/internal/exprutils"
	"exusiai.dev/roguestats-backend/internal/model"
)

type Event struct {
	fx.In

	Ent         *ent.Client
	AuthService Auth
}

func (s Event) CreateEventFromInput(ctx context.Context, input model.CreateEventInput) (*ent.Event, error) {
	client := ent.FromContext(ctx)

	user, err := s.AuthService.CurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	// get schema from research
	research, err := client.Research.Get(ctx, input.ResearchID)
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

	return client.Event.Create().
		SetContent(input.Content).
		SetUserAgent(input.UserAgent).
		SetResearchID(input.ResearchID).
		SetUserID(user.ID).
		Save(ctx)
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
func (s Event) getEventsWithFilter(ctx context.Context, researchID string, filterInput string) ([]*ent.Event, error) {
	events, err := s.Ent.Event.Query().
		Where(event.HasResearchWith(research.ID(researchID))).
		All(ctx)
	if err != nil {
		return nil, err
	}
	if filterInput == "" {
		return events, nil
	}
	exprRunner := exprutils.GetExprRunner()
	filteredEvents := make([]*ent.Event, 0)
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

func (s Event) mapEventToResult(event *ent.Event, resultMappingInput string) ([]interface{}, error) {
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
