package agentcommands

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/mocks"
	"github.com/firesworder/password_saver/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func TestAgentCommands_UpdateTextData(t *testing.T) {
	ac, r, w := createMockAgentCommands(t)
	newTD := storage.TextData{ID: 50, TextData: "new text data", MetaInfo: "MI upd"}
	ac.state.Set(storage.TextData{ID: 50, TextData: "text data", MetaInfo: "MI1"})

	ac.isAuthorized = false
	ac.UpdateTextData(newTD.ID)
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true
	mockGA := ac.grpcAgent.(*mocks.GrpcAgent)

	// корректная td запись
	// готовим инпут
	r.WriteString(fmt.Sprintf("%s\n%s\n", newTD.TextData, newTD.MetaInfo))
	ac.UpdateTextData(newTD.ID)
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", enterTextData, enterMetaInfo), w.String())

	// проверяем пришедшее в grpc агент
	gotTD, ok := mockGA.InputArgs[0].(storage.TextData)
	require.Equal(t, true, ok)
	assert.Equal(t, newTD, gotTD)

	// проверяем обновление стейта
	v, err := ac.state.Get(newTD.ID, "text")
	require.NoError(t, err)
	vTD, ok := v.(storage.TextData)
	require.Equal(t, true, ok)
	assert.Equal(t, newTD, vTD)

	// td запись, ошибка на вводе text data
	r.Reset()
	w.Reset()
	ac.UpdateTextData(newTD.ID)
	assert.Equal(t, fmt.Sprintf("%s\nerr: %s\n", enterTextData, io.EOF), w.String())

	// td запись, ошибка на вводе meta info
	r.Reset()
	w.Reset()
	r.WriteString("text data number 2\n")
	ac.UpdateTextData(newTD.ID)
	assert.Equal(t, fmt.Sprintf("%s\n%s\nerr: %s\n", enterTextData, enterMetaInfo, io.EOF), w.String())

	// td запись с ошибкой в grpc вызове
	r.Reset()
	w.Reset()
	r.WriteString("\n" + "\n")
	ac.UpdateTextData(newTD.ID)
	assert.Equal(t, fmt.Sprintf("%s\n%s\n%s\n", enterTextData, enterMetaInfo, "err: data invalid"), w.String())
}

func TestAgentCommands_UpdateBankData(t *testing.T) {
	ac, r, w := createMockAgentCommands(t)
	newBD := storage.BankData{
		ID: 50, CardNumber: "9988776655443322", CardExpire: "11/25", CVV: "444", MetaInfo: "MI upd",
	}
	ac.state.Set(storage.BankData{
		ID: 50, CardNumber: "0011223344556677", CardExpire: "09/23", CVV: "612", MetaInfo: "MI",
	})

	ac.isAuthorized = false
	ac.UpdateBankData(newBD.ID)
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true
	mockGA := ac.grpcAgent.(*mocks.GrpcAgent)

	// корректная bd запись
	// готовим инпут
	r.WriteString(fmt.Sprintf("%s %s %s\n%s\n", newBD.CardNumber, newBD.CardExpire, newBD.CVV, newBD.MetaInfo))
	ac.UpdateBankData(newBD.ID)
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", enterBankData, enterMetaInfo), w.String())

	// проверяем пришедшее в grpc агент
	gotBD, ok := mockGA.InputArgs[0].(storage.BankData)
	require.Equal(t, true, ok)
	assert.Equal(t, newBD, gotBD)

	// проверяем обновление стейта
	v, err := ac.state.Get(newBD.ID, "bank")
	require.NoError(t, err)
	vBD, ok := v.(storage.BankData)
	require.Equal(t, true, ok)
	assert.Equal(t, newBD, vBD)

	// bd запись, ошибка на вводе bank data
	r.Reset()
	w.Reset()
	r.WriteString("")
	ac.UpdateBankData(newBD.ID)
	assert.Equal(t, fmt.Sprintf("%s\nerr: %s\n", enterBankData, io.EOF), w.String())

	// bd запись, ошибка на кол-ве введенных полей(меньше 3х)
	r.Reset()
	w.Reset()
	r.WriteString(fmt.Sprintf("%s %s\n\n", newBD.CardNumber, newBD.CardExpire))
	ac.UpdateBankData(newBD.ID)
	assert.Equal(t, fmt.Sprintf("%s\nerr: %s\n", enterBankData, "input error"), w.String())

	// bd запись, ошибка на вводе meta info
	ac, r, w = createMockAgentCommands(t)
	ac.isAuthorized = true
	r.WriteString(fmt.Sprintf("%s %s %s\n", newBD.CardNumber, newBD.CardExpire, newBD.CVV))
	ac.UpdateBankData(newBD.ID)
	assert.Equal(t, fmt.Sprintf("%s\n%s\nerr: %s\n", enterBankData, enterMetaInfo, io.EOF), w.String())

	// bd запись с ошибкой в grpc вызове
	r.Reset()
	w.Reset()
	// CVV == 000 - спец. кейс для триггера ошибки
	r.WriteString(fmt.Sprintf("%s %s %s\n\n", newBD.CardNumber, newBD.CardExpire, "000"))
	ac.UpdateBankData(newBD.ID)
	assert.Equal(t, fmt.Sprintf("%s\n%s\n%s\n", enterBankData, enterMetaInfo, "err: data invalid"), w.String())
}

func TestAgentCommands_UpdateBinaryData(t *testing.T) {
	ac, r, w := createMockAgentCommands(t)
	newBD := storage.BinaryData{ID: 50, BinaryData: []byte("Ayaka"), MetaInfo: "MI upd"}
	ac.state.Set(storage.BinaryData{ID: 50, BinaryData: []byte("Qiqi"), MetaInfo: "MI"})

	ac.isAuthorized = false
	ac.UpdateBinaryData(newBD.ID)
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true
	mockGA := ac.grpcAgent.(*mocks.GrpcAgent)
	bindTDir := "binary_test"

	// корректная bd запись
	// готовим инпут
	r.WriteString(fmt.Sprintf("%s/%s\n%s\n", bindTDir, "binary_test.txt", newBD.MetaInfo))
	ac.UpdateBinaryData(newBD.ID)
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", enterBinaryData, enterMetaInfo), w.String())

	// проверяем пришедшее в grpc агент
	gotBD, ok := mockGA.InputArgs[0].(storage.BinaryData)
	require.Equal(t, true, ok)
	assert.Equal(t, newBD, gotBD)

	// проверяем обновление стейта
	v, err := ac.state.Get(newBD.ID, "binary")
	require.NoError(t, err)
	vBD, ok := v.(storage.BinaryData)
	require.Equal(t, true, ok)
	assert.Equal(t, newBD, vBD)

	// bd запись, ошибка на вводе пути к файлу
	r.Reset()
	w.Reset()
	ac.UpdateBinaryData(newBD.ID)
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", enterBinaryData, "err: EOF"), w.String())

	// bd запись, ошибка пути к файлу(файл не существует)
	r.Reset()
	w.Reset()
	r.WriteString(fmt.Sprintf("%s/%s\n", bindTDir, "demo_file.txt"))
	ac.UpdateBinaryData(newBD.ID)
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", enterBinaryData, "err: file does not exist"), w.String())

	// bd запись, ошибка на вводе meta info
	r.Reset()
	w.Reset()
	r.WriteString(fmt.Sprintf("%s/%s\n", bindTDir, "binary_test.txt"))
	ac.UpdateBinaryData(newBD.ID)
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\nerr: %s\n", enterBinaryData, enterMetaInfo, io.EOF), w.String())

	// bd запись, ошибка в grpc агенте
	r.Reset()
	w.Reset()
	r.WriteString(fmt.Sprintf("%s/%s\n%s\n", bindTDir, "empty_test.txt", newBD.MetaInfo))
	ac.UpdateBinaryData(newBD.ID)
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\nerr: %s\n", enterBinaryData, enterMetaInfo, "data invalid"), w.String())
}
