// Code generated by ent, DO NOT EDIT.

package event

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"exusiai.dev/roguestats-backend/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// UserAgent applies equality check predicate on the "user_agent" field. It's identical to UserAgentEQ.
func UserAgent(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUserAgent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtContains applies the Contains predicate on the "created_at" field.
func CreatedAtContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldCreatedAt, v))
}

// CreatedAtHasPrefix applies the HasPrefix predicate on the "created_at" field.
func CreatedAtHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldCreatedAt, v))
}

// CreatedAtHasSuffix applies the HasSuffix predicate on the "created_at" field.
func CreatedAtHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldCreatedAt, v))
}

// CreatedAtEqualFold applies the EqualFold predicate on the "created_at" field.
func CreatedAtEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldCreatedAt, v))
}

// CreatedAtContainsFold applies the ContainsFold predicate on the "created_at" field.
func CreatedAtContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldCreatedAt, v))
}

// UserAgentEQ applies the EQ predicate on the "user_agent" field.
func UserAgentEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUserAgent, v))
}

// UserAgentNEQ applies the NEQ predicate on the "user_agent" field.
func UserAgentNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldUserAgent, v))
}

// UserAgentIn applies the In predicate on the "user_agent" field.
func UserAgentIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldUserAgent, vs...))
}

// UserAgentNotIn applies the NotIn predicate on the "user_agent" field.
func UserAgentNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldUserAgent, vs...))
}

// UserAgentGT applies the GT predicate on the "user_agent" field.
func UserAgentGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldUserAgent, v))
}

// UserAgentGTE applies the GTE predicate on the "user_agent" field.
func UserAgentGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldUserAgent, v))
}

// UserAgentLT applies the LT predicate on the "user_agent" field.
func UserAgentLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldUserAgent, v))
}

// UserAgentLTE applies the LTE predicate on the "user_agent" field.
func UserAgentLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldUserAgent, v))
}

// UserAgentContains applies the Contains predicate on the "user_agent" field.
func UserAgentContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldUserAgent, v))
}

// UserAgentHasPrefix applies the HasPrefix predicate on the "user_agent" field.
func UserAgentHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldUserAgent, v))
}

// UserAgentHasSuffix applies the HasSuffix predicate on the "user_agent" field.
func UserAgentHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldUserAgent, v))
}

// UserAgentEqualFold applies the EqualFold predicate on the "user_agent" field.
func UserAgentEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldUserAgent, v))
}

// UserAgentContainsFold applies the ContainsFold predicate on the "user_agent" field.
func UserAgentContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldUserAgent, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasResearch applies the HasEdge predicate on the "research" edge.
func HasResearch() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ResearchTable, ResearchColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResearchWith applies the HasEdge predicate on the "research" edge with a given conditions (other predicates).
func HasResearchWith(preds ...predicate.Research) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newResearchStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		p(s.Not())
	})
}
