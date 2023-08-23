package textdata

import (
	"context"
	"errors"
	"password_saver/internal/storage"
)

var ErrNotFound = errors.New("element not found")
var ErrEmptyData = errors.New("element can not be empty")

type MockTextData struct {
	textDataMap map[int]storage.TextData
	lastUsedID  int
}

func (m *MockTextData) AddTextData(ctx context.Context, td storage.TextData) error {
	if len(td.TextData) == 0 {
		return ErrEmptyData
	}

	m.lastUsedID++
	td.ID = m.lastUsedID
	m.textDataMap[td.ID] = td
	return nil
}

func (m *MockTextData) UpdateTextData(ctx context.Context, td storage.TextData) error {
	if len(td.TextData) == 0 {
		return ErrEmptyData
	}
	if _, ok := m.textDataMap[td.ID]; !ok {
		return ErrNotFound
	}
	m.textDataMap[td.ID] = td
	return nil
}

func (m *MockTextData) DeleteTextData(ctx context.Context, td storage.TextData) error {
	if _, ok := m.textDataMap[td.ID]; !ok {
		return ErrNotFound
	}
	delete(m.textDataMap, td.ID)
	return nil
}
