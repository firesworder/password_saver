package agentcommands

import (
	"github.com/firesworder/password_saver/internal/storage"
	"io"
	"os"
	"strings"
)

func (ac *AgentCommands) UpdateTextData(ID int) {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}

	var err error
	textData := storage.TextData{ID: ID}
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

	if err = ac.grpcAgent.UpdateTextDataRecord(textData); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	ac.state.Set(textData)
}

func (ac *AgentCommands) UpdateBankData(ID int) {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}

	var err error
	bankData := storage.BankData{ID: ID}
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

	err = ac.grpcAgent.UpdateBankDataRecord(bankData)
	if err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	ac.state.Set(bankData)
}

func (ac *AgentCommands) UpdateBinaryData(ID int) {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}

	binaryData := storage.BinaryData{ID: ID}
	ac.writer.WriteString(enterBinaryData)
	binaryFP, err := ac.reader.GetUserInput()
	if err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}

	f, err := os.Open(strings.TrimSpace(binaryFP))
	if err != nil {
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

	if err = ac.grpcAgent.UpdateBinaryDataRecord(binaryData); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	ac.state.Set(binaryData)
}
