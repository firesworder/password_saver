package main

import (
	"github.com/firesworder/password_saver/internal/grpcserver"
	"github.com/firesworder/password_saver/internal/server"
	"github.com/firesworder/password_saver/internal/server/env"
	"log"
)

func main() {
	env.ParseEnvArgs()

	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	grpcService, err := grpcserver.NewGRPCServer(s)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(grpcService.Serve(&env.Env))
}
