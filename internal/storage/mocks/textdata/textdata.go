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

func (m *MockTextData) AddTextData(ctx context.Context, td storage.TextData) error {
	if len(td.TextData) == 0 {
		return ErrEmptyData
	}

	m.LastUsedID++
	td.ID = m.LastUsedID
	m.TextDataMap[td.ID] = td
	return nil
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
