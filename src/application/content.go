package application

import (
	"context"

	"github.com/sky0621/wolf-bff/src/application/model"
)

type Content struct {
	ContentRepository ContentRepository
}

func (w Content) CreateContent(ctx context.Context, in model.Content) error {
	return w.ContentRepository.CreateContent(ctx, in)
}

type ContentRepository interface {
	CreateContent(ctx context.Context, in model.Content) error
}
