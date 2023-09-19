package mocks

import (
	"context"
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
)

type UserRepository struct {
}

func (ur *UserRepository) CreateUser(ctx context.Context, u storage.User) (*storage.User, error) {
	if u.Login == "demo" {
		return nil, fmt.Errorf("test error")
	}
	u.ID = 1
	return &u, nil
}

func (ur *UserRepository) GetUser(ctx context.Context, u storage.User) (*storage.User, error) {
	if u.Login == "demo" && u.HashedPassword == "demo" {
		return nil, fmt.Errorf("test error")
	}
	// для пароля "uspass"
	return &storage.User{ID: u.ID, Login: u.Login, HashedPassword: "$2a$08$4IS/HXcjs27KmlB08EJRn.51YB3c6cokW.xZZdVOrw0VtcEG/opBa"}, nil
}

type TextDataRepository struct {
}

func (t TextDataRepository) AddTextData(ctx context.Context, td storage.TextData, u *storage.User) (int, error) {
	if td.TextData == "" {
		return 0, fmt.Errorf("test error")
	}
	return 1, nil
}

func (t TextDataRepository) UpdateTextData(ctx context.Context, td storage.TextData, u *storage.User) error {
	if td.ID == 100 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (t TextDataRepository) DeleteTextData(ctx context.Context, td storage.TextData, u *storage.User) error {
	if td.ID == 100 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (t TextDataRepository) GetAllRecords(ctx context.Context, u *storage.User) ([]storage.TextData, error) {
	return []storage.TextData{
		{ID: 150, TextData: "td1", MetaInfo: "mi1", UserID: 100},
	}, nil
}

type BankDataRepository struct {
}

func (b BankDataRepository) AddBankData(ctx context.Context, bd storage.BankData, u *storage.User) (int, error) {
	if bd.CardNumber == "" || bd.CardExpire == "" || bd.CVV == "" {
		return 0, fmt.Errorf("test error")
	}
	return 1, nil
}

func (b BankDataRepository) UpdateBankData(ctx context.Context, bd storage.BankData, u *storage.User) error {
	if bd.ID == 100 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (b BankDataRepository) DeleteBankData(ctx context.Context, bd storage.BankData, u *storage.User) error {
	if bd.ID == 100 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (b BankDataRepository) GetAllRecords(ctx context.Context, u *storage.User) ([]storage.BankData, error) {
	return []storage.BankData{
		{ID: 150, CardNumber: "0011223344559900", CardExpire: "12/24", CVV: "333", MetaInfo: "mi1", UserID: 100},
	}, nil
}

type BinaryDataRepository struct {
}

func (b BinaryDataRepository) AddBinaryData(ctx context.Context, bd storage.BinaryData, u *storage.User) (int, error) {
	if len(bd.BinaryData) == 0 {
		return 0, fmt.Errorf("test error")
	}
	return 1, nil
}

func (b BinaryDataRepository) UpdateBinaryData(ctx context.Context, bd storage.BinaryData, u *storage.User) error {
	if bd.ID == 100 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (b BinaryDataRepository) DeleteBinaryData(ctx context.Context, bd storage.BinaryData, u *storage.User) error {
	if bd.ID == 100 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (b BinaryDataRepository) GetAllRecords(ctx context.Context, u *storage.User) ([]storage.BinaryData, error) {
	return []storage.BinaryData{
		{ID: 150, BinaryData: []byte("binary data 1"), MetaInfo: "mi1", UserID: 100},
	}, nil
}
