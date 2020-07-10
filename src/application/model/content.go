package model

// コンテンツ
type Content interface {
	IsContent()
}

// テキストコンテンツ
type TextContent struct {
	ID string `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// コンテンツタイプ
	ContentType ContentType `json:"contentType"`
	// テキスト
	Text string `json:"text"`
}

func (c TextContent) IsContent() {}

// 画像コンテンツ
type ImageContent struct {
	ID string `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// コンテンツタイプ
	ContentType ContentType `json:"contentType"`
	// 画像パス
	Path string `json:"path"`
}

func (c ImageContent) IsContent() {}

// 音声コンテンツ
type VoiceContent struct {
	ID string `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// コンテンツタイプ
	ContentType ContentType `json:"contentType"`
}

func (c VoiceContent) IsContent() {}

// 動画コンテンツ
type MovieContent struct {
	ID string `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// コンテンツタイプ
	ContentType ContentType `json:"contentType"`
}

func (c MovieContent) IsContent() {}

// その他コンテンツ
type OtherContent struct {
	ID string `json:"id"`
	// コンテンツ名
	Name *string `json:"name"`
	// コンテンツタイプ
	ContentType ContentType `json:"contentType"`
}

func (c OtherContent) IsContent() {}

// ------------------------------------------------------------------
// Input
// ------------------------------------------------------------------

// コンテンツインプット
type ContentInput struct {
	// コンテンツ名
	Name *string `json:"name"`
	// コンテンツタイプ
	ContentType ContentType `json:"contentType"`
}

// 「今日こと」テキストインプット
type WhtTextContentInput struct {
	// タイトル
	Title *string `json:"title"`
	// コンテンツ名
	Name *string `json:"name"`
	// テキスト
	Text string `json:"text"`
}

func (i WhtTextContentInput) IsContent() {}

// 「今日こと」画像インプット
type WhtImageContentInput struct {
	// タイトル
	Title *string `json:"title"`
	// コンテンツ名
	Name *string `json:"name"`
	// 画像パス
	Path string `json:"path"`
}

func (i WhtImageContentInput) IsContent() {}

// 「今日こと」音声インプット
type WhtVoiceContentInput struct {
	// タイトル
	Title *string `json:"title"`
	// コンテンツ名
	Name *string `json:"name"`
}

func (i WhtVoiceContentInput) IsContent() {}

// 「今日こと」動画インプット
type WhtMovieContentInput struct {
	// タイトル
	Title *string `json:"title"`
	// コンテンツ名
	Name *string `json:"name"`
}

func (i WhtMovieContentInput) IsContent() {}

// 「今日こと」その他インプット
type WhtOtherContentInput struct {
	// タイトル
	Title *string `json:"title"`
	// コンテンツ名
	Name *string `json:"name"`
}

func (i WhtOtherContentInput) IsContent() {}
