// Package grpcagent реализует grpc агент.
package grpcagent

import (
	pb "github.com/firesworder/password_saver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// GRPCAgent экземпляр агента.
type GRPCAgent struct {
	conn       *grpc.ClientConn
	grpcClient pb.PasswordSaverClient
	userToken  string
}

// NewGRPCAgent конструктор grpc агента. Создает соединение к серверу.
func NewGRPCAgent(serverAddress string, caCertFp string) (*GRPCAgent, error) {
	var err error
	agent := GRPCAgent{}

	creds, err := credentials.NewClientTLSFromFile(caCertFp, "127.0.0.1")
	if err != nil {
		return nil, err
	}

	if agent.conn, err = grpc.Dial(serverAddress, grpc.WithTransportCredentials(creds)); err != nil {
		return nil, err
	}
	agent.grpcClient = pb.NewPasswordSaverClient(agent.conn)
	return &agent, nil
}

// Close завершает соединение с сервером.
func (a *GRPCAgent) Close() error {
	return a.conn.Close()
}
