// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
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

// OrgUnitMemberQuery is the builder for querying OrgUnitMember entities.
type OrgUnitMemberQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrgUnitMember
	// eager-loading edges.
	withCreateBy        *UserQuery
	withUpdateBy        *UserQuery
	withUser            *UserQuery
	withPosition        *OrgUnitPositionQuery
	withBelongToOrgUnit *OrgUnitQuery
	modifiers           []func(s *sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrgUnitMemberQuery builder.
func (oumq *OrgUnitMemberQuery) Where(ps ...predicate.OrgUnitMember) *OrgUnitMemberQuery {
	oumq.predicates = append(oumq.predicates, ps...)
	return oumq
}

// Limit adds a limit step to the query.
func (oumq *OrgUnitMemberQuery) Limit(limit int) *OrgUnitMemberQuery {
	oumq.limit = &limit
	return oumq
}

// Offset adds an offset step to the query.
func (oumq *OrgUnitMemberQuery) Offset(offset int) *OrgUnitMemberQuery {
	oumq.offset = &offset
	return oumq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oumq *OrgUnitMemberQuery) Unique(unique bool) *OrgUnitMemberQuery {
	oumq.unique = &unique
	return oumq
}

// Order adds an order step to the query.
func (oumq *OrgUnitMemberQuery) Order(o ...OrderFunc) *OrgUnitMemberQuery {
	oumq.order = append(oumq.order, o...)
	return oumq
}

// QueryCreateBy chains the current query on the "create_by" edge.
func (oumq *OrgUnitMemberQuery) QueryCreateBy() *UserQuery {
	query := &UserQuery{config: oumq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oumq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oumq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgunitmember.Table, orgunitmember.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orgunitmember.CreateByTable, orgunitmember.CreateByColumn),
		)
		schemaConfig := oumq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.OrgUnitMember
		fromU = sqlgraph.SetNeighbors(oumq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUpdateBy chains the current query on the "update_by" edge.
func (oumq *OrgUnitMemberQuery) QueryUpdateBy() *UserQuery {
	query := &UserQuery{config: oumq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oumq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oumq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgunitmember.Table, orgunitmember.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orgunitmember.UpdateByTable, orgunitmember.UpdateByColumn),
		)
		schemaConfig := oumq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.OrgUnitMember
		fromU = sqlgraph.SetNeighbors(oumq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (oumq *OrgUnitMemberQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: oumq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oumq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oumq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgunitmember.Table, orgunitmember.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orgunitmember.UserTable, orgunitmember.UserColumn),
		)
		schemaConfig := oumq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.OrgUnitMember
		fromU = sqlgraph.SetNeighbors(oumq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPosition chains the current query on the "position" edge.
func (oumq *OrgUnitMemberQuery) QueryPosition() *OrgUnitPositionQuery {
	query := &OrgUnitPositionQuery{config: oumq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oumq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oumq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgunitmember.Table, orgunitmember.FieldID, selector),
			sqlgraph.To(orgunitposition.Table, orgunitposition.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orgunitmember.PositionTable, orgunitmember.PositionColumn),
		)
		schemaConfig := oumq.schemaConfig
		step.To.Schema = schemaConfig.OrgUnitPosition
		step.Edge.Schema = schemaConfig.OrgUnitMember
		fromU = sqlgraph.SetNeighbors(oumq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBelongToOrgUnit chains the current query on the "belongToOrgUnit" edge.
func (oumq *OrgUnitMemberQuery) QueryBelongToOrgUnit() *OrgUnitQuery {
	query := &OrgUnitQuery{config: oumq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oumq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oumq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgunitmember.Table, orgunitmember.FieldID, selector),
			sqlgraph.To(orgunit.Table, orgunit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orgunitmember.BelongToOrgUnitTable, orgunitmember.BelongToOrgUnitColumn),
		)
		schemaConfig := oumq.schemaConfig
		step.To.Schema = schemaConfig.OrgUnit
		step.Edge.Schema = schemaConfig.OrgUnitMember
		fromU = sqlgraph.SetNeighbors(oumq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrgUnitMember entity from the query.
// Returns a *NotFoundError when no OrgUnitMember was found.
func (oumq *OrgUnitMemberQuery) First(ctx context.Context) (*OrgUnitMember, error) {
	nodes, err := oumq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orgunitmember.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oumq *OrgUnitMemberQuery) FirstX(ctx context.Context) *OrgUnitMember {
	node, err := oumq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrgUnitMember ID from the query.
// Returns a *NotFoundError when no OrgUnitMember ID was found.
func (oumq *OrgUnitMemberQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oumq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orgunitmember.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oumq *OrgUnitMemberQuery) FirstIDX(ctx context.Context) int {
	id, err := oumq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrgUnitMember entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one OrgUnitMember entity is not found.
// Returns a *NotFoundError when no OrgUnitMember entities are found.
func (oumq *OrgUnitMemberQuery) Only(ctx context.Context) (*OrgUnitMember, error) {
	nodes, err := oumq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orgunitmember.Label}
	default:
		return nil, &NotSingularError{orgunitmember.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oumq *OrgUnitMemberQuery) OnlyX(ctx context.Context) *OrgUnitMember {
	node, err := oumq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrgUnitMember ID in the query.
// Returns a *NotSingularError when exactly one OrgUnitMember ID is not found.
// Returns a *NotFoundError when no entities are found.
func (oumq *OrgUnitMemberQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oumq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orgunitmember.Label}
	default:
		err = &NotSingularError{orgunitmember.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oumq *OrgUnitMemberQuery) OnlyIDX(ctx context.Context) int {
	id, err := oumq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrgUnitMembers.
func (oumq *OrgUnitMemberQuery) All(ctx context.Context) ([]*OrgUnitMember, error) {
	if err := oumq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oumq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oumq *OrgUnitMemberQuery) AllX(ctx context.Context) []*OrgUnitMember {
	nodes, err := oumq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrgUnitMember IDs.
func (oumq *OrgUnitMemberQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oumq.Select(orgunitmember.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oumq *OrgUnitMemberQuery) IDsX(ctx context.Context) []int {
	ids, err := oumq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oumq *OrgUnitMemberQuery) Count(ctx context.Context) (int, error) {
	if err := oumq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oumq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oumq *OrgUnitMemberQuery) CountX(ctx context.Context) int {
	count, err := oumq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oumq *OrgUnitMemberQuery) Exist(ctx context.Context) (bool, error) {
	if err := oumq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oumq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oumq *OrgUnitMemberQuery) ExistX(ctx context.Context) bool {
	exist, err := oumq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrgUnitMemberQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oumq *OrgUnitMemberQuery) Clone() *OrgUnitMemberQuery {
	if oumq == nil {
		return nil
	}
	return &OrgUnitMemberQuery{
		config:              oumq.config,
		limit:               oumq.limit,
		offset:              oumq.offset,
		order:               append([]OrderFunc{}, oumq.order...),
		predicates:          append([]predicate.OrgUnitMember{}, oumq.predicates...),
		withCreateBy:        oumq.withCreateBy.Clone(),
		withUpdateBy:        oumq.withUpdateBy.Clone(),
		withUser:            oumq.withUser.Clone(),
		withPosition:        oumq.withPosition.Clone(),
		withBelongToOrgUnit: oumq.withBelongToOrgUnit.Clone(),
		// clone intermediate query.
		sql:  oumq.sql.Clone(),
		path: oumq.path,
	}
}

// WithCreateBy tells the query-builder to eager-load the nodes that are connected to
// the "create_by" edge. The optional arguments are used to configure the query builder of the edge.
func (oumq *OrgUnitMemberQuery) WithCreateBy(opts ...func(*UserQuery)) *OrgUnitMemberQuery {
	query := &UserQuery{config: oumq.config}
	for _, opt := range opts {
		opt(query)
	}
	oumq.withCreateBy = query
	return oumq
}

// WithUpdateBy tells the query-builder to eager-load the nodes that are connected to
// the "update_by" edge. The optional arguments are used to configure the query builder of the edge.
func (oumq *OrgUnitMemberQuery) WithUpdateBy(opts ...func(*UserQuery)) *OrgUnitMemberQuery {
	query := &UserQuery{config: oumq.config}
	for _, opt := range opts {
		opt(query)
	}
	oumq.withUpdateBy = query
	return oumq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (oumq *OrgUnitMemberQuery) WithUser(opts ...func(*UserQuery)) *OrgUnitMemberQuery {
	query := &UserQuery{config: oumq.config}
	for _, opt := range opts {
		opt(query)
	}
	oumq.withUser = query
	return oumq
}

// WithPosition tells the query-builder to eager-load the nodes that are connected to
// the "position" edge. The optional arguments are used to configure the query builder of the edge.
func (oumq *OrgUnitMemberQuery) WithPosition(opts ...func(*OrgUnitPositionQuery)) *OrgUnitMemberQuery {
	query := &OrgUnitPositionQuery{config: oumq.config}
	for _, opt := range opts {
		opt(query)
	}
	oumq.withPosition = query
	return oumq
}

// WithBelongToOrgUnit tells the query-builder to eager-load the nodes that are connected to
// the "belongToOrgUnit" edge. The optional arguments are used to configure the query builder of the edge.
func (oumq *OrgUnitMemberQuery) WithBelongToOrgUnit(opts ...func(*OrgUnitQuery)) *OrgUnitMemberQuery {
	query := &OrgUnitQuery{config: oumq.config}
	for _, opt := range opts {
		opt(query)
	}
	oumq.withBelongToOrgUnit = query
	return oumq
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
//	client.OrgUnitMember.Query().
//		GroupBy(orgunitmember.FieldCreateByUser).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oumq *OrgUnitMemberQuery) GroupBy(field string, fields ...string) *OrgUnitMemberGroupBy {
	group := &OrgUnitMemberGroupBy{config: oumq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oumq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oumq.sqlQuery(ctx), nil
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
//	client.OrgUnitMember.Query().
//		Select(orgunitmember.FieldCreateByUser).
//		Scan(ctx, &v)
//
func (oumq *OrgUnitMemberQuery) Select(fields ...string) *OrgUnitMemberSelect {
	oumq.fields = append(oumq.fields, fields...)
	return &OrgUnitMemberSelect{OrgUnitMemberQuery: oumq}
}

func (oumq *OrgUnitMemberQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oumq.fields {
		if !orgunitmember.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oumq.path != nil {
		prev, err := oumq.path(ctx)
		if err != nil {
			return err
		}
		oumq.sql = prev
	}
	return nil
}

func (oumq *OrgUnitMemberQuery) sqlAll(ctx context.Context) ([]*OrgUnitMember, error) {
	var (
		nodes       = []*OrgUnitMember{}
		_spec       = oumq.querySpec()
		loadedTypes = [5]bool{
			oumq.withCreateBy != nil,
			oumq.withUpdateBy != nil,
			oumq.withUser != nil,
			oumq.withPosition != nil,
			oumq.withBelongToOrgUnit != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &OrgUnitMember{config: oumq.config}
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
	if len(oumq.modifiers) > 0 {
		_spec.Modifiers = oumq.modifiers
	}
	_spec.Node.Schema = oumq.schemaConfig.OrgUnitMember
	ctx = internal.NewSchemaConfigContext(ctx, oumq.schemaConfig)
	if err := sqlgraph.QueryNodes(ctx, oumq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oumq.withCreateBy; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrgUnitMember)
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

	if query := oumq.withUpdateBy; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrgUnitMember)
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

	if query := oumq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrgUnitMember)
		for i := range nodes {
			fk := nodes[i].UserID
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := oumq.withPosition; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrgUnitMember)
		for i := range nodes {
			fk := nodes[i].OrgUnitPositionID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(orgunitposition.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "org_unit_position_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Position = n
			}
		}
	}

	if query := oumq.withBelongToOrgUnit; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*OrgUnitMember)
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

func (oumq *OrgUnitMemberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oumq.querySpec()
	if len(oumq.modifiers) > 0 {
		_spec.Modifiers = oumq.modifiers
	}
	_spec.Node.Schema = oumq.schemaConfig.OrgUnitMember
	ctx = internal.NewSchemaConfigContext(ctx, oumq.schemaConfig)
	return sqlgraph.CountNodes(ctx, oumq.driver, _spec)
}

func (oumq *OrgUnitMemberQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oumq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oumq *OrgUnitMemberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orgunitmember.Table,
			Columns: orgunitmember.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orgunitmember.FieldID,
			},
		},
		From:   oumq.sql,
		Unique: true,
	}
	if unique := oumq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oumq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgunitmember.FieldID)
		for i := range fields {
			if fields[i] != orgunitmember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oumq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oumq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oumq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oumq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oumq *OrgUnitMemberQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oumq.driver.Dialect())
	t1 := builder.Table(orgunitmember.Table)
	columns := oumq.fields
	if len(columns) == 0 {
		columns = orgunitmember.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oumq.sql != nil {
		selector = oumq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, m := range oumq.modifiers {
		m(selector)
	}
	t1.Schema(oumq.schemaConfig.OrgUnitMember)
	ctx = internal.NewSchemaConfigContext(ctx, oumq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range oumq.predicates {
		p(selector)
	}
	for _, p := range oumq.order {
		p(selector)
	}
	if offset := oumq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oumq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (oumq *OrgUnitMemberQuery) ForUpdate(opts ...sql.LockOption) *OrgUnitMemberQuery {
	if oumq.driver.Dialect() == dialect.Postgres {
		oumq.Unique(false)
	}
	oumq.modifiers = append(oumq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return oumq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (oumq *OrgUnitMemberQuery) ForShare(opts ...sql.LockOption) *OrgUnitMemberQuery {
	if oumq.driver.Dialect() == dialect.Postgres {
		oumq.Unique(false)
	}
	oumq.modifiers = append(oumq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return oumq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (oumq *OrgUnitMemberQuery) Modify(modifiers ...func(s *sql.Selector)) *OrgUnitMemberSelect {
	oumq.modifiers = append(oumq.modifiers, modifiers...)
	return oumq.Select()
}

// OrgUnitMemberGroupBy is the group-by builder for OrgUnitMember entities.
type OrgUnitMemberGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oumgb *OrgUnitMemberGroupBy) Aggregate(fns ...AggregateFunc) *OrgUnitMemberGroupBy {
	oumgb.fns = append(oumgb.fns, fns...)
	return oumgb
}

// Scan applies the group-by query and scans the result into the given value.
func (oumgb *OrgUnitMemberGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := oumgb.path(ctx)
	if err != nil {
		return err
	}
	oumgb.sql = query
	return oumgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oumgb *OrgUnitMemberGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := oumgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (oumgb *OrgUnitMemberGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(oumgb.fields) > 1 {
		return nil, errors.New("ent: OrgUnitMemberGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := oumgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oumgb *OrgUnitMemberGroupBy) StringsX(ctx context.Context) []string {
	v, err := oumgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oumgb *OrgUnitMemberGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oumgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitmember.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitMemberGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oumgb *OrgUnitMemberGroupBy) StringX(ctx context.Context) string {
	v, err := oumgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (oumgb *OrgUnitMemberGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(oumgb.fields) > 1 {
		return nil, errors.New("ent: OrgUnitMemberGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := oumgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oumgb *OrgUnitMemberGroupBy) IntsX(ctx context.Context) []int {
	v, err := oumgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oumgb *OrgUnitMemberGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oumgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitmember.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitMemberGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oumgb *OrgUnitMemberGroupBy) IntX(ctx context.Context) int {
	v, err := oumgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (oumgb *OrgUnitMemberGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(oumgb.fields) > 1 {
		return nil, errors.New("ent: OrgUnitMemberGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := oumgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oumgb *OrgUnitMemberGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := oumgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oumgb *OrgUnitMemberGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oumgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitmember.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitMemberGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oumgb *OrgUnitMemberGroupBy) Float64X(ctx context.Context) float64 {
	v, err := oumgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (oumgb *OrgUnitMemberGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(oumgb.fields) > 1 {
		return nil, errors.New("ent: OrgUnitMemberGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := oumgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oumgb *OrgUnitMemberGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := oumgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (oumgb *OrgUnitMemberGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oumgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitmember.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitMemberGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oumgb *OrgUnitMemberGroupBy) BoolX(ctx context.Context) bool {
	v, err := oumgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oumgb *OrgUnitMemberGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range oumgb.fields {
		if !orgunitmember.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oumgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oumgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oumgb *OrgUnitMemberGroupBy) sqlQuery() *sql.Selector {
	selector := oumgb.sql.Select()
	aggregation := make([]string, 0, len(oumgb.fns))
	for _, fn := range oumgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oumgb.fields)+len(oumgb.fns))
		for _, f := range oumgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oumgb.fields...)...)
}

// OrgUnitMemberSelect is the builder for selecting fields of OrgUnitMember entities.
type OrgUnitMemberSelect struct {
	*OrgUnitMemberQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oums *OrgUnitMemberSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oums.prepareQuery(ctx); err != nil {
		return err
	}
	oums.sql = oums.OrgUnitMemberQuery.sqlQuery(ctx)
	return oums.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (oums *OrgUnitMemberSelect) ScanX(ctx context.Context, v interface{}) {
	if err := oums.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (oums *OrgUnitMemberSelect) Strings(ctx context.Context) ([]string, error) {
	if len(oums.fields) > 1 {
		return nil, errors.New("ent: OrgUnitMemberSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := oums.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (oums *OrgUnitMemberSelect) StringsX(ctx context.Context) []string {
	v, err := oums.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (oums *OrgUnitMemberSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = oums.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitmember.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitMemberSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (oums *OrgUnitMemberSelect) StringX(ctx context.Context) string {
	v, err := oums.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (oums *OrgUnitMemberSelect) Ints(ctx context.Context) ([]int, error) {
	if len(oums.fields) > 1 {
		return nil, errors.New("ent: OrgUnitMemberSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := oums.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (oums *OrgUnitMemberSelect) IntsX(ctx context.Context) []int {
	v, err := oums.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (oums *OrgUnitMemberSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = oums.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitmember.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitMemberSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (oums *OrgUnitMemberSelect) IntX(ctx context.Context) int {
	v, err := oums.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (oums *OrgUnitMemberSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(oums.fields) > 1 {
		return nil, errors.New("ent: OrgUnitMemberSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := oums.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (oums *OrgUnitMemberSelect) Float64sX(ctx context.Context) []float64 {
	v, err := oums.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (oums *OrgUnitMemberSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = oums.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitmember.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitMemberSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (oums *OrgUnitMemberSelect) Float64X(ctx context.Context) float64 {
	v, err := oums.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (oums *OrgUnitMemberSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(oums.fields) > 1 {
		return nil, errors.New("ent: OrgUnitMemberSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := oums.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (oums *OrgUnitMemberSelect) BoolsX(ctx context.Context) []bool {
	v, err := oums.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (oums *OrgUnitMemberSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = oums.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orgunitmember.Label}
	default:
		err = fmt.Errorf("ent: OrgUnitMemberSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (oums *OrgUnitMemberSelect) BoolX(ctx context.Context) bool {
	v, err := oums.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oums *OrgUnitMemberSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oums.sql.Query()
	if err := oums.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (oums *OrgUnitMemberSelect) Modify(modifiers ...func(s *sql.Selector)) *OrgUnitMemberSelect {
	oums.modifiers = append(oums.modifiers, modifiers...)
	return oums
}
