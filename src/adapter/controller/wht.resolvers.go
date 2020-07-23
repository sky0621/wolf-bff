package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sky0621/wolf-bff/src/adapter"
	"github.com/sky0621/wolf-bff/src/adapter/controller/gqlmodel"
	"github.com/sky0621/wolf-bff/src/adapter/gateway"
	"github.com/sky0621/wolf-bff/src/application"
	"github.com/sky0621/wolf-bff/src/application/domain"
	"github.com/sky0621/wolf-bff/src/internal"
)

// ------------------------------------------------------------------
// Mutation
// ------------------------------------------------------------------

func (r *mutationResolver) CreateWht(ctx context.Context, wht gqlmodel.WhtInput) (*gqlmodel.MutationResponse, error) {
	res, err := adapter.Tx(ctx, r.db, func(ctx context.Context, txx *sqlx.Tx) (*adapter.TxResponse, error) {
		id, err := application.NewWht(gateway.NewWhtRepository(txx), gateway.NewContentRepository(txx)).
			CreateWht(ctx, domain.Wht{RecordDate: wht.RecordDate, Title: wht.Title})
		return &adapter.TxResponse{CreatedID: id}, err
	})
	if err != nil {
		fmt.Printf("%#+v", err) // TODO: use custom logger
		return nil, err
	}
	return &gqlmodel.MutationResponse{ID: internal.FromInt64ToPStr(res.CreatedID)}, nil
}

func (r *mutationResolver) CreateTextContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.TextContentInput) (*gqlmodel.MutationResponse, error) {
	res, err := adapter.Tx(ctx, r.db, func(ctx context.Context, txx *sqlx.Tx) (*adapter.TxResponse, error) {
		var contents []domain.TextContent
		for _, in := range inputs {
			contents = append(contents, domain.NewTextContent(in.Name, in.Text))
		}
		err := application.NewWht(gateway.NewWhtRepository(txx), gateway.NewContentRepository(txx)).
			CreateTextContent(ctx, recordDate, contents)
		return &adapter.TxResponse{CreatedID: 0}, err
	})
	if err != nil {
		fmt.Printf("%#+v", err) // TODO: use custom logger
		return nil, err
	}
	return &gqlmodel.MutationResponse{ID: internal.FromInt64ToPStr(res.CreatedID)}, nil
}

func (r *mutationResolver) CreateImageContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.ImageContentInput) (*gqlmodel.MutationResponse, error) {
	// FIXME:
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateVoiceContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.VoiceContentInput) (*gqlmodel.MutationResponse, error) {
	// FIXME:
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMovieContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.MovieContentInput) (*gqlmodel.MutationResponse, error) {
	// FIXME:
	panic(fmt.Errorf("not implemented"))
}

// ------------------------------------------------------------------
// Query
// ------------------------------------------------------------------

func (r *queryResolver) FindWht(ctx context.Context, condition *gqlmodel.WhtConditionInput) ([]gqlmodel.Wht, error) {
	c := &domain.WhtCondition{}
	if condition != nil {
		id := condition.ID.DBUniqueID()
		c.ID = &id
	}

	records, err := application.NewWht(gateway.NewWhtRepository(r.db)).ReadWht(ctx, c)
	if err != nil {
		// TODO: logger
		fmt.Printf("%#+v", err)
		return nil, err
	}

	var results []gqlmodel.Wht
	for _, r := range records {
		if r.ID == nil || r.RecordDate == internal.NilTime {
			// TODO: logger
			err := errors.New("id or recordDate is nil")
			fmt.Printf("%#+v", err)
			return nil, err
		}
		results = append(results, gqlmodel.Wht{
			ID:         gqlmodel.WhtID(*r.ID),
			RecordDate: r.RecordDate,
			Title:      r.Title,
		})
	}
	return results, nil
}

func (r *whtResolver) Contents(ctx context.Context, obj *gqlmodel.Wht) ([]gqlmodel.Content, error) {
	contents, err := For(ctx).contentLoader.Load(obj.ID.DBUniqueID())
	if err != nil {
		// TODO: logger
		fmt.Printf("%#+v", err)
		return nil, err
	}
	return contents, nil
}

// Wht returns WhtResolver implementation.
func (r *Resolver) Wht() WhtResolver { return &whtResolver{r} }

type whtResolver struct{ *Resolver }
