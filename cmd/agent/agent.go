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
