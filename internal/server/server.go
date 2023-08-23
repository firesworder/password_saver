package server

import (
	"context"
	pb "github.com/firesworder/password_saver/proto"
)

type Server struct {
	pb.UnimplementedPasswordSaverServer
}

func (s *Server) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) AddTextDataRecord(ctx context.Context, request *pb.AddTextDataRequest) (*pb.AddTextDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateTextDataRecord(ctx context.Context, request *pb.UpdateTextDataRequest) (*pb.UpdateTextDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteTextDataRecord(ctx context.Context, request *pb.DeleteTextDataRequest) (*pb.DeleteTextDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) AddBankDataRecord(ctx context.Context, request *pb.AddBankDataRequest) (*pb.AddBankDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateBankDataRecord(ctx context.Context, request *pb.UpdateBankDataRequest) (*pb.UpdateBankDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteBankDataRecord(ctx context.Context, request *pb.DeleteBankDataRequest) (*pb.DeleteBankDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) AddBinaryDataRecord(ctx context.Context, request *pb.AddBinaryDataRequest) (*pb.AddBinaryDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateBinaryDataRecord(ctx context.Context, request *pb.UpdateBinaryDataRequest) (*pb.UpdateBinaryDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeleteBinaryDataRecord(ctx context.Context, request *pb.DeleteBinaryDataRequest) (*pb.DeleteBinaryDataResponse, error) {
	//TODO implement me
	panic("implement me")
}
