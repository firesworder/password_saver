package mocks

import (
	"context"
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
)

type Server struct {
}

func (s *Server) RegisterUser(ctx context.Context, user storage.User) (string, error) {
	if user.Login == "admin" && user.HashedPassword == "admin" {
		return "", fmt.Errorf("test_error")
	}
	return "some_token", nil
}

func (s *Server) LoginUser(ctx context.Context, user storage.User) (string, error) {
	if user.Login == "admin" && user.HashedPassword == "admin" {
		return "", fmt.Errorf("test_error")
	}
	return "some_token", nil
}

func (s *Server) AddTextData(ctx context.Context, textData storage.TextData) (int, error) {
	if textData.TextData == "" {
		return 0, fmt.Errorf("test_error")
	}
	return 1, nil
}

func (s *Server) UpdateTextData(ctx context.Context, textData storage.TextData) error {
	if textData.ID == 100 {
		return fmt.Errorf("test_error")
	}
	return nil
}

func (s *Server) DeleteTextData(ctx context.Context, textData storage.TextData) error {
	if textData.ID == 100 {
		return fmt.Errorf("test_error")
	}
	return nil
}

func (s *Server) AddBankData(ctx context.Context, bankData storage.BankData) (int, error) {
	if bankData.CardNumber == "" || bankData.CardExpire == "" || bankData.CVV == "" || bankData.CVV == "000" {
		return 0, fmt.Errorf("test_error")
	}
	return 1, nil
}

func (s *Server) UpdateBankData(ctx context.Context, bankData storage.BankData) error {
	if bankData.ID == 100 {
		return fmt.Errorf("test_error")
	}
	return nil
}

func (s *Server) DeleteBankData(ctx context.Context, bankData storage.BankData) error {
	if bankData.ID == 100 {
		return fmt.Errorf("test_error")
	}
	return nil
}

func (s *Server) AddBinaryData(ctx context.Context, binaryData storage.BinaryData) (int, error) {
	if len(binaryData.BinaryData) == 0 {
		return 0, fmt.Errorf("test_error")
	}
	return 1, nil
}

func (s *Server) UpdateBinaryData(ctx context.Context, binaryData storage.BinaryData) error {
	if binaryData.ID == 100 {
		return fmt.Errorf("test_error")
	}
	return nil
}

func (s *Server) DeleteBinaryData(ctx context.Context, binaryData storage.BinaryData) error {
	if binaryData.ID == 100 {
		return fmt.Errorf("test_error")
	}
	return nil
}

func (s *Server) GetAllRecords(ctx context.Context) (*storage.RecordsList, error) {
	result := &storage.RecordsList{
		TextDataList: []storage.TextData{{ID: 150, TextData: "td", MetaInfo: "mi1", UserID: 15}},
		BankDataList: []storage.BankData{
			{ID: 150, CardNumber: "1122334466778899", CardExpire: "09/12", CVV: "456", UserID: 15},
			{ID: 160, CardNumber: "0055662233661122", CardExpire: "09/24", CVV: "264", UserID: 17},
		},
		BinaryDataList: []storage.BinaryData{{ID: 200, BinaryData: []byte("demo byte"), UserID: 15}},
	}
	return result, nil
}
