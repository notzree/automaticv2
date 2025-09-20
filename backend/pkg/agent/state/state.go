package state

import (
	"github.com/google/uuid"
	"github.com/notzree/automaticv2/v2/pkg/agent/message"
)

// AgentStateManager represents the state of an agent
// including things like
// message history
// available tools
// environment variables (todo)
// permissions (todo)
type AgentStateManager interface {
	//todo (richard): think about these methods lol
	Load(sessionID uuid.UUID) (*AgentState, error)
	Snapshot(AgentState) error
}

type AgentState struct {
	SessionID   uuid.UUID
	RootMessage message.Message
	// TODO (richard): create required structs
	ConnectedTools []string
	// TODO (richard): Should we be fetching environment variables each time we fetch state (security?)
	EnvironmentVariables map[string]string
	Permissions          []string
}
