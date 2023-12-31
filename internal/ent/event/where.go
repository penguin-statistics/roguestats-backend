// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

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
func CreatedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUserID, v))
}

// ResearchID applies equality check predicate on the "research_id" field. It's identical to ResearchIDEQ.
func ResearchID(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldResearchID, v))
}

// UserAgent applies equality check predicate on the "user_agent" field. It's identical to UserAgentEQ.
func UserAgent(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUserAgent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldUserID, v))
}

// ResearchIDEQ applies the EQ predicate on the "research_id" field.
func ResearchIDEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldResearchID, v))
}

// ResearchIDNEQ applies the NEQ predicate on the "research_id" field.
func ResearchIDNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldResearchID, v))
}

// ResearchIDIn applies the In predicate on the "research_id" field.
func ResearchIDIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldResearchID, vs...))
}

// ResearchIDNotIn applies the NotIn predicate on the "research_id" field.
func ResearchIDNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldResearchID, vs...))
}

// ResearchIDGT applies the GT predicate on the "research_id" field.
func ResearchIDGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldResearchID, v))
}

// ResearchIDGTE applies the GTE predicate on the "research_id" field.
func ResearchIDGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldResearchID, v))
}

// ResearchIDLT applies the LT predicate on the "research_id" field.
func ResearchIDLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldResearchID, v))
}

// ResearchIDLTE applies the LTE predicate on the "research_id" field.
func ResearchIDLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldResearchID, v))
}

// ResearchIDContains applies the Contains predicate on the "research_id" field.
func ResearchIDContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldResearchID, v))
}

// ResearchIDHasPrefix applies the HasPrefix predicate on the "research_id" field.
func ResearchIDHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldResearchID, v))
}

// ResearchIDHasSuffix applies the HasSuffix predicate on the "research_id" field.
func ResearchIDHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldResearchID, v))
}

// ResearchIDEqualFold applies the EqualFold predicate on the "research_id" field.
func ResearchIDEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldResearchID, v))
}

// ResearchIDContainsFold applies the ContainsFold predicate on the "research_id" field.
func ResearchIDContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldResearchID, v))
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
