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

func (r *jobHistoryMutationResolver) Create(ctx context.Context, obj *ent.JobHistoryMutation, input ent.CreateJobHistoryInput) (bool, error) {
	txModel := ent.FromContext(ctx)

	err := txModel.JobHistory.Create().SetInput(input).Exec(ctx)
	if err != nil {
		return false, apierror.Transform(ctx, err)
	}
	return true, nil
}

func (r *jobHistoryMutationResolver) Delete(ctx context.Context, obj *ent.JobHistoryMutation, jobHistoryID int) (bool, error) {
	txModel := ent.FromContext(ctx)

	err := txModel.JobHistory.DeleteOneID(jobHistoryID).Exec(ctx)
	if err != nil {
		return false, apierror.Transform(ctx, err)
	}
	return true, nil
}

func (r *jobHistoryQueryResolver) JobHistories(ctx context.Context, obj *ent.JobHistoryQuery, filters *ent.JobHistoryWhereInput, orders []*ent.JobHistoryOrder) (*ent.JobHistoryConnection, error) {
	// TODO 等GO 1.18出来带泛型了。把所有的获取重构一遍，逻辑是一样的。

	// 处理排序和筛选条件
	orderOptions := make([]ent.JobHistoryPaginateOption, 0)

	for _, order := range orders {
		orderOptions = append(orderOptions, ent.WithJobHistoryOrder(order))
	}
	orderOptions = append(orderOptions, ent.WithJobHistoryFilter(filters.Filter))

	cnn, err := ent.FromContext(ctx).JobHistory.Query().
		Paginate(ctx, nil, nil, nil, nil, orderOptions...)

	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	return cnn, nil
}

func (r *mutationResolver) JobHistoryMutation(ctx context.Context) (*ent.JobHistoryMutation, error) {
	return &ent.JobHistoryMutation{}, middleware.CheckAuthInfo(ctx, ent.FromContext(ctx))
}

func (r *queryResolver) JobHistoryQuery(ctx context.Context) (*ent.JobHistoryQuery, error) {
	return &ent.JobHistoryQuery{}, middleware.CheckAuthInfo(ctx, ent.FromContext(ctx))
}

// JobHistoryMutation returns generated.JobHistoryMutationResolver implementation.
func (r *Resolver) JobHistoryMutation() generated.JobHistoryMutationResolver {
	return &jobHistoryMutationResolver{r}
}

// JobHistoryQuery returns generated.JobHistoryQueryResolver implementation.
func (r *Resolver) JobHistoryQuery() generated.JobHistoryQueryResolver {
	return &jobHistoryQueryResolver{r}
}

type jobHistoryMutationResolver struct{ *Resolver }
type jobHistoryQueryResolver struct{ *Resolver }
