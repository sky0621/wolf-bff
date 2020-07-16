package gqlmodel

import "github.com/sky0621/wolf-bff/src/application/model"

// 今日こと
type Wht struct {
	// ID
	ID string `json:"id"`
	// 記録日
	RecordDate string `json:"recordDate"`
	// タイトル
	Title *string `json:"title"`
	// コンテンツリスト
	//Contents []Content `json:"contents"`
}

func (Wht) IsNode() {}

// 今日ことインプット
type WhtInput struct {
	// タイトル
	Title *string `json:"title"`
}

func (i *WhtInput) ToModel() model.WhtInput {
	return model.WhtInput{
		Title: i.Title,
	}
}
