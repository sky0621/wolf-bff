package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/wolf-bff/src/adapter/controller"
	"github.com/sky0621/wolf-bff/src/adapter/controller/gqlmodel"
)

func (r *mutationResolver) CreateWht(ctx context.Context, wht controller.WhtInput) (*controller.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTextContents(ctx context.Context, inputs []controller.TextContentInput) (*controller.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateImageContents(ctx context.Context, inputs []controller.ImageContentInput) (*controller.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateVoiceContents(ctx context.Context, inputs []controller.VoiceContentInput) (*controller.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMovieContents(ctx context.Context, inputs []controller.MovieContentInput) (*controller.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FindWht(ctx context.Context, condition *controller.WhtConditionInput) ([]gqlmodel.Wht, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *whtResolver) Contents(ctx context.Context, obj *gqlmodel.Wht) ([]gqlmodel.Content, error) {
	panic(fmt.Errorf("not implemented"))
}

// Wht returns controller.WhtResolver implementation.
func (r *Resolver) Wht() controller.WhtResolver { return &whtResolver{r} }

type whtResolver struct{ *Resolver }
