package env

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

var testEnvVars = []string{"ADDRESS"}

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
			},
		},

		// 2 server address
		{
			name:    "Test 2.1. 'ServerAddress'. Set by cmd",
			cmdStr:  "file.exe -a=cmd.site",
			envVars: map[string]string{},
			wantEnv: Environment{
				ServerAddress: "cmd.site",
			},
		},
		{
			name:    "Test 2.2. 'ServerAddress'. Set by ENV",
			cmdStr:  "file.exe",
			envVars: map[string]string{"ADDRESS": "env.site"},
			wantEnv: Environment{
				ServerAddress: "env.site",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Устанавливаю env в дефолтные значения
			Env = Environment{
				ServerAddress: "localhost:8080",
			}

			UpdateOSEnvState(t, testEnvVars, tt.envVars)
			// устанавливаю os.Args как эмулятор вызванной команды
			os.Args = strings.Split(tt.cmdStr, " ")
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
			initCmdArgs()

			// сама проверка корректности парсинга
			require.NotPanics(t, ParseEnvArgs)
			assert.Equal(t, tt.wantEnv, Env)
		})
	}

	Env = envBefore
	UpdateOSEnvState(t, testEnvVars, savedState)
}
