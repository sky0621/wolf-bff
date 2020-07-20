package gateway

import (
	"context"

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

func (r *contentRepository) CreateContent(ctx context.Context, in domain.Content) error {
	// FIXME:
	return nil
}
