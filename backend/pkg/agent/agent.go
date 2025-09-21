package agent

import "github.com/notzree/automaticv2/v2/pkg/agent/state"

type Agent struct {
	// required behaviour
	// ability to store / retrieve context (?)(semantic/vector search, fulltext search (turbopuffer))
	// ability to execute MCP tools
	// ability to recall its tool execution history (part of retrieving memory)

	// NEW STUFF:
	// Ability to create workflows
	// Detect if user intends to execute workflow ->  Ability to execute workflows

	// TODO list stuff for v0
	// need: Chat memory
	// need: MCP tool execution
	// need: Concept of workflow to the agent (may be difficult to do with libraries like genkit...?)
	// genkit + mcp + custom tools + workflow tool?
	// concept of turns
	// user -> http server + sessionID -> Fetch history -> generate next turn -> save turn -> return turn output
	context AgentContext
	state   state.AgentStateManager
}

type AgentContext struct {
}

func NewAgent(aCtx AgentContext, aState state.AgentStateManager) *Agent {
	return &Agent{
		context: aCtx,
		state:   aState,
	}
}

func (a *Agent) LoadFromSession(sessionID string) error {
	// err := a.state.LoadStateFromSession(sessionID)
	// if err != nil {
	// 	return err
	// }
	return nil
}
