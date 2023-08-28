// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"exusiai.dev/roguestats-backend/internal/ent/event"
	"exusiai.dev/roguestats-backend/internal/ent/research"
	"exusiai.dev/roguestats-backend/internal/ent/user"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (e *EventQuery) CollectFields(ctx context.Context, satisfies ...string) (*EventQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return e, nil
	}
	if err := e.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return e, nil
}

func (e *EventQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(event.Columns))
		selectedFields = []string{event.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: e.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			e.withUser = query
			if _, ok := fieldSeen[event.FieldUserID]; !ok {
				selectedFields = append(selectedFields, event.FieldUserID)
				fieldSeen[event.FieldUserID] = struct{}{}
			}
		case "research":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ResearchClient{config: e.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, researchImplementors)...); err != nil {
				return err
			}
			e.withResearch = query
			if _, ok := fieldSeen[event.FieldResearchID]; !ok {
				selectedFields = append(selectedFields, event.FieldResearchID)
				fieldSeen[event.FieldResearchID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[event.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, event.FieldCreatedAt)
				fieldSeen[event.FieldCreatedAt] = struct{}{}
			}
		case "userID":
			if _, ok := fieldSeen[event.FieldUserID]; !ok {
				selectedFields = append(selectedFields, event.FieldUserID)
				fieldSeen[event.FieldUserID] = struct{}{}
			}
		case "researchID":
			if _, ok := fieldSeen[event.FieldResearchID]; !ok {
				selectedFields = append(selectedFields, event.FieldResearchID)
				fieldSeen[event.FieldResearchID] = struct{}{}
			}
		case "userAgent":
			if _, ok := fieldSeen[event.FieldUserAgent]; !ok {
				selectedFields = append(selectedFields, event.FieldUserAgent)
				fieldSeen[event.FieldUserAgent] = struct{}{}
			}
		case "content":
			if _, ok := fieldSeen[event.FieldContent]; !ok {
				selectedFields = append(selectedFields, event.FieldContent)
				fieldSeen[event.FieldContent] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		e.Select(selectedFields...)
	}
	return nil
}

type eventPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []EventPaginateOption
}

func newEventPaginateArgs(rv map[string]any) *eventPaginateArgs {
	args := &eventPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &EventOrder{Field: &EventOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithEventOrder(order))
			}
		case *EventOrder:
			if v != nil {
				args.opts = append(args.opts, WithEventOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *ResearchQuery) CollectFields(ctx context.Context, satisfies ...string) (*ResearchQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return r, nil
	}
	if err := r.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *ResearchQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(research.Columns))
		selectedFields = []string{research.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "events":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&EventClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, eventImplementors)...); err != nil {
				return err
			}
			r.WithNamedEvents(alias, func(wq *EventQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[research.FieldName]; !ok {
				selectedFields = append(selectedFields, research.FieldName)
				fieldSeen[research.FieldName] = struct{}{}
			}
		case "schema":
			if _, ok := fieldSeen[research.FieldSchema]; !ok {
				selectedFields = append(selectedFields, research.FieldSchema)
				fieldSeen[research.FieldSchema] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		r.Select(selectedFields...)
	}
	return nil
}

type researchPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ResearchPaginateOption
}

func newResearchPaginateArgs(rv map[string]any) *researchPaginateArgs {
	args := &researchPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &ResearchOrder{Field: &ResearchOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithResearchOrder(order))
			}
		case *ResearchOrder:
			if v != nil {
				args.opts = append(args.opts, WithResearchOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "events":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&EventClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, eventImplementors)...); err != nil {
				return err
			}
			u.WithNamedEvents(alias, func(wq *EventQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[user.FieldName]; !ok {
				selectedFields = append(selectedFields, user.FieldName)
				fieldSeen[user.FieldName] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "credential":
			if _, ok := fieldSeen[user.FieldCredential]; !ok {
				selectedFields = append(selectedFields, user.FieldCredential)
				fieldSeen[user.FieldCredential] = struct{}{}
			}
		case "attributes":
			if _, ok := fieldSeen[user.FieldAttributes]; !ok {
				selectedFields = append(selectedFields, user.FieldAttributes)
				fieldSeen[user.FieldAttributes] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
