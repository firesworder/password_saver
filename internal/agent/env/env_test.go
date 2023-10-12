package env

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testEnvVars = []string{
	"ADDRESS", "CA_CERT_FILE",
}

func SaveOSVarsState(testEnvVars []string) map[string]string {
	osEnvVarsState := map[string]string{}
	for _, key := range testEnvVars {
		if v, ok := os.LookupEnv(key); ok {
			osEnvVarsState[key] = v
		}
	}
	return osEnvVarsState
}

func UpdateOSEnvState(t *testing.T, testEnvVars []string, newState map[string]string) {
	// удаляю переменные окружения, если они были до этого установлены
	for _, key := range testEnvVars {
		err := os.Unsetenv(key)
		require.NoError(t, err)
	}
	// устанавливаю переменные окружения использованные для теста
	for key, value := range newState {
		err := os.Setenv(key, value)
		require.NoError(t, err)
	}
}

func TestParseEnvArgs(t *testing.T) {
	savedState := SaveOSVarsState(testEnvVars)
	envBefore := Env

	tests := []struct {
		name    string
		cmdStr  string
		envVars map[string]string
		wantEnv Environment
	}{
		{
			name:    "Test 1. Empty cmd args and env vars.",
			cmdStr:  "file.exe",
			envVars: map[string]string{},
			wantEnv: Environment{
				ServerAddress: "localhost:8080",
				CACert:        "",
			},
		},
		{
			name:    "Test 2. CMD args",
			cmdStr:  "file.exe -a=localhost:3030 -ca_c=cert.pem",
			envVars: map[string]string{},
			wantEnv: Environment{
				ServerAddress: "localhost:3030",
				CACert:        "cert.pem",
			},
		},
		{
			name:    "Test 3. CMD and ENV args",
			cmdStr:  "file.exe -a=localhost:3030 -ca_c=dada",
			envVars: map[string]string{"ADDRESS": "localhost:4545", "CA_CERT_FILE": "demo.pem"},
			wantEnv: Environment{
				ServerAddress: "localhost:4545",
				CACert:        "demo.pem",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Env = Environment{
				ServerAddress: "",
				CACert:        "",
			}

			UpdateOSEnvState(t, testEnvVars, tt.envVars)
			// устанавливаю os.Args как эмулятор вызванной команды
			os.Args = strings.Split(tt.cmdStr, " ")
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
			initCmdArgs()

			// сама проверка корректности парсинга
			require.NoError(t, ParseEnvArgs())
			assert.Equal(t, tt.wantEnv, Env)
		})
	}

	Env = envBefore
	UpdateOSEnvState(t, testEnvVars, savedState)
}
