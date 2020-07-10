package model

// 今日こと
type Wht struct {
	// ID
	ID string `json:"id"`
	// 記録日
	RecordDate string `json:"recordDate"`
	// タイトル
	Title *string `json:"title"`
}

func (w Wht) IsNode() {}

// ------------------------------------------------------------------
// Input
// ------------------------------------------------------------------

// 今日ことインプット
type WhtInput struct {
	// タイトル
	Title *string `json:"title"`
	// コンテンツリスト
	Contents []ContentInput `json:"contents"`
}

// 「今日こと」検索条件
type WhtConditionInput struct {
	// ID
	ID *string `json:"id"`
}
