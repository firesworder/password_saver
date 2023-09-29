package mocks

import (
	"context"
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
)

type UserRepository struct {
	users map[string]storage.User
}

func NewUR() *UserRepository {
	return &UserRepository{users: map[string]storage.User{}}
}

func (ur *UserRepository) CreateUser(ctx context.Context, u storage.User) (*storage.User, error) {
	if u.Login == "demo" {
		return nil, fmt.Errorf("test error")
	}
	ur.users[u.Login] = u
	return &u, nil
}

func (ur *UserRepository) GetUser(ctx context.Context, u storage.User) (*storage.User, error) {
	if u.Login == "demo" {
		return nil, fmt.Errorf("test error")
	}
	uInBD, ok := ur.users[u.Login]
	if !ok {
		return nil, storage.ErrLoginNotExist
	}
	return &uInBD, nil
}

type RecordRepository struct {
	records map[int]storage.Record
	id      int
}

func NewRR() *RecordRepository {
	return &RecordRepository{records: map[int]storage.Record{}, id: 0}
}

func (rr *RecordRepository) AddRecord(ctx context.Context, r storage.Record, uid int) (int, error) {
	if uid == -1 {
		return 0, fmt.Errorf("test error")
	}
	rr.id++
	r.ID = rr.id
	rr.records[rr.id] = r
	return rr.id, nil
}

func (rr *RecordRepository) UpdateRecord(ctx context.Context, r storage.Record, uid int) error {
	if uid == -1 {
		return fmt.Errorf("test error")
	}
	if _, ok := rr.records[r.ID]; !ok {
		return storage.ErrElementNotFound
	}
	rr.records[r.ID] = r
	return nil
}

func (rr *RecordRepository) DeleteRecord(ctx context.Context, r storage.Record, uid int) error {
	if uid == -1 {
		return fmt.Errorf("test error")
	}
	if _, ok := rr.records[r.ID]; !ok {
		return storage.ErrElementNotFound
	}
	delete(rr.records, r.ID)
	return nil
}

func (rr *RecordRepository) GetAll(ctx context.Context, uid int) ([]storage.Record, error) {
	if uid == -1 {
		return nil, fmt.Errorf("test error")
	}
	r := make([]storage.Record, 0)
	for _, elem := range rr.records {
		r = append(r, elem)
	}
	return r, nil
}
