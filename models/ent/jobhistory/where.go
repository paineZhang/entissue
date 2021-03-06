// Code generated by entc, DO NOT EDIT.

package jobhistory

import (
	"time"
	"wing/models/ent/internal"
	"wing/models/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateByUser applies equality check predicate on the "create_by_user" field. It's identical to CreateByUserEQ.
func CreateByUser(v int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateByUser), v))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// CreateByUserEQ applies the EQ predicate on the "create_by_user" field.
func CreateByUserEQ(v int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateByUser), v))
	})
}

// CreateByUserNEQ applies the NEQ predicate on the "create_by_user" field.
func CreateByUserNEQ(v int) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateByUser), v))
	})
}

// CreateByUserIn applies the In predicate on the "create_by_user" field.
func CreateByUserIn(vs ...int) predicate.JobHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobHistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateByUser), v...))
	})
}

// CreateByUserNotIn applies the NotIn predicate on the "create_by_user" field.
func CreateByUserNotIn(vs ...int) predicate.JobHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobHistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateByUser), v...))
	})
}

// CreateByUserIsNil applies the IsNil predicate on the "create_by_user" field.
func CreateByUserIsNil() predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateByUser)))
	})
}

// CreateByUserNotNil applies the NotNil predicate on the "create_by_user" field.
func CreateByUserNotNil() predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateByUser)))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.JobHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobHistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.JobHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobHistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.JobHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobHistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.JobHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobHistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// JobEntryLeaveTypeEQ applies the EQ predicate on the "job_entry_leave_type" field.
func JobEntryLeaveTypeEQ(v JobEntryLeaveType) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldJobEntryLeaveType), v))
	})
}

// JobEntryLeaveTypeNEQ applies the NEQ predicate on the "job_entry_leave_type" field.
func JobEntryLeaveTypeNEQ(v JobEntryLeaveType) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldJobEntryLeaveType), v))
	})
}

// JobEntryLeaveTypeIn applies the In predicate on the "job_entry_leave_type" field.
func JobEntryLeaveTypeIn(vs ...JobEntryLeaveType) predicate.JobHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobHistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldJobEntryLeaveType), v...))
	})
}

// JobEntryLeaveTypeNotIn applies the NotIn predicate on the "job_entry_leave_type" field.
func JobEntryLeaveTypeNotIn(vs ...JobEntryLeaveType) predicate.JobHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.JobHistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldJobEntryLeaveType), v...))
	})
}

// HasCreateBy applies the HasEdge predicate on the "create_by" edge.
func HasCreateBy() predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreateByTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CreateByTable, CreateByColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.JobHistory
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreateByWith applies the HasEdge predicate on the "create_by" edge with a given conditions (other predicates).
func HasCreateByWith(preds ...predicate.User) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreateByInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CreateByTable, CreateByColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.JobHistory
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBelongTo applies the HasEdge predicate on the "belong_to" edge.
func HasBelongTo() predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BelongToTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BelongToTable, BelongToColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.JobHistory
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBelongToWith applies the HasEdge predicate on the "belong_to" edge with a given conditions (other predicates).
func HasBelongToWith(preds ...predicate.User) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BelongToInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BelongToTable, BelongToColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.JobHistory
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobHistory) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobHistory) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
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
func Not(p predicate.JobHistory) predicate.JobHistory {
	return predicate.JobHistory(func(s *sql.Selector) {
		p(s.Not())
	})
}
