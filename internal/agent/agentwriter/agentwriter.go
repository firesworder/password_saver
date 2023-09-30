// Package agentwriter обертка над bufio.Writer для избежания дублирования функций записи в буфер writer.
package agentwriter

import (
	"bufio"
	"log"
)

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
	var err error
	if _, err = a.writer.WriteString(str + "\n"); err != nil {
		log.Fatal(err)
	}
	if err = a.writer.Flush(); err != nil {
		log.Fatal(err)
	}
}

// WriteErrorString записывает строку в буфер добавляя префикс "err: " + перенос строки,
// а затем сохраняет изменения в буфере.
func (a *AgentWriter) WriteErrorString(errStr string) {
	var err error
	if _, err = a.writer.WriteString("err: " + errStr + "\n"); err != nil {
		log.Fatal(err)
	}
	if err = a.writer.Flush(); err != nil {
		log.Fatal(err)
	}
}
