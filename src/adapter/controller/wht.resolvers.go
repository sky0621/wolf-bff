package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/wolf-bff/src/adapter/controller/graphqlmodel"
	"github.com/sky0621/wolf-bff/src/adapter/gateway"
	"github.com/sky0621/wolf-bff/src/application"
	"github.com/sky0621/wolf-bff/src/application/model"
	"golang.org/x/xerrors"
)

func (r *mutationResolver) CreateWht(ctx context.Context, wht model.WhtInput) (*graphqlmodel.MutationResponse, error) {
	id, err := application.NewWht(gateway.NewWhtRepository(r.db)).CreateWht(ctx, wht)
	if err != nil {
		return nil, xerrors.Errorf("failed to CreateWht: %w", err)
	}
	return &graphqlmodel.MutationResponse{
		ID: &id,
	}, nil
}

func (r *queryResolver) FindWht(ctx context.Context, condition *model.WhtConditionInput) ([]model.Wht, error) {
	// FIXME:
	panic(fmt.Errorf("not implemented"))
}

func (r *whtResolver) Contents(ctx context.Context, obj *model.Wht) ([]model.Content, error) {
	// FIXME:
	panic(fmt.Errorf("not implemented"))
}

// Wht returns WhtResolver implementation.
func (r *Resolver) Wht() WhtResolver { return &whtResolver{r} }

type whtResolver struct{ *Resolver }
