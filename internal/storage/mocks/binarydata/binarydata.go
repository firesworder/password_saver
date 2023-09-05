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

func (m *MockBinaryData) AddBinaryData(ctx context.Context, bd storage.BinaryData) (int, error) {
	if len(bd.BinaryData) == 0 {
		return 0, ErrEmptyData
	}

	m.LastUsedID++
	bd.ID = m.LastUsedID
	m.BinaryData[bd.ID] = bd
	return bd.ID, nil
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

func (m *MockBinaryData) GetAllRecords(ctx context.Context) ([]storage.BinaryData, error) {
	result := make([]storage.BinaryData, 0, len(m.BinaryData))
	for _, v := range m.BinaryData {
		result = append(result, v)
	}
	return result, nil
}
