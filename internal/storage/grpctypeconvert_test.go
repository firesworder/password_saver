package storage

import (
	pb "github.com/firesworder/password_saver/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBankDataFromGRPC(t *testing.T) {
	wantS := BankData{
		ID: 150, CardNumber: "0011 2233 4455 6677", CardExpire: "09/23", CVV: "345", MetaInfo: "MI1",
	}
	inputPb := &pb.BankData{
		Id: 150, CardNumber: "0011 2233 4455 6677", CardExpiry: "09/23", Cvv: "345", MetaInfo: "MI1",
	}
	assert.Equal(t, wantS, BankDataFromGRPC(inputPb))
}

func TestBankDataToGRPC(t *testing.T) {
	wantPb := &pb.BankData{
		Id: 150, CardNumber: "0011 2233 4455 6677", CardExpiry: "09/23", Cvv: "345", MetaInfo: "MI1",
	}
	inputS := BankData{
		ID: 150, CardNumber: "0011 2233 4455 6677", CardExpire: "09/23", CVV: "345", MetaInfo: "MI1",
	}
	assert.Equal(t, wantPb, BankDataToGRPC(inputS))
}

func TestBinaryDataFromGRPC(t *testing.T) {
	wantS := BinaryData{
		ID: 150, BinaryData: []byte("binary data"), MetaInfo: "MI1",
	}
	inputPb := &pb.BinaryData{
		Id: 150, BinaryData: []byte("binary data"), MetaInfo: "MI1",
	}
	assert.Equal(t, wantS, BinaryDataFromGRPC(inputPb))
}

func TestBinaryDataToGRPC(t *testing.T) {
	wantPb := &pb.BinaryData{
		Id: 150, BinaryData: []byte("binary data"), MetaInfo: "MI1",
	}
	inputS := BinaryData{
		ID: 150, BinaryData: []byte("binary data"), MetaInfo: "MI1",
	}
	assert.Equal(t, wantPb, BinaryDataToGRPC(inputS))
}

func TestTextDataFromGRPC(t *testing.T) {
	wantS := TextData{
		ID: 150, TextData: "text content", MetaInfo: "MI1",
	}
	inputPb := &pb.TextData{
		Id: 150, TextData: "text content", MetaInfo: "MI1",
	}
	assert.Equal(t, wantS, TextDataFromGRPC(inputPb))
}

func TestTextDataToGRPC(t *testing.T) {
	wantPb := &pb.TextData{
		Id: 150, TextData: "text content", MetaInfo: "MI1",
	}
	inputS := TextData{
		ID: 150, TextData: "text content", MetaInfo: "MI1",
	}
	assert.Equal(t, wantPb, TextDataToGRPC(inputS))
}
