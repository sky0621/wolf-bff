package gateway

import (
	"context"

	"github.com/volatiletech/null/v8"

	"golang.org/x/xerrors"

	"github.com/sky0621/wolf-bff/src/adapter/gateway/sqlboilermodel"

	"github.com/sky0621/wolf-bff/src/application"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/sky0621/wolf-bff/src/application/domain"
)

func NewContentRepository(db boil.ContextExecutor) application.ContentRepository {
	return &contentRepository{db: db}
}

type contentRepository struct {
	db boil.ContextExecutor
}

func (r *contentRepository) CreateTextContents(ctx context.Context, whtID int64, inputs []domain.TextContent) error {
	// TODO: バッチ形式を検討！
	for _, in := range inputs {
		mdl := sqlboilermodel.ContentText{
			WHTID: whtID,
			Name:  null.StringFromPtr(in.Name),
			Text:  in.Text,
		}
		if err := mdl.Insert(ctx, r.db, boil.Infer()); err != nil {
			return xerrors.Errorf("failed to insert content_text[whtID:%d][in:%#+v]: %w", whtID, in, err)
		}
	}
	return nil
}
