package application

import (
	"context"

	"github.com/sky0621/wolf-bff/src/application/domain"
)

func NewContent(repository ContentRepository) *Content {
	return &Content{repository: repository}
}

type Content struct {
	repository ContentRepository
}

func (w Content) CreateContent(ctx context.Context, in domain.Content) error {
	return w.repository.CreateContent(ctx, in)
}

type ContentRepository interface {
	CreateContent(ctx context.Context, in domain.Content) error
}
