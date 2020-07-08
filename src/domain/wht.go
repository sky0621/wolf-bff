package domain

import "context"

type Wht struct {
	ID int64
}

type WhtContent interface {
	IsWht() bool
}

type WhtTextContent struct {
}

type WhtLogic interface {
	FindWht(ctx context.Context) ([]Wht, error)
	CreateWht(ctx context.Context) (Wht, error)
}

type whtLogic struct {
}

//func (w *whtLogic)
