package state

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/notzree/automaticv2/v2/pkg/agent/message"
)

type InMemoryState struct {
	MessageStore map[uuid.UUID]message.Message
	Loaded       *message.Message
}

func NewInMemoryState() *InMemoryState {

	return &InMemoryState{
		MessageStore: make(map[uuid.UUID]message.Message),
		Loaded:       nil,
	}
}

func (s *InMemoryState) Load(sessionID uuid.UUID) (*AgentState, error) {
	rootMsg, exist := s.MessageStore[sessionID]
	if !exist {
		return nil, fmt.Errorf("couldn't find sessionID: %v in storage", sessionID)
	}
	s.Loaded = &rootMsg
	// TODO (richard): implement this
	return nil, nil
}
