// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"wing/models/ent/internal"
	"wing/models/ent/jobhistory"
	"wing/models/ent/predicate"
	"wing/models/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobHistoryQuery is the builder for querying JobHistory entities.
type JobHistoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.JobHistory
	// eager-loading edges.
	withCreateBy *UserQuery
	withBelongTo *UserQuery
	withFKs      bool
	modifiers    []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobHistoryQuery builder.
func (jhq *JobHistoryQuery) Where(ps ...predicate.JobHistory) *JobHistoryQuery {
	jhq.predicates = append(jhq.predicates, ps...)
	return jhq
}

// Limit adds a limit step to the query.
func (jhq *JobHistoryQuery) Limit(limit int) *JobHistoryQuery {
	jhq.limit = &limit
	return jhq
}

// Offset adds an offset step to the query.
func (jhq *JobHistoryQuery) Offset(offset int) *JobHistoryQuery {
	jhq.offset = &offset
	return jhq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jhq *JobHistoryQuery) Unique(unique bool) *JobHistoryQuery {
	jhq.unique = &unique
	return jhq
}

// Order adds an order step to the query.
func (jhq *JobHistoryQuery) Order(o ...OrderFunc) *JobHistoryQuery {
	jhq.order = append(jhq.order, o...)
	return jhq
}

// QueryCreateBy chains the current query on the "create_by" edge.
func (jhq *JobHistoryQuery) QueryCreateBy() *UserQuery {
	query := &UserQuery{config: jhq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jhq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jhq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(jobhistory.Table, jobhistory.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, jobhistory.CreateByTable, jobhistory.CreateByColumn),
		)
		schemaConfig := jhq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.JobHistory
		fromU = sqlgraph.SetNeighbors(jhq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBelongTo chains the current query on the "belong_to" edge.
func (jhq *JobHistoryQuery) QueryBelongTo() *UserQuery {
	query := &UserQuery{config: jhq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jhq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jhq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(jobhistory.Table, jobhistory.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, jobhistory.BelongToTable, jobhistory.BelongToColumn),
		)
		schemaConfig := jhq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.JobHistory
		fromU = sqlgraph.SetNeighbors(jhq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first JobHistory entity from the query.
// Returns a *NotFoundError when no JobHistory was found.
func (jhq *JobHistoryQuery) First(ctx context.Context) (*JobHistory, error) {
	nodes, err := jhq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{jobhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jhq *JobHistoryQuery) FirstX(ctx context.Context) *JobHistory {
	node, err := jhq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JobHistory ID from the query.
// Returns a *NotFoundError when no JobHistory ID was found.
func (jhq *JobHistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jhq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{jobhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jhq *JobHistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := jhq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single JobHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one JobHistory entity is not found.
// Returns a *NotFoundError when no JobHistory entities are found.
func (jhq *JobHistoryQuery) Only(ctx context.Context) (*JobHistory, error) {
	nodes, err := jhq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{jobhistory.Label}
	default:
		return nil, &NotSingularError{jobhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jhq *JobHistoryQuery) OnlyX(ctx context.Context) *JobHistory {
	node, err := jhq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only JobHistory ID in the query.
// Returns a *NotSingularError when exactly one JobHistory ID is not found.
// Returns a *NotFoundError when no entities are found.
func (jhq *JobHistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jhq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{jobhistory.Label}
	default:
		err = &NotSingularError{jobhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jhq *JobHistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := jhq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JobHistories.
func (jhq *JobHistoryQuery) All(ctx context.Context) ([]*JobHistory, error) {
	if err := jhq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return jhq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (jhq *JobHistoryQuery) AllX(ctx context.Context) []*JobHistory {
	nodes, err := jhq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JobHistory IDs.
func (jhq *JobHistoryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := jhq.Select(jobhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jhq *JobHistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := jhq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jhq *JobHistoryQuery) Count(ctx context.Context) (int, error) {
	if err := jhq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return jhq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (jhq *JobHistoryQuery) CountX(ctx context.Context) int {
	count, err := jhq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jhq *JobHistoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := jhq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return jhq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (jhq *JobHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := jhq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jhq *JobHistoryQuery) Clone() *JobHistoryQuery {
	if jhq == nil {
		return nil
	}
	return &JobHistoryQuery{
		config:       jhq.config,
		limit:        jhq.limit,
		offset:       jhq.offset,
		order:        append([]OrderFunc{}, jhq.order...),
		predicates:   append([]predicate.JobHistory{}, jhq.predicates...),
		withCreateBy: jhq.withCreateBy.Clone(),
		withBelongTo: jhq.withBelongTo.Clone(),
		// clone intermediate query.
		sql:  jhq.sql.Clone(),
		path: jhq.path,
	}
}

// WithCreateBy tells the query-builder to eager-load the nodes that are connected to
// the "create_by" edge. The optional arguments are used to configure the query builder of the edge.
func (jhq *JobHistoryQuery) WithCreateBy(opts ...func(*UserQuery)) *JobHistoryQuery {
	query := &UserQuery{config: jhq.config}
	for _, opt := range opts {
		opt(query)
	}
	jhq.withCreateBy = query
	return jhq
}

// WithBelongTo tells the query-builder to eager-load the nodes that are connected to
// the "belong_to" edge. The optional arguments are used to configure the query builder of the edge.
func (jhq *JobHistoryQuery) WithBelongTo(opts ...func(*UserQuery)) *JobHistoryQuery {
	query := &UserQuery{config: jhq.config}
	for _, opt := range opts {
		opt(query)
	}
	jhq.withBelongTo = query
	return jhq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateByUser int `json:"create_by_user,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.JobHistory.Query().
//		GroupBy(jobhistory.FieldCreateByUser).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (jhq *JobHistoryQuery) GroupBy(field string, fields ...string) *JobHistoryGroupBy {
	group := &JobHistoryGroupBy{config: jhq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := jhq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return jhq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateByUser int `json:"create_by_user,omitempty"`
//	}
//
//	client.JobHistory.Query().
//		Select(jobhistory.FieldCreateByUser).
//		Scan(ctx, &v)
//
func (jhq *JobHistoryQuery) Select(fields ...string) *JobHistorySelect {
	jhq.fields = append(jhq.fields, fields...)
	return &JobHistorySelect{JobHistoryQuery: jhq}
}

func (jhq *JobHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range jhq.fields {
		if !jobhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jhq.path != nil {
		prev, err := jhq.path(ctx)
		if err != nil {
			return err
		}
		jhq.sql = prev
	}
	return nil
}

func (jhq *JobHistoryQuery) sqlAll(ctx context.Context) ([]*JobHistory, error) {
	var (
		nodes       = []*JobHistory{}
		withFKs     = jhq.withFKs
		_spec       = jhq.querySpec()
		loadedTypes = [2]bool{
			jhq.withCreateBy != nil,
			jhq.withBelongTo != nil,
		}
	)
	if jhq.withCreateBy != nil || jhq.withBelongTo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, jobhistory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &JobHistory{config: jhq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(jhq.modifiers) > 0 {
		_spec.Modifiers = jhq.modifiers
	}
	_spec.Node.Schema = jhq.schemaConfig.JobHistory
	ctx = internal.NewSchemaConfigContext(ctx, jhq.schemaConfig)
	if err := sqlgraph.QueryNodes(ctx, jhq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := jhq.withCreateBy; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*JobHistory)
		for i := range nodes {
			fk := nodes[i].CreateByUser
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "create_by_user" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.CreateBy = n
			}
		}
	}

	if query := jhq.withBelongTo; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*JobHistory)
		for i := range nodes {
			if nodes[i].user_job_histories == nil {
				continue
			}
			fk := *nodes[i].user_job_histories
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_job_histories" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.BelongTo = n
			}
		}
	}

	return nodes, nil
}

func (jhq *JobHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jhq.querySpec()
	if len(jhq.modifiers) > 0 {
		_spec.Modifiers = jhq.modifiers
	}
	_spec.Node.Schema = jhq.schemaConfig.JobHistory
	ctx = internal.NewSchemaConfigContext(ctx, jhq.schemaConfig)
	return sqlgraph.CountNodes(ctx, jhq.driver, _spec)
}

func (jhq *JobHistoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := jhq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (jhq *JobHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   jobhistory.Table,
			Columns: jobhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jobhistory.FieldID,
			},
		},
		From:   jhq.sql,
		Unique: true,
	}
	if unique := jhq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := jhq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobhistory.FieldID)
		for i := range fields {
			if fields[i] != jobhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := jhq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jhq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jhq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jhq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jhq *JobHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jhq.driver.Dialect())
	t1 := builder.Table(jobhistory.Table)
	columns := jhq.fields
	if len(columns) == 0 {
		columns = jobhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jhq.sql != nil {
		selector = jhq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, m := range jhq.modifiers {
		m(selector)
	}
	t1.Schema(jhq.schemaConfig.JobHistory)
	ctx = internal.NewSchemaConfigContext(ctx, jhq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range jhq.predicates {
		p(selector)
	}
	for _, p := range jhq.order {
		p(selector)
	}
	if offset := jhq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jhq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (jhq *JobHistoryQuery) ForUpdate(opts ...sql.LockOption) *JobHistoryQuery {
	if jhq.driver.Dialect() == dialect.Postgres {
		jhq.Unique(false)
	}
	jhq.modifiers = append(jhq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return jhq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (jhq *JobHistoryQuery) ForShare(opts ...sql.LockOption) *JobHistoryQuery {
	if jhq.driver.Dialect() == dialect.Postgres {
		jhq.Unique(false)
	}
	jhq.modifiers = append(jhq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return jhq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (jhq *JobHistoryQuery) Modify(modifiers ...func(s *sql.Selector)) *JobHistorySelect {
	jhq.modifiers = append(jhq.modifiers, modifiers...)
	return jhq.Select()
}

// JobHistoryGroupBy is the group-by builder for JobHistory entities.
type JobHistoryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jhgb *JobHistoryGroupBy) Aggregate(fns ...AggregateFunc) *JobHistoryGroupBy {
	jhgb.fns = append(jhgb.fns, fns...)
	return jhgb
}

// Scan applies the group-by query and scans the result into the given value.
func (jhgb *JobHistoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := jhgb.path(ctx)
	if err != nil {
		return err
	}
	jhgb.sql = query
	return jhgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jhgb *JobHistoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := jhgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (jhgb *JobHistoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(jhgb.fields) > 1 {
		return nil, errors.New("ent: JobHistoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := jhgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jhgb *JobHistoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := jhgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jhgb *JobHistoryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jhgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobhistory.Label}
	default:
		err = fmt.Errorf("ent: JobHistoryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jhgb *JobHistoryGroupBy) StringX(ctx context.Context) string {
	v, err := jhgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (jhgb *JobHistoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(jhgb.fields) > 1 {
		return nil, errors.New("ent: JobHistoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := jhgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jhgb *JobHistoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := jhgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jhgb *JobHistoryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jhgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobhistory.Label}
	default:
		err = fmt.Errorf("ent: JobHistoryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jhgb *JobHistoryGroupBy) IntX(ctx context.Context) int {
	v, err := jhgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (jhgb *JobHistoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(jhgb.fields) > 1 {
		return nil, errors.New("ent: JobHistoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := jhgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jhgb *JobHistoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := jhgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jhgb *JobHistoryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jhgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobhistory.Label}
	default:
		err = fmt.Errorf("ent: JobHistoryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jhgb *JobHistoryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := jhgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (jhgb *JobHistoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(jhgb.fields) > 1 {
		return nil, errors.New("ent: JobHistoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := jhgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jhgb *JobHistoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := jhgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jhgb *JobHistoryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jhgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobhistory.Label}
	default:
		err = fmt.Errorf("ent: JobHistoryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jhgb *JobHistoryGroupBy) BoolX(ctx context.Context) bool {
	v, err := jhgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jhgb *JobHistoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range jhgb.fields {
		if !jobhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := jhgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jhgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (jhgb *JobHistoryGroupBy) sqlQuery() *sql.Selector {
	selector := jhgb.sql.Select()
	aggregation := make([]string, 0, len(jhgb.fns))
	for _, fn := range jhgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(jhgb.fields)+len(jhgb.fns))
		for _, f := range jhgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(jhgb.fields...)...)
}

// JobHistorySelect is the builder for selecting fields of JobHistory entities.
type JobHistorySelect struct {
	*JobHistoryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (jhs *JobHistorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := jhs.prepareQuery(ctx); err != nil {
		return err
	}
	jhs.sql = jhs.JobHistoryQuery.sqlQuery(ctx)
	return jhs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jhs *JobHistorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := jhs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (jhs *JobHistorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(jhs.fields) > 1 {
		return nil, errors.New("ent: JobHistorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := jhs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jhs *JobHistorySelect) StringsX(ctx context.Context) []string {
	v, err := jhs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (jhs *JobHistorySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jhs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobhistory.Label}
	default:
		err = fmt.Errorf("ent: JobHistorySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jhs *JobHistorySelect) StringX(ctx context.Context) string {
	v, err := jhs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (jhs *JobHistorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(jhs.fields) > 1 {
		return nil, errors.New("ent: JobHistorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := jhs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jhs *JobHistorySelect) IntsX(ctx context.Context) []int {
	v, err := jhs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (jhs *JobHistorySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jhs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobhistory.Label}
	default:
		err = fmt.Errorf("ent: JobHistorySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jhs *JobHistorySelect) IntX(ctx context.Context) int {
	v, err := jhs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (jhs *JobHistorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(jhs.fields) > 1 {
		return nil, errors.New("ent: JobHistorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := jhs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jhs *JobHistorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := jhs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (jhs *JobHistorySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jhs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobhistory.Label}
	default:
		err = fmt.Errorf("ent: JobHistorySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jhs *JobHistorySelect) Float64X(ctx context.Context) float64 {
	v, err := jhs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (jhs *JobHistorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(jhs.fields) > 1 {
		return nil, errors.New("ent: JobHistorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := jhs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jhs *JobHistorySelect) BoolsX(ctx context.Context) []bool {
	v, err := jhs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (jhs *JobHistorySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jhs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{jobhistory.Label}
	default:
		err = fmt.Errorf("ent: JobHistorySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jhs *JobHistorySelect) BoolX(ctx context.Context) bool {
	v, err := jhs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jhs *JobHistorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := jhs.sql.Query()
	if err := jhs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (jhs *JobHistorySelect) Modify(modifiers ...func(s *sql.Selector)) *JobHistorySelect {
	jhs.modifiers = append(jhs.modifiers, modifiers...)
	return jhs
}
