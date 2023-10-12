// Точка сборкиа агента.
// Доступные команды можно увидеть набрав help в консоли(после запуска агента)
// Кратко процесс работы: авториз.\регистрация и затем работа с данными пользователя

// Переменные окружения для агента:
// cmd: `ca_c` - путь к файлу к ca сертификату, `a` - адрес сервера
// перем.окружения: `CA_CERT_FILE` - путь к файлу к ca сертификату, `ADDRESS` - адрес сервера
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/firesworder/password_saver/internal/agent"
	"github.com/firesworder/password_saver/internal/agent/agentcommands/grpcagent"
	"github.com/firesworder/password_saver/internal/agent/env"
)

var (
	BuildVersion = "N/A"
	BuildDate    = "N/A"
	BuildCommit  = "N/A"
)

func main() {
	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n", BuildVersion, BuildDate, BuildCommit)
	if err := env.ParseEnvArgs(); err != nil {
		log.Fatal(err)
	}

	grpcAgent, err := grpcagent.NewGRPCAgent(env.Env.ServerAddress, env.Env.CACert)
	if err != nil {
		log.Fatal(err)
	}

	a := agent.NewAgent(grpcAgent)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	a.Serve(ctx)
}
