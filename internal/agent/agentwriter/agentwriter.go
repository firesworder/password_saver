// Package agentwriter обертка над bufio.Writer для избежания дублирования функций записи в буфер writer.
package agentwriter

import "bufio"

// AgentWriter тип-обертка над bufio.Writer, для добавления повт. функций.
type AgentWriter struct {
	writer *bufio.Writer
}

// NewAgentWriter создает экземпляр AgentWriter на основе переданного writer(без доп. логики внутри)
func NewAgentWriter(writer *bufio.Writer) *AgentWriter {
	return &AgentWriter{writer: writer}
}

// WriteString записывает строку в буфер + добавляет перенос строки, а затем сохраняет изменения в буфере.
func (a *AgentWriter) WriteString(str string) {
	a.writer.WriteString(str + "\n")
	a.writer.Flush()
}

// WriteErrorString записывает строку в буфер добавляя префикс "err: " + перенос строки,
// а затем сохраняет изменения в буфере.
func (a *AgentWriter) WriteErrorString(errStr string) {
	a.writer.WriteString("err: " + errStr + "\n")
	a.writer.Flush()
}
