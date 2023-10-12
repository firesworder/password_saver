package mocks

import (
	"context"
	"fmt"

	pb "github.com/firesworder/password_saver/proto"
)

// GRPCServer экземпляр grpc сервера.
type GRPCServer struct {
	pb.UnimplementedPasswordSaverServer
}

func (gs *GRPCServer) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest,
) (*pb.RegisterUserResponse, error) {
	login, password := request.Login, request.Password
	if login == "admin" && password == "admin" {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.RegisterUserResponse{Token: "demo_token"}, nil
}

func (gs *GRPCServer) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	login, password := request.Login, request.Password
	if login == "admin" && password == "admin" {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.LoginUserResponse{Token: "demo_token"}, nil
}

func (gs *GRPCServer) AddTextDataRecord(ctx context.Context, request *pb.AddTextDataRequest,
) (*pb.AddTextDataResponse, error) {
	if request.TextData.MetaInfo == "skip" {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.AddTextDataResponse{Id: 1}, nil
}

func (gs *GRPCServer) UpdateTextDataRecord(ctx context.Context, request *pb.UpdateTextDataRequest,
) (*pb.UpdateTextDataResponse, error) {
	if request.TextData.MetaInfo == "skip" {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.UpdateTextDataResponse{}, nil
}

func (gs *GRPCServer) DeleteTextDataRecord(ctx context.Context, request *pb.DeleteTextDataRequest,
) (*pb.DeleteTextDataResponse, error) {
	if request.Id == 100 {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.DeleteTextDataResponse{}, nil
}

func (gs *GRPCServer) AddBankDataRecord(ctx context.Context, request *pb.AddBankDataRequest,
) (*pb.AddBankDataResponse, error) {
	if request.BankData.MetaInfo == "skip" {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.AddBankDataResponse{Id: 1}, nil
}

func (gs *GRPCServer) UpdateBankDataRecord(ctx context.Context, request *pb.UpdateBankDataRequest,
) (*pb.UpdateBankDataResponse, error) {
	if request.BankData.MetaInfo == "skip" {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.UpdateBankDataResponse{}, nil
}

func (gs *GRPCServer) DeleteBankDataRecord(ctx context.Context, request *pb.DeleteBankDataRequest,
) (*pb.DeleteBankDataResponse, error) {
	if request.Id == 100 {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.DeleteBankDataResponse{}, nil
}

func (gs *GRPCServer) AddBinaryDataRecord(ctx context.Context, request *pb.AddBinaryDataRequest,
) (*pb.AddBinaryDataResponse, error) {
	if request.BinaryData.MetaInfo == "skip" {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.AddBinaryDataResponse{Id: 1}, nil
}

func (gs *GRPCServer) UpdateBinaryDataRecord(ctx context.Context, request *pb.UpdateBinaryDataRequest,
) (*pb.UpdateBinaryDataResponse, error) {
	if request.BinaryData.MetaInfo == "skip" {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.UpdateBinaryDataResponse{}, nil
}

func (gs *GRPCServer) DeleteBinaryDataRecord(ctx context.Context, request *pb.DeleteBinaryDataRequest,
) (*pb.DeleteBinaryDataResponse, error) {
	if request.Id == 100 {
		return nil, fmt.Errorf("demo_error")
	}
	return &pb.DeleteBinaryDataResponse{}, nil
}

func (gs *GRPCServer) GetAllRecords(ctx context.Context, request *pb.GetAllRecordsRequest,
) (*pb.GetAllRecordsResponse, error) {
	return &pb.GetAllRecordsResponse{TextDataList: []*pb.TextData{{Id: 100, TextData: "test data"}}}, nil
}
