package agent

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/firesworder/password_saver/internal/agent/agentreader"
	"github.com/firesworder/password_saver/internal/agent/agentwriter"
	"github.com/firesworder/password_saver/internal/mocks"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func createMockAgent(t *testing.T) (*Agent, *bytes.Buffer, *bytes.Buffer) {
	mockGA := &mocks.GrpcAgent{}
	agent := NewAgent(mockGA)

	bufR := bytes.NewBufferString("")
	agent.reader = agentreader.NewAgentReader(bufio.NewReader(bufR))
	bufW := bytes.NewBufferString("")
	agent.writer = agentwriter.NewAgentWriter(bufio.NewWriter(bufW))
	agent.commands = &mocks.AgentCommands{}
	return agent, bufR, bufW
}

func TestNewAgent(t *testing.T) {
	mockGA := &mocks.GrpcAgent{}
	agent := NewAgent(mockGA)
	assert.NotEmpty(t, agent)
}

func TestAgent_Serve(t *testing.T) {
	agent, bufR, bufW := createMockAgent(t)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	bufR.WriteString("some_command\n")
	agent.Serve(ctx)
	hintMsg := "Enter command\n"
	assert.Equal(t, hintMsg+"err: unknown command\n"+hintMsg, bufW.String())
}

func TestAgent_controller(t *testing.T) {
	agent, bufR, bufW := createMockAgent(t)
	mac := agent.commands.(*mocks.AgentCommands)

	tests := []struct {
		name       string
		command    string
		input      string
		wantFC     string
		wantOutput string
	}{
		{
			name:       "Test 1. Register user.",
			command:    "register_user",
			input:      "",
			wantFC:     "RegisterUser",
			wantOutput: "",
		},
		{
			name:       "Test 2. Login user.",
			command:    "login_user",
			input:      "",
			wantFC:     "LoginUser",
			wantOutput: "",
		},
		{
			name:       "Test 3. Create text record.",
			command:    "create_record",
			input:      "text\n",
			wantFC:     "CreateTextData",
			wantOutput: fmt.Sprintf("%s\n", enterDT),
		},
		{
			name:       "Test 4. Create bank record.",
			command:    "create_record",
			input:      "bank\n",
			wantFC:     "CreateBankData",
			wantOutput: fmt.Sprintf("%s\n", enterDT),
		},
		{
			name:       "Test 5. Create binary record.",
			command:    "create_record",
			input:      "binary\n",
			wantFC:     "CreateBinaryData",
			wantOutput: fmt.Sprintf("%s\n", enterDT),
		},
		{
			name:       "Test 6. Open text record.",
			command:    "open_record",
			input:      "10 text\n",
			wantFC:     "OpenTextData",
			wantOutput: fmt.Sprintf("%s\n", enterIDaDT),
		},
		{
			name:       "Test 7. Open bank record.",
			command:    "open_record",
			input:      "10 bank\n",
			wantFC:     "OpenBankData",
			wantOutput: fmt.Sprintf("%s\n", enterIDaDT),
		},
		{
			name:       "Test 8. Open binary record.",
			command:    "open_record",
			input:      "10 binary\n",
			wantFC:     "OpenBinaryData",
			wantOutput: fmt.Sprintf("%s\n", enterIDaDT),
		},
		{
			name:       "Test 9. Update text record.",
			command:    "update_record",
			input:      "10 text\n",
			wantFC:     "UpdateTextData",
			wantOutput: fmt.Sprintf("%s\n", enterIDaDT),
		},
		{
			name:       "Test 10. Update bank record.",
			command:    "update_record",
			input:      "10 bank\n",
			wantFC:     "UpdateBankData",
			wantOutput: fmt.Sprintf("%s\n", enterIDaDT),
		},
		{
			name:       "Test 11. Update binary record.",
			command:    "update_record",
			input:      "10 binary\n",
			wantFC:     "UpdateBinaryData",
			wantOutput: fmt.Sprintf("%s\n", enterIDaDT),
		},
		{
			name:       "Test 12. Delete text record.",
			command:    "delete_record",
			input:      "10 text\n",
			wantFC:     "DeleteTextData",
			wantOutput: fmt.Sprintf("%s\n", enterIDaDT),
		},
		{
			name:       "Test 13. Delete bank record.",
			command:    "delete_record",
			input:      "10 bank\n",
			wantFC:     "DeleteBankData",
			wantOutput: fmt.Sprintf("%s\n", enterIDaDT),
		},
		{
			name:       "Test 14. Delete binary record.",
			command:    "delete_record",
			input:      "10 binary\n",
			wantFC:     "DeleteBinaryData",
			wantOutput: fmt.Sprintf("%s\n", enterIDaDT),
		},
		{
			name:       "Test 15. Show all records.",
			command:    "show_all_records",
			input:      "",
			wantFC:     "ShowAllRecords",
			wantOutput: "",
		},
		{
			name:    "Test 16. Help.",
			command: "help",
			input:   "",
			wantFC:  "",
			wantOutput: "Commands:\nAuth methods:\n- register_user, login_user\n\n" +
				"User data methods(required auth!):\n- create_record, open_record, update_record, delete_record\n" +
				"- show_all_records\n\n",
		},
		{
			name:       "Test 17. Unknown command.",
			command:    "some",
			wantFC:     "",
			wantOutput: "err: unknown command\n",
		},

		// calls with errors
		// create_record
		{
			name:       "Test 18. Create record. EOF input.",
			command:    "create_record",
			input:      "text",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterDT, io.EOF),
		},
		{
			name:       "Test 19. Create record. Unknown datatype",
			command:    "create_record",
			input:      "ddd\n",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterDT, "unknown data type"),
		},
		// open_record
		{
			name:       "Test 20. Open record. Input only ID, without DT.",
			command:    "open_record",
			input:      "10\n",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterIDaDT, "input error"),
		},
		{
			name:       "Test 21. Open record. Input id not number.",
			command:    "open_record",
			input:      "id text\n",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterIDaDT, "strconv.Atoi: parsing \"id\": invalid syntax"),
		},
		{
			name:       "Test 22. Open record. Input unknown DT.",
			command:    "open_record",
			input:      "10 someType\n",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterIDaDT, "unknown data type"),
		},
		// update_record
		{
			name:       "Test 23. Update record. Input only ID, without DT.",
			command:    "update_record",
			input:      "10\n",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterIDaDT, "input error"),
		},
		{
			name:       "Test 24. Update record. Input id not number.",
			command:    "update_record",
			input:      "id text\n",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterIDaDT, "strconv.Atoi: parsing \"id\": invalid syntax"),
		},
		{
			name:       "Test 25. Update record. Input unknown DT.",
			command:    "update_record",
			input:      "10 someType\n",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterIDaDT, "unknown data type"),
		},
		// delete_record
		{
			name:       "Test 26. Delete record. Input only ID, without DT.",
			command:    "delete_record",
			input:      "10\n",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterIDaDT, "input error"),
		},
		{
			name:       "Test 27. Delete record. Input id not number.",
			command:    "delete_record",
			input:      "id text\n",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterIDaDT, "strconv.Atoi: parsing \"id\": invalid syntax"),
		},
		{
			name:       "Test 28. Delete record. Input unknown DT.",
			command:    "delete_record",
			input:      "10 someType\n",
			wantFC:     "",
			wantOutput: fmt.Sprintf("%s\nerr: %s\n", enterIDaDT, "unknown data type"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mac.LastFuncCalled = ""
			bufR.Reset()
			bufW.Reset()
			bufR.WriteString(tt.input)

			agent.controller(tt.command)
			assert.Equal(t, tt.wantFC, mac.LastFuncCalled)
			assert.Equal(t, tt.wantOutput, bufW.String())
		})
	}
}
