package agentwriter

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAgentWriter(t *testing.T) {
	buf := bytes.NewBufferString("")
	writer := bufio.NewWriter(buf)
	aw := NewAgentWriter(writer)
	assert.Equal(t, writer, aw.writer)
}

func TestAgentWriter_WriteString(t *testing.T) {
	buf := bytes.NewBufferString("")
	writer := bufio.NewWriter(buf)
	aw := NewAgentWriter(writer)

	dStr := "demo_str"
	aw.WriteString(dStr)
	assert.Equal(t, dStr, strings.TrimSpace(buf.String()))
}

func TestAgentWriter_WriteErrorString(t *testing.T) {
	buf := bytes.NewBufferString("")
	writer := bufio.NewWriter(buf)
	aw := NewAgentWriter(writer)

	dErrStr := "demo_err_str"
	aw.WriteErrorString(dErrStr)
	assert.Equal(t, fmt.Sprintf("err: %s", dErrStr), strings.TrimSpace(buf.String()))
}
