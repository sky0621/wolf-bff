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

func (w Content) CreateTextContent(ctx context.Context, whtID int64, in domain.TextContent) error {
	return w.repository.CreateTextContent(ctx, whtID, in)
}

type ContentRepository interface {
	CreateTextContent(ctx context.Context, whtID int64, in domain.TextContent) error
}
