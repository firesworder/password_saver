// Package agentreader обертка над bufio.Reader для избежания дублирования функций чтения из буфера reader.
package agentreader

import (
	"bufio"
	"fmt"
	"strings"
)

// AgentReader тип-обертка над bufio.Reader, для добавления повт. функций.
type AgentReader struct {
	reader *bufio.Reader
}

// NewAgentReader создает экземпляр AgentReader на основе переданного reader(без доп. логики внутри)
func NewAgentReader(reader *bufio.Reader) *AgentReader {
	return &AgentReader{reader: reader}
}

// GetUserFields считывает строку из буфера(до первой встречи символа \n), удаляет пробелы перед и после символов
// и разбивает строку на подстроки по разделителю пробела(т.е. на слова).
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

// GetUserInput считывает строку из буфера(до первой встречи символа \n) и удаляет пробелы перед и после символов.
func (ar *AgentReader) GetUserInput() (string, error) {
	input, err := ar.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}
