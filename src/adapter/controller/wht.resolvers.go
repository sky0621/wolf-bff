package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/sky0621/wolf-bff/src/adapter/controller/gqlmodel"
)

func (r *mutationResolver) CreateWht(ctx context.Context, wht WhtInput) (*MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTextContents(ctx context.Context, inputs []TextContentInput) (*MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateImageContents(ctx context.Context, inputs []ImageContentInput) (*MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateVoiceContents(ctx context.Context, inputs []VoiceContentInput) (*MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMovieContents(ctx context.Context, inputs []MovieContentInput) (*MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FindWht(ctx context.Context, condition *WhtConditionInput) ([]gqlmodel.Wht, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *whtResolver) Contents(ctx context.Context, obj *gqlmodel.Wht) ([]gqlmodel.Content, error) {
	contents, err := For(ctx).ContentLoader.Load(obj.ID)
	if err != nil {
		// TODO: logger
		log.Println(err)
		return nil, err
	}
	return contents, nil
}

// Wht returns WhtResolver implementation.
func (r *Resolver) Wht() WhtResolver { return &whtResolver{r} }

type whtResolver struct{ *Resolver }
