package gqlmodel

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
