package agentcommands

import (
	"github.com/firesworder/password_saver/internal/agent/agentcommands/agentstate"
	"github.com/firesworder/password_saver/internal/agent/agentcommands/grpcagent"
	"github.com/firesworder/password_saver/internal/agent/agentreader"
	"github.com/firesworder/password_saver/internal/agent/agentwriter"
)

type AgentCommands struct {
	state        *agentstate.State
	grpcAgent    grpcagent.IGRPCAgent
	reader       *agentreader.AgentReader
	writer       *agentwriter.AgentWriter
	isAuthorized bool
}

func NewAgentCommands(
	grpcAgent grpcagent.IGRPCAgent, reader *agentreader.AgentReader, writer *agentwriter.AgentWriter,
) *AgentCommands {
	return &AgentCommands{grpcAgent: grpcAgent, reader: reader, writer: writer, state: agentstate.NewState()}
}
