// Точка сборкиа агента.
// Доступные команды можно увидеть набрав help в консоли(после запуска агента)
// Кратко процесс работы: авториз.\регистрация и затем работа с данными пользователя

// Переменные окружения для агента:
// cmd: `ca_c` - путь к файлу к ca сертификату, `a` - адрес сервера
// перем.окружения: `CA_CERT_FILE` - путь к файлу к ca сертификату, `ADDRESS` - адрес сервера
package main

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/agent"
	"github.com/firesworder/password_saver/internal/agent/env"
	"log"
)

var (
	buildVersion = "0.0.1"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n", buildVersion, buildDate, buildCommit)

	env.ParseEnvArgs()

	a, err := agent.NewAgent(&env.Env)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(a.Serve())
}
