package agentcommands

import (
	"errors"
	"github.com/firesworder/password_saver/internal/storage"
	"io"
	"io/fs"
	"os"
	"strings"
)

// CreateTextData создает текстовую запись на сервере.
// При успешной записи - добавляет созданную запись(с получ. от сервера ID) в стейт.
func (ac *AgentCommands) CreateTextData() {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}

	var textData storage.TextData
	var err error
	ac.writer.WriteString(enterTextData)
	if textData.TextData, err = ac.reader.GetUserInput(); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}

	ac.writer.WriteString(enterMetaInfo)
	if textData.MetaInfo, err = ac.reader.GetUserInput(); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}

	if textData.ID, err = ac.grpcAgent.CreateTextDataRecord(textData); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	ac.state.Set(textData)
}

// CreateBankData создает банковскую запись на сервере.
// При успешной записи - добавляет созданную запись(с получ. от сервера ID) в стейт.
func (ac *AgentCommands) CreateBankData() {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}

	var bankData storage.BankData
	ac.writer.WriteString(enterBankData)
	fields, err := ac.reader.GetUserFields()
	if err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	if len(fields) < 3 {
		ac.writer.WriteErrorString("input error")
		return
	}
	bankData.CardNumber, bankData.CardExpire, bankData.CVV = fields[0], fields[1], fields[2]

	ac.writer.WriteString(enterMetaInfo)
	if bankData.MetaInfo, err = ac.reader.GetUserInput(); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}

	if bankData.ID, err = ac.grpcAgent.CreateBankDataRecord(bankData); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	ac.state.Set(bankData)
}

// CreateBinaryData создает бинарную запись на сервере.
// На вход будет запрошен файл с бинарным содержимым(которое и будет сохранено на сервере).
// При успешной записи - добавляет созданную запись(с получ. от сервера ID) в стейт.
func (ac *AgentCommands) CreateBinaryData() {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}

	var binaryData storage.BinaryData
	ac.writer.WriteString(enterBinaryData)
	binaryFP, err := ac.reader.GetUserInput()
	if err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}

	f, err := os.Open(strings.TrimSpace(binaryFP))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			ac.writer.WriteErrorString("file does not exist")
			return
		}
		ac.writer.WriteErrorString(err.Error())
		return
	}

	if binaryData.BinaryData, err = io.ReadAll(f); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}

	ac.writer.WriteString(enterMetaInfo)
	if binaryData.MetaInfo, err = ac.reader.GetUserInput(); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}

	if binaryData.ID, err = ac.grpcAgent.CreateBinaryDataRecord(binaryData); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	ac.state.Set(binaryData)
}
