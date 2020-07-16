package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/wolf-bff/src/adapter/controller"
	"github.com/sky0621/wolf-bff/src/adapter/controller/gqlmodel"
)

func (r *mutationResolver) Noop(ctx context.Context, input *gqlmodel.NoopInput) (*gqlmodel.NoopPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id string) (gqlmodel.Node, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns controller.MutationResolver implementation.
func (r *Resolver) Mutation() controller.MutationResolver { return &mutationResolver{r} }

// Query returns controller.QueryResolver implementation.
func (r *Resolver) Query() controller.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
