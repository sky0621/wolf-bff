package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/sky0621/wolf-bff/src/internal"

	"github.com/sky0621/wolf-bff/src/application/domain"

	"github.com/jmoiron/sqlx"
	"github.com/sky0621/wolf-bff/src/adapter"
	"github.com/sky0621/wolf-bff/src/adapter/controller/gqlmodel"
	"github.com/sky0621/wolf-bff/src/adapter/gateway"
	"github.com/sky0621/wolf-bff/src/application"
	"golang.org/x/xerrors"
)

// ------------------------------------------------------------------
// Mutation
// ------------------------------------------------------------------

func (r *mutationResolver) CreateWht(ctx context.Context, wht gqlmodel.WhtInput) (*gqlmodel.MutationResponse, error) {
	res, err := adapter.Tx(ctx, r.db, func(ctx context.Context, txx *sqlx.Tx) (*adapter.TxResponse, error) {
		id, err := application.NewWhtOnly(gateway.NewWhtRepository(txx)).CreateWht(ctx, wht.ToModel())
		if err != nil {
			return nil, xerrors.Errorf("failed to CreateWht: %w", err)
		}
		return &adapter.TxResponse{CreatedID: id}, nil
	})
	if err != nil {
		return nil, xerrors.Errorf("failed to Tx: %w", err)
	}
	id := strconv.Itoa(int(res.CreatedID))
	return &gqlmodel.MutationResponse{ID: &id}, nil
}

func (r *mutationResolver) CreateTextContents(ctx context.Context, inputs []gqlmodel.TextContentInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateImageContents(ctx context.Context, inputs []gqlmodel.ImageContentInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateVoiceContents(ctx context.Context, inputs []gqlmodel.VoiceContentInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMovieContents(ctx context.Context, inputs []gqlmodel.MovieContentInput) (*gqlmodel.MutationResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// ------------------------------------------------------------------
// Query
// ------------------------------------------------------------------

func (r *queryResolver) FindWht(ctx context.Context, condition *gqlmodel.WhtConditionInput) ([]gqlmodel.Wht, error) {
	c := &domain.WhtCondition{}
	if condition != nil {
		id := domain.ID(condition.ID.DBUniqueID())
		c.ID = &id
	}
	records, err := application.NewWhtOnly(gateway.NewWhtRepository(r.db)).ReadWht(ctx, c)
	if err != nil {
		return nil, xerrors.Errorf("failed to ReadWht: %w", err)
	}
	var results []gqlmodel.Wht
	for _, r := range records {
		results = append(results, gqlmodel.Wht{
			ID:         gqlmodel.WhtID(r.ID),
			RecordDate: time.Time(r.RecordDate),
			Title:      internal.PStr(r.Title.String()),
		})
	}
	return results, nil
}

func (r *whtResolver) Contents(ctx context.Context, obj *gqlmodel.Wht) ([]gqlmodel.Content, error) {
	contents, err := For(ctx).contentLoader.Load(obj.ID.DBUniqueID())
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
