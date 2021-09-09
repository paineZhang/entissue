package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"wing/models/ent"
	apierror "wing/resolver/error"
	"wing/resolver/generated"
	"wing/resolver/middleware"
	model1 "wing/resolver/model"
)

func (r *mutationResolver) OrganizationMutation(ctx context.Context) (*ent.OrganizationMutation, error) {
	return &ent.OrganizationMutation{}, middleware.CheckAuthInfo(ctx, ent.FromContext(ctx))
}

func (r *orgUnitResolver) Positons(ctx context.Context, obj *ent.OrgUnit) ([]*ent.OrgUnitPosition, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *organizationMutationResolver) CreateOrganization(ctx context.Context, obj *ent.OrganizationMutation, input ent.CreateOrganizationInput) (*ent.Organization, error) {
	model := ent.FromContext(ctx)

	count, err := model.Organization.Query().Count(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	if count != 0 {
		return nil, apierror.New(ctx, model1.ErrorCodeAlreadyExist, "已经存在企业组织实例，当前系统没有涉及多租户")
	}

	org, err := model.Organization.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}
	return org, nil
}

func (r *organizationMutationResolver) UpdateOrganization(ctx context.Context, obj *ent.OrganizationMutation, id int, input ent.UpdateOrganizationInput) (bool, error) {
	model := ent.FromContext(ctx)

	err := model.Organization.UpdateOneID(id).SetInput(input).Exec(ctx)
	if err != nil {
		return false, apierror.Transform(ctx, err)
	}
	return true, nil
}

func (r *organizationMutationResolver) CreateOrgUnit(ctx context.Context, obj *ent.OrganizationMutation, input ent.CreateOrgUnitInput) (*ent.OrgUnit, error) {
	org, err := ent.FromContext(ctx).OrgUnit.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}
	return org, nil
}

func (r *organizationMutationResolver) UpdateOrgUnit(ctx context.Context, obj *ent.OrganizationMutation, id int, input ent.UpdateOrgUnitInput) (bool, error) {
	err := ent.FromContext(ctx).OrgUnit.UpdateOneID(id).SetInput(input).Exec(ctx)
	if err != nil {
		return false, apierror.Transform(ctx, err)
	}
	return true, nil
}

func (r *organizationMutationResolver) CreateOrgUnitMember(ctx context.Context, obj *ent.OrganizationMutation, input ent.CreateOrgUnitMemberInput) (*ent.OrgUnitMember, error) {
	org, err := ent.FromContext(ctx).OrgUnitMember.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}
	return org, nil
}

func (r *organizationMutationResolver) UpdateOrgUnitMember(ctx context.Context, obj *ent.OrganizationMutation, id int, input ent.UpdateOrgUnitMemberInput) (bool, error) {
	err := ent.FromContext(ctx).OrgUnitMember.UpdateOneID(id).SetInput(input).Exec(ctx)
	if err != nil {
		return false, apierror.Transform(ctx, err)
	}
	return true, nil
}

func (r *organizationMutationResolver) CreateOrgUnitPosition(ctx context.Context, obj *ent.OrganizationMutation, input ent.CreateOrgUnitPositionInput) (*ent.OrgUnitPosition, error) {
	org, err := ent.FromContext(ctx).OrgUnitPosition.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}
	return org, nil
}

func (r *organizationMutationResolver) UpdateOrgUnitPosition(ctx context.Context, obj *ent.OrganizationMutation, id int, input ent.UpdateOrgUnitPositionInput) (bool, error) {
	err := ent.FromContext(ctx).OrgUnitPosition.UpdateOneID(id).SetInput(input).Exec(ctx)
	if err != nil {
		return false, apierror.Transform(ctx, err)
	}
	return true, nil
}

func (r *organizationQueryResolver) Organization(ctx context.Context, obj *ent.OrganizationQuery) (*ent.Organization, error) {
	org, err := ent.FromContext(ctx).Organization.Query().First(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}
	return org, err
}

func (r *organizationQueryResolver) OrgUnits(ctx context.Context, obj *ent.OrganizationQuery, filters *ent.OrgUnitWhereInput) ([]*ent.OrgUnit, error) {
	cnn, err := ent.FromContext(ctx).OrgUnit.Query().
		Paginate(ctx, nil, nil, nil, nil, ent.WithOrgUnitFilter(filters.Filter))

	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	if cnn == nil || cnn.TotalCount == 0 {
		return nil, nil
	}

	ret := make([]*ent.OrgUnit, cnn.TotalCount)

	for _, edge := range cnn.Edges {
		ret = append(ret, edge.Node)
	}
	return ret, nil
}

func (r *organizationQueryResolver) OrgUnitMembers(ctx context.Context, obj *ent.OrganizationQuery, filters *ent.OrgUnitMemberWhereInput) ([]*ent.OrgUnitMember, error) {
	cnn, err := ent.FromContext(ctx).OrgUnitMember.Query().
		Paginate(ctx, nil, nil, nil, nil, ent.WithOrgUnitMemberFilter(filters.Filter))

	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	if cnn == nil || cnn.TotalCount == 0 {
		return nil, nil
	}

	ret := make([]*ent.OrgUnitMember, cnn.TotalCount)

	for _, edge := range cnn.Edges {
		ret = append(ret, edge.Node)
	}
	return ret, nil
}

func (r *organizationQueryResolver) OrgUnitPositions(ctx context.Context, obj *ent.OrganizationQuery, filters *ent.OrgUnitPositionWhereInput) ([]*ent.OrgUnitPosition, error) {
	cnn, err := ent.FromContext(ctx).OrgUnitPosition.Query().
		Paginate(ctx, nil, nil, nil, nil, ent.WithOrgUnitPositionFilter(filters.Filter))

	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	if cnn == nil || cnn.TotalCount == 0 {
		return nil, nil
	}

	ret := make([]*ent.OrgUnitPosition, cnn.TotalCount)

	for _, edge := range cnn.Edges {
		ret = append(ret, edge.Node)
	}
	return ret, nil
}

func (r *queryResolver) OrganizationQuery(ctx context.Context) (*ent.OrganizationQuery, error) {
	return &ent.OrganizationQuery{}, middleware.CheckAuthInfo(ctx, ent.FromContext(ctx))
}

// OrgUnit returns generated.OrgUnitResolver implementation.
func (r *Resolver) OrgUnit() generated.OrgUnitResolver { return &orgUnitResolver{r} }

// OrganizationMutation returns generated.OrganizationMutationResolver implementation.
func (r *Resolver) OrganizationMutation() generated.OrganizationMutationResolver {
	return &organizationMutationResolver{r}
}

// OrganizationQuery returns generated.OrganizationQueryResolver implementation.
func (r *Resolver) OrganizationQuery() generated.OrganizationQueryResolver {
	return &organizationQueryResolver{r}
}

type orgUnitResolver struct{ *Resolver }
type organizationMutationResolver struct{ *Resolver }
type organizationQueryResolver struct{ *Resolver }
