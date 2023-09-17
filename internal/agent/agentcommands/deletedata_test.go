package agentcommands

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/mocks"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAgentCommands_DeleteTextData(t *testing.T) {
	ac, _, w := createMockAgentCommands(t)
	testTD := storage.TextData{ID: 50, TextData: "text data", MetaInfo: "MI 1"}
	ac.state.Set(testTD)

	ac.isAuthorized = false
	ac.DeleteTextData(100)
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true
	mockGA := ac.grpcAgent.(*mocks.GrpcAgent)

	// корректное удаление
	ac.DeleteTextData(testTD.ID)
	// проверяем вывод команды
	assert.Equal(t, "", w.String())
	// проверяем пришедшее в grpc агент
	gotTD, ok := mockGA.InputArgs[0].(storage.TextData)
	require.Equal(t, true, ok)
	wantTD := storage.TextData{ID: testTD.ID}
	assert.Equal(t, wantTD, gotTD)
	// проверяем обновление стейта
	_, err := ac.state.Get(testTD.ID, "text")
	require.NotEqual(t, nil, err)

	// bd запись, ошибка на grpc агенте
	w.Reset()
	ac.DeleteTextData(100)
	assert.Equal(t, "err: test error\n", w.String())

	// td запись, ошибка при удалении из стейта
	w.Reset()
	ac.DeleteTextData(150)
	assert.Equal(t, "err: record was not found\n", w.String())
}

func TestAgentCommands_DeleteBankData(t *testing.T) {
	ac, _, w := createMockAgentCommands(t)
	testBD := storage.BankData{
		ID: 50, CardNumber: "0011223344556677", CardExpire: "09/24", CVV: "521", MetaInfo: "MI 1",
	}
	ac.state.Set(testBD)

	ac.isAuthorized = false
	ac.DeleteBankData(100)
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true
	mockGA := ac.grpcAgent.(*mocks.GrpcAgent)

	// корректное удаление
	ac.DeleteBankData(testBD.ID)
	// проверяем вывод команды
	assert.Equal(t, "", w.String())
	// проверяем пришедшее в grpc агент
	gotBD, ok := mockGA.InputArgs[0].(storage.BankData)
	require.Equal(t, true, ok)
	wantBD := storage.BankData{ID: testBD.ID}
	assert.Equal(t, wantBD, gotBD)
	// проверяем обновление стейта
	_, err := ac.state.Get(testBD.ID, "bank")
	require.NotEqual(t, nil, err)

	// bd запись, ошибка на grpc агенте
	w.Reset()
	ac.DeleteBankData(100)
	assert.Equal(t, "err: test error\n", w.String())

	// td запись, ошибка при удалении из стейта
	w.Reset()
	ac.DeleteBankData(150)
	assert.Equal(t, "err: record was not found\n", w.String())
}

func TestAgentCommands_DeleteBinaryData(t *testing.T) {
	ac, _, w := createMockAgentCommands(t)
	testBD := storage.BinaryData{ID: 50, BinaryData: []byte("binary data"), MetaInfo: "MI 1"}
	ac.state.Set(testBD)

	ac.isAuthorized = false
	ac.DeleteBinaryData(100)
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true
	mockGA := ac.grpcAgent.(*mocks.GrpcAgent)

	// корректное удаление
	ac.DeleteBinaryData(testBD.ID)
	// проверяем вывод команды
	assert.Equal(t, "", w.String())
	// проверяем пришедшее в grpc агент
	gotBD, ok := mockGA.InputArgs[0].(storage.BinaryData)
	require.Equal(t, true, ok)
	wantBD := storage.BinaryData{ID: testBD.ID}
	assert.Equal(t, wantBD, gotBD)
	// проверяем обновление стейта
	_, err := ac.state.Get(testBD.ID, "binary")
	require.NotEqual(t, nil, err)

	// bd запись, ошибка на grpc агенте
	w.Reset()
	ac.DeleteBinaryData(100)
	assert.Equal(t, "err: test error\n", w.String())

	// td запись, ошибка при удалении из стейта
	w.Reset()
	ac.DeleteBinaryData(150)
	assert.Equal(t, "err: record was not found\n", w.String())
}
