package grpcagent

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/firesworder/password_saver/proto"
)

// GRPCAgent основной тип пакета, реализующий функционал grpc агента.
// В себе хранит помимо подключения к grpc серверу, токен авторизации на этом сервере, используемый
// для методов работы с данными.
type GRPCAgent struct {
	conn       *grpc.ClientConn
	grpcClient pb.PasswordSaverClient
	userToken  string
}

// NewGRPCAgent создает защищенное(TLS) подключение к серверу.
// В аргументах требуется передать адрес к серверу и удостов.сертификат.
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
