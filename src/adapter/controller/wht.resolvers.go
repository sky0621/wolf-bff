package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/wolf-bff/src/adapter/controller/graphqlmodel"
)

func (r *mutationResolver) CreateWhtText(ctx context.Context, content graphqlmodel.WhtTextInput) (*graphqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateWhtImage(ctx context.Context, content graphqlmodel.WhtImageInput) (*graphqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FindWht(ctx context.Context, condition *graphqlmodel.WhtConditionInput) ([]graphqlmodel.Wht, error) {
	panic(fmt.Errorf("not implemented"))
}
