package domain

// コンテンツ
type Content interface {
	IsContent()
}

func NewTextContent(name *string, text string) Content {
	return &TextContent{
		ContentType: ContentTypeText,
		Name:        name,
		Text:        text,
	}
}

type TextContent struct {
	ContentType ContentType
	Name        *string
	Text        string
}

func (c *TextContent) IsContent() {}

func NewImageContent(name *string, path string) Content {
	return &ImageContent{
		ContentType: ContentTypeImage,
		Name:        name,
		Path:        path,
	}
}

type ImageContent struct {
	ContentType ContentType
	Name        *string
	Path        string
}

func (c *ImageContent) IsContent() {}

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
)
