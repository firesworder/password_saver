package agentcommands

type IAgentCommands interface {
	RegisterUser()
	LoginUser()

	CreateTextData()
	CreateBankData()
	CreateBinaryData()

	OpenTextData(recordID int)
	OpenBankData(recordID int)
	OpenBinaryData(recordID int)
	ShowAllRecords()

	UpdateTextData(recordID int)
	UpdateBankData(recordID int)
	UpdateBinaryData(recordID int)

	DeleteTextData(recordID int)
	DeleteBankData(recordID int)
	DeleteBinaryData(recordID int)
}