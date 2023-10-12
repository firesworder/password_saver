// Package env реализует получение переменных окружения из ком.строки(при запуске) или перем.окружения системы.
// Данный пакет получает значения переменных необх. для работы агента системы.
package env

import (
	"flag"

	"github.com/caarlos0/env/v7"
)

// Инициализирует параметры командной строки.
func init() {
	initCmdArgs()
}

// Environment тип для хранения значения переменных необходимых для работы агента.
type Environment struct {
	ServerAddress string `env:"ADDRESS"`
	CACert        string `env:"CA_CERT_FILE"`
}

// Env объект с переменными окружения(из ENV и cmd args).
var Env Environment

func initCmdArgs() {
	flag.StringVar(&Env.ServerAddress, "a", "localhost:8080", "server address")
	flag.StringVar(&Env.CACert, "ca_c", "", "ca_cert file")
}

// ParseEnvArgs Парсит значения полей Env. Сначала из cmd аргументов, затем из перем-х окружения.
func ParseEnvArgs() error {
	// Парсинг аргументов cmd
	flag.Parse()

	// Парсинг перем окружения
	err := env.Parse(&Env)
	if err != nil {
		return err
	}
	return nil
}
