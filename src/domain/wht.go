package domain

import "context"

type Wht interface {
	FindWht(ctx context.Context)
	CreateWht(ctx context.Context)
}
