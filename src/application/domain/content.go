package domain

func NewTextContent(name *string, text string) TextContent {
	return TextContent{
		Name: name,
		Text: text,
	}
}

type TextContent struct {
	Name *string
	Text string
}

func NewImageContent(name *string, path string) ImageContent {
	return ImageContent{
		Name: name,
		Path: path,
	}
}

type ImageContent struct {
	Name *string
	Path string
}
