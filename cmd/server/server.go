package main

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/grpcserver"
	"github.com/firesworder/password_saver/internal/server"
	"github.com/firesworder/password_saver/internal/server/env"
	"log"
	"net"
)

var (
	buildVersion = "0.0.1"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	fmt.Printf("Build version: %s\nBuild date: %s\nBuild commit: %s\n", buildVersion, buildDate, buildCommit)

	env.ParseEnvArgs()

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
	log.Fatal(serverGRPC.Serve(listen))
}
