package main

import (
	"github.com/firesworder/password_saver/internal/grpcserver"
	"github.com/firesworder/password_saver/internal/server"
	"github.com/firesworder/password_saver/internal/server/env"
	"log"
)

func main() {
	env.ParseEnvArgs()

	s, err := server.NewServer(&env.Env)
	if err != nil {
		log.Fatal(err)
	}

	grpcService, err := grpcserver.NewGRPCServer(s, &env.Env)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(grpcService.Serve())
}
