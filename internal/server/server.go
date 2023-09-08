// Package server реализует сервер, как прослойку между grpcserver и sqlstorage.
package server

import (
	"context"
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/firesworder/password_saver/internal/storage/sqlstorage"
	"google.golang.org/grpc/metadata"
)

const ctxTokenParam = "userToken"

type Server struct {
	authUsers map[string]storage.User

	uRep    storage.UserRepository
	tRep    storage.TextDataRepository
	bankRep storage.BankDataRepository
	binRep  storage.BinaryDataRepository

	genToken []byte
}

func NewServer() (*Server, error) {
	ssql, err := sqlstorage.NewStorage(storage.DevDSN)
	if err != nil {
		return nil, err
	}

	genToken, err := generateRandom(32)
	if err != nil {
		return nil, err
	}

	s := &Server{
		authUsers: map[string]storage.User{},

		uRep:    ssql.UserRep,
		tRep:    ssql.TextRep,
		bankRep: ssql.BankRep,
		binRep:  ssql.BinaryRep,

		genToken: genToken,
	}
	return s, nil
}

func (s *Server) getUserFromContext(ctx context.Context) (*storage.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("can not access request metadata")
	}

	var token string
	if tokenParam := md.Get(ctxTokenParam); len(tokenParam) != 0 {
		token = tokenParam[0]
	} else {
		return nil, fmt.Errorf("userToken is not set")
	}

	user, ok := s.authUsers[token]
	if !ok {
		return nil, fmt.Errorf("user is not auth")
	}
	return &user, nil
}

func (s *Server) AddTextData(ctx context.Context, textData storage.TextData) (int, error) {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return 0, err
	}

	id, err := s.tRep.AddTextData(ctx, textData, u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Server) UpdateTextData(ctx context.Context, textData storage.TextData) error {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return err
	}

	err = s.tRep.UpdateTextData(ctx, textData, u)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) DeleteTextData(ctx context.Context, textData storage.TextData) error {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return err
	}

	err = s.tRep.DeleteTextData(ctx, textData, u)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) AddBankData(ctx context.Context, bankData storage.BankData) (int, error) {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return 0, err
	}

	id, err := s.bankRep.AddBankData(ctx, bankData, u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Server) UpdateBankData(ctx context.Context, bankData storage.BankData) error {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return err
	}

	err = s.bankRep.UpdateBankData(ctx, bankData, u)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) DeleteBankData(ctx context.Context, bankData storage.BankData) error {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return err
	}

	err = s.bankRep.DeleteBankData(ctx, bankData, u)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) AddBinaryData(ctx context.Context, binaryData storage.BinaryData) (int, error) {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return 0, err
	}

	id, err := s.binRep.AddBinaryData(ctx, binaryData, u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Server) UpdateBinaryData(ctx context.Context, binaryData storage.BinaryData) error {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return err
	}

	err = s.binRep.UpdateBinaryData(ctx, binaryData, u)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) DeleteBinaryData(ctx context.Context, binaryData storage.BinaryData) error {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return err
	}

	err = s.binRep.DeleteBinaryData(ctx, binaryData, u)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) GetAllRecords(ctx context.Context) (*storage.RecordsList, error) {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	recList := &storage.RecordsList{}
	recList.TextDataList, err = s.tRep.GetAllRecords(ctx, u)
	if err != nil {
		return nil, err
	}
	recList.BankDataList, err = s.bankRep.GetAllRecords(ctx, u)
	if err != nil {
		return nil, err
	}
	recList.BinaryDataList, err = s.binRep.GetAllRecords(ctx, u)
	if err != nil {
		return nil, err
	}
	return recList, nil
}
