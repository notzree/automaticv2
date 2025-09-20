package block

type TextBlock struct {
	content string
}

func NewTextBlock(content string) TextBlock {
	return TextBlock{
		content: content,
	}
}

func (tb TextBlock) GetDisplayContent() (string, error) {
	return tb.content, nil
}
