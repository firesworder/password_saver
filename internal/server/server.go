package server

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"google.golang.org/grpc/metadata"

	"github.com/firesworder/password_saver/internal/crypt"
	"github.com/firesworder/password_saver/internal/server/env"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/firesworder/password_saver/internal/storage/sqlstorage"
)

const ctxTokenParam = "userToken"

// Server основной тип пакета, реализующий функционал пакета.
// Объект типа хранит в себе хранилище(map) токенов и ассоц. с этими токенами авториз. пользователями,
// ссылки на объекты репозиториев данных(от SQL подключения) и сгенерированная соль для генерации новых токенов.
type Server struct {
	authUsers sync.Map
	encoder   *crypt.Encoder
	decoder   *crypt.Decoder
	ssql      *sqlstorage.Storage
	genToken  []byte
}

// NewServer создает подключение к БД(SQL в д.с.) по переданному в env DNS адресу.
// Также генерирует соль для токенов и возвращает в итоге инициал. объект Server.
func NewServer(env *env.Environment) (*Server, error) {
	if env == nil {
		return nil, fmt.Errorf("env can't be nil")
	}

	var err error
	s := &Server{}
	if s.ssql, err = sqlstorage.NewStorage(env.DSN); err != nil {
		return nil, err
	}

	if s.encoder, err = crypt.NewEncoder(env.CertFile); err != nil {
		return nil, err
	}

	if s.decoder, err = crypt.NewDecoder(env.PrivateKeyFile); err != nil {
		return nil, err
	}

	if s.genToken, err = generateRandom(genTokenSize); err != nil {
		return nil, err
	}
	return s, nil
}

// getUserFromContext получает из контекста(метаданных контекста) токен пользователя.
// Если параметр токена в контексте отствует или токен не найден среди авториз. пользователей - возвр. ошибка.
func (s *Server) getUserFromContext(ctx context.Context) (*storage.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("can not access request metadata")
	}

	var token string
	if tokenParam := md.Get(ctxTokenParam); len(tokenParam) != 0 {
		token = tokenParam[0]
	} else {
		return nil, fmt.Errorf("userToken is not set")
	}

	v, ok := s.authUsers.Load(token)
	if !ok {
		return nil, fmt.Errorf("user is not auth")
	}
	user, ok := v.(storage.User)
	if !ok {
		return nil, fmt.Errorf("type assertion error: can not cast %T to %T", v, user)
	}
	return &user, nil
}

// функция возвращает объект storage.Record, получая данные из объекта типа данных rawRecord.
// Функции известны только типы перечисленные в пакете `storage`.
func (s *Server) getRecordFromData(rawRecord interface{}) (r *storage.Record, err error) {
	switch v := rawRecord.(type) {
	case storage.TextData:
		r = &storage.Record{ID: v.ID, RecordType: "text", Content: []byte(v.TextData), MetaInfo: v.MetaInfo}
	case storage.BankData:
		r = &storage.Record{ID: v.ID, RecordType: "bank",
			Content: []byte(strings.Join([]string{v.CardNumber, v.CardExpire, v.CVV}, ",")), MetaInfo: v.MetaInfo}
	case storage.BinaryData:
		r = &storage.Record{ID: v.ID, RecordType: "binary", Content: v.BinaryData, MetaInfo: v.MetaInfo}
	default:
		return nil, fmt.Errorf("unknown datatype")
	}

	if r.Content, err = s.encoder.Encode(r.Content); err != nil {
		return nil, err
	}
	return r, nil
}

// AddRecord добавляет запись данных в БД(после шифрования) данного пользователя.
// В rawRecord нужно передать объект данных.
// Пользователь определяется по токену полученному из контекста.
func (s *Server) AddRecord(ctx context.Context, rawRecord interface{}) (int, error) {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return 0, err
	}

	r, err := s.getRecordFromData(rawRecord)
	if err != nil {
		return 0, err
	}
	return s.ssql.RecordRep.AddRecord(ctx, *r, u.ID)
}

// UpdateRecord обновляет запись данных в БД(после шифрования) данного пользователя.
// В rawRecord нужно передать id обновляемой записи, а также новые поля данных.
// Пользователь определяется по токену полученному из контекста.
func (s *Server) UpdateRecord(ctx context.Context, rawRecord interface{}) error {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return err
	}

	r, err := s.getRecordFromData(rawRecord)
	if err != nil {
		return err
	}
	return s.ssql.RecordRep.UpdateRecord(ctx, *r, u.ID)
}

// DeleteRecord удаляет запись данных из БД данного пользователя.
// В rawRecord нужно передать id записи.
// Пользователь определяется по токену полученному из контекста.
func (s *Server) DeleteRecord(ctx context.Context, rawRecord interface{}) error {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return err
	}

	r, err := s.getRecordFromData(rawRecord)
	if err != nil {
		return err
	}
	return s.ssql.RecordRep.DeleteRecord(ctx, *r, u.ID)
}

// GetAllRecords дешифрует и возвращает все записи данного пользователя.
func (s *Server) GetAllRecords(ctx context.Context) (*storage.RecordsList, error) {
	u, err := s.getUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	recList := &storage.RecordsList{
		TextDataList:   make([]storage.TextData, 0),
		BankDataList:   make([]storage.BankData, 0),
		BinaryDataList: make([]storage.BinaryData, 0),
	}

	rSlice, err := s.ssql.RecordRep.GetAll(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	for _, r := range rSlice {
		content, err := s.decoder.Decode(r.Content)
		if err != nil {
			return nil, err
		}

		switch r.RecordType {
		case "text":
			recList.TextDataList = append(recList.TextDataList,
				storage.TextData{ID: r.ID, TextData: string(content), MetaInfo: r.MetaInfo},
			)
		case "bank":
			bankData := strings.Split(string(content), ",")
			recList.BankDataList = append(recList.BankDataList, storage.BankData{
				ID: r.ID, CardNumber: bankData[0], CardExpire: bankData[1], CVV: bankData[2], MetaInfo: r.MetaInfo,
			})
		case "binary":
			recList.BinaryDataList = append(recList.BinaryDataList,
				storage.BinaryData{ID: r.ID, BinaryData: content, MetaInfo: r.MetaInfo},
			)
		}
	}
	return recList, nil
}
