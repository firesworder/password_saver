package grpcagent

import (
	pb "github.com/firesworder/password_saver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCAgent struct {
	conn       *grpc.ClientConn
	grpcClient pb.PasswordSaverClient
}

func NewGRPCAgent(serverAddr string) (*GRPCAgent, error) {
	var err error
	agent := GRPCAgent{}
	if agent.conn, err = grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	}
	agent.grpcClient = pb.NewPasswordSaverClient(agent.conn)
	return &agent, nil
}

func (a *GRPCAgent) Close() error {
	return a.conn.Close()
}
