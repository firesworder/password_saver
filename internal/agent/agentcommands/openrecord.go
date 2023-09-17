package agentcommands

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
	"os"
	"strings"
)

func (ac *AgentCommands) OpenTextData(recordID int) {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}

	v, err := ac.state.Get(recordID, "text")
	if err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	td := v.(storage.TextData)

	ac.writer.WriteString(
		fmt.Sprintf("Text data record:\nID: %d\nContent: %s\nMetaInfo: %s\n", td.ID, td.TextData, td.MetaInfo),
	)
}
func (ac *AgentCommands) OpenBankData(recordID int) {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}

	v, err := ac.state.Get(recordID, "bank")
	if err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	bd := v.(storage.BankData)

	ac.writer.WriteString(
		fmt.Sprintf("Bank data record:\nID: %d\nCardNumber: %s\nCardExpiry: %s | CVV: %s\nMetaInfo: %s\n",
			bd.ID, bd.CardNumber, bd.CardExpire, bd.CVV, bd.MetaInfo,
		),
	)
}
func (ac *AgentCommands) OpenBinaryData(recordID int) {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}

	ac.writer.WriteString("Enter filepath to save binary content")
	fp, err := ac.reader.GetUserInput()
	if err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}

	f, err := os.OpenFile(strings.TrimSpace(fp), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}

	v, err := ac.state.Get(recordID, "binary")
	if err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	bd := v.(storage.BinaryData)

	if _, err = f.Write(bd.BinaryData); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	ac.writer.WriteString("writing complete")
}

func (ac *AgentCommands) ShowAllRecords() {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString("auth required")
		return
	}

	currentState, err := ac.grpcAgent.ShowAllRecords()
	if err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}

	ac.writer.WriteString("Text data records:")
	ac.writer.WriteString("ID MetaInfo")
	for _, d := range currentState.TextDataList {
		ac.state.Set(d)
		ac.writer.WriteString(fmt.Sprintf("%d %s", d.ID, d.MetaInfo))
	}

	ac.writer.WriteString("Bank data records:")
	ac.writer.WriteString("ID MetaInfo")
	for _, d := range currentState.BankDataList {
		ac.state.Set(d)
		ac.writer.WriteString(fmt.Sprintf("%d %s", d.ID, d.MetaInfo))
	}

	ac.writer.WriteString("Binary data records:")
	ac.writer.WriteString("ID MetaInfo")
	for _, d := range currentState.BinaryDataList {
		ac.state.Set(d)
		ac.writer.WriteString(fmt.Sprintf("%d %s", d.ID, d.MetaInfo))
	}
}
