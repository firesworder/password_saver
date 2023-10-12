package agentcommands

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/firesworder/password_saver/internal/agent/agentreader"
	"github.com/firesworder/password_saver/internal/agent/agentwriter"
	"github.com/firesworder/password_saver/internal/mocks"
)

func createMockAgentCommands(t *testing.T) (*AgentCommands, *bytes.Buffer, *bytes.Buffer) {
	mockGA := mocks.NewGRPCAgent()

	bufR := bytes.NewBufferString("")
	reader := agentreader.NewAgentReader(bufio.NewReader(bufR))
	bufW := bytes.NewBufferString("")
	writer := agentwriter.NewAgentWriter(bufio.NewWriter(bufW))
	agentCommands := NewAgentCommands(mockGA, reader, writer)
	return agentCommands, bufR, bufW
}

func TestNewAgentCommands(t *testing.T) {
	buf := bytes.NewBufferString("")
	reader := agentreader.NewAgentReader(bufio.NewReader(buf))
	writer := agentwriter.NewAgentWriter(bufio.NewWriter(buf))
	ga := mocks.NewGRPCAgent()

	ac := NewAgentCommands(ga, reader, writer)
	assert.NotEmpty(t, ac.state)
	assert.NotEmpty(t, ac.grpcAgent)
	assert.NotEmpty(t, ac.reader)
	assert.NotEmpty(t, ac.writer)
	assert.Equal(t, false, ac.isAuthorized)
}
