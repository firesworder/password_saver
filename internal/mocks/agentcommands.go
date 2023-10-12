package mocks

type AgentCommands struct {
	LastFuncCalled string
}

func (a *AgentCommands) RegisterUser() {
	a.LastFuncCalled = "RegisterUser"
}

func (a *AgentCommands) LoginUser() {
	a.LastFuncCalled = "LoginUser"
}

func (a *AgentCommands) CreateTextData() {
	a.LastFuncCalled = "CreateTextData"
}

func (a *AgentCommands) CreateBankData() {
	a.LastFuncCalled = "CreateBankData"
}

func (a *AgentCommands) CreateBinaryData() {
	a.LastFuncCalled = "CreateBinaryData"
}

func (a *AgentCommands) OpenTextData(recordID int) {
	a.LastFuncCalled = "OpenTextData"
}

func (a *AgentCommands) OpenBankData(recordID int) {
	a.LastFuncCalled = "OpenBankData"
}

func (a *AgentCommands) OpenBinaryData(recordID int) {
	a.LastFuncCalled = "OpenBinaryData"
}

func (a *AgentCommands) ShowAllRecords() {
	a.LastFuncCalled = "ShowAllRecords"
}

func (a *AgentCommands) UpdateTextData(recordID int) {
	a.LastFuncCalled = "UpdateTextData"
}

func (a *AgentCommands) UpdateBankData(recordID int) {
	a.LastFuncCalled = "UpdateBankData"
}

func (a *AgentCommands) UpdateBinaryData(recordID int) {
	a.LastFuncCalled = "UpdateBinaryData"
}

func (a *AgentCommands) DeleteTextData(recordID int) {
	a.LastFuncCalled = "DeleteTextData"
}

func (a *AgentCommands) DeleteBankData(recordID int) {
	a.LastFuncCalled = "DeleteBankData"
}

func (a *AgentCommands) DeleteBinaryData(recordID int) {
	a.LastFuncCalled = "DeleteBinaryData"
}
