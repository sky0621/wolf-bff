package gateway

import (
	"context"

	"github.com/sky0621/wolf-bff/src/application"

	"github.com/jmoiron/sqlx"
	"github.com/sky0621/wolf-bff/src/application/model"
)

func NewWhtRepository(db *sqlx.DB) application.WhtRepository {
	return &whtRepository{db: db}
}

type whtRepository struct {
	db *sqlx.DB
}

func (r *whtRepository) CreateWht(ctx context.Context, in model.WhtInput) (string, error) {
	// FIXME:
	//mdl := sqlboilermodel.WHT{
	//	Recorddate: time.Now(),
	//	//Title:      null.StringFromPtr(in.Title),
	//}
	//mdl.Insert(ctx, r.db, boil.Infer())
	return "", nil
}
