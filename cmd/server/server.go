package main

import (
	"github.com/firesworder/password_saver/internal/grpcserver"
	"github.com/firesworder/password_saver/internal/server"
	pb "github.com/firesworder/password_saver/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcService, err := grpcserver.NewGRPCServer(s)
	if err != nil {
		log.Fatal(err)
	}
	serverGRPC := grpc.NewServer()
	pb.RegisterPasswordSaverServer(serverGRPC, grpcService)

	log.Fatal(serverGRPC.Serve(listen))
}
