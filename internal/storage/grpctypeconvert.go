package storage

import pb "github.com/firesworder/password_saver/proto"

func TextDataToGRPC(input TextData) *pb.TextData {
	return &pb.TextData{Id: int64(input.ID), TextData: input.TextData, MetaInfo: input.MetaInfo}
}

func TextDataFromGRPC(input *pb.TextData) TextData {
	return TextData{ID: int(input.Id), TextData: input.TextData, MetaInfo: input.MetaInfo}
}

func BankDataToGRPC(input BankData) *pb.BankData {
	return &pb.BankData{
		Id:         int64(input.ID),
		CardNumber: input.CardNumber,
		CardExpiry: input.CardExpire,
		Cvv:        input.CVV,
		MetaInfo:   input.MetaInfo,
	}
}

func BankDataFromGRPC(input *pb.BankData) BankData {
	return BankData{
		ID:         int(input.Id),
		CardNumber: input.CardNumber,
		CardExpire: input.CardExpiry,
		CVV:        input.Cvv,
		MetaInfo:   input.MetaInfo,
	}
}

func BinaryDataToGRPC(input BinaryData) *pb.BinaryData {
	return &pb.BinaryData{Id: int64(input.ID), BinaryData: input.BinaryData, MetaInfo: input.MetaInfo}
}

func BinaryDataFromGRPC(input *pb.BinaryData) BinaryData {
	return BinaryData{ID: int(input.Id), BinaryData: input.BinaryData, MetaInfo: input.MetaInfo}
}
