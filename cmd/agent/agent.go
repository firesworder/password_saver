package main

import (
	"github.com/firesworder/password_saver/internal/agent"
	"github.com/firesworder/password_saver/internal/agent/env"
	"log"
)

func main() {
	env.ParseEnvArgs()

	a, err := agent.NewAgent(&env.Env)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(a.Serve())
}
