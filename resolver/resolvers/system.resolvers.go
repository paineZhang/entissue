package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"wing/models/ent"
	apierror "wing/resolver/error"
	"wing/resolver/generated"
	"wing/resolver/middleware"
)

func (r *mutationResolver) SystemMutation(ctx context.Context) (*ent.SystemMutation, error) {
	return &ent.SystemMutation{}, middleware.CheckAuthInfo(ctx, ent.FromContext(ctx))
}

func (r *queryResolver) SystemQuery(ctx context.Context) (*ent.SystemQuery, error) {
	return &ent.SystemQuery{}, nil
}

func (r *systemMutationResolver) SetName(ctx context.Context, obj *ent.SystemMutation, name string) (bool, error) {
	txModel := ent.FromContext(ctx)
	err := txModel.System.Update().SetName(name).Exec(ctx)
	if err != nil {
		return false, apierror.Transform(ctx, err)
	}
	return true, nil
}

func (r *systemQueryResolver) System(ctx context.Context, obj *ent.SystemQuery) (*ent.System, error) {
	system, err := ent.FromContext(ctx).System.Query().First(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}
	return system, nil
}

// SystemMutation returns generated.SystemMutationResolver implementation.
func (r *Resolver) SystemMutation() generated.SystemMutationResolver {
	return &systemMutationResolver{r}
}

// SystemQuery returns generated.SystemQueryResolver implementation.
func (r *Resolver) SystemQuery() generated.SystemQueryResolver { return &systemQueryResolver{r} }

type systemMutationResolver struct{ *Resolver }
type systemQueryResolver struct{ *Resolver }
