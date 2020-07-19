package domain

import "time"

func NewWht(id int64, recordDate time.Time, title string) *Wht {
	t := Title(title)
	return &Wht{
		ID:         ID(id),
		RecordDate: RecordDate(recordDate),
		Title:      &t,
	}
}

// 今日こと
type Wht struct {
	// ID
	ID ID
	// 記録日
	RecordDate RecordDate
	// タイトル
	Title *Title
}

// 「今日こと」検索条件
type WhtCondition struct {
	// ID
	ID *ID
}

// ID
type ID int64

func (v *ID) Validate() bool {
	if v == nil {
		return false
	}
	if *v == 0 {
		return false
	}
	return true
}

func (v *ID) ToPrimitive() int64 {
	if v == nil {
		return 0
	}
	return int64(*v)
}

// 記録日
type RecordDate time.Time

func (v *RecordDate) Validate() bool {
	if v == nil {
		return false
	}
	return true
}

// タイトル
type Title string

func (v *Title) Validate() bool {
	if v == nil {
		return false
	}
	if *v == "" {
		return false
	}
	return true
}
