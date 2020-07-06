package gateway

import (
	"context"

	"github.com/sky0621/wolf-bff/src/domain"
	"github.com/sky0621/wolf-bff/src/driver"
)

func NewWht(db driver.RDB) domain.Wht {
	return &wht{db: db}
}

type wht struct {
	db driver.RDB
}

func (w *wht) FindWht(ctx context.Context) {
	// FIXME:
}

func (w *wht) CreateWht(ctx context.Context) {
	// FIXME:
	//in := mdl.WHT{}
	//in.Insert()
}
