package grpcserver

import (
	"context"
	"github.com/firesworder/password_saver/internal/server"
	pb "github.com/firesworder/password_saver/proto"
)

type GRPCServer struct {
	pb.UnimplementedPasswordSaverServer

	s *server.Server
}

func NewGRPCServer(s *server.Server) (*GRPCServer, error) {
	gs := &GRPCServer{s: s}
	return gs, nil
}

func (s *GRPCServer) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) AddTextDataRecord(ctx context.Context, request *pb.AddTextDataRequest) (*pb.AddTextDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) UpdateTextDataRecord(ctx context.Context, request *pb.UpdateTextDataRequest) (*pb.UpdateTextDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) DeleteTextDataRecord(ctx context.Context, request *pb.DeleteTextDataRequest) (*pb.DeleteTextDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) AddBankDataRecord(ctx context.Context, request *pb.AddBankDataRequest) (*pb.AddBankDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) UpdateBankDataRecord(ctx context.Context, request *pb.UpdateBankDataRequest) (*pb.UpdateBankDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) DeleteBankDataRecord(ctx context.Context, request *pb.DeleteBankDataRequest) (*pb.DeleteBankDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) AddBinaryDataRecord(ctx context.Context, request *pb.AddBinaryDataRequest) (*pb.AddBinaryDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) UpdateBinaryDataRecord(ctx context.Context, request *pb.UpdateBinaryDataRequest) (*pb.UpdateBinaryDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) DeleteBinaryDataRecord(ctx context.Context, request *pb.DeleteBinaryDataRequest) (*pb.DeleteBinaryDataResponse, error) {
	//TODO implement me
	panic("implement me")
}
