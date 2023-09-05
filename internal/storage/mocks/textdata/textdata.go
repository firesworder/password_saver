package textdata

import (
	"context"
	"errors"
	"github.com/firesworder/password_saver/internal/storage"
)

var ErrNotFound = errors.New("element not found")
var ErrEmptyData = errors.New("element can not be empty")

type MockTextData struct {
	TextDataMap map[int]storage.TextData
	LastUsedID  int
}

func (m *MockTextData) AddTextData(ctx context.Context, td storage.TextData) (int, error) {
	if len(td.TextData) == 0 {
		return 0, ErrEmptyData
	}

	m.LastUsedID++
	td.ID = m.LastUsedID
	m.TextDataMap[td.ID] = td
	return td.ID, nil
}

func (m *MockTextData) UpdateTextData(ctx context.Context, td storage.TextData) error {
	if len(td.TextData) == 0 {
		return ErrEmptyData
	}
	if _, ok := m.TextDataMap[td.ID]; !ok {
		return ErrNotFound
	}
	m.TextDataMap[td.ID] = td
	return nil
}

func (m *MockTextData) DeleteTextData(ctx context.Context, td storage.TextData) error {
	if _, ok := m.TextDataMap[td.ID]; !ok {
		return ErrNotFound
	}
	delete(m.TextDataMap, td.ID)
	return nil
}

func (m *MockTextData) GetAllRecords(ctx context.Context) ([]storage.TextData, error) {
	result := make([]storage.TextData, 0, len(m.TextDataMap))
	for _, v := range m.TextDataMap {
		result = append(result, v)
	}
	return result, nil
}
