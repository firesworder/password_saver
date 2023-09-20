// Package env реализует получение переменных окружения из ком.строки(при запуске) или перем.окружения системы.
// Данный пакет получает значения переменных необх. для работы сервера системы.
package env

import (
	"flag"
	"github.com/caarlos0/env/v7"
	"log"
)

// Инициализирует параметры командной строки.
func init() {
	initCmdArgs()
}

// Environment для получения(из ENV и cmd) и хранения переменных окружения сервера.
type Environment struct {
	ServerAddress  string `env:"ADDRESS"`
	DSN            string `env:"DATABASE_DSN"`
	CertFile       string `env:"CERT_FILE"`
	PrivateKeyFile string `env:"PRIVATE_KEY_FILE"`
}

// Env объект с переменными окружения(из ENV и cmd args).
var Env Environment

func initCmdArgs() {
	flag.StringVar(&Env.ServerAddress, "a", "localhost:8080", "server address")
	flag.StringVar(&Env.DSN, "d", "", "database dsn")
	flag.StringVar(&Env.CertFile, "c", "", "cert file")
	flag.StringVar(&Env.PrivateKeyFile, "pk", "", "private key file")
}

// ParseEnvArgs Парсит значения полей Env. Сначала из cmd аргументов, затем из перем-х окружения.
func ParseEnvArgs() {
	// Парсинг аргументов cmd
	flag.Parse()

	// Парсинг перем окружения
	err := env.Parse(&Env)
	if err != nil {
		log.Fatal(err)
	}
}
