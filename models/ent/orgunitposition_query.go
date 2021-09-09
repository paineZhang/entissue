// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"wing/models/ent/internal"
	"wing/models/ent/orgunit"
	"wing/models/ent/orgunitmember"
	"wing/models/ent/orgunitposition"
	"wing/models/ent/predicate"
	"wing/models/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrgUnitPositionQuery is the builder for querying OrgUnitPosition entities.
type OrgUnitPositionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrgUnitPosition
	// eager-loading edges.
	withCreateBy               *UserQuery
	withUpdateBy               *UserQuery
	withBelongToOrgUnitMembers *OrgUnitMemberQuery
	withBelongToOrgUnit        *OrgUnitQuery
	modifiers                  []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrgUnitPositionQuery builder.
func (oupq *OrgUnitPositionQuery) Where(ps ...predicate.OrgUnitPosition) *OrgUnitPositionQuery {
	oupq.predicates = append(oupq.predicates, ps...)
	return oupq
}

// Limit adds a limit step to the query.
func (oupq *OrgUnitPositionQuery) Limit(limit int) *OrgUnitPositionQuery {
	oupq.limit = &limit
	return oupq
}

// Offset adds an offset step to the query.
func (oupq *OrgUnitPositionQuery) Offset(offset int) *OrgUnitPositionQuery {
	oupq.offset = &offset
	return oupq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oupq *OrgUnitPositionQuery) Unique(unique bool) *OrgUnitPositionQuery {
	oupq.unique = &unique
	return oupq
}

// Order adds an order step to the query.
func (oupq *OrgUnitPositionQuery) Order(o ...OrderFunc) *OrgUnitPositionQuery {
	oupq.order = append(oupq.order, o...)
	return oupq
}

// QueryCreateBy chains the current query on the "create_by" edge.
func (oupq *OrgUnitPositionQuery) QueryCreateBy() *UserQuery {
	query := &UserQuery{config: oupq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oupq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oupq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgunitposition.Table, orgunitposition.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orgunitposition.CreateByTable, orgunitposition.CreateByColumn),
		)
		schemaConfig := oupq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.OrgUnitPosition
		fromU = sqlgraph.SetNeighbors(oupq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUpdateBy chains the current query on the "update_by" edge.
func (oupq *OrgUnitPositionQuery) QueryUpdateBy() *UserQuery {
	query := &UserQuery{config: oupq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oupq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oupq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgunitposition.Table, orgunitposition.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orgunitposition.UpdateByTable, orgunitposition.UpdateByColumn),
		)
		schemaConfig := oupq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.OrgUnitPosition
		fromU = sqlgraph.SetNeighbors(oupq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBelongToOrgUnitMembers chains the current query on the "belongToOrgUnitMembers" edge.
func (oupq *OrgUnitPositionQuery) QueryBelongToOrgUnitMembers() *OrgUnitMemberQuery {
	query := &OrgUnitMemberQuery{config: oupq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oupq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oupq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgunitposition.Table, orgunitposition.FieldID, selector),
			sqlgraph.To(orgunitmember.Table, orgunitmember.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, orgunitposition.BelongToOrgUnitMembersTable, orgunitposition.BelongToOrgUnitMembersColumn),
		)
		schemaConfig := oupq.schemaConfig
		step.To.Schema = schemaConfig.OrgUnitMember
		step.Edge.Schema = schemaConfig.OrgUnitMember
		fromU = sqlgraph.SetNeighbors(oupq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBelongToOrgUnit chains the current query on the "belongToOrgUnit" edge.
func (oupq *OrgUnitPositionQuery) QueryBelongToOrgUnit() *OrgUnitQuery {
	query := &OrgUnitQuery{config: oupq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oupq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oupq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgunitposition.Table, orgunitposition.FieldID, selector),
			sqlgraph.To(orgunit.Table, orgunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orgunitposition.BelongToOrgUnitTable, orgunitposition.BelongToOrgUnitColumn),
		)
		schemaConfig := oupq.schemaConfig
		step.To.Schema = schemaConfig.OrgUnit
		step.Edge.Schema = schemaConfig.OrgUnitPosition
		fromU = sqlgraph.SetNeighbors(oupq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrgUnitPosition entity from the query.
// Returns a *NotFoundError when no OrgUnitPosition was found.
func (oupq *OrgUnitPositionQuery) First(ctx context.Context) (*OrgUnitPosition, error) {
	nodes, err := oupq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orgunitposition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oupq *OrgUnitPositionQuery) FirstX(ctx context.Context) *OrgUnitPosition {
	node, err := oupq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrgUnitPosition ID from the query.
// Returns a *NotFoundError when no OrgUnitPosition ID was found.
func (oupq *OrgUnitPositionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oupq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orgunitposition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oupq *OrgUnitPositionQuery) FirstIDX(ctx context.Context) int {
	id, err := oupq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrgUnitPosition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OrgUnitPosition entity is not found.
// Returns a *NotFoundError when no OrgUnitPosition entities are found.
func (oupq *OrgUnitPositionQuery) Only(ctx context.Context) (*OrgUnitPosition, error) {
	nodes, err := oupq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orgunitposition.Label}
	default:
		return nil, &NotSingularError{orgunitposition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oupq *OrgUnitPositionQuery) OnlyX(ctx context.Context) *OrgUnitPosition {
	node, err := oupq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrgUnitPosition ID in the query.
// Returns a *NotSingularError when exactly one OrgUnitPosition ID is not found.
// Returns a *NotFoundError when no entities are found.
func (oupq *OrgUnitPositionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oupq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orgunitposition.Label}
	default:
		err = &NotSingularError{orgunitposition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oupq *OrgUnitPositionQuery) OnlyIDX(ctx context.Context) int {
	id, err := oupq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrgUnitPositions.
func (oupq *OrgUnitPositionQuery) All(ctx context.Context) ([]*OrgUnitPosition, error) {
	if err := oupq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oupq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oupq *OrgUnitPositionQuery) AllX(ctx context.Context) []*OrgUnitPosition {
	nodes, err := oupq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrgUnitPosition IDs.
func (oupq *OrgUnitPositionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oupq.Select(orgunitposition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oupq *OrgUnitPositionQuery) IDsX(ctx context.Context) []int {
	ids, err := oupq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oupq *OrgUnitPositionQuery) Count(ctx context.Context) (int, error) {
	if err := oupq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oupq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oupq *OrgUnitPositionQuery) CountX(ctx context.Context) int {
	count, err := oupq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oupq *OrgUnitPositionQuery) Exist(ctx context.Context) (bool, error) {
	if err := oupq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oupq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oupq *OrgUnitPositionQuery) ExistX(ctx context.Context) bool {
	exist, err := oupq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrgUnitPositionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oupq *OrgUnitPositionQuery) Clone() *OrgUnitPositionQuery {
	if oupq == nil {
		return nil
	}
	return &OrgUnitPositionQuery{
		config:                     oupq.config,
		limit:                      oupq.limit,
		offset:                     oupq.offset,
		order:                      append([]OrderFunc{}, oupq.order...),
		predicates:                 append([]predicate.OrgUnitPosition{}, oupq.predicates...),
		withCreateBy:               oupq.withCreateBy.Clone(),
		withUpdateBy:               oupq.withUpdateBy.Clone(),
		withBelongToOrgUnitMembers: oupq.withBelongToOrgUnitMembers.Clone(),
		withBelongToOrgUnit:        oupq.withBelongToOrgUnit.Clone(),
		// clone intermediate query.
		sql:  oupq.sql.Clone(),
		path: oupq.path,
	}
}

// WithCreateBy tells the query-builder to eager-load the nodes that are connected to
// the "create_by" edge. The optional arguments are used to configure the query builder of the edge.
func (oupq *OrgUnitPositionQuery) WithCreateBy(opts ...func(*UserQuery)) *OrgUnitPositionQuery {
	query := &UserQuery{config: oupq.config}
	for _, opt := range opts {
		opt(query)
	}
	oupq.withCreateBy = query
	return oupq
}

// WithUpdateBy tells the query-builder to eager-load the nodes that are connected to
// the "update_by" edge. The optional arguments are used to configure the query builder of the edge.
func (oupq *OrgUnitPositionQuery) WithUpdateBy(opts ...func(*UserQuery)) *OrgUnitPositionQuery {
	query := &UserQuery{config: oupq.config}
	for _, opt := range opts {
		opt(query)
	}
	oupq.withUpdateBy = query
	return oupq
}

// WithBelongToOrgUnitMembers tells the query-builder to eager-load the nodes that are connected to
// the "belongToOrgUnitMembers" edge. The optional arguments are used to configure the query builder of the edge.
func (oupq *OrgUnitPositionQuery) WithBelongToOrgUnitMembers(opts ...func(*OrgUnitMemberQuery)) *OrgUnitPositionQuery {
	query := &OrgUnitMemberQuery{config: oupq.config}
	for _, opt := range opts {
		opt(query)
	}
	oupq.withBelongToOrgUnitMembers = query
	return oupq
}

// WithBelongToOrgUnit tells the query-builder to eager-load the nodes that are connected to
// the "belongToOrgUnit" edge. The optional arguments are used to configure the query builder of the edge.
func (oupq *OrgUnitPositionQuery) WithBelongToOrgUnit(opts ...func(*OrgUnitQuery)) *OrgUnitPositionQuery {
	query := &OrgUnitQuery{config: oupq.config}
	for _, opt := range opts {
		opt(query)
	}
	oupq.withBelongToOrgUnit = query
	return oupq
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
//	client.OrgUnitPosition.Query().
//		GroupBy(orgunitposition.FieldCreateByUser).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oupq *OrgUnitPositionQuery) GroupBy(field string, fields ...string) *OrgUnitPositionGroupBy {
	group := &OrgUnitPositionGroupBy{config: oupq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oupq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oupq.sqlQuery(ctx), nil
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
//	client.OrgUnitPosition.Query().
//		Select(orgunitposition.FieldCreateByUser).
//		Scan(ctx, &v)
//
func (oupq *OrgUnitPositionQuery) Select(fields ...string) *OrgUnitPositionSelect {
	oupq.fields = append(oupq.fields, fields...)
	return &OrgUnitPositionSelect{OrgUnitPositionQuery: oupq}
}

func (oupq *OrgUnitPositionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oupq.fields {
		if !orgunitposition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oupq.path != nil {
		prev, err := oupq.path(ctx)
		if err != nil {
			return err
		}
		oupq.sql = prev
	}
	return nil
}

func (oupq *OrgUnitPositionQuery) sqlAll(ctx context.Context) ([]*OrgUnitPosition, error) {
	var (
		nodes       = []*OrgUnitPosition{}
		_spec       = oupq.querySpec()
		loadedTypes = [4]bool{
			oupq.withCreateBy != nil,
			oupq.withUpdateBy != nil,
			oupq.withBelongToOrgUnitMembers != nil,
			oupq.withBelongToOrgUnit != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrgUnitPosition{config: oupq.config}
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
	if len(oupq.modifiers) > 0 {
		_spec.Modifiers = oupq.modifiers
	}
	_spec.Node.Schema = oupq.schemaConfig.OrgUnitPosition
	ctx = internal.NewSchemaConfigContext(ctx, oupq.schemaConfig)
	if err := sqlgraph.QueryNodes(ctx, oupq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oupq.withCreateBy; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrgUnitPosition)
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

	if query := oupq.withUpdateBy; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrgUnitPosition)
		for i := range nodes {
			fk := nodes[i].UpdateByUser
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
				return nil, fmt.Errorf(`unexpected foreign-key "update_by_user" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.UpdateBy = n
			}
		}
	}

	if query := oupq.withBelongToOrgUnitMembers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*OrgUnitPosition)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.BelongToOrgUnitMembers = []*OrgUnitMember{}
		}
		query.Where(predicate.OrgUnitMember(func(s *sql.Selector) {
			s.Where(sql.InValues(orgunitposition.BelongToOrgUnitMembersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.OrgUnitPositionID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "org_unit_position_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.BelongToOrgUnitMembers = append(node.Edges.BelongToOrgUnitMembers, n)
		}
	}

	if query := oupq.withBelongToOrgUnit; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrgUnitPosition)
		for i := range nodes {
			fk := nodes[i].OrgUnitID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(orgunit.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "org_unit_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.BelongToOrgUnit = n
			}
		}
	}

	return nodes, nil
}

func (oupq *OrgUnitPositionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oupq.querySpec()
	if len(oupq.modifiers) > 0 {
		_spec.Modifiers = oupq.modifiers
	}
	_spec.Node.Schema = oupq.schemaConfig.OrgUnitPosition
	ctx = internal.NewSchemaConfigContext(ctx, oupq.schemaConfig)
	return sqlgraph.CountNodes(ctx, oupq.driver, _spec)
}

func (oupq *OrgUnitPositionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oupq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oupq *OrgUnitPositionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orgunitposition.Table,
			Columns: orgunitposition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orgunitposition.FieldID,
			},
		},
		From:   oupq.sql,
		Unique: true,
	}
	if unique := oupq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oupq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgunitposition.FieldID)
		for i := range fields {
			if fields[i] != orgunitposition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oupq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oupq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oupq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oupq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oupq *OrgUnitPositionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oupq.driver.Dialect())
	t1 := builder.Table(orgunitposition.Table)
	columns := oupq.fields
	if len(columns) == 0 {
		columns = orgunitposition.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oupq.sql != nil {
		selector = oupq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, m := range oupq.modifiers {
		m(selector)
	}
	t1.Schema(oupq.schemaConfig.OrgUnitPosition)
	ctx = internal.NewSchemaConfigContext(ctx, oupq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range oupq.predicates {
		p(selector)
	}
	for _, p := range oupq.order {
		p(selector)
	}
	if offset := oupq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oupq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (oupq *OrgUnitPositionQuery) ForUpdate(opts ...sql.LockOption) *OrgUnitPositionQuery {
	if oupq.driver.Dialect() == dialect.Postgres {
		oupq.Unique(false)
	}
	oupq.modifiers = append(oupq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return oupq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (oupq *OrgUnitPositionQuery) ForShare(opts ...sql.LockOption) *OrgUnitPositionQuery {
	if oupq.driver.Dialect() == dialect.Postgres {
		oupq.Unique(false)
	}
	oupq.modifiers = append(oupq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return oupq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (oupq *OrgUnitPositionQuery) Modify(modifiers ...func(s *sql.Selector)) *OrgUnitPositionSelect {
	oupq.modifiers = append(oupq.modifiers, modifiers...)
	return oupq.Select()
}

// OrgUnitPositionGroupBy is the group-by builder for OrgUnitPosition entities.
type OrgUnitPositionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oupgb *OrgUnitPositionGroupBy) Aggregate(fns ...AggregateFunc) *OrgUnitPositionGroupBy {
	oupgb.fns = append(oupgb.fns, fns...)
	return oupgb
}

// Scan applies the group-by query and scans the result into the given value.
func (oupgb *OrgUnitPositionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oupgb.path(ctx)
	if err != nil {
		return err
	}
	oupgb.sql = query
	return oupgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oupgb *OrgUnitPositionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := oupgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (oupgb *OrgUnitPositionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(oupgb.fields) > 1 {
		return nil, errors.New("ent: OrgUnitPositionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := oupgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oupgb *OrgUnitPositionGroupBy) StringsX(ctx context.Context) []string {
	v, err := oupgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oupgb *OrgUnitPositionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oupgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitposition.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitPositionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oupgb *OrgUnitPositionGroupBy) StringX(ctx context.Context) string {
	v, err := oupgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (oupgb *OrgUnitPositionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(oupgb.fields) > 1 {
		return nil, errors.New("ent: OrgUnitPositionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := oupgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oupgb *OrgUnitPositionGroupBy) IntsX(ctx context.Context) []int {
	v, err := oupgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oupgb *OrgUnitPositionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oupgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitposition.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitPositionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oupgb *OrgUnitPositionGroupBy) IntX(ctx context.Context) int {
	v, err := oupgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (oupgb *OrgUnitPositionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(oupgb.fields) > 1 {
		return nil, errors.New("ent: OrgUnitPositionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := oupgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oupgb *OrgUnitPositionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := oupgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oupgb *OrgUnitPositionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oupgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitposition.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitPositionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oupgb *OrgUnitPositionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := oupgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (oupgb *OrgUnitPositionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(oupgb.fields) > 1 {
		return nil, errors.New("ent: OrgUnitPositionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := oupgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oupgb *OrgUnitPositionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := oupgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oupgb *OrgUnitPositionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oupgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitposition.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitPositionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oupgb *OrgUnitPositionGroupBy) BoolX(ctx context.Context) bool {
	v, err := oupgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oupgb *OrgUnitPositionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oupgb.fields {
		if !orgunitposition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oupgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oupgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oupgb *OrgUnitPositionGroupBy) sqlQuery() *sql.Selector {
	selector := oupgb.sql.Select()
	aggregation := make([]string, 0, len(oupgb.fns))
	for _, fn := range oupgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oupgb.fields)+len(oupgb.fns))
		for _, f := range oupgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oupgb.fields...)...)
}

// OrgUnitPositionSelect is the builder for selecting fields of OrgUnitPosition entities.
type OrgUnitPositionSelect struct {
	*OrgUnitPositionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oups *OrgUnitPositionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oups.prepareQuery(ctx); err != nil {
		return err
	}
	oups.sql = oups.OrgUnitPositionQuery.sqlQuery(ctx)
	return oups.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oups *OrgUnitPositionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := oups.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (oups *OrgUnitPositionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(oups.fields) > 1 {
		return nil, errors.New("ent: OrgUnitPositionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := oups.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oups *OrgUnitPositionSelect) StringsX(ctx context.Context) []string {
	v, err := oups.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (oups *OrgUnitPositionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oups.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitposition.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitPositionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oups *OrgUnitPositionSelect) StringX(ctx context.Context) string {
	v, err := oups.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (oups *OrgUnitPositionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(oups.fields) > 1 {
		return nil, errors.New("ent: OrgUnitPositionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := oups.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oups *OrgUnitPositionSelect) IntsX(ctx context.Context) []int {
	v, err := oups.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (oups *OrgUnitPositionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oups.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitposition.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitPositionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oups *OrgUnitPositionSelect) IntX(ctx context.Context) int {
	v, err := oups.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (oups *OrgUnitPositionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(oups.fields) > 1 {
		return nil, errors.New("ent: OrgUnitPositionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := oups.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oups *OrgUnitPositionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := oups.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (oups *OrgUnitPositionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oups.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitposition.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitPositionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oups *OrgUnitPositionSelect) Float64X(ctx context.Context) float64 {
	v, err := oups.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (oups *OrgUnitPositionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(oups.fields) > 1 {
		return nil, errors.New("ent: OrgUnitPositionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := oups.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oups *OrgUnitPositionSelect) BoolsX(ctx context.Context) []bool {
	v, err := oups.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (oups *OrgUnitPositionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oups.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitposition.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitPositionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oups *OrgUnitPositionSelect) BoolX(ctx context.Context) bool {
	v, err := oups.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oups *OrgUnitPositionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oups.sql.Query()
	if err := oups.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (oups *OrgUnitPositionSelect) Modify(modifiers ...func(s *sql.Selector)) *OrgUnitPositionSelect {
	oups.modifiers = append(oups.modifiers, modifiers...)
	return oups
}