package grpcagent

import (
	"github.com/firesworder/password_saver/internal/grpcserver"
	"github.com/firesworder/password_saver/internal/server"
	pb "github.com/firesworder/password_saver/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"testing"
)

const devTestAddr = "127.0.0.1:3030"

func startTestGRPCServer(t *testing.T, s *server.Server) *grpc.Server {
	serverStarted := make(chan struct{})
	// определяем порт для сервера
	listen, err := net.Listen("tcp", devTestAddr)
	require.NoError(t, err)

	// инциал. сервис
	service, err := grpcserver.NewGRPCServer(s)
	require.NoError(t, err)
	creds, err := credentials.NewServerTLSFromFile(
		"C:\\Users\\person\\GolandProjects\\password_saver\\cert.pem",
		"C:\\Users\\person\\GolandProjects\\password_saver\\privKey.pem",
	)
	if err != nil {
		log.Fatal(err)
	}
	// создаем пустой grpc сервер, без опций
	serverGRPC := grpc.NewServer(grpc.Creds(creds))
	// регистрируем сервис на сервере
	pb.RegisterPasswordSaverServer(serverGRPC, service)

	// запуск сервера в горутине
	go func() {
		serverStarted <- struct{}{}
		// запускаем grpc сервер на выделенном порту 'listen'
		if err := serverGRPC.Serve(listen); err != nil {
			require.NoError(t, err)
		}
	}()
	<-serverStarted

	return serverGRPC
}

func TestNewGRPCAgent(t *testing.T) {
	s, err := server.NewServer()
	require.NoError(t, err)
	testGS := startTestGRPCServer(t, s)

	tests := []struct {
		name string
	}{
		{
			name: "Test 1. Correct agent creation.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGAgent, err := NewGRPCAgent(devTestAddr)
			require.NoError(t, err)
			assert.NotEmpty(t, gotGAgent)
			err = gotGAgent.Close()
			require.NoError(t, err)
		})
	}
	testGS.GracefulStop()
}
