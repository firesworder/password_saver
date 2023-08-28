package grpcserver

import (
	"context"
	"fmt"
	"github.com/firesworder/password_saver/internal/server"
	"github.com/firesworder/password_saver/internal/storage"
	pb "github.com/firesworder/password_saver/proto"
)

type GRPCServer struct {
	pb.UnimplementedPasswordSaverServer

	serv *server.Server
}

func NewGRPCServer(s *server.Server) (*GRPCServer, error) {
	gs := &GRPCServer{serv: s}
	return gs, nil
}

func (gs *GRPCServer) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	if request.Login == "" || request.Password == "" {
		return nil, fmt.Errorf("login and password fields can not be empty")
	}
	// todo: добавить хеширование пароля на сервере
	rUser := storage.User{Login: request.Login, HashedPassword: request.Password}

	// todo: отсюда нужно вернуть токен
	err := gs.serv.RegisterUser(ctx, rUser)
	if err != nil {
		return nil, err
	}

	resp := pb.RegisterUserResponse{Token: "token_template_reg"}
	return &resp, nil
}

func (gs *GRPCServer) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	if request.Login == "" || request.Password == "" {
		return nil, fmt.Errorf("login and password fields can not be empty")
	}

	rUser := storage.User{Login: request.Login, HashedPassword: request.Password}
	err := gs.serv.LoginUser(ctx, rUser)
	if err != nil {
		return nil, err
	}

	resp := pb.LoginUserResponse{Token: "token_template_login"}
	return &resp, nil
}

func (gs *GRPCServer) AddTextDataRecord(ctx context.Context, request *pb.AddTextDataRequest) (*pb.AddTextDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (gs *GRPCServer) UpdateTextDataRecord(ctx context.Context, request *pb.UpdateTextDataRequest) (*pb.UpdateTextDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (gs *GRPCServer) DeleteTextDataRecord(ctx context.Context, request *pb.DeleteTextDataRequest) (*pb.DeleteTextDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (gs *GRPCServer) AddBankDataRecord(ctx context.Context, request *pb.AddBankDataRequest) (*pb.AddBankDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (gs *GRPCServer) UpdateBankDataRecord(ctx context.Context, request *pb.UpdateBankDataRequest) (*pb.UpdateBankDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (gs *GRPCServer) DeleteBankDataRecord(ctx context.Context, request *pb.DeleteBankDataRequest) (*pb.DeleteBankDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (gs *GRPCServer) AddBinaryDataRecord(ctx context.Context, request *pb.AddBinaryDataRequest) (*pb.AddBinaryDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (gs *GRPCServer) UpdateBinaryDataRecord(ctx context.Context, request *pb.UpdateBinaryDataRequest) (*pb.UpdateBinaryDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (gs *GRPCServer) DeleteBinaryDataRecord(ctx context.Context, request *pb.DeleteBinaryDataRequest) (*pb.DeleteBinaryDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (gs *GRPCServer) GetAllRecords(ctx context.Context, request *pb.GetAllRecordsRequest) (*pb.GetAllRecordsResponse, error) {
	//TODO implement me
	panic("implement me")
}
