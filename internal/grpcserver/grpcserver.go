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

func (gs *GRPCServer) UpdateTextDataRecord(ctx context.Context, request *pb.UpdateTextDataRequest) (*pb.UpdateTextDataResponse, error) {
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

func (gs *GRPCServer) DeleteTextDataRecord(ctx context.Context, request *pb.DeleteTextDataRequest) (*pb.DeleteTextDataResponse, error) {
	err := gs.serv.DeleteTextData(ctx, storage.TextData{
		ID: int(request.Id),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteTextDataResponse{}
	return resp, nil
}

func (gs *GRPCServer) AddBankDataRecord(ctx context.Context, request *pb.AddBankDataRequest) (*pb.AddBankDataResponse, error) {
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

func (gs *GRPCServer) UpdateBankDataRecord(ctx context.Context, request *pb.UpdateBankDataRequest) (*pb.UpdateBankDataResponse, error) {
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

func (gs *GRPCServer) DeleteBankDataRecord(ctx context.Context, request *pb.DeleteBankDataRequest) (*pb.DeleteBankDataResponse, error) {
	err := gs.serv.DeleteBankData(ctx, storage.BankData{
		ID: int(request.Id),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteBankDataResponse{}
	return resp, nil
}

func (gs *GRPCServer) AddBinaryDataRecord(ctx context.Context, request *pb.AddBinaryDataRequest) (*pb.AddBinaryDataResponse, error) {
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

func (gs *GRPCServer) UpdateBinaryDataRecord(ctx context.Context, request *pb.UpdateBinaryDataRequest) (*pb.UpdateBinaryDataResponse, error) {
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

func (gs *GRPCServer) DeleteBinaryDataRecord(ctx context.Context, request *pb.DeleteBinaryDataRequest) (*pb.DeleteBinaryDataResponse, error) {
	err := gs.serv.DeleteBinaryData(ctx, storage.BinaryData{
		ID: int(request.Id),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.DeleteBinaryDataResponse{}
	return resp, nil
}

func (gs *GRPCServer) GetAllRecords(ctx context.Context, request *pb.GetAllRecordsRequest) (*pb.GetAllRecordsResponse, error) {
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
