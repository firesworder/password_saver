package agentreader

import (
	"bufio"
	"fmt"
	"strings"
)

type AgentReader struct {
	reader *bufio.Reader
}

func NewAgentReader(reader *bufio.Reader) *AgentReader {
	return &AgentReader{reader: reader}
}

func (ar *AgentReader) GetUserFields() ([]string, error) {
	input, err := ar.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	fields := strings.Fields(input)
	if len(fields) < 1 {
		return nil, fmt.Errorf("input error")
	}
	return fields, nil
}

func (ar *AgentReader) GetUserInput() (string, error) {
	input, err := ar.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}
