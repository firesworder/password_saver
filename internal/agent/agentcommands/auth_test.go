package agentcommands

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/mocks"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAgentCommands_RegisterUser(t *testing.T) {
	ac, r, w := createMockAgentCommands(t)
	helpStr := enterAuthData + "\n"

	type wantArgs struct {
		login, password string
	}

	tests := []struct {
		name       string
		input      string
		wantArgs   wantArgs
		wantErrMsg string
	}{
		{
			name:  "Test 1. Correct input.",
			input: "demoLog demoPass\n",
			wantArgs: wantArgs{
				login:    "demoLog",
				password: "demoPass",
			},
			wantErrMsg: "",
		},
		{
			name:  "Test 2. Correct input, but user already exist.",
			input: "demoLog demoPass\n",
			wantArgs: wantArgs{
				login:    "demoLog",
				password: "demoPass",
			},
			wantErrMsg: "err: login already exist\n",
		},
		{
			name:       "Test 3. Incorrect input(only 1 field filled).",
			input:      "demoLog\n",
			wantArgs:   wantArgs{},
			wantErrMsg: "err: input error, required 2 fields\n",
		},
		{
			name:       "Test 4. Incorrect input(empty string).",
			input:      "\n",
			wantArgs:   wantArgs{},
			wantErrMsg: "err: input error\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.WriteString(tt.input)

			ac.RegisterUser()
			assert.Equal(t, helpStr+tt.wantErrMsg, w.String())
			defer func() {
				w.Reset()
			}()

			// input args
			if tt.wantErrMsg != "" {
				return
			}
			mockGA := ac.grpcAgent.(*mocks.GrpcAgent)

			gotLogin, ok := mockGA.InputArgs[0].(string)
			require.Equal(t, true, ok)
			gotPass, ok := mockGA.InputArgs[1].(string)
			require.Equal(t, true, ok)
			assert.Equal(t, gotLogin, tt.wantArgs.login)
			assert.Equal(t, gotPass, tt.wantArgs.password)
		})
	}
}

func TestAgentCommands_LoginUser(t *testing.T) {
	ac, r, w := createMockAgentCommands(t)
	helpStr := enterAuthData + "\n"

	regLog, regPass := "regLog", "regPass"
	ac.grpcAgent.(*mocks.GrpcAgent).Users = map[string]storage.User{
		regLog + regPass: {Login: regLog, HashedPassword: regPass},
	}

	type wantArgs struct {
		login, password string
	}

	tests := []struct {
		name       string
		input      string
		wantArgs   wantArgs
		wantErrMsg string
	}{
		{
			name:  "Test 1. Correct input.",
			input: fmt.Sprintf("%s %s\n", regLog, regPass),
			wantArgs: wantArgs{
				login:    regLog,
				password: regPass,
			},
			wantErrMsg: "",
		},
		{
			name:  "Test 2. Correct input, but user not exist.",
			input: "demoLog demoPass\n",
			wantArgs: wantArgs{
				login:    "demoLog",
				password: "demoPass",
			},
			wantErrMsg: "err: login not exist\n",
		},
		{
			name:       "Test 3. Incorrect input.",
			input:      "demoLog\n",
			wantArgs:   wantArgs{},
			wantErrMsg: "err: input error, required 2 fields\n",
		},
		{
			name:       "Test 4. Incorrect input(empty string).",
			input:      "\n",
			wantArgs:   wantArgs{},
			wantErrMsg: "err: input error\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r.WriteString(tt.input)

			ac.LoginUser()
			assert.Equal(t, helpStr+tt.wantErrMsg, w.String())
			defer func() {
				w.Reset()
			}()

			// input args
			if tt.wantErrMsg != "" {
				return
			}
			mockGA := ac.grpcAgent.(*mocks.GrpcAgent)

			gotLogin, ok := mockGA.InputArgs[0].(string)
			require.Equal(t, true, ok)
			gotPass, ok := mockGA.InputArgs[1].(string)
			require.Equal(t, true, ok)
			assert.Equal(t, gotLogin, tt.wantArgs.login)
			assert.Equal(t, gotPass, tt.wantArgs.password)
		})
	}
}
