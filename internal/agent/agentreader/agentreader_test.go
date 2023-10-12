package agentreader

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAgentReader(t *testing.T) {
	buf := bytes.NewBufferString("")
	reader := bufio.NewReader(buf)
	ar := NewAgentReader(reader)
	assert.NotEmpty(t, ar)
}

func TestAgentReader_GetUserFields(t *testing.T) {
	buf := bytes.NewBufferString("")
	reader := bufio.NewReader(buf)
	ar := NewAgentReader(reader)
	assert.NotEmpty(t, ar)

	tests := []struct {
		name         string
		input        string
		wantCommands []string
		wantErr      error
	}{
		{
			name:         "Test 1. Not empty input with \\n",
			input:        "dada tata\n",
			wantCommands: []string{"dada", "tata"},
			wantErr:      nil,
		},
		{
			name:         "Test 2. Empty line without \\n",
			input:        "",
			wantCommands: nil,
			wantErr:      io.EOF,
		},
		{
			name:         "Test 3. Empty line with \\n",
			input:        "\n",
			wantCommands: nil,
			wantErr:      fmt.Errorf("input error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.WriteString(tt.input)
			gotCommands, gotErr := ar.GetUserFields()
			assert.Equal(t, tt.wantErr, gotErr)
			assert.Equal(t, tt.wantCommands, gotCommands)
		})
	}
}

func TestAgentReader_GetUserInput(t *testing.T) {
	buf := bytes.NewBufferString("")
	reader := bufio.NewReader(buf)
	ar := NewAgentReader(reader)
	assert.NotEmpty(t, ar)

	tests := []struct {
		name      string
		input     string
		wantInput string
		wantErr   error
	}{
		{
			name:      "Test 1. Not empty input with \\n",
			input:     "dada tata\n",
			wantInput: "dada tata",
			wantErr:   nil,
		},
		{
			name:      "Test 2. Empty line without \\n",
			input:     "",
			wantInput: "",
			wantErr:   io.EOF,
		},
		{
			name:      "Test 3. Empty line with \\n",
			input:     "\n",
			wantInput: "",
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.WriteString(tt.input)
			gotInput, gotErr := ar.GetUserInput()
			assert.Equal(t, tt.wantErr, gotErr)
			assert.Equal(t, tt.wantInput, gotInput)
		})
	}
}
