package agentcommands

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/firesworder/password_saver/internal/storage"
)

func TestAgentCommands_OpenTextData(t *testing.T) {
	ac, _, w := createMockAgentCommands(t)
	testTD := storage.TextData{ID: 50, TextData: "text data", MetaInfo: "MI 1"}
	ac.state.Set(testTD)

	ac.isAuthorized = false
	ac.OpenTextData(100)
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true

	// корректное открытие
	ac.OpenTextData(testTD.ID)
	// проверяем вывод команды
	assert.Equal(t,
		fmt.Sprintf("Text data record:\nID: %d\nContent: %s\nMetaInfo: %s\n\n",
			testTD.ID, testTD.TextData, testTD.MetaInfo),
		w.String(),
	)

	// td запись, ошибка отсутствия элемента в стейте
	w.Reset()
	ac.OpenTextData(100)
	// проверяем вывод команды
	assert.Equal(t, "err: record was not found\n", w.String())
}

func TestAgentCommands_OpenBankData(t *testing.T) {
	ac, _, w := createMockAgentCommands(t)
	testBD := storage.BankData{ID: 50, CardNumber: "0011223344556677", CardExpire: "09/24", CVV: "555", MetaInfo: "MI 1"}
	ac.state.Set(testBD)

	ac.isAuthorized = false
	ac.OpenBankData(100)
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true

	// корректное открытие
	ac.OpenBankData(testBD.ID)
	// проверяем вывод команды
	assert.Equal(t,
		fmt.Sprintf("Bank data record:\nID: %d\nCardNumber: %s\nCardExpiry: %s | CVV: %s\nMetaInfo: %s\n\n",
			testBD.ID, testBD.CardNumber, testBD.CardExpire, testBD.CVV, testBD.MetaInfo),
		w.String(),
	)

	// td запись, ошибка отсутствия элемента в стейте
	w.Reset()
	ac.OpenBankData(100)
	// проверяем вывод команды
	assert.Equal(t, "err: record was not found\n", w.String())
}

func TestAgentCommands_OpenBinaryData(t *testing.T) {
	ac, r, w := createMockAgentCommands(t)
	testBD := storage.BinaryData{ID: 50, BinaryData: []byte("binary data"), MetaInfo: "MI 1"}
	ac.state.Set(testBD)
	hint1, hint2 := "Enter filepath to save binary content", "writing complete"
	tmpFile := "temp_binary.txt"

	ac.isAuthorized = false
	ac.OpenBinaryData(100)
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true

	// корректное открытие
	r.WriteString(tmpFile + "\n")
	ac.OpenBinaryData(testBD.ID)
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", hint1, hint2), w.String())

	content, err := os.ReadFile(tmpFile)
	require.NoError(t, err)
	assert.Equal(t, testBD.BinaryData, content)
	require.NoError(t, os.Remove(tmpFile))

	// bd запись, ошибка пути к файлу
	r.Reset()
	w.Reset()
	ac.OpenBinaryData(testBD.ID)
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", hint1, "err: EOF"), w.String())

	// bd запись, ошибка из grpc агента
	r.Reset()
	w.Reset()
	r.WriteString(tmpFile + "\n")
	ac.OpenBinaryData(100)
	// проверяем вывод команды
	assert.Equal(t, fmt.Sprintf("%s\n%s\n", hint1, "err: record was not found"), w.String())
	require.NoError(t, os.Remove(tmpFile))
}

func TestAgentCommands_showAllRecordsCommand(t *testing.T) {
	ac, _, w := createMockAgentCommands(t)
	testTD := storage.TextData{ID: 1, TextData: "Aranara", MetaInfo: "meta info", UserID: 1}
	testTD2 := storage.TextData{ID: 10, TextData: "Ararakalari", MetaInfo: "meta info", UserID: 1}
	testBankD := storage.BankData{
		ID: 2, CardNumber: "0011223344556677", CardExpire: "12/23", CVV: "453", MetaInfo: "meta info", UserID: 1,
	}
	testBinD := storage.BinaryData{ID: 3, BinaryData: []byte("Aranara"), MetaInfo: "meta info", UserID: 1}

	// без авторизации
	ac.isAuthorized = false
	ac.ShowAllRecords()
	assert.Equal(t, fmt.Sprintf("err: %s\n", authReqErr), w.String())
	w.Reset()

	// с авторизацией
	ac.isAuthorized = true
	ac.ShowAllRecords()

	wantResult := fmt.Sprintf("%s\n%s\n", "Text data records:", "ID MetaInfo")
	wantResult += fmt.Sprintf("%d %s\n", testTD.ID, testTD.MetaInfo)
	wantResult += fmt.Sprintf("%d %s\n", testTD2.ID, testTD2.MetaInfo)

	wantResult += fmt.Sprintf("%s\n%s\n", "Bank data records:", "ID MetaInfo")
	wantResult += fmt.Sprintf("%d %s\n", testBankD.ID, testBankD.MetaInfo)

	wantResult += fmt.Sprintf("%s\n%s\n", "Binary data records:", "ID MetaInfo")
	wantResult += fmt.Sprintf("%d %s\n", testBinD.ID, testBinD.MetaInfo)

	assert.Equal(t, wantResult, w.String())
}
