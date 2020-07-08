package graphqlmodel

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
