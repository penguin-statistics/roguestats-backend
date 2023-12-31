// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"exusiai.dev/roguestats-backend/internal/ent/event"
	"exusiai.dev/roguestats-backend/internal/ent/querypreset"
	"exusiai.dev/roguestats-backend/internal/ent/research"
	"exusiai.dev/roguestats-backend/internal/ent/user"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[string]
	PageInfo       = entgql.PageInfo[string]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// EventEdge is the edge representation of Event.
type EventEdge struct {
	Node   *Event `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// EventConnection is the connection containing edges to Event.
type EventConnection struct {
	Edges      []*EventEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *EventConnection) build(nodes []*Event, pager *eventPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Event
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Event {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Event {
			return nodes[i]
		}
	}
	c.Edges = make([]*EventEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &EventEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// EventPaginateOption enables pagination customization.
type EventPaginateOption func(*eventPager) error

// WithEventOrder configures pagination ordering.
func WithEventOrder(order *EventOrder) EventPaginateOption {
	if order == nil {
		order = DefaultEventOrder
	}
	o := *order
	return func(pager *eventPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultEventOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithEventFilter configures pagination filter.
func WithEventFilter(filter func(*EventQuery) (*EventQuery, error)) EventPaginateOption {
	return func(pager *eventPager) error {
		if filter == nil {
			return errors.New("EventQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type eventPager struct {
	reverse bool
	order   *EventOrder
	filter  func(*EventQuery) (*EventQuery, error)
}

func newEventPager(opts []EventPaginateOption, reverse bool) (*eventPager, error) {
	pager := &eventPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultEventOrder
	}
	return pager, nil
}

func (p *eventPager) applyFilter(query *EventQuery) (*EventQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *eventPager) toCursor(e *Event) Cursor {
	return p.order.Field.toCursor(e)
}

func (p *eventPager) applyCursors(query *EventQuery, after, before *Cursor) (*EventQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultEventOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *eventPager) applyOrder(query *EventQuery) *EventQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultEventOrder.Field {
		query = query.Order(DefaultEventOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *eventPager) orderExpr(query *EventQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultEventOrder.Field {
			b.Comma().Ident(DefaultEventOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Event.
func (e *EventQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...EventPaginateOption,
) (*EventConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newEventPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if e, err = pager.applyFilter(e); err != nil {
		return nil, err
	}
	conn := &EventConnection{Edges: []*EventEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := e.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if e, err = pager.applyCursors(e, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		e.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := e.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	e = pager.applyOrder(e)
	nodes, err := e.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// EventOrderFieldID orders Event by id.
	EventOrderFieldID = &EventOrderField{
		Value: func(e *Event) (ent.Value, error) {
			return e.ID, nil
		},
		column: event.FieldID,
		toTerm: event.ByID,
		toCursor: func(e *Event) Cursor {
			return Cursor{
				ID:    e.ID,
				Value: e.ID,
			}
		},
	}
	// EventOrderFieldCreatedAt orders Event by created_at.
	EventOrderFieldCreatedAt = &EventOrderField{
		Value: func(e *Event) (ent.Value, error) {
			return e.CreatedAt, nil
		},
		column: event.FieldCreatedAt,
		toTerm: event.ByCreatedAt,
		toCursor: func(e *Event) Cursor {
			return Cursor{
				ID:    e.ID,
				Value: e.CreatedAt,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f EventOrderField) String() string {
	var str string
	switch f.column {
	case EventOrderFieldID.column:
		str = "ID"
	case EventOrderFieldCreatedAt.column:
		str = "CREATED_AT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f EventOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *EventOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("EventOrderField %T must be a string", v)
	}
	switch str {
	case "ID":
		*f = *EventOrderFieldID
	case "CREATED_AT":
		*f = *EventOrderFieldCreatedAt
	default:
		return fmt.Errorf("%s is not a valid EventOrderField", str)
	}
	return nil
}

// EventOrderField defines the ordering field of Event.
type EventOrderField struct {
	// Value extracts the ordering value from the given Event.
	Value    func(*Event) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) event.OrderOption
	toCursor func(*Event) Cursor
}

// EventOrder defines the ordering of Event.
type EventOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *EventOrderField `json:"field"`
}

// DefaultEventOrder is the default ordering of Event.
var DefaultEventOrder = &EventOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &EventOrderField{
		Value: func(e *Event) (ent.Value, error) {
			return e.ID, nil
		},
		column: event.FieldID,
		toTerm: event.ByID,
		toCursor: func(e *Event) Cursor {
			return Cursor{ID: e.ID}
		},
	},
}

// ToEdge converts Event into EventEdge.
func (e *Event) ToEdge(order *EventOrder) *EventEdge {
	if order == nil {
		order = DefaultEventOrder
	}
	return &EventEdge{
		Node:   e,
		Cursor: order.Field.toCursor(e),
	}
}

// QueryPresetEdge is the edge representation of QueryPreset.
type QueryPresetEdge struct {
	Node   *QueryPreset `json:"node"`
	Cursor Cursor       `json:"cursor"`
}

// QueryPresetConnection is the connection containing edges to QueryPreset.
type QueryPresetConnection struct {
	Edges      []*QueryPresetEdge `json:"edges"`
	PageInfo   PageInfo           `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

func (c *QueryPresetConnection) build(nodes []*QueryPreset, pager *querypresetPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *QueryPreset
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *QueryPreset {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *QueryPreset {
			return nodes[i]
		}
	}
	c.Edges = make([]*QueryPresetEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &QueryPresetEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// QueryPresetPaginateOption enables pagination customization.
type QueryPresetPaginateOption func(*querypresetPager) error

// WithQueryPresetOrder configures pagination ordering.
func WithQueryPresetOrder(order *QueryPresetOrder) QueryPresetPaginateOption {
	if order == nil {
		order = DefaultQueryPresetOrder
	}
	o := *order
	return func(pager *querypresetPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultQueryPresetOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithQueryPresetFilter configures pagination filter.
func WithQueryPresetFilter(filter func(*QueryPresetQuery) (*QueryPresetQuery, error)) QueryPresetPaginateOption {
	return func(pager *querypresetPager) error {
		if filter == nil {
			return errors.New("QueryPresetQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type querypresetPager struct {
	reverse bool
	order   *QueryPresetOrder
	filter  func(*QueryPresetQuery) (*QueryPresetQuery, error)
}

func newQueryPresetPager(opts []QueryPresetPaginateOption, reverse bool) (*querypresetPager, error) {
	pager := &querypresetPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultQueryPresetOrder
	}
	return pager, nil
}

func (p *querypresetPager) applyFilter(query *QueryPresetQuery) (*QueryPresetQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *querypresetPager) toCursor(qp *QueryPreset) Cursor {
	return p.order.Field.toCursor(qp)
}

func (p *querypresetPager) applyCursors(query *QueryPresetQuery, after, before *Cursor) (*QueryPresetQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultQueryPresetOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *querypresetPager) applyOrder(query *QueryPresetQuery) *QueryPresetQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultQueryPresetOrder.Field {
		query = query.Order(DefaultQueryPresetOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *querypresetPager) orderExpr(query *QueryPresetQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultQueryPresetOrder.Field {
			b.Comma().Ident(DefaultQueryPresetOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to QueryPreset.
func (qp *QueryPresetQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...QueryPresetPaginateOption,
) (*QueryPresetConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newQueryPresetPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if qp, err = pager.applyFilter(qp); err != nil {
		return nil, err
	}
	conn := &QueryPresetConnection{Edges: []*QueryPresetEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := qp.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if qp, err = pager.applyCursors(qp, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		qp.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := qp.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	qp = pager.applyOrder(qp)
	nodes, err := qp.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// QueryPresetOrderFieldID orders QueryPreset by id.
	QueryPresetOrderFieldID = &QueryPresetOrderField{
		Value: func(qp *QueryPreset) (ent.Value, error) {
			return qp.ID, nil
		},
		column: querypreset.FieldID,
		toTerm: querypreset.ByID,
		toCursor: func(qp *QueryPreset) Cursor {
			return Cursor{
				ID:    qp.ID,
				Value: qp.ID,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f QueryPresetOrderField) String() string {
	var str string
	switch f.column {
	case QueryPresetOrderFieldID.column:
		str = "ID"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f QueryPresetOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *QueryPresetOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("QueryPresetOrderField %T must be a string", v)
	}
	switch str {
	case "ID":
		*f = *QueryPresetOrderFieldID
	default:
		return fmt.Errorf("%s is not a valid QueryPresetOrderField", str)
	}
	return nil
}

// QueryPresetOrderField defines the ordering field of QueryPreset.
type QueryPresetOrderField struct {
	// Value extracts the ordering value from the given QueryPreset.
	Value    func(*QueryPreset) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) querypreset.OrderOption
	toCursor func(*QueryPreset) Cursor
}

// QueryPresetOrder defines the ordering of QueryPreset.
type QueryPresetOrder struct {
	Direction OrderDirection         `json:"direction"`
	Field     *QueryPresetOrderField `json:"field"`
}

// DefaultQueryPresetOrder is the default ordering of QueryPreset.
var DefaultQueryPresetOrder = &QueryPresetOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &QueryPresetOrderField{
		Value: func(qp *QueryPreset) (ent.Value, error) {
			return qp.ID, nil
		},
		column: querypreset.FieldID,
		toTerm: querypreset.ByID,
		toCursor: func(qp *QueryPreset) Cursor {
			return Cursor{ID: qp.ID}
		},
	},
}

// ToEdge converts QueryPreset into QueryPresetEdge.
func (qp *QueryPreset) ToEdge(order *QueryPresetOrder) *QueryPresetEdge {
	if order == nil {
		order = DefaultQueryPresetOrder
	}
	return &QueryPresetEdge{
		Node:   qp,
		Cursor: order.Field.toCursor(qp),
	}
}

// ResearchEdge is the edge representation of Research.
type ResearchEdge struct {
	Node   *Research `json:"node"`
	Cursor Cursor    `json:"cursor"`
}

// ResearchConnection is the connection containing edges to Research.
type ResearchConnection struct {
	Edges      []*ResearchEdge `json:"edges"`
	PageInfo   PageInfo        `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

func (c *ResearchConnection) build(nodes []*Research, pager *researchPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Research
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Research {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Research {
			return nodes[i]
		}
	}
	c.Edges = make([]*ResearchEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &ResearchEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// ResearchPaginateOption enables pagination customization.
type ResearchPaginateOption func(*researchPager) error

// WithResearchOrder configures pagination ordering.
func WithResearchOrder(order *ResearchOrder) ResearchPaginateOption {
	if order == nil {
		order = DefaultResearchOrder
	}
	o := *order
	return func(pager *researchPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultResearchOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithResearchFilter configures pagination filter.
func WithResearchFilter(filter func(*ResearchQuery) (*ResearchQuery, error)) ResearchPaginateOption {
	return func(pager *researchPager) error {
		if filter == nil {
			return errors.New("ResearchQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type researchPager struct {
	reverse bool
	order   *ResearchOrder
	filter  func(*ResearchQuery) (*ResearchQuery, error)
}

func newResearchPager(opts []ResearchPaginateOption, reverse bool) (*researchPager, error) {
	pager := &researchPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultResearchOrder
	}
	return pager, nil
}

func (p *researchPager) applyFilter(query *ResearchQuery) (*ResearchQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *researchPager) toCursor(r *Research) Cursor {
	return p.order.Field.toCursor(r)
}

func (p *researchPager) applyCursors(query *ResearchQuery, after, before *Cursor) (*ResearchQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultResearchOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *researchPager) applyOrder(query *ResearchQuery) *ResearchQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultResearchOrder.Field {
		query = query.Order(DefaultResearchOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *researchPager) orderExpr(query *ResearchQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultResearchOrder.Field {
			b.Comma().Ident(DefaultResearchOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Research.
func (r *ResearchQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...ResearchPaginateOption,
) (*ResearchConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newResearchPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if r, err = pager.applyFilter(r); err != nil {
		return nil, err
	}
	conn := &ResearchConnection{Edges: []*ResearchEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := r.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if r, err = pager.applyCursors(r, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		r.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := r.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	r = pager.applyOrder(r)
	nodes, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// ResearchOrderFieldID orders Research by id.
	ResearchOrderFieldID = &ResearchOrderField{
		Value: func(r *Research) (ent.Value, error) {
			return r.ID, nil
		},
		column: research.FieldID,
		toTerm: research.ByID,
		toCursor: func(r *Research) Cursor {
			return Cursor{
				ID:    r.ID,
				Value: r.ID,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f ResearchOrderField) String() string {
	var str string
	switch f.column {
	case ResearchOrderFieldID.column:
		str = "ID"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f ResearchOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *ResearchOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("ResearchOrderField %T must be a string", v)
	}
	switch str {
	case "ID":
		*f = *ResearchOrderFieldID
	default:
		return fmt.Errorf("%s is not a valid ResearchOrderField", str)
	}
	return nil
}

// ResearchOrderField defines the ordering field of Research.
type ResearchOrderField struct {
	// Value extracts the ordering value from the given Research.
	Value    func(*Research) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) research.OrderOption
	toCursor func(*Research) Cursor
}

// ResearchOrder defines the ordering of Research.
type ResearchOrder struct {
	Direction OrderDirection      `json:"direction"`
	Field     *ResearchOrderField `json:"field"`
}

// DefaultResearchOrder is the default ordering of Research.
var DefaultResearchOrder = &ResearchOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &ResearchOrderField{
		Value: func(r *Research) (ent.Value, error) {
			return r.ID, nil
		},
		column: research.FieldID,
		toTerm: research.ByID,
		toCursor: func(r *Research) Cursor {
			return Cursor{ID: r.ID}
		},
	},
}

// ToEdge converts Research into ResearchEdge.
func (r *Research) ToEdge(order *ResearchOrder) *ResearchEdge {
	if order == nil {
		order = DefaultResearchOrder
	}
	return &ResearchEdge{
		Node:   r,
		Cursor: order.Field.toCursor(r),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	reverse bool
	order   *UserOrder
	filter  func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption, reverse bool) (*userPager, error) {
	pager := &userPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) (*UserQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultUserOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *userPager) applyOrder(query *UserQuery) *UserQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(DefaultUserOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *userPager) orderExpr(query *UserQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := u.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if u, err = pager.applyCursors(u, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	u = pager.applyOrder(u)
	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	// Value extracts the ordering value from the given User.
	Value    func(*User) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) user.OrderOption
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.ID, nil
		},
		column: user.FieldID,
		toTerm: user.ByID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
