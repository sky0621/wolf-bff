package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/wolf-bff/src/adapter/controller/graphqlmodel"
)

func (r *mutationResolver) CreateWhtText(ctx context.Context, content graphqlmodel.WhtTextInput) (*graphqlmodel.MutationResponse, error) {
	r.wht.CreateWht(ctx)
	// FIXME:
	return nil, nil
}

func (r *mutationResolver) CreateWhtImage(ctx context.Context, content graphqlmodel.WhtImageInput) (*graphqlmodel.MutationResponse, error) {
	r.wht.CreateWht(ctx)
	// FIXME:
	return nil, nil
}

func (r *queryResolver) FindWht(ctx context.Context, condition *graphqlmodel.WhtConditionInput) ([]graphqlmodel.Wht, error) {
	r.wht.FindWht(ctx)
	// FIXME:
	return nil, nil
}

func (r *whtResolver) Contents(ctx context.Context, obj *graphqlmodel.Wht) ([]graphqlmodel.Content, error) {
	// FIXME:
	panic(fmt.Errorf("not implemented"))
}

// Wht returns WhtResolver implementation.
func (r *Resolver) Wht() WhtResolver { return &whtResolver{r} }

type whtResolver struct{ *Resolver }
