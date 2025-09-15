package state

import "github.com/google/uuid"

// AgentState represents the state of an agent
// including things like
// message history
// available tools
// environment variables (todo)
// permissions (todo)
type AgentState interface {
	//todo (richard): think about these methods lol
	Load(sessionID uuid.UUID) error
	Snapshot() error
}
