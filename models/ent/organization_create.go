// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wing/models/ent/organization"
	"wing/models/ent/orgunit"
	"wing/models/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrganizationCreate is the builder for creating a Organization entity.
type OrganizationCreate struct {
	config
	mutation *OrganizationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateByUser sets the "create_by_user" field.
func (oc *OrganizationCreate) SetCreateByUser(i int) *OrganizationCreate {
	oc.mutation.SetCreateByUser(i)
	return oc
}

// SetNillableCreateByUser sets the "create_by_user" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreateByUser(i *int) *OrganizationCreate {
	if i != nil {
		oc.SetCreateByUser(*i)
	}
	return oc
}

// SetUpdateByUser sets the "update_by_user" field.
func (oc *OrganizationCreate) SetUpdateByUser(i int) *OrganizationCreate {
	oc.mutation.SetUpdateByUser(i)
	return oc
}

// SetNillableUpdateByUser sets the "update_by_user" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdateByUser(i *int) *OrganizationCreate {
	if i != nil {
		oc.SetUpdateByUser(*i)
	}
	return oc
}

// SetCreateTime sets the "create_time" field.
func (oc *OrganizationCreate) SetCreateTime(t time.Time) *OrganizationCreate {
	oc.mutation.SetCreateTime(t)
	return oc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreateTime(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetCreateTime(*t)
	}
	return oc
}

// SetUpdateTime sets the "update_time" field.
func (oc *OrganizationCreate) SetUpdateTime(t time.Time) *OrganizationCreate {
	oc.mutation.SetUpdateTime(t)
	return oc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdateTime(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetUpdateTime(*t)
	}
	return oc
}

// SetName sets the "name" field.
func (oc *OrganizationCreate) SetName(s string) *OrganizationCreate {
	oc.mutation.SetName(s)
	return oc
}

// SetCreateByID sets the "create_by" edge to the User entity by ID.
func (oc *OrganizationCreate) SetCreateByID(id int) *OrganizationCreate {
	oc.mutation.SetCreateByID(id)
	return oc
}

// SetNillableCreateByID sets the "create_by" edge to the User entity by ID if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreateByID(id *int) *OrganizationCreate {
	if id != nil {
		oc = oc.SetCreateByID(*id)
	}
	return oc
}

// SetCreateBy sets the "create_by" edge to the User entity.
func (oc *OrganizationCreate) SetCreateBy(u *User) *OrganizationCreate {
	return oc.SetCreateByID(u.ID)
}

// SetUpdateByID sets the "update_by" edge to the User entity by ID.
func (oc *OrganizationCreate) SetUpdateByID(id int) *OrganizationCreate {
	oc.mutation.SetUpdateByID(id)
	return oc
}

// SetNillableUpdateByID sets the "update_by" edge to the User entity by ID if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdateByID(id *int) *OrganizationCreate {
	if id != nil {
		oc = oc.SetUpdateByID(*id)
	}
	return oc
}

// SetUpdateBy sets the "update_by" edge to the User entity.
func (oc *OrganizationCreate) SetUpdateBy(u *User) *OrganizationCreate {
	return oc.SetUpdateByID(u.ID)
}

// AddUnitIDs adds the "units" edge to the OrgUnit entity by IDs.
func (oc *OrganizationCreate) AddUnitIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddUnitIDs(ids...)
	return oc
}

// AddUnits adds the "units" edges to the OrgUnit entity.
func (oc *OrganizationCreate) AddUnits(o ...*OrgUnit) *OrganizationCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddUnitIDs(ids...)
}

// AddStaffIDs adds the "staffs" edge to the User entity by IDs.
func (oc *OrganizationCreate) AddStaffIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddStaffIDs(ids...)
	return oc
}

// AddStaffs adds the "staffs" edges to the User entity.
func (oc *OrganizationCreate) AddStaffs(u ...*User) *OrganizationCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return oc.AddStaffIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (oc *OrganizationCreate) Mutation() *OrganizationMutation {
	return oc.mutation
}

// Save creates the Organization in the database.
func (oc *OrganizationCreate) Save(ctx context.Context) (*Organization, error) {
	var (
		err  error
		node *Organization
	)
	if err := oc.defaults(); err != nil {
		return nil, err
	}
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrganizationCreate) SaveX(ctx context.Context) *Organization {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrganizationCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrganizationCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrganizationCreate) defaults() error {
	if _, ok := oc.mutation.CreateTime(); !ok {
		if organization.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized organization.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := organization.DefaultCreateTime()
		oc.mutation.SetCreateTime(v)
	}
	if _, ok := oc.mutation.UpdateTime(); !ok {
		if organization.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized organization.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := organization.DefaultUpdateTime()
		oc.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrganizationCreate) check() error {
	if _, ok := oc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := oc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	return nil
}

func (oc *OrganizationCreate) sqlSave(ctx context.Context) (*Organization, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *OrganizationCreate) createSpec() (*Organization, *sqlgraph.CreateSpec) {
	var (
		_node = &Organization{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: organization.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: organization.FieldID,
			},
		}
	)
	_spec.Schema = oc.schemaConfig.Organization
	_spec.OnConflict = oc.conflict
	if value, ok := oc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organization.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := oc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: organization.FieldUpdateTime,
		})
		_node.UpdateTime = &value
	}
	if value, ok := oc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: organization.FieldName,
		})
		_node.Name = value
	}
	if nodes := oc.mutation.CreateByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   organization.CreateByTable,
			Columns: []string{organization.CreateByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		edge.Schema = oc.schemaConfig.Organization
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CreateByUser = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.UpdateByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   organization.UpdateByTable,
			Columns: []string{organization.UpdateByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		edge.Schema = oc.schemaConfig.Organization
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UpdateByUser = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.UnitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.UnitsTable,
			Columns: []string{organization.UnitsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orgunit.FieldID,
				},
			},
		}
		edge.Schema = oc.schemaConfig.OrgUnit
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.StaffsTable,
			Columns: organization.StaffsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		edge.Schema = oc.schemaConfig.OrganizationStaffs
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Organization.Create().
//		SetCreateByUser(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrganizationUpsert) {
//			SetCreateByUser(v+v).
//		}).
//		Exec(ctx)
//
func (oc *OrganizationCreate) OnConflict(opts ...sql.ConflictOption) *OrganizationUpsertOne {
	oc.conflict = opts
	return &OrganizationUpsertOne{
		create: oc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Organization.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (oc *OrganizationCreate) OnConflictColumns(columns ...string) *OrganizationUpsertOne {
	oc.conflict = append(oc.conflict, sql.ConflictColumns(columns...))
	return &OrganizationUpsertOne{
		create: oc,
	}
}

type (
	// OrganizationUpsertOne is the builder for "upsert"-ing
	//  one Organization node.
	OrganizationUpsertOne struct {
		create *OrganizationCreate
	}

	// OrganizationUpsert is the "OnConflict" setter.
	OrganizationUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateByUser sets the "create_by_user" field.
func (u *OrganizationUpsert) SetCreateByUser(v int) *OrganizationUpsert {
	u.Set(organization.FieldCreateByUser, v)
	return u
}

// UpdateCreateByUser sets the "create_by_user" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateCreateByUser() *OrganizationUpsert {
	u.SetExcluded(organization.FieldCreateByUser)
	return u
}

// ClearCreateByUser clears the value of the "create_by_user" field.
func (u *OrganizationUpsert) ClearCreateByUser() *OrganizationUpsert {
	u.SetNull(organization.FieldCreateByUser)
	return u
}

// SetUpdateByUser sets the "update_by_user" field.
func (u *OrganizationUpsert) SetUpdateByUser(v int) *OrganizationUpsert {
	u.Set(organization.FieldUpdateByUser, v)
	return u
}

// UpdateUpdateByUser sets the "update_by_user" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateUpdateByUser() *OrganizationUpsert {
	u.SetExcluded(organization.FieldUpdateByUser)
	return u
}

// ClearUpdateByUser clears the value of the "update_by_user" field.
func (u *OrganizationUpsert) ClearUpdateByUser() *OrganizationUpsert {
	u.SetNull(organization.FieldUpdateByUser)
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *OrganizationUpsert) SetCreateTime(v time.Time) *OrganizationUpsert {
	u.Set(organization.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateCreateTime() *OrganizationUpsert {
	u.SetExcluded(organization.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *OrganizationUpsert) SetUpdateTime(v time.Time) *OrganizationUpsert {
	u.Set(organization.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateUpdateTime() *OrganizationUpsert {
	u.SetExcluded(organization.FieldUpdateTime)
	return u
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *OrganizationUpsert) ClearUpdateTime() *OrganizationUpsert {
	u.SetNull(organization.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *OrganizationUpsert) SetName(v string) *OrganizationUpsert {
	u.Set(organization.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrganizationUpsert) UpdateName() *OrganizationUpsert {
	u.SetExcluded(organization.FieldName)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Organization.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *OrganizationUpsertOne) UpdateNewValues() *OrganizationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Organization.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *OrganizationUpsertOne) Ignore() *OrganizationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrganizationUpsertOne) DoNothing() *OrganizationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrganizationCreate.OnConflict
// documentation for more info.
func (u *OrganizationUpsertOne) Update(set func(*OrganizationUpsert)) *OrganizationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrganizationUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateByUser sets the "create_by_user" field.
func (u *OrganizationUpsertOne) SetCreateByUser(v int) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetCreateByUser(v)
	})
}

// UpdateCreateByUser sets the "create_by_user" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateCreateByUser() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateCreateByUser()
	})
}

// ClearCreateByUser clears the value of the "create_by_user" field.
func (u *OrganizationUpsertOne) ClearCreateByUser() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearCreateByUser()
	})
}

// SetUpdateByUser sets the "update_by_user" field.
func (u *OrganizationUpsertOne) SetUpdateByUser(v int) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetUpdateByUser(v)
	})
}

// UpdateUpdateByUser sets the "update_by_user" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateUpdateByUser() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateUpdateByUser()
	})
}

// ClearUpdateByUser clears the value of the "update_by_user" field.
func (u *OrganizationUpsertOne) ClearUpdateByUser() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearUpdateByUser()
	})
}

// SetCreateTime sets the "create_time" field.
func (u *OrganizationUpsertOne) SetCreateTime(v time.Time) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateCreateTime() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *OrganizationUpsertOne) SetUpdateTime(v time.Time) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateUpdateTime() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *OrganizationUpsertOne) ClearUpdateTime() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *OrganizationUpsertOne) SetName(v string) *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrganizationUpsertOne) UpdateName() *OrganizationUpsertOne {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *OrganizationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrganizationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrganizationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrganizationUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrganizationUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrganizationCreateBulk is the builder for creating many Organization entities in bulk.
type OrganizationCreateBulk struct {
	config
	builders []*OrganizationCreate
	conflict []sql.ConflictOption
}

// Save creates the Organization entities in the database.
func (ocb *OrganizationCreateBulk) Save(ctx context.Context) ([]*Organization, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Organization, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ocb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) SaveX(ctx context.Context) []*Organization {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrganizationCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Organization.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrganizationUpsert) {
//			SetCreateByUser(v+v).
//		}).
//		Exec(ctx)
//
func (ocb *OrganizationCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrganizationUpsertBulk {
	ocb.conflict = opts
	return &OrganizationUpsertBulk{
		create: ocb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Organization.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ocb *OrganizationCreateBulk) OnConflictColumns(columns ...string) *OrganizationUpsertBulk {
	ocb.conflict = append(ocb.conflict, sql.ConflictColumns(columns...))
	return &OrganizationUpsertBulk{
		create: ocb,
	}
}

// OrganizationUpsertBulk is the builder for "upsert"-ing
// a bulk of Organization nodes.
type OrganizationUpsertBulk struct {
	create *OrganizationCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Organization.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *OrganizationUpsertBulk) UpdateNewValues() *OrganizationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Organization.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *OrganizationUpsertBulk) Ignore() *OrganizationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrganizationUpsertBulk) DoNothing() *OrganizationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrganizationCreateBulk.OnConflict
// documentation for more info.
func (u *OrganizationUpsertBulk) Update(set func(*OrganizationUpsert)) *OrganizationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrganizationUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateByUser sets the "create_by_user" field.
func (u *OrganizationUpsertBulk) SetCreateByUser(v int) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetCreateByUser(v)
	})
}

// UpdateCreateByUser sets the "create_by_user" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateCreateByUser() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateCreateByUser()
	})
}

// ClearCreateByUser clears the value of the "create_by_user" field.
func (u *OrganizationUpsertBulk) ClearCreateByUser() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearCreateByUser()
	})
}

// SetUpdateByUser sets the "update_by_user" field.
func (u *OrganizationUpsertBulk) SetUpdateByUser(v int) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetUpdateByUser(v)
	})
}

// UpdateUpdateByUser sets the "update_by_user" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateUpdateByUser() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateUpdateByUser()
	})
}

// ClearUpdateByUser clears the value of the "update_by_user" field.
func (u *OrganizationUpsertBulk) ClearUpdateByUser() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearUpdateByUser()
	})
}

// SetCreateTime sets the "create_time" field.
func (u *OrganizationUpsertBulk) SetCreateTime(v time.Time) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateCreateTime() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *OrganizationUpsertBulk) SetUpdateTime(v time.Time) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateUpdateTime() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateUpdateTime()
	})
}

// ClearUpdateTime clears the value of the "update_time" field.
func (u *OrganizationUpsertBulk) ClearUpdateTime() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.ClearUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *OrganizationUpsertBulk) SetName(v string) *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *OrganizationUpsertBulk) UpdateName() *OrganizationUpsertBulk {
	return u.Update(func(s *OrganizationUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *OrganizationUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrganizationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrganizationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrganizationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
