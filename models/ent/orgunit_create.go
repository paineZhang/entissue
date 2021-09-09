// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wing/models/ent/organization"
	"wing/models/ent/orgunit"
	"wing/models/ent/orgunitmember"
	"wing/models/ent/orgunitposition"
	"wing/models/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrgUnitCreate is the builder for creating a OrgUnit entity.
type OrgUnitCreate struct {
	config
	mutation *OrgUnitMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateByUser sets the "create_by_user" field.
func (ouc *OrgUnitCreate) SetCreateByUser(i int) *OrgUnitCreate {
	ouc.mutation.SetCreateByUser(i)
	return ouc
}

// SetNillableCreateByUser sets the "create_by_user" field if the given value is not nil.
func (ouc *OrgUnitCreate) SetNillableCreateByUser(i *int) *OrgUnitCreate {
	if i != nil {
		ouc.SetCreateByUser(*i)
	}
	return ouc
}

// SetUpdateByUser sets the "update_by_user" field.
func (ouc *OrgUnitCreate) SetUpdateByUser(i int) *OrgUnitCreate {
	ouc.mutation.SetUpdateByUser(i)
	return ouc
}

// SetNillableUpdateByUser sets the "update_by_user" field if the given value is not nil.
func (ouc *OrgUnitCreate) SetNillableUpdateByUser(i *int) *OrgUnitCreate {
	if i != nil {
		ouc.SetUpdateByUser(*i)
	}
	return ouc
}

// SetCreateTime sets the "create_time" field.
func (ouc *OrgUnitCreate) SetCreateTime(t time.Time) *OrgUnitCreate {
	ouc.mutation.SetCreateTime(t)
	return ouc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ouc *OrgUnitCreate) SetNillableCreateTime(t *time.Time) *OrgUnitCreate {
	if t != nil {
		ouc.SetCreateTime(*t)
	}
	return ouc
}

// SetUpdateTime sets the "update_time" field.
func (ouc *OrgUnitCreate) SetUpdateTime(t time.Time) *OrgUnitCreate {
	ouc.mutation.SetUpdateTime(t)
	return ouc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ouc *OrgUnitCreate) SetNillableUpdateTime(t *time.Time) *OrgUnitCreate {
	if t != nil {
		ouc.SetUpdateTime(*t)
	}
	return ouc
}

// SetName sets the "name" field.
func (ouc *OrgUnitCreate) SetName(s string) *OrgUnitCreate {
	ouc.mutation.SetName(s)
	return ouc
}

// SetDuty sets the "duty" field.
func (ouc *OrgUnitCreate) SetDuty(s string) *OrgUnitCreate {
	ouc.mutation.SetDuty(s)
	return ouc
}

// SetNillableDuty sets the "duty" field if the given value is not nil.
func (ouc *OrgUnitCreate) SetNillableDuty(s *string) *OrgUnitCreate {
	if s != nil {
		ouc.SetDuty(*s)
	}
	return ouc
}

// SetCreateByID sets the "create_by" edge to the User entity by ID.
func (ouc *OrgUnitCreate) SetCreateByID(id int) *OrgUnitCreate {
	ouc.mutation.SetCreateByID(id)
	return ouc
}

// SetNillableCreateByID sets the "create_by" edge to the User entity by ID if the given value is not nil.
func (ouc *OrgUnitCreate) SetNillableCreateByID(id *int) *OrgUnitCreate {
	if id != nil {
		ouc = ouc.SetCreateByID(*id)
	}
	return ouc
}

// SetCreateBy sets the "create_by" edge to the User entity.
func (ouc *OrgUnitCreate) SetCreateBy(u *User) *OrgUnitCreate {
	return ouc.SetCreateByID(u.ID)
}

// SetUpdateByID sets the "update_by" edge to the User entity by ID.
func (ouc *OrgUnitCreate) SetUpdateByID(id int) *OrgUnitCreate {
	ouc.mutation.SetUpdateByID(id)
	return ouc
}

// SetNillableUpdateByID sets the "update_by" edge to the User entity by ID if the given value is not nil.
func (ouc *OrgUnitCreate) SetNillableUpdateByID(id *int) *OrgUnitCreate {
	if id != nil {
		ouc = ouc.SetUpdateByID(*id)
	}
	return ouc
}

// SetUpdateBy sets the "update_by" edge to the User entity.
func (ouc *OrgUnitCreate) SetUpdateBy(u *User) *OrgUnitCreate {
	return ouc.SetUpdateByID(u.ID)
}

// AddMemberIDs adds the "members" edge to the OrgUnitMember entity by IDs.
func (ouc *OrgUnitCreate) AddMemberIDs(ids ...int) *OrgUnitCreate {
	ouc.mutation.AddMemberIDs(ids...)
	return ouc
}

// AddMembers adds the "members" edges to the OrgUnitMember entity.
func (ouc *OrgUnitCreate) AddMembers(o ...*OrgUnitMember) *OrgUnitCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouc.AddMemberIDs(ids...)
}

// AddPositionIDs adds the "positions" edge to the OrgUnitPosition entity by IDs.
func (ouc *OrgUnitCreate) AddPositionIDs(ids ...int) *OrgUnitCreate {
	ouc.mutation.AddPositionIDs(ids...)
	return ouc
}

// AddPositions adds the "positions" edges to the OrgUnitPosition entity.
func (ouc *OrgUnitCreate) AddPositions(o ...*OrgUnitPosition) *OrgUnitCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouc.AddPositionIDs(ids...)
}

// SetSupUnitID sets the "supUnit" edge to the OrgUnit entity by ID.
func (ouc *OrgUnitCreate) SetSupUnitID(id int) *OrgUnitCreate {
	ouc.mutation.SetSupUnitID(id)
	return ouc
}

// SetNillableSupUnitID sets the "supUnit" edge to the OrgUnit entity by ID if the given value is not nil.
func (ouc *OrgUnitCreate) SetNillableSupUnitID(id *int) *OrgUnitCreate {
	if id != nil {
		ouc = ouc.SetSupUnitID(*id)
	}
	return ouc
}

// SetSupUnit sets the "supUnit" edge to the OrgUnit entity.
func (ouc *OrgUnitCreate) SetSupUnit(o *OrgUnit) *OrgUnitCreate {
	return ouc.SetSupUnitID(o.ID)
}

// AddSubUnitIDs adds the "subUnits" edge to the OrgUnit entity by IDs.
func (ouc *OrgUnitCreate) AddSubUnitIDs(ids ...int) *OrgUnitCreate {
	ouc.mutation.AddSubUnitIDs(ids...)
	return ouc
}

// AddSubUnits adds the "subUnits" edges to the OrgUnit entity.
func (ouc *OrgUnitCreate) AddSubUnits(o ...*OrgUnit) *OrgUnitCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouc.AddSubUnitIDs(ids...)
}

// SetBelongToOrgID sets the "belongToOrg" edge to the Organization entity by ID.
func (ouc *OrgUnitCreate) SetBelongToOrgID(id int) *OrgUnitCreate {
	ouc.mutation.SetBelongToOrgID(id)
	return ouc
}

// SetBelongToOrg sets the "belongToOrg" edge to the Organization entity.
func (ouc *OrgUnitCreate) SetBelongToOrg(o *Organization) *OrgUnitCreate {
	return ouc.SetBelongToOrgID(o.ID)
}

// Mutation returns the OrgUnitMutation object of the builder.
func (ouc *OrgUnitCreate) Mutation() *OrgUnitMutation {
	return ouc.mutation
}

// Save creates the OrgUnit in the database.
func (ouc *OrgUnitCreate) Save(ctx context.Context) (*OrgUnit, error) {
	var (
		err  error
		node *OrgUnit
	)
	if err := ouc.defaults(); err != nil {
		return nil, err
	}
	if len(ouc.hooks) == 0 {
		if err = ouc.check(); err != nil {
			return nil, err
		}
		node, err = ouc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrgUnitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouc.check(); err != nil {
				return nil, err
			}
			ouc.mutation = mutation
			if node, err = ouc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ouc.hooks) - 1; i >= 0; i-- {
			if ouc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ouc *OrgUnitCreate) SaveX(ctx context.Context) *OrgUnit {
	v, err := ouc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ouc *OrgUnitCreate) Exec(ctx context.Context) error {
	_, err := ouc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouc *OrgUnitCreate) ExecX(ctx context.Context) {
	if err := ouc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouc *OrgUnitCreate) defaults() error {
	if _, ok := ouc.mutation.CreateTime(); !ok {
		if orgunit.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized orgunit.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := orgunit.DefaultCreateTime()
		ouc.mutation.SetCreateTime(v)
	}
	if _, ok := ouc.mutation.UpdateTime(); !ok {
		if orgunit.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized orgunit.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := orgunit.DefaultUpdateTime()
		ouc.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ouc *OrgUnitCreate) check() error {
	if _, ok := ouc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := ouc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if _, ok := ouc.mutation.BelongToOrgID(); !ok {
		return &ValidationError{Name: "belongToOrg", err: errors.New("ent: missing required edge \"belongToOrg\"")}
	}
	return nil
}

func (ouc *OrgUnitCreate) sqlSave(ctx context.Context) (*OrgUnit, error) {
	_node, _spec := ouc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ouc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ouc *OrgUnitCreate) createSpec() (*OrgUnit, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgUnit{config: ouc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orgunit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orgunit.FieldID,
			},
		}
	)
	_spec.Schema = ouc.schemaConfig.OrgUnit
	_spec.OnConflict = ouc.conflict
	if value, ok := ouc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orgunit.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ouc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orgunit.FieldUpdateTime,
		})
		_node.UpdateTime = &value
	}
	if value, ok := ouc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orgunit.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ouc.mutation.Duty(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orgunit.FieldDuty,
		})
		_node.Duty = value
	}
	if nodes := ouc.mutation.CreateByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgunit.CreateByTable,
			Columns: []string{orgunit.CreateByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		edge.Schema = ouc.schemaConfig.OrgUnit
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreateByUser = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.UpdateByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgunit.UpdateByTable,
			Columns: []string{orgunit.UpdateByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		edge.Schema = ouc.schemaConfig.OrgUnit
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UpdateByUser = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orgunit.MembersTable,
			Columns: []string{orgunit.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orgunitmember.FieldID,
				},
			},
		}
		edge.Schema = ouc.schemaConfig.OrgUnitMember
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.PositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orgunit.PositionsTable,
			Columns: []string{orgunit.PositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orgunitposition.FieldID,
				},
			},
		}
		edge.Schema = ouc.schemaConfig.OrgUnitPosition
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.SupUnitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orgunit.SupUnitTable,
			Columns: []string{orgunit.SupUnitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orgunit.FieldID,
				},
			},
		}
		edge.Schema = ouc.schemaConfig.OrgUnit
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.org_unit_sub_units = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.SubUnitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orgunit.SubUnitsTable,
			Columns: []string{orgunit.SubUnitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orgunit.FieldID,
				},
			},
		}
		edge.Schema = ouc.schemaConfig.OrgUnit
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.BelongToOrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orgunit.BelongToOrgTable,
			Columns: []string{orgunit.BelongToOrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		edge.Schema = ouc.schemaConfig.OrgUnit
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_units = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgUnit.Create().
//		SetCreateByUser(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgUnitUpsert) {
//			SetCreateByUser(v+v).
//		}).
//		Exec(ctx)
//
func (ouc *OrgUnitCreate) OnConflict(opts ...sql.ConflictOption) *OrgUnitUpsertOne {
	ouc.conflict = opts
	return &OrgUnitUpsertOne{
		create: ouc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgUnit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ouc *OrgUnitCreate) OnConflictColumns(columns ...string) *OrgUnitUpsertOne {
	ouc.conflict = append(ouc.conflict, sql.ConflictColumns(columns...))
	return &OrgUnitUpsertOne{
		create: ouc,
	}
}

type (
	// OrgUnitUpsertOne is the builder for "upsert"-ing
	//  one OrgUnit node.
	OrgUnitUpsertOne struct {
		create *OrgUnitCreate
	}

	// OrgUnitUpsert is the "OnConflict" setter.
	OrgUnitUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateByUser sets the "create_by_user" field.
func (u *OrgUnitUpsert) SetCreateByUser(v int) *OrgUnitUpsert {
	u.Set(orgunit.FieldCreateByUser, v)
	return u
}

// UpdateCreateByUser sets the "create_by_user" field to the value that was provided on create.
func (u *OrgUnitUpsert) UpdateCreateByUser() *OrgUnitUpsert {
	u.SetExcluded(orgunit.FieldCreateByUser)
	return u
}

// ClearCreateByUser clears the value of the "create_by_user" field.
func (u *OrgUnitUpsert) ClearCreateByUser() *OrgUnitUpsert {
	u.SetNull(orgunit.FieldCreateByUser)
	return u
}

// SetUpdateByUser sets the "update_by_user" field.
func (u *OrgUnitUpsert) SetUpdateByUser(v int) *OrgUnitUpsert {
	u.Set(orgunit.FieldUpdateByUser, v)
	return u
}

// UpdateUpdateByUser sets the "update_by_user" field to the value that was provided on create.
func (u *OrgUnitUpsert) UpdateUpdateByUser() *OrgUnitUpsert {
	u.SetExcluded(orgunit.FieldUpdateByUser)
	return u
}

// ClearUpdateByUser clears the value of the "update_by_user" field.
func (u *OrgUnitUpsert) ClearUpdateByUser() *OrgUnitUpsert {
	u.SetNull(orgunit.FieldUpdateByUser)
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *OrgUnitUpsert) SetCreateTime(v time.Time) *OrgUnitUpsert {
	u.Set(orgunit.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *OrgUnitUpsert) UpdateCreateTime() *OrgUnitUpsert {
	u.SetExcluded(orgunit.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *OrgUnitUpsert) SetUpdateTime(v time.Time) *OrgUnitUpsert {
	u.Set(orgunit.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *OrgUnitUpsert) UpdateUpdateTime() *OrgUnitUpsert {
	u.SetExcluded(orgunit.FieldUpdateTime)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *OrgUnitUpsert) ClearUpdateTime() *OrgUnitUpsert {
	u.SetNull(orgunit.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *OrgUnitUpsert) SetName(v string) *OrgUnitUpsert {
	u.Set(orgunit.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrgUnitUpsert) UpdateName() *OrgUnitUpsert {
	u.SetExcluded(orgunit.FieldName)
	return u
}

// SetDuty sets the "duty" field.
func (u *OrgUnitUpsert) SetDuty(v string) *OrgUnitUpsert {
	u.Set(orgunit.FieldDuty, v)
	return u
}

// UpdateDuty sets the "duty" field to the value that was provided on create.
func (u *OrgUnitUpsert) UpdateDuty() *OrgUnitUpsert {
	u.SetExcluded(orgunit.FieldDuty)
	return u
}

// ClearDuty clears the value of the "duty" field.
func (u *OrgUnitUpsert) ClearDuty() *OrgUnitUpsert {
	u.SetNull(orgunit.FieldDuty)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.OrgUnit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *OrgUnitUpsertOne) UpdateNewValues() *OrgUnitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.OrgUnit.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *OrgUnitUpsertOne) Ignore() *OrgUnitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgUnitUpsertOne) DoNothing() *OrgUnitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgUnitCreate.OnConflict
// documentation for more info.
func (u *OrgUnitUpsertOne) Update(set func(*OrgUnitUpsert)) *OrgUnitUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgUnitUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateByUser sets the "create_by_user" field.
func (u *OrgUnitUpsertOne) SetCreateByUser(v int) *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetCreateByUser(v)
	})
}

// UpdateCreateByUser sets the "create_by_user" field to the value that was provided on create.
func (u *OrgUnitUpsertOne) UpdateCreateByUser() *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateCreateByUser()
	})
}

// ClearCreateByUser clears the value of the "create_by_user" field.
func (u *OrgUnitUpsertOne) ClearCreateByUser() *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.ClearCreateByUser()
	})
}

// SetUpdateByUser sets the "update_by_user" field.
func (u *OrgUnitUpsertOne) SetUpdateByUser(v int) *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetUpdateByUser(v)
	})
}

// UpdateUpdateByUser sets the "update_by_user" field to the value that was provided on create.
func (u *OrgUnitUpsertOne) UpdateUpdateByUser() *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateUpdateByUser()
	})
}

// ClearUpdateByUser clears the value of the "update_by_user" field.
func (u *OrgUnitUpsertOne) ClearUpdateByUser() *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.ClearUpdateByUser()
	})
}

// SetCreateTime sets the "create_time" field.
func (u *OrgUnitUpsertOne) SetCreateTime(v time.Time) *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *OrgUnitUpsertOne) UpdateCreateTime() *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *OrgUnitUpsertOne) SetUpdateTime(v time.Time) *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *OrgUnitUpsertOne) UpdateUpdateTime() *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *OrgUnitUpsertOne) ClearUpdateTime() *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.ClearUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *OrgUnitUpsertOne) SetName(v string) *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrgUnitUpsertOne) UpdateName() *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateName()
	})
}

// SetDuty sets the "duty" field.
func (u *OrgUnitUpsertOne) SetDuty(v string) *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetDuty(v)
	})
}

// UpdateDuty sets the "duty" field to the value that was provided on create.
func (u *OrgUnitUpsertOne) UpdateDuty() *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateDuty()
	})
}

// ClearDuty clears the value of the "duty" field.
func (u *OrgUnitUpsertOne) ClearDuty() *OrgUnitUpsertOne {
	return u.Update(func(s *OrgUnitUpsert) {
		s.ClearDuty()
	})
}

// Exec executes the query.
func (u *OrgUnitUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrgUnitCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgUnitUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrgUnitUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrgUnitUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrgUnitCreateBulk is the builder for creating many OrgUnit entities in bulk.
type OrgUnitCreateBulk struct {
	config
	builders []*OrgUnitCreate
	conflict []sql.ConflictOption
}

// Save creates the OrgUnit entities in the database.
func (oucb *OrgUnitCreateBulk) Save(ctx context.Context) ([]*OrgUnit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oucb.builders))
	nodes := make([]*OrgUnit, len(oucb.builders))
	mutators := make([]Mutator, len(oucb.builders))
	for i := range oucb.builders {
		func(i int, root context.Context) {
			builder := oucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgUnitMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, oucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = oucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oucb *OrgUnitCreateBulk) SaveX(ctx context.Context) []*OrgUnit {
	v, err := oucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oucb *OrgUnitCreateBulk) Exec(ctx context.Context) error {
	_, err := oucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oucb *OrgUnitCreateBulk) ExecX(ctx context.Context) {
	if err := oucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgUnit.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgUnitUpsert) {
//			SetCreateByUser(v+v).
//		}).
//		Exec(ctx)
//
func (oucb *OrgUnitCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrgUnitUpsertBulk {
	oucb.conflict = opts
	return &OrgUnitUpsertBulk{
		create: oucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgUnit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (oucb *OrgUnitCreateBulk) OnConflictColumns(columns ...string) *OrgUnitUpsertBulk {
	oucb.conflict = append(oucb.conflict, sql.ConflictColumns(columns...))
	return &OrgUnitUpsertBulk{
		create: oucb,
	}
}

// OrgUnitUpsertBulk is the builder for "upsert"-ing
// a bulk of OrgUnit nodes.
type OrgUnitUpsertBulk struct {
	create *OrgUnitCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrgUnit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *OrgUnitUpsertBulk) UpdateNewValues() *OrgUnitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgUnit.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *OrgUnitUpsertBulk) Ignore() *OrgUnitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgUnitUpsertBulk) DoNothing() *OrgUnitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgUnitCreateBulk.OnConflict
// documentation for more info.
func (u *OrgUnitUpsertBulk) Update(set func(*OrgUnitUpsert)) *OrgUnitUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgUnitUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateByUser sets the "create_by_user" field.
func (u *OrgUnitUpsertBulk) SetCreateByUser(v int) *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetCreateByUser(v)
	})
}

// UpdateCreateByUser sets the "create_by_user" field to the value that was provided on create.
func (u *OrgUnitUpsertBulk) UpdateCreateByUser() *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateCreateByUser()
	})
}

// ClearCreateByUser clears the value of the "create_by_user" field.
func (u *OrgUnitUpsertBulk) ClearCreateByUser() *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.ClearCreateByUser()
	})
}

// SetUpdateByUser sets the "update_by_user" field.
func (u *OrgUnitUpsertBulk) SetUpdateByUser(v int) *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetUpdateByUser(v)
	})
}

// UpdateUpdateByUser sets the "update_by_user" field to the value that was provided on create.
func (u *OrgUnitUpsertBulk) UpdateUpdateByUser() *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateUpdateByUser()
	})
}

// ClearUpdateByUser clears the value of the "update_by_user" field.
func (u *OrgUnitUpsertBulk) ClearUpdateByUser() *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.ClearUpdateByUser()
	})
}

// SetCreateTime sets the "create_time" field.
func (u *OrgUnitUpsertBulk) SetCreateTime(v time.Time) *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *OrgUnitUpsertBulk) UpdateCreateTime() *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *OrgUnitUpsertBulk) SetUpdateTime(v time.Time) *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *OrgUnitUpsertBulk) UpdateUpdateTime() *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *OrgUnitUpsertBulk) ClearUpdateTime() *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.ClearUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *OrgUnitUpsertBulk) SetName(v string) *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrgUnitUpsertBulk) UpdateName() *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateName()
	})
}

// SetDuty sets the "duty" field.
func (u *OrgUnitUpsertBulk) SetDuty(v string) *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.SetDuty(v)
	})
}

// UpdateDuty sets the "duty" field to the value that was provided on create.
func (u *OrgUnitUpsertBulk) UpdateDuty() *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.UpdateDuty()
	})
}

// ClearDuty clears the value of the "duty" field.
func (u *OrgUnitUpsertBulk) ClearDuty() *OrgUnitUpsertBulk {
	return u.Update(func(s *OrgUnitUpsert) {
		s.ClearDuty()
	})
}

// Exec executes the query.
func (u *OrgUnitUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrgUnitCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrgUnitCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgUnitUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}