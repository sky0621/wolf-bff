package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/wolf-bff/src/graph/model"
)

func (r *mutationResolver) CreateWhtText(ctx context.Context, content model.WhtTextInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateWhtImage(ctx context.Context, content model.WhtImageInput) (*model.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FindWht(ctx context.Context, condition *model.WhtConditionInput) ([]*model.Wht, error) {
	panic(fmt.Errorf("not implemented"))
}
