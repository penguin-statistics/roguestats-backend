// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"time"

	"exusiai.dev/roguestats-backend/internal/ent/event"
	"exusiai.dev/roguestats-backend/internal/ent/predicate"
	"exusiai.dev/roguestats-backend/internal/ent/querypreset"
	"exusiai.dev/roguestats-backend/internal/ent/research"
	"exusiai.dev/roguestats-backend/internal/ent/user"
)

// EventWhereInput represents a where input for filtering Event queries.
type EventWhereInput struct {
	Predicates []predicate.Event  `json:"-"`
	Not        *EventWhereInput   `json:"not,omitempty"`
	Or         []*EventWhereInput `json:"or,omitempty"`
	And        []*EventWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID             *string  `json:"id,omitempty"`
	IDNEQ          *string  `json:"idNEQ,omitempty"`
	IDIn           []string `json:"idIn,omitempty"`
	IDNotIn        []string `json:"idNotIn,omitempty"`
	IDGT           *string  `json:"idGT,omitempty"`
	IDGTE          *string  `json:"idGTE,omitempty"`
	IDLT           *string  `json:"idLT,omitempty"`
	IDLTE          *string  `json:"idLTE,omitempty"`
	IDEqualFold    *string  `json:"idEqualFold,omitempty"`
	IDContainsFold *string  `json:"idContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "user" edge predicates.
	HasUser     *bool             `json:"hasUser,omitempty"`
	HasUserWith []*UserWhereInput `json:"hasUserWith,omitempty"`

	// "research" edge predicates.
	HasResearch     *bool                 `json:"hasResearch,omitempty"`
	HasResearchWith []*ResearchWhereInput `json:"hasResearchWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *EventWhereInput) AddPredicates(predicates ...predicate.Event) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the EventWhereInput filter on the EventQuery builder.
func (i *EventWhereInput) Filter(q *EventQuery) (*EventQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyEventWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyEventWhereInput is returned in case the EventWhereInput is empty.
var ErrEmptyEventWhereInput = errors.New("ent: empty predicate EventWhereInput")

// P returns a predicate for filtering events.
// An error is returned if the input is empty or invalid.
func (i *EventWhereInput) P() (predicate.Event, error) {
	var predicates []predicate.Event
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, event.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Event, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, event.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Event, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, event.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, event.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, event.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, event.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, event.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, event.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, event.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, event.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, event.IDLTE(*i.IDLTE))
	}
	if i.IDEqualFold != nil {
		predicates = append(predicates, event.IDEqualFold(*i.IDEqualFold))
	}
	if i.IDContainsFold != nil {
		predicates = append(predicates, event.IDContainsFold(*i.IDContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, event.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, event.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, event.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, event.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, event.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, event.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, event.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, event.CreatedAtLTE(*i.CreatedAtLTE))
	}

	if i.HasUser != nil {
		p := event.HasUser()
		if !*i.HasUser {
			p = event.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUserWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasUserWith))
		for _, w := range i.HasUserWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasUserWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, event.HasUserWith(with...))
	}
	if i.HasResearch != nil {
		p := event.HasResearch()
		if !*i.HasResearch {
			p = event.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasResearchWith) > 0 {
		with := make([]predicate.Research, 0, len(i.HasResearchWith))
		for _, w := range i.HasResearchWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasResearchWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, event.HasResearchWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyEventWhereInput
	case 1:
		return predicates[0], nil
	default:
		return event.And(predicates...), nil
	}
}

// QueryPresetWhereInput represents a where input for filtering QueryPreset queries.
type QueryPresetWhereInput struct {
	Predicates []predicate.QueryPreset  `json:"-"`
	Not        *QueryPresetWhereInput   `json:"not,omitempty"`
	Or         []*QueryPresetWhereInput `json:"or,omitempty"`
	And        []*QueryPresetWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID             *string  `json:"id,omitempty"`
	IDNEQ          *string  `json:"idNEQ,omitempty"`
	IDIn           []string `json:"idIn,omitempty"`
	IDNotIn        []string `json:"idNotIn,omitempty"`
	IDGT           *string  `json:"idGT,omitempty"`
	IDGTE          *string  `json:"idGTE,omitempty"`
	IDLT           *string  `json:"idLT,omitempty"`
	IDLTE          *string  `json:"idLTE,omitempty"`
	IDEqualFold    *string  `json:"idEqualFold,omitempty"`
	IDContainsFold *string  `json:"idContainsFold,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "mapping" field predicates.
	Mapping             *string  `json:"mapping,omitempty"`
	MappingNEQ          *string  `json:"mappingNEQ,omitempty"`
	MappingIn           []string `json:"mappingIn,omitempty"`
	MappingNotIn        []string `json:"mappingNotIn,omitempty"`
	MappingGT           *string  `json:"mappingGT,omitempty"`
	MappingGTE          *string  `json:"mappingGTE,omitempty"`
	MappingLT           *string  `json:"mappingLT,omitempty"`
	MappingLTE          *string  `json:"mappingLTE,omitempty"`
	MappingContains     *string  `json:"mappingContains,omitempty"`
	MappingHasPrefix    *string  `json:"mappingHasPrefix,omitempty"`
	MappingHasSuffix    *string  `json:"mappingHasSuffix,omitempty"`
	MappingEqualFold    *string  `json:"mappingEqualFold,omitempty"`
	MappingContainsFold *string  `json:"mappingContainsFold,omitempty"`

	// "research" edge predicates.
	HasResearch     *bool                 `json:"hasResearch,omitempty"`
	HasResearchWith []*ResearchWhereInput `json:"hasResearchWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *QueryPresetWhereInput) AddPredicates(predicates ...predicate.QueryPreset) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the QueryPresetWhereInput filter on the QueryPresetQuery builder.
func (i *QueryPresetWhereInput) Filter(q *QueryPresetQuery) (*QueryPresetQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyQueryPresetWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyQueryPresetWhereInput is returned in case the QueryPresetWhereInput is empty.
var ErrEmptyQueryPresetWhereInput = errors.New("ent: empty predicate QueryPresetWhereInput")

// P returns a predicate for filtering querypresets.
// An error is returned if the input is empty or invalid.
func (i *QueryPresetWhereInput) P() (predicate.QueryPreset, error) {
	var predicates []predicate.QueryPreset
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, querypreset.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.QueryPreset, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, querypreset.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.QueryPreset, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, querypreset.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, querypreset.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, querypreset.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, querypreset.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, querypreset.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, querypreset.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, querypreset.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, querypreset.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, querypreset.IDLTE(*i.IDLTE))
	}
	if i.IDEqualFold != nil {
		predicates = append(predicates, querypreset.IDEqualFold(*i.IDEqualFold))
	}
	if i.IDContainsFold != nil {
		predicates = append(predicates, querypreset.IDContainsFold(*i.IDContainsFold))
	}
	if i.Name != nil {
		predicates = append(predicates, querypreset.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, querypreset.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, querypreset.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, querypreset.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, querypreset.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, querypreset.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, querypreset.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, querypreset.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, querypreset.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, querypreset.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, querypreset.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, querypreset.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, querypreset.NameContainsFold(*i.NameContainsFold))
	}
	if i.Mapping != nil {
		predicates = append(predicates, querypreset.MappingEQ(*i.Mapping))
	}
	if i.MappingNEQ != nil {
		predicates = append(predicates, querypreset.MappingNEQ(*i.MappingNEQ))
	}
	if len(i.MappingIn) > 0 {
		predicates = append(predicates, querypreset.MappingIn(i.MappingIn...))
	}
	if len(i.MappingNotIn) > 0 {
		predicates = append(predicates, querypreset.MappingNotIn(i.MappingNotIn...))
	}
	if i.MappingGT != nil {
		predicates = append(predicates, querypreset.MappingGT(*i.MappingGT))
	}
	if i.MappingGTE != nil {
		predicates = append(predicates, querypreset.MappingGTE(*i.MappingGTE))
	}
	if i.MappingLT != nil {
		predicates = append(predicates, querypreset.MappingLT(*i.MappingLT))
	}
	if i.MappingLTE != nil {
		predicates = append(predicates, querypreset.MappingLTE(*i.MappingLTE))
	}
	if i.MappingContains != nil {
		predicates = append(predicates, querypreset.MappingContains(*i.MappingContains))
	}
	if i.MappingHasPrefix != nil {
		predicates = append(predicates, querypreset.MappingHasPrefix(*i.MappingHasPrefix))
	}
	if i.MappingHasSuffix != nil {
		predicates = append(predicates, querypreset.MappingHasSuffix(*i.MappingHasSuffix))
	}
	if i.MappingEqualFold != nil {
		predicates = append(predicates, querypreset.MappingEqualFold(*i.MappingEqualFold))
	}
	if i.MappingContainsFold != nil {
		predicates = append(predicates, querypreset.MappingContainsFold(*i.MappingContainsFold))
	}

	if i.HasResearch != nil {
		p := querypreset.HasResearch()
		if !*i.HasResearch {
			p = querypreset.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasResearchWith) > 0 {
		with := make([]predicate.Research, 0, len(i.HasResearchWith))
		for _, w := range i.HasResearchWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasResearchWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, querypreset.HasResearchWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyQueryPresetWhereInput
	case 1:
		return predicates[0], nil
	default:
		return querypreset.And(predicates...), nil
	}
}

// ResearchWhereInput represents a where input for filtering Research queries.
type ResearchWhereInput struct {
	Predicates []predicate.Research  `json:"-"`
	Not        *ResearchWhereInput   `json:"not,omitempty"`
	Or         []*ResearchWhereInput `json:"or,omitempty"`
	And        []*ResearchWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID             *string  `json:"id,omitempty"`
	IDNEQ          *string  `json:"idNEQ,omitempty"`
	IDIn           []string `json:"idIn,omitempty"`
	IDNotIn        []string `json:"idNotIn,omitempty"`
	IDGT           *string  `json:"idGT,omitempty"`
	IDGTE          *string  `json:"idGTE,omitempty"`
	IDLT           *string  `json:"idLT,omitempty"`
	IDLTE          *string  `json:"idLTE,omitempty"`
	IDEqualFold    *string  `json:"idEqualFold,omitempty"`
	IDContainsFold *string  `json:"idContainsFold,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "events" edge predicates.
	HasEvents     *bool              `json:"hasEvents,omitempty"`
	HasEventsWith []*EventWhereInput `json:"hasEventsWith,omitempty"`

	// "query_presets" edge predicates.
	HasQueryPresets     *bool                    `json:"hasQueryPresets,omitempty"`
	HasQueryPresetsWith []*QueryPresetWhereInput `json:"hasQueryPresetsWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *ResearchWhereInput) AddPredicates(predicates ...predicate.Research) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the ResearchWhereInput filter on the ResearchQuery builder.
func (i *ResearchWhereInput) Filter(q *ResearchQuery) (*ResearchQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyResearchWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyResearchWhereInput is returned in case the ResearchWhereInput is empty.
var ErrEmptyResearchWhereInput = errors.New("ent: empty predicate ResearchWhereInput")

// P returns a predicate for filtering researches.
// An error is returned if the input is empty or invalid.
func (i *ResearchWhereInput) P() (predicate.Research, error) {
	var predicates []predicate.Research
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, research.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Research, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, research.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Research, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, research.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, research.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, research.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, research.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, research.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, research.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, research.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, research.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, research.IDLTE(*i.IDLTE))
	}
	if i.IDEqualFold != nil {
		predicates = append(predicates, research.IDEqualFold(*i.IDEqualFold))
	}
	if i.IDContainsFold != nil {
		predicates = append(predicates, research.IDContainsFold(*i.IDContainsFold))
	}
	if i.Name != nil {
		predicates = append(predicates, research.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, research.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, research.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, research.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, research.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, research.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, research.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, research.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, research.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, research.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, research.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, research.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, research.NameContainsFold(*i.NameContainsFold))
	}

	if i.HasEvents != nil {
		p := research.HasEvents()
		if !*i.HasEvents {
			p = research.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasEventsWith) > 0 {
		with := make([]predicate.Event, 0, len(i.HasEventsWith))
		for _, w := range i.HasEventsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasEventsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, research.HasEventsWith(with...))
	}
	if i.HasQueryPresets != nil {
		p := research.HasQueryPresets()
		if !*i.HasQueryPresets {
			p = research.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasQueryPresetsWith) > 0 {
		with := make([]predicate.QueryPreset, 0, len(i.HasQueryPresetsWith))
		for _, w := range i.HasQueryPresetsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasQueryPresetsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, research.HasQueryPresetsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyResearchWhereInput
	case 1:
		return predicates[0], nil
	default:
		return research.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID             *string  `json:"id,omitempty"`
	IDNEQ          *string  `json:"idNEQ,omitempty"`
	IDIn           []string `json:"idIn,omitempty"`
	IDNotIn        []string `json:"idNotIn,omitempty"`
	IDGT           *string  `json:"idGT,omitempty"`
	IDGTE          *string  `json:"idGTE,omitempty"`
	IDLT           *string  `json:"idLT,omitempty"`
	IDLTE          *string  `json:"idLTE,omitempty"`
	IDEqualFold    *string  `json:"idEqualFold,omitempty"`
	IDContainsFold *string  `json:"idContainsFold,omitempty"`

	// "events" edge predicates.
	HasEvents     *bool              `json:"hasEvents,omitempty"`
	HasEventsWith []*EventWhereInput `json:"hasEventsWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserWhereInput is returned in case the UserWhereInput is empty.
var ErrEmptyUserWhereInput = errors.New("ent: empty predicate UserWhereInput")

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.IDEqualFold != nil {
		predicates = append(predicates, user.IDEqualFold(*i.IDEqualFold))
	}
	if i.IDContainsFold != nil {
		predicates = append(predicates, user.IDContainsFold(*i.IDContainsFold))
	}

	if i.HasEvents != nil {
		p := user.HasEvents()
		if !*i.HasEvents {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasEventsWith) > 0 {
		with := make([]predicate.Event, 0, len(i.HasEventsWith))
		for _, w := range i.HasEventsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasEventsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasEventsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}
