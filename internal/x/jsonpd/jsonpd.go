package jsonpd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/lib/pq"
	"github.com/samber/lo"

	"exusiai.dev/roguestats-backend/internal/ent/predicate"
)

type Primitive = any

// Predicate is a map of string to any.
// It can contain one and ONLY one of the following keys:
// - $and: an array of Predicates
// - $or: an array of Predicates
// - $not: a Predicate
// - (any key that does not start with $. this is to be the name of the field to be compared): another object layer, which contains:
//   - a $ comparison operator (e.g. $eq, $ne, $gt, $ge, $lt, $le, $in, $contains)
//   - a Primitive value, or an array of Primitive values (for $in)
type Predicate map[string]any

func Parse(jsonBytes []byte) (*Predicate, error) {
	var predicate Predicate
	err := json.Unmarshal(jsonBytes, &predicate)
	if err != nil {
		return nil, err
	}
	return &predicate, nil
}

func (p Predicate) SQL(column string) (string, error) {
	return p.sql(column, "")
}

func (p Predicate) EntEventPredicate(column string) (predicate.Event, error) {
	sqlStr, err := p.SQL(column)
	if err != nil {
		return nil, err
	}
	return predicate.Event(func(s *sql.Selector) {
		s.Where((&sql.Predicate{}).Append(func(b *sql.Builder) {
			b.WriteString(sqlStr)
		}))
	}), nil
}

func primitiveSql(primitive Primitive) (string, error) {
	var v string
	switch primitive.(type) {
	// values
	case string:
		v = fmt.Sprintf("%q", primitive.(string))
	case json.Number:
		n, err := convertToNumber(primitive.(json.Number))
		if err != nil {
			return "", err
		}
		v = fmt.Sprintf("%v", n)
	case bool:
		v = strconv.FormatBool(primitive.(bool))
	case nil:
		v = "null"
	// // arrays
	// case []string:
	// 	// build a string of the form ['a','b','c']
	// 	var sb strings.Builder
	// 	sb.WriteRune('[')
	// 	for i, s := range primitive.([]string) {
	// 		if i > 0 {
	// 			sb.WriteRune(',')
	// 		}
	// 		sb.WriteString(fmt.Sprintf("%q", s))
	// 	}
	// 	sb.WriteRune(']')
	// 	return sb.String(), nil
	// case []int:
	// 	// build a string of the form [1,2,3]
	// 	var sb strings.Builder
	// 	sb.WriteRune('[')
	// 	for i, n := range primitive.([]int) {
	// 		if i > 0 {
	// 			sb.WriteRune(',')
	// 		}
	// 		sb.WriteString(strconv.Itoa(n))
	// 	}
	// 	sb.WriteRune(']')
	// 	return sb.String(), nil
	// case []float64:
	// 	// build a string of the form [1.1,2.2,3.3]
	// 	var sb strings.Builder
	// 	sb.WriteRune('[')
	// 	for i, n := range primitive.([]float64) {
	// 		if i > 0 {
	// 			sb.WriteRune(',')
	// 		}
	// 		sb.WriteString(strconv.FormatFloat(n, 'f', -1, 64))
	// 	}
	// 	sb.WriteRune(']')
	// 	return sb.String(), nil
	default:
		return "", fmt.Errorf("unsupported primitive type: %T", primitive)
	}
	return pq.QuoteLiteral(v) + "::jsonb", nil
}

// sql recursively builds the SQL query from the predicate.
// The SQL query is running on a PostgreSQL database against a JSONB column `column`.
// The SQL query is returned as a string.
// You should use the proper PostgreSQL JSONB operators to query the JSONB column.
func (p Predicate) sql(column string, field string) (string, error) {
	var sb sql.Predicate
	if len(p) == 0 {
		return "", fmt.Errorf("empty predicate")
	}
	if len(p) > 1 {
		return "", fmt.Errorf("predicate should contain only one key")
	}

	key := ""
	value := any(nil)
	for k, v := range p {
		key = k
		value = v
	}

	switch key {
	case "$and", "$or":
		joinClause := lo.Ternary(key == "$and", " AND ", " OR ")
		predicates := value.([]any)
		for i, predicate := range predicates {
			if i > 0 {
				sb.S(joinClause)
			}
			subPredicate := Predicate(predicate.(map[string]any))
			subSql, err := subPredicate.sql(column, "")
			if err != nil {
				return "", err
			}
			sb.Wrap(func(b *sql.Builder) {
				b.WriteString(subSql)
			})
		}
	case "$not":
		subPredicate := Predicate(value.(map[string]any))
		subSql, err := subPredicate.sql(column, "")
		if err != nil {
			return "", err
		}
		sb.WriteString("NOT (")
		sb.WriteString(subSql)
		sb.WriteByte(')')
	case "$eq", "$ne", "$gt", "$ge", "$lt", "$le", "$in", "$contains":
		if field == "" {
			return "", fmt.Errorf("field is empty")
		}
		// build the SQL query for the comparison operator
		// the value is a Primitive
		primitive := Primitive(value)
		primitiveSql, err := primitiveSql(primitive)
		if err != nil {
			return "", err
		}
		fieldSel := pq.QuoteIdentifier(column) + "->" + pq.QuoteLiteral(field)
		switch key {
		case "$eq":
			sb.S(fieldSel).S("=").S(primitiveSql)
		case "$ne":
			sb.S(fieldSel).S("<>").S(primitiveSql)
		case "$gt":
			sb.S(fieldSel).S(">").S(primitiveSql)
		case "$ge":
			sb.S(fieldSel).S(">=").S(primitiveSql)
		case "$lt":
			sb.S(fieldSel).S("<").S(primitiveSql)
		case "$le":
			sb.S(fieldSel).S("<=").S(primitiveSql)
		case "$in":
			sb.S(fieldSel).S("IN").S(primitiveSql)
		case "$contains":
			sb.S(fieldSel).S("@>").S(primitiveSql)
		}
	default:
		// build the SQL query for the field
		// the value is a Predicate
		subPredicate := Predicate(value.(map[string]any))
		subSql, err := subPredicate.sql(column, key)
		if err != nil {
			return "", err
		}
		sb.WriteString(subSql)
	}

	return sb.String(), nil
}

func convertToNumber(input json.Number) (interface{}, error) {
	if i, err := input.Int64(); err == nil {
		return i, nil
	}
	if f, err := input.Float64(); err == nil {
		return f, nil
	}
	return nil, strconv.ErrSyntax
}
