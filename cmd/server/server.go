package main

import (
	"fmt"
	"log"
	"net"

	"github.com/firesworder/password_saver/internal/grpcserver"
	"github.com/firesworder/password_saver/internal/server"
	"github.com/firesworder/password_saver/internal/server/env"
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

	s, err := server.NewServer(&env.Env)
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.Listen("tcp", env.Env.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}
	grpcService, err := grpcserver.NewGRPCService(s)
	if err != nil {
		log.Fatal(err)
	}
	serverGRPC, err := grpcService.PrepareServer(&env.Env)
	if err = serverGRPC.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
