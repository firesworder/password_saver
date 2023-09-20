// Package grpcserver реализует grpc сервер.
package grpcserver

import (
	"context"
	"fmt"
	"github.com/firesworder/password_saver/internal/server"
	"github.com/firesworder/password_saver/internal/server/env"
	"github.com/firesworder/password_saver/internal/storage"
	pb "github.com/firesworder/password_saver/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// GRPCService экземпляр grpc сервера.
type GRPCService struct {
	pb.UnimplementedPasswordSaverServer

	serv     server.IServer
	grpcSObj *grpc.Server
}

// NewGRPCService конструктор grpc сервера(обертка над server.Server).
func NewGRPCService(s server.IServer) (*GRPCService, error) {
	grpcService := &GRPCService{serv: s}
	return grpcService, nil
}

// PrepareServer запускает grpcserver + создает TLS соединение.
func (gs *GRPCService) PrepareServer(env *env.Environment) (*grpc.Server, error) {
	creds, err := credentials.NewServerTLSFromFile(env.CertFile, env.PrivateKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	serverGRPC := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterPasswordSaverServer(serverGRPC, gs)
	return serverGRPC, nil
}

func (gs *GRPCService) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	if request.Login == "" || request.Password == "" {
		return nil, fmt.Errorf("login and password fields can not be empty")
	}

	rUser := storage.User{Login: request.Login, HashedPassword: request.Password}
	userToken, err := gs.serv.RegisterUser(ctx, rUser)
	if err != nil {
		return nil, err
	}

	resp := pb.RegisterUserResponse{Token: userToken}
	return &resp, nil
}

func (gs *GRPCService) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	if request.Login == "" || request.Password == "" {
		return nil, fmt.Errorf("login and password fields can not be empty")
	}

	rUser := storage.User{Login: request.Login, HashedPassword: request.Password}
	userToken, err := gs.serv.LoginUser(ctx, rUser)
	if err != nil {
		return nil, err
	}

	resp := pb.LoginUserResponse{Token: userToken}
	return &resp, nil
}

func (gs *GRPCService) AddTextDataRecord(ctx context.Context, request *pb.AddTextDataRequest) (*pb.AddTextDataResponse, error) {
	id, err := gs.serv.AddTextData(ctx, storage.TextData{
		TextData: request.TextData.TextData,
		MetaInfo: request.TextData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.AddTextDataResponse{Id: int64(id)}
	return resp, nil
}

func (gs *GRPCService) UpdateTextDataRecord(ctx context.Context, request *pb.UpdateTextDataRequest) (*pb.UpdateTextDataResponse, error) {
	err := gs.serv.UpdateTextData(ctx, storage.TextData{
		ID:       int(request.TextData.Id),
		TextData: request.TextData.TextData,
		MetaInfo: request.TextData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateTextDataResponse{}
	return resp, nil
}

func (gs *GRPCService) DeleteTextDataRecord(ctx context.Context, request *pb.DeleteTextDataRequest) (*pb.DeleteTextDataResponse, error) {
	err := gs.serv.DeleteTextData(ctx, storage.TextData{
		ID: int(request.Id),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteTextDataResponse{}
	return resp, nil
}

func (gs *GRPCService) AddBankDataRecord(ctx context.Context, request *pb.AddBankDataRequest) (*pb.AddBankDataResponse, error) {
	id, err := gs.serv.AddBankData(ctx, storage.BankData{
		CardNumber: request.BankData.CardNumber,
		CardExpire: request.BankData.CardExpiry,
		CVV:        request.BankData.Cvv,
		MetaInfo:   request.BankData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.AddBankDataResponse{Id: int64(id)}
	return resp, nil
}

func (gs *GRPCService) UpdateBankDataRecord(ctx context.Context, request *pb.UpdateBankDataRequest) (*pb.UpdateBankDataResponse, error) {
	err := gs.serv.UpdateBankData(ctx, storage.BankData{
		ID:         int(request.BankData.Id),
		CardNumber: request.BankData.CardNumber,
		CardExpire: request.BankData.CardExpiry,
		CVV:        request.BankData.Cvv,
		MetaInfo:   request.BankData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateBankDataResponse{}
	return resp, nil
}

func (gs *GRPCService) DeleteBankDataRecord(ctx context.Context, request *pb.DeleteBankDataRequest) (*pb.DeleteBankDataResponse, error) {
	err := gs.serv.DeleteBankData(ctx, storage.BankData{
		ID: int(request.Id),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteBankDataResponse{}
	return resp, nil
}

func (gs *GRPCService) AddBinaryDataRecord(ctx context.Context, request *pb.AddBinaryDataRequest) (*pb.AddBinaryDataResponse, error) {
	id, err := gs.serv.AddBinaryData(ctx, storage.BinaryData{
		BinaryData: request.BinaryData.BinaryData,
		MetaInfo:   request.BinaryData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.AddBinaryDataResponse{Id: int64(id)}
	return resp, nil
}

func (gs *GRPCService) UpdateBinaryDataRecord(ctx context.Context, request *pb.UpdateBinaryDataRequest) (*pb.UpdateBinaryDataResponse, error) {
	err := gs.serv.UpdateBinaryData(ctx, storage.BinaryData{
		ID:         int(request.BinaryData.Id),
		BinaryData: request.BinaryData.BinaryData,
		MetaInfo:   request.BinaryData.MetaInfo,
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.UpdateBinaryDataResponse{}
	return resp, nil
}

func (gs *GRPCService) DeleteBinaryDataRecord(ctx context.Context, request *pb.DeleteBinaryDataRequest) (*pb.DeleteBinaryDataResponse, error) {
	err := gs.serv.DeleteBinaryData(ctx, storage.BinaryData{
		ID: int(request.Id),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteBinaryDataResponse{}
	return resp, nil
}

func (gs *GRPCService) GetAllRecords(ctx context.Context, request *pb.GetAllRecordsRequest) (*pb.GetAllRecordsResponse, error) {
	records, err := gs.serv.GetAllRecords(ctx)
	if err != nil {
		return nil, err
	}

	textDL := make([]*pb.TextData, 0, len(records.TextDataList))
	bankDL := make([]*pb.BankData, 0, len(records.BankDataList))
	binaryDL := make([]*pb.BinaryData, 0, len(records.BinaryDataList))

	for _, v := range records.TextDataList {
		textDL = append(textDL, &pb.TextData{
			Id:       int64(v.ID),
			TextData: v.TextData,
			MetaInfo: v.MetaInfo,
		})
	}

	for _, v := range records.BankDataList {
		bankDL = append(bankDL, &pb.BankData{
			Id:         int64(v.ID),
			CardNumber: v.CardNumber,
			CardExpiry: v.CardExpire,
			Cvv:        v.CVV,
			MetaInfo:   v.MetaInfo,
		})
	}

	for _, v := range records.BinaryDataList {
		binaryDL = append(binaryDL, &pb.BinaryData{
			Id:         int64(v.ID),
			BinaryData: v.BinaryData,
			MetaInfo:   v.MetaInfo,
		})
	}

	resp := &pb.GetAllRecordsResponse{
		TextDataList:   textDL,
		BankDataList:   bankDL,
		BinaryDataList: binaryDL,
	}
	return resp, nil
}
