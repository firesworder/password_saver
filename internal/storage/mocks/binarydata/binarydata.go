package binarydata

import (
	"context"
	"errors"
	"github.com/firesworder/password_saver/internal/storage"
)

var ErrNotFound = errors.New("element not found")
var ErrEmptyData = errors.New("element can not be empty")

type MockBinaryData struct {
	BinaryData map[int]storage.BinaryData
	LastUsedID int
}

func (m *MockBinaryData) AddBinaryData(ctx context.Context, bd storage.BinaryData) error {
	if len(bd.BinaryData) == 0 {
		return ErrEmptyData
	}

	m.LastUsedID++
	bd.ID = m.LastUsedID
	m.BinaryData[bd.ID] = bd
	return nil
}

func (m *MockBinaryData) UpdateBinaryData(ctx context.Context, bd storage.BinaryData) error {
	if len(bd.BinaryData) == 0 {
		return ErrEmptyData
	}
	if _, ok := m.BinaryData[bd.ID]; !ok {
		return ErrNotFound
	}
	m.BinaryData[bd.ID] = bd
	return nil
}

func (m *MockBinaryData) DeleteBinaryData(ctx context.Context, bd storage.BinaryData) error {
	if _, ok := m.BinaryData[bd.ID]; !ok {
		return ErrNotFound
	}
	delete(m.BinaryData, bd.ID)
	return nil
}
