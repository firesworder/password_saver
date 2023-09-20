// Package agentstate реализует локальный стейт записей пользователя.
// В рамках пакета(и спец.объекта State) реализовано как само локальное хранилище, так и методы доступа к нему:
// - получение записи, сохранение записи и ее удаление(из стейта)
package agentstate

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
)

// State хранилище данных пользователя(текстовых, банковских и бинарных) с методами непрямого доступа к ним.
type State struct {
	textDL   map[int]storage.TextData
	bankDL   map[int]storage.BankData
	binaryDL map[int]storage.BinaryData
}

// NewState инициализирует пустые map для каждого из типа данных.
func NewState() *State {
	return &State{
		textDL:   map[int]storage.TextData{},
		bankDL:   map[int]storage.BankData{},
		binaryDL: map[int]storage.BinaryData{},
	}
}

// Get возвращает запись с указанным id и типом записи.
// Если запись не найдена - возвращается ошибка.
func (s *State) Get(id int, dataType string) (interface{}, error) {
	var v interface{}
	var ok bool
	if dataType == "text" {
		if v, ok = s.textDL[id]; ok {
			return v, nil
		}
	} else if dataType == "bank" {
		if v, ok = s.bankDL[id]; ok {
			return v, nil
		}
	} else if dataType == "binary" {
		if v, ok = s.binaryDL[id]; ok {
			return v, nil
		}
	}
	return nil, fmt.Errorf("record was not found")
}

// Set сохраняет запись в стейте(если она представлена одним из типов данных пол-ля), иначе ничего не делает.
func (s *State) Set(record interface{}) {
	switch v := record.(type) {
	case storage.TextData:
		s.textDL[v.ID] = v
	case storage.BankData:
		s.bankDL[v.ID] = v
	case storage.BinaryData:
		s.binaryDL[v.ID] = v
	}
}

// Delete удаляет запись с указанным в ID.
// Если запись не найдена с таким ID - возвращается ошибка.
func (s *State) Delete(id int) error {
	var ok bool
	if _, ok = s.textDL[id]; ok {
		delete(s.textDL, id)
		return nil
	}
	if _, ok = s.bankDL[id]; ok {
		delete(s.bankDL, id)
		return nil
	}
	if _, ok = s.binaryDL[id]; ok {
		delete(s.binaryDL, id)
		return nil
	}
	return fmt.Errorf("record was not found")
}
