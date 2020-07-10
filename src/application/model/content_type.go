package model

import (
	"fmt"
	"io"
	"strconv"
)

// コンテンツタイプ
type ContentType string

const (
	// テキスト
	ContentTypeText ContentType = "Text"
	// 画像
	ContentTypeImage ContentType = "Image"
	// 音声
	ContentTypeVoice ContentType = "Voice"
	// 動画
	ContentTypeMovie ContentType = "Movie"
	// その他
	ContentTypeOther ContentType = "Other"
)

var AllContentType = []ContentType{
	ContentTypeText,
	ContentTypeImage,
	ContentTypeVoice,
	ContentTypeMovie,
	ContentTypeOther,
}

func (e ContentType) IsValid() bool {
	switch e {
	case ContentTypeText, ContentTypeImage, ContentTypeVoice, ContentTypeMovie, ContentTypeOther:
		return true
	}
	return false
}

func (e ContentType) String() string {
	return string(e)
}

func (e *ContentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContentType", str)
	}
	return nil
}

func (e ContentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
