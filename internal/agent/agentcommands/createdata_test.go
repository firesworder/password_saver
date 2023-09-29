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

func TestAgentCommands_CreateTextDataCommand(t *testing.T) {
	ac, r, w := createMockAgentCommands(t)

	ac.isAuthorized = false
	ac.CreateTextData()
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true
	tdtMsg, tdmMsg := fmt.Sprintf("%s\n", enterTextData), fmt.Sprintf("%s\n", enterMetaInfo)
	mockGA := ac.grpcAgent.(*mocks.GrpcAgent)

	// корректная td запись
	// готовим инпут
	r.WriteString("text data number 1\n" + "meta info number 1\n")
	ac.CreateTextData()
	// проверяем вывод команды
	assert.Equal(t, tdtMsg+tdmMsg, w.String())

	// проверяем пришедшее в grpc агент
	gotTD, ok := mockGA.InputArgs[0].(storage.TextData)
	require.Equal(t, true, ok)
	wantTD := storage.TextData{TextData: "text data number 1", MetaInfo: "meta info number 1"}
	assert.Equal(t, wantTD, gotTD)

	// проверяем обновление стейта
	v, err := ac.state.Get(1, "text")
	require.NoError(t, err)
	vTD, ok := v.(storage.TextData)
	require.Equal(t, true, ok)
	wantTD.ID = 1
	assert.Equal(t, wantTD, vTD)

	// td запись, ошибка на вводе text data
	r.Reset()
	w.Reset()
	ac.CreateTextData()
	assert.Equal(t, tdtMsg+"err: EOF\n", w.String())

	// td запись, ошибка на вводе meta info
	r.Reset()
	w.Reset()
	r.WriteString("text data number 2\n")
	ac.CreateTextData()
	assert.Equal(t, tdtMsg+tdmMsg+"err: EOF\n", w.String())

	// td запись с ошибкой в grpc вызове
	r.Reset()
	w.Reset()
	r.WriteString("\n" + "\n")
	ac.CreateTextData()
	assert.Equal(t, tdtMsg+tdmMsg+"err: data invalid\n", w.String())
}

func TestAgentCommands_CreateBankDataCommand(t *testing.T) {
	ac, r, w := createMockAgentCommands(t)

	ac.isAuthorized = false
	ac.CreateBankData()
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true
	mockGA := ac.grpcAgent.(*mocks.GrpcAgent)

	// корректная bd запись
	// готовим инпут
	testCardNumber, testCardExpiry, testCVV := "0011223344556677", "09/23", "612"
	testMI := "meta info number 1"
	r.WriteString(fmt.Sprintf("%s %s %s\n%s\n", testCardNumber, testCardExpiry, testCVV, testMI))
	ac.CreateBankData()
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", enterBankData, enterMetaInfo), w.String())

	// проверяем пришедшее в grpc агент
	gotBD, ok := mockGA.InputArgs[0].(storage.BankData)
	require.Equal(t, true, ok)
	wantBD := storage.BankData{
		CardNumber: testCardNumber, CardExpire: testCardExpiry, CVV: testCVV, MetaInfo: testMI,
	}
	assert.Equal(t, wantBD, gotBD)

	// проверяем обновление стейта
	v, err := ac.state.Get(1, "bank")
	require.NoError(t, err)
	vBD, ok := v.(storage.BankData)
	require.Equal(t, true, ok)
	wantBD.ID = 1
	assert.Equal(t, wantBD, vBD)

	// bd запись, ошибка на вводе bank data
	r.Reset()
	w.Reset()
	r.WriteString("")
	ac.CreateBankData()
	assert.Equal(t, fmt.Sprintf("%s\nerr: EOF\n", enterBankData), w.String())

	// bd запись, ошибка на кол-ве введенных полей(меньше 3х)
	r.Reset()
	w.Reset()
	// CVV == 000 - спец. кейс для триггера ошибки
	r.WriteString(fmt.Sprintf("%s %s\n\n", testCardNumber, testCardExpiry))
	ac.CreateBankData()
	assert.Equal(t, fmt.Sprintf("%s\nerr: %s\n", enterBankData, "input error"), w.String())

	// bd запись, ошибка на вводе meta info
	ac, r, w = createMockAgentCommands(t)
	ac.isAuthorized = true
	r.WriteString(fmt.Sprintf("%s %s %s\n", testCardNumber, testCardExpiry, testCVV))
	ac.CreateBankData()
	assert.Equal(t, fmt.Sprintf("%s\n%s\nerr: EOF\n", enterBankData, enterMetaInfo), w.String())

	// bd запись с ошибкой в grpc вызове
	r.Reset()
	w.Reset()
	// CVV == 000 - спец. кейс для триггера ошибки
	r.WriteString(fmt.Sprintf("%s %s %s\n\n", testCardNumber, testCardExpiry, "000"))
	ac.CreateBankData()
	assert.Equal(t, fmt.Sprintf("%s\n%s\n%s\n", enterBankData, enterMetaInfo, "err: data invalid"), w.String())
}

func TestAgentCommands_CreateBinaryDataCommand(t *testing.T) {
	ac, r, w := createMockAgentCommands(t)

	ac.isAuthorized = false
	ac.CreateBinaryData()
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true
	mockGA := ac.grpcAgent.(*mocks.GrpcAgent)
	bindMI := "meta info number 1"
	bindTDir := "binary_test"

	// корректная bd запись
	// готовим инпут
	r.WriteString(fmt.Sprintf("%s\n%s\n", fmt.Sprintf("%s/%s", bindTDir, "binary_test.txt"), bindMI))
	ac.CreateBinaryData()
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", enterBinaryData, enterMetaInfo), w.String())

	// проверяем пришедшее в grpc агент
	gotBD, ok := mockGA.InputArgs[0].(storage.BinaryData)
	require.Equal(t, true, ok)
	wantBD := storage.BinaryData{BinaryData: []byte("Ayaka"), MetaInfo: bindMI}
	assert.Equal(t, wantBD, gotBD)

	// проверяем обновление стейта
	v, err := ac.state.Get(1, "binary")
	require.NoError(t, err)
	vBD, ok := v.(storage.BinaryData)
	require.Equal(t, true, ok)
	wantBD.ID = 1
	assert.Equal(t, wantBD, vBD)

	// bd запись, ошибка на вводе пути к файлу
	r.Reset()
	w.Reset()
	ac.CreateBinaryData()
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", enterBinaryData, "err: EOF"), w.String())

	// bd запись, ошибка пути к файлу(файл не существует)
	r.Reset()
	w.Reset()
	r.WriteString(fmt.Sprintf("%s/%s\n", bindTDir, "demo_file.txt"))
	ac.CreateBinaryData()
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", enterBinaryData, "err: file does not exist"), w.String())

	// bd запись, ошибка на вводе meta info
	r.Reset()
	w.Reset()
	r.WriteString(fmt.Sprintf("%s\n", fmt.Sprintf("%s/%s", bindTDir, "binary_test.txt")))
	ac.CreateBinaryData()
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\nerr: %s\n", enterBinaryData, enterMetaInfo, io.EOF), w.String())

	// bd запись, ошибка в grpc агенте
	r.Reset()
	w.Reset()
	r.WriteString(fmt.Sprintf("%s\n%s\n", fmt.Sprintf("%s/%s", bindTDir, "empty_test.txt"), bindMI))
	ac.CreateBinaryData()
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\nerr: %s\n", enterBinaryData, enterMetaInfo, "data invalid"), w.String())
}
