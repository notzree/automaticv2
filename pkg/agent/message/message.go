package message

import (
	"slices"

	"github.com/google/uuid"
	"github.com/notzree/automaticv2/v2/pkg/agent/block"
	types "github.com/notzree/automaticv2/v2/pkg/consts"
)

// who sent the message, where in the chat does it belong, which session does it belong to, etc.
type Message struct {
	ID            uuid.UUID
	ParentMessage *Message
	ChildMessages []*Message
	sessionID     uuid.UUID
	Owner         types.Owner
	Block         block.Block
}

func NewSystemMessage(systemBlock block.TextBlock) *Message {
	return &Message{
		ID:            uuid.Max,
		ParentMessage: nil,
		ChildMessages: make([]*Message, 0),
		sessionID:     uuid.New(),
		Owner:         types.OwnerSystem,
		Block:         systemBlock,
	}
}

func GetPathFromMessageID(root *Message, messageID uuid.UUID) ([]Message, error) {
	target := FindMessage(root, messageID)
	if target == nil {
		return nil, ErrMsgNotFound
	}

	// Build path from target to root
	path := []Message{}
	current := target

	for current != nil {
		path = append(path, *current)
		current = current.ParentMessage
	}

	slices.Reverse(path)
	return path, nil
}

func FindMessage(node *Message, targetID uuid.UUID) *Message {
	if node == nil {
		return nil
	}

	if node.ID == targetID {
		return node
	}

	for _, childPtr := range node.ChildMessages {
		if found := FindMessage(childPtr, targetID); found != nil {
			return found
		}
	}
	return nil
}
