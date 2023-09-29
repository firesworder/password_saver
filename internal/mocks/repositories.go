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
	if u.Login == "demo" {
		return nil, fmt.Errorf("test error")
	}
	// для пароля "uspass"
	return &storage.User{ID: u.ID, Login: u.Login, HashedPassword: "$2a$08$4IS/HXcjs27KmlB08EJRn.51YB3c6cokW.xZZdVOrw0VtcEG/opBa"}, nil
}

type RecordRepository struct {
}

func (rr *RecordRepository) AddRecord(ctx context.Context, r storage.Record, uid int) (int, error) {
	if uid == -1 {
		return 0, fmt.Errorf("test error")
	}
	return 100, nil
}

func (rr *RecordRepository) UpdateRecord(ctx context.Context, r storage.Record, uid int) error {
	if uid == -1 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (rr *RecordRepository) DeleteRecord(ctx context.Context, r storage.Record, uid int) error {
	if uid == -1 {
		return fmt.Errorf("test error")
	}
	return nil
}

func (rr *RecordRepository) GetAll(ctx context.Context, uid int) ([]storage.Record, error) {
	if uid == -1 {
		return nil, fmt.Errorf("test error")
	}
	return []storage.Record{
		{ID: 100, RecordType: "text", Content: []byte("test content 1"), MetaInfo: "MI1"},
		{ID: 150, RecordType: "text", Content: []byte("test content 2"), MetaInfo: "MI2"},
		{ID: 200, RecordType: "bank", Content: []byte("bank content 1"), MetaInfo: "MI3"},
		{ID: 250, RecordType: "binary", Content: []byte("binary content 1"), MetaInfo: "MI4"},
	}, nil
}
