package storage

import (
	pb "github.com/firesworder/password_saver/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBankDataFromGRPC(t *testing.T) {
	got := BankDataFromGRPC(&pb.BankData{
		Id:         10,
		CardNumber: "0011 2233 4455 6677",
		CardExpiry: "10/24",
		Cvv:        "456",
		MetaInfo:   "mi1",
	})
	want := BankData{
		ID:         10,
		CardNumber: "0011 2233 4455 6677",
		CardExpire: "10/24",
		CVV:        "456",
		MetaInfo:   "mi1",
	}
	assert.Equal(t, want, got)
}

func TestBankDataToGRPC(t *testing.T) {
	got := BankDataToGRPC(BankData{
		ID:         10,
		CardNumber: "0011 2233 4455 6677",
		CardExpire: "10/24",
		CVV:        "456",
		MetaInfo:   "mi1",
	})
	want := &pb.BankData{
		Id:         10,
		CardNumber: "0011 2233 4455 6677",
		CardExpiry: "10/24",
		Cvv:        "456",
		MetaInfo:   "mi1",
	}
	assert.Equal(t, want, got)
}

func TestBinaryDataFromGRPC(t *testing.T) {
	got := BinaryDataFromGRPC(&pb.BinaryData{
		Id:         15,
		BinaryData: []byte("dadada"),
		MetaInfo:   "mr1",
	})
	want := BinaryData{
		ID:         15,
		BinaryData: []byte("dadada"),
		MetaInfo:   "mr1",
	}
	assert.Equal(t, want, got)
}

func TestBinaryDataToGRPC(t *testing.T) {
	got := BinaryDataToGRPC(BinaryData{
		ID:         15,
		BinaryData: []byte("dadada"),
		MetaInfo:   "mr1",
	})
	want := &pb.BinaryData{
		Id:         15,
		BinaryData: []byte("dadada"),
		MetaInfo:   "mr1",
	}
	assert.Equal(t, want, got)
}

func TestTextDataFromGRPC(t *testing.T) {
	got := TextDataFromGRPC(&pb.TextData{
		Id:       10,
		TextData: "dadada",
		MetaInfo: "mi1",
	})
	want := TextData{
		ID:       10,
		TextData: "dadada",
		MetaInfo: "mi1",
	}
	assert.Equal(t, want, got)
}

func TestTextDataToGRPC(t *testing.T) {
	got := TextDataToGRPC(TextData{
		ID:       10,
		TextData: "dadada",
		MetaInfo: "mi1",
	})
	want := &pb.TextData{
		Id:       10,
		TextData: "dadada",
		MetaInfo: "mi1",
	}
	assert.Equal(t, want, got)
}
