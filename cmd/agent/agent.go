package main

import (
	"github.com/firesworder/password_saver/internal/agent"
	"log"
)

func main() {
	a, err := agent.NewAgent()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(a.Serve())
}
