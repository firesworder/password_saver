package agentcommands

import "github.com/firesworder/password_saver/internal/storage"

func (ac *AgentCommands) DeleteTextData(recordID int) {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}
	var err error
	if err = ac.grpcAgent.DeleteTextDataRecord(storage.TextData{ID: recordID}); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	if err = ac.state.Delete(recordID); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
}
func (ac *AgentCommands) DeleteBankData(recordID int) {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}
	var err error
	if err = ac.grpcAgent.DeleteBankDataRecord(storage.BankData{ID: recordID}); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	if err = ac.state.Delete(recordID); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
}
func (ac *AgentCommands) DeleteBinaryData(recordID int) {
	if !ac.isAuthorized {
		ac.writer.WriteErrorString(authReqErr)
		return
	}
	var err error
	if err = ac.grpcAgent.DeleteBinaryDataRecord(storage.BinaryData{ID: recordID}); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
	if err = ac.state.Delete(recordID); err != nil {
		ac.writer.WriteErrorString(err.Error())
		return
	}
}
