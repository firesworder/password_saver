package storage

import pb "github.com/firesworder/password_saver/proto"

// TextDataToGRPC возвращает объект proto.TextData на основе storage.TextData.
func TextDataToGRPC(input TextData) *pb.TextData {
	return &pb.TextData{Id: int64(input.ID), TextData: input.TextData, MetaInfo: input.MetaInfo}
}

// TextDataFromGRPC возвращает объект storage.TextData на основе proto.TextData.
func TextDataFromGRPC(input *pb.TextData) TextData {
	return TextData{ID: int(input.Id), TextData: input.TextData, MetaInfo: input.MetaInfo}
}

// BankDataToGRPC возвращает объект proto.BankData на основе storage.BankData.
func BankDataToGRPC(input BankData) *pb.BankData {
	return &pb.BankData{
		Id:         int64(input.ID),
		CardNumber: input.CardNumber,
		CardExpiry: input.CardExpire,
		Cvv:        input.CVV,
		MetaInfo:   input.MetaInfo,
	}
}

// BankDataFromGRPC возвращает объект storage.BankData на основе proto.BankData.
func BankDataFromGRPC(input *pb.BankData) BankData {
	return BankData{
		ID:         int(input.Id),
		CardNumber: input.CardNumber,
		CardExpire: input.CardExpiry,
		CVV:        input.Cvv,
		MetaInfo:   input.MetaInfo,
	}
}

// BinaryDataToGRPC возвращает объект proto.BinaryData на основе storage.BinaryData.
func BinaryDataToGRPC(input BinaryData) *pb.BinaryData {
	return &pb.BinaryData{Id: int64(input.ID), BinaryData: input.BinaryData, MetaInfo: input.MetaInfo}
}

// BinaryDataFromGRPC возвращает объект storage.BinaryData на основе proto.BinaryData.
func BinaryDataFromGRPC(input *pb.BinaryData) BinaryData {
	return BinaryData{ID: int(input.Id), BinaryData: input.BinaryData, MetaInfo: input.MetaInfo}
}
