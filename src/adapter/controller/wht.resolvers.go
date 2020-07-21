package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

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
		id, err := application.NewWht(gateway.NewWhtRepository(txx)).CreateWht(ctx, domain.Wht{
			RecordDate: &wht.RecordDate,
			Title:      wht.Title,
		})
		if err != nil {
			return nil, xerrors.Errorf("failed to CreateWht: %w", err)
		}
		return &adapter.TxResponse{CreatedID: id}, nil
	})
	if err != nil {
		// TODO: logger
		fmt.Printf("%#+v", err)
		return nil, err
	}

	id := strconv.Itoa(int(res.CreatedID))
	return &gqlmodel.MutationResponse{ID: &id}, nil
}

func (r *mutationResolver) CreateTextContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.TextContentInput) (*gqlmodel.MutationResponse, error) {
	res, err := adapter.Tx(ctx, r.db, func(ctx context.Context, txx *sqlx.Tx) (*adapter.TxResponse, error) {
		// 該当日の「今日こと」作成済みチェック
		wht, err := application.NewWht(gateway.NewWhtRepository(txx)).GetWhtByRecordDate(ctx, recordDate)
		if err != nil {
			return nil, xerrors.Errorf("failed to GetWhtByRecordDate[recordDate:%#+v]: %w", recordDate, err)
		}

		var whtID int64
		if wht == nil {
			var err error
			// まず、該当日の分の「今日こと」を作成
			whtID, err = application.NewWht(gateway.NewWhtRepository(txx)).CreateWht(ctx, domain.Wht{
				RecordDate: &recordDate,
				Title:      wht.Title,
			})
			if err != nil {
				return nil, xerrors.Errorf("failed to CreateWht: %w", err)
			}
		} else {
			whtID = *wht.ID
		}

		// テキストコンテンツを作成
		contentApp := application.NewContent(gateway.NewContentRepository(txx))
		for _, in := range inputs {
			// TODO: バッチ形式を検討！
			if err := contentApp.CreateTextContent(ctx, whtID, domain.NewTextContent(in.Name, in.Text)); err != nil {
				return nil, xerrors.Errorf("failed to CreateTextContent[whtID:%d][in:%#+v]: %w", whtID, in, err)
			}
		}
		return nil, nil
	})
	if err != nil {
		// TODO: logger
		fmt.Printf("%#+v", err)
		return nil, err
	}

	id := strconv.Itoa(int(res.CreatedID))
	return &gqlmodel.MutationResponse{ID: &id}, nil
}

func (r *mutationResolver) CreateImageContents(ctx context.Context, recordDate time.Time, inputs []gqlmodel.ImageContentInput) (*gqlmodel.MutationResponse, error) {
	res, err := adapter.Tx(ctx, r.db, func(ctx context.Context, txx *sqlx.Tx) (*adapter.TxResponse, error) {
		// FIXME:
		return nil, nil
	})
	if err != nil {
		// TODO: logger
		fmt.Printf("%#+v", err)
		return nil, err
	}

	id := strconv.Itoa(int(res.CreatedID))
	return &gqlmodel.MutationResponse{ID: &id}, nil
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
		if r.ID == nil || r.RecordDate == nil {
			// TODO: logger
			err := errors.New("id or recordDate is nil")
			fmt.Printf("%#+v", err)
			return nil, err
		}
		results = append(results, gqlmodel.Wht{
			ID:         gqlmodel.WhtID(*r.ID),
			RecordDate: *r.RecordDate,
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
