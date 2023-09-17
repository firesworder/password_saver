package agentwriter

import "bufio"

type AgentWriter struct {
	writer *bufio.Writer
}

func NewAgentWriter(writer *bufio.Writer) *AgentWriter {
	return &AgentWriter{writer: writer}
}

func (a *AgentWriter) WriteString(str string) {
	a.writer.WriteString(str + "\n")
	a.writer.Flush()
}

func (a *AgentWriter) WriteErrorString(errStr string) {
	a.writer.WriteString("err: " + errStr + "\n")
	a.writer.Flush()
}
