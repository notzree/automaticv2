package block

// actual content of a message (whether its text, image, tool call, tool result, etc)
type Block interface {
	GetDisplayContent() (string, error)
}
