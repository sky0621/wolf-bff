package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sky0621/wolf-bff/src/graph/model"
)

func (r *mutationResolver) CreateWhtText(ctx context.Context, content model.WhtTextInput) (*model.MutationResponse, error) {
	// FIXME:
	fmt.Println("CreateWhtText")
	fmt.Println(content)
	ret := "9abd8293-dc5c-493d-85f2-e81fb86efd07"
	return &model.MutationResponse{
		ID: &ret,
	}, nil
}

func (r *mutationResolver) CreateWhtImage(ctx context.Context, content model.WhtImageInput) (*model.MutationResponse, error) {
	// FIXME:
	fmt.Println("CreateWhtImage")
	fmt.Println(content)
	ret := "8a4013bd-e802-4398-bc0d-40d3b1e4b5da"
	return &model.MutationResponse{
		ID: &ret,
	}, nil
}

func (r *queryResolver) FindWht(ctx context.Context, condition *model.WhtConditionInput) ([]*model.Wht, error) {
	// FIXME:
	fmt.Println("FindWht")
	fmt.Println(condition)

	tt := "「今日こと」のタイトル"
	nm := "コンテンツの名前"
	return []*model.Wht{
		{
			ID:          "d1780890-b103-4edb-88a5-39b0b30ba2b1",
			RecordDate:  "2020-07-04T15:04:05Z09:00",
			Title:       &tt,
			ContentType: model.ContentTypeText,
			Content: &model.TextContent{
				Name: &nm,
				Text: "今日は、遠足に行きました。",
			},
		},
		{
			ID:          "d817a009-5691-4552-bc9c-7611f3d68032",
			RecordDate:  "2020-07-04T20:16:22Z09:00",
			Title:       &tt,
			ContentType: model.ContentTypeImage,
			Content: &model.ImageContent{
				Name: &nm,
				Path: "https://~~~~/01.png",
			},
		},
	}, nil
}
