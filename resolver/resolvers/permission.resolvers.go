package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	casbinadapter "wing/adapters/casbin"
	"wing/models/ent"
	apierror "wing/resolver/error"
	"wing/resolver/generated"
	"wing/resolver/middleware"
	"wing/resolver/model"
)

func (r *mutationResolver) PermissionMutation(ctx context.Context) (*model.PermissionMutation, error) {
	return &model.PermissionMutation{}, middleware.CheckAuthInfo(ctx, ent.FromContext(ctx))
}

func (r *permissionMutationResolver) UpdatePermissionPoliciesForRole(ctx context.Context, obj *model.PermissionMutation, orgPositionID int, allPolicies []*model.UpdatePermissionPolicyInput) (bool, error) {
	// TODO 确保所有的resource都有
	client := ent.FromContext(ctx)

	orgId := client.Organization.Query().FirstIDX(ctx)

	var cb *casbinadapter.CasbinAdapter
	containerResolve(&cb)

	polices := make([]casbinadapter.Policy, len(allPolicies))

	for idx, p := range allPolicies {
		polices[idx].ResourceID = p.ResourceID
		polices[idx].Action = casbinadapter.Action(p.Action)
		polices[idx].Effect = casbinadapter.Effect(p.Effect)
	}

	err := cb.UpdatePoliciesForRoleInDomain(OrgDomainPrefix, orgId, orgPositionID, polices)
	if err != nil {
		return false, apierror.Transform(ctx, err)
	}

	return true, nil
}

func (r *permissionQueryResolver) Resources(ctx context.Context, obj *model.PermissionQuery, filter *model.PermissionResourceWhereInput) ([]*ent.Resource, error) {
	resources, err := ent.FromContext(ctx).Resource.Query().All(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}
	return resources, nil
}

func (r *queryResolver) PermissionQuery(ctx context.Context) (*model.PermissionQuery, error) {
	return &model.PermissionQuery{}, middleware.CheckAuthInfo(ctx, ent.FromContext(ctx))
}

func (r *resourceResolver) Policies(ctx context.Context, obj *ent.Resource) ([]*model.PermissionPolicy, error) {
	var cb *casbinadapter.CasbinAdapter
	containerResolve(&cb)

	client := ent.FromContext(ctx)

	orgId := client.Organization.Query().FirstIDX(ctx)

	rps := cb.GetPoliciesForResourceInDomain(OrgDomainPrefix, orgId, obj.ID)

	pps := make([]*model.PermissionPolicy, len(rps))

	for idx, rp := range rps {

		oup, err := client.OrgUnitPosition.Get(ctx, rp.RoleID)
		if err != nil {
			return nil, apierror.Transform(ctx, err)
		}

		pps[idx] = &model.PermissionPolicy{
			OrgPosition: oup,
			Action:      model.PermissionAction(rp.Action),
			Effect:      model.PermissionEffect(rp.Effect),
		}
	}

	return pps, nil
}

// PermissionMutation returns generated.PermissionMutationResolver implementation.
func (r *Resolver) PermissionMutation() generated.PermissionMutationResolver {
	return &permissionMutationResolver{r}
}

// PermissionQuery returns generated.PermissionQueryResolver implementation.
func (r *Resolver) PermissionQuery() generated.PermissionQueryResolver {
	return &permissionQueryResolver{r}
}

// Resource returns generated.ResourceResolver implementation.
func (r *Resolver) Resource() generated.ResourceResolver { return &resourceResolver{r} }

type permissionMutationResolver struct{ *Resolver }
type permissionQueryResolver struct{ *Resolver }
type resourceResolver struct{ *Resolver }
