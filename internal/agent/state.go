package agent

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
)

type state struct {
	textDL   map[int]storage.TextData
	bankDL   map[int]storage.BankData
	binaryDL map[int]storage.BinaryData
}

func newState() *state {
	return &state{
		textDL:   map[int]storage.TextData{},
		bankDL:   map[int]storage.BankData{},
		binaryDL: map[int]storage.BinaryData{},
	}
}

func (s *state) get(id int, dataType string) (interface{}, error) {
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

func (s *state) set(record interface{}) {
	switch v := record.(type) {
	case storage.TextData:
		s.textDL[v.ID] = v
	case storage.BankData:
		s.bankDL[v.ID] = v
	case storage.BinaryData:
		s.binaryDL[v.ID] = v
	}
}

func (s *state) delete(id int) error {
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
