package grpcagent

import (
	pb "github.com/firesworder/password_saver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GRPCAgent struct {
	conn       *grpc.ClientConn
	grpcClient pb.PasswordSaverClient
	userToken  string
}

func NewGRPCAgent(serverAddr string, pkFile string) (*GRPCAgent, error) {
	var err error
	agent := GRPCAgent{}

	creds, err := credentials.NewClientTLSFromFile(pkFile, "127.0.0.1")
	if err != nil {
		return nil, err
	}

	if agent.conn, err = grpc.Dial(serverAddr, grpc.WithTransportCredentials(creds)); err != nil {
		return nil, err
	}
	agent.grpcClient = pb.NewPasswordSaverClient(agent.conn)
	return &agent, nil
}

func (a *GRPCAgent) Close() error {
	return a.conn.Close()
}
