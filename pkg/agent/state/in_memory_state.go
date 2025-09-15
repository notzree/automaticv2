package state

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/notzree/automaticv2/v2/pkg/agent/message"
)

type InMemoryState struct {
	Store  map[uuid.UUID]message.Message
	Loaded *message.Message
}

func NewInMemoryState() *InMemoryState {

	return &InMemoryState{
		Store:  make(map[uuid.UUID]message.Message),
		Loaded: nil,
	}
}

func (s *InMemoryState) Load(sessionID uuid.UUID) error {
	rootMsg, exist := s.Store[sessionID]
	if !exist {
		return fmt.Errorf("couldn't find sessionID: %v in storage", sessionID)
	}
	s.Loaded = &rootMsg
	return nil
}
