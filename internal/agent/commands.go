package agent

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
	"io"
	"log"
	"os"
)

func (a *Agent) RegisterUserCommand() {
	var err error
	var login, password string
	fmt.Println("Enter login and password separated by space")
	_, err = fmt.Scan(&login, &password)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = a.grpcAgent.RegisterUser(login, password)
	if err != nil {
		log.Println(err)
		return
	}
}

func (a *Agent) LoginUserCommand() {
	var err error
	var login, password string
	fmt.Println("Enter login and password separated by space")
	_, err = fmt.Scan(&login, &password)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = a.grpcAgent.LoginUser(login, password)
	if err != nil {
		log.Println(err)
		return
	}
}

// create commands

func (a *Agent) CreateRecordCommand() {
	var err error
	var dataType string
	fmt.Println("Choose data type(enter name type): text, bank or binary")
	_, err = fmt.Scan(&dataType)
	if err != nil {
		log.Println(err)
		return
	}

	switch dataType {
	case "text":
		a.CreateTextDataCommand()
	case "bank":
		a.CreateBankDataCommand()
	case "binary":
		a.CreateBinaryDataCommand()
	default:
		fmt.Println("unknown data type")
		return
	}
}

func (a *Agent) CreateTextDataCommand() {
	var err error
	var textData storage.TextData
	fmt.Println("Enter text data")
	_, err = fmt.Scan(&textData.TextData)
	if err != nil {
		log.Println(err)
		return
	}

	if textData.MetaInfo, err = scanMetaInfo(); err != nil {
		log.Println(err)
		return
	}

	textData.ID, err = a.grpcAgent.CreateTextDataRecord(textData)
	if err != nil {
		log.Println(err)
		return
	}
	a.state.set(textData)
}

func (a *Agent) CreateBankDataCommand() {
	var err error
	var bankData storage.BankData
	fmt.Println("Enter bank data separated by spaces: CardNumber(without spaces), CardExpiry, CVV")
	_, err = fmt.Scan(&bankData.CardNumber, &bankData.CardExpire, &bankData.CVV)
	if err != nil {
		log.Println(err)
		return
	}

	if bankData.MetaInfo, err = scanMetaInfo(); err != nil {
		log.Println(err)
		return
	}

	bankData.ID, err = a.grpcAgent.CreateBankDataRecord(bankData)
	if err != nil {
		log.Println(err)
		return
	}
	a.state.set(bankData)
}

func (a *Agent) CreateBinaryDataCommand() {
	var err error
	var binaryFP string
	var binaryData storage.BinaryData
	fmt.Println("Enter binary data filepath")
	_, err = fmt.Scan(&binaryFP)
	if err != nil {
		log.Println(err)
		return
	}

	f, err := os.Open(binaryFP)
	if err != nil {
		log.Println(err)
		return
	}
	binaryData.BinaryData, err = io.ReadAll(f)
	if err != nil {
		log.Println(err)
		return
	}

	if binaryData.MetaInfo, err = scanMetaInfo(); err != nil {
		log.Println(err)
		return
	}

	binaryData.ID, err = a.grpcAgent.CreateBinaryDataRecord(binaryData)
	if err != nil {
		log.Println(err)
		return
	}
	a.state.set(binaryData)
}

// open commands

func (a *Agent) OpenRecordCommand() {
	var err error
	var recordID int
	fmt.Println("Choose record ID")
	_, err = fmt.Scan(&recordID)
	if err != nil {
		log.Println(err)
		return
	}

	// todo: заглушка
	recordDT := "text"
	switch recordDT {
	case "text":
		a.OpenTextDataRecordCommand(recordID)
	case "bank":
		a.OpenBankDataRecordCommand(recordID)
	case "binary":
		a.OpenBinaryDataRecordCommand(recordID)
	}
}

func (a *Agent) OpenTextDataRecordCommand(ID int) {
	textDataExample := storage.TextData{
		ID:       ID,
		TextData: "Text data example",
		MetaInfo: "td1",
	}

	fmt.Println("Text content:")
	fmt.Println(textDataExample.TextData)
}

func (a *Agent) OpenBankDataRecordCommand(ID int) {
	textDataExample := storage.BankData{
		ID:         ID,
		CardNumber: "5566 7788 9900 1122",
		CardExpire: "12/23",
		CVV:        "465",
		MetaInfo:   "bd2",
	}

	fmt.Println("Bank content:")
	fmt.Printf("CardNumber: %s\n", textDataExample.CardNumber)
	fmt.Printf("CardExpiry: %s CVV:%s\n", textDataExample.CardNumber, textDataExample.CVV)
}

func (a *Agent) OpenBinaryDataRecordCommand(ID int) {
	textDataExample := storage.BinaryData{
		ID:         ID,
		BinaryData: []byte("Binary data"),
		MetaInfo:   "binD2",
	}

	var binaryFP string
	fmt.Println("Enter binary data filepath to save")
	_, err := fmt.Scan(&binaryFP)
	if err != nil {
		log.Println(err)
		return
	}

	f, err := os.Open(binaryFP)
	if err != nil {
		log.Println(err)
	}
	_, err = f.Write(textDataExample.BinaryData)
	if err != nil {
		log.Println(err)
	}
}

// update commands

func (a *Agent) UpdateRecordCommand() {
	var err error
	var recordID int
	fmt.Println("Choose record ID")
	_, err = fmt.Scan(&recordID)
	if err != nil {
		log.Println(err)
		return
	}
	record, err := a.state.get(recordID)
	if err != nil {
		log.Println(err)
		return
	}

	switch v := record.(type) {
	case storage.TextData:
		a.UpdateTextDataCommand(v.ID)
	case storage.BankData:
		a.UpdateBankDataCommand(v.ID)
	case storage.BinaryData:
		a.UpdateBinaryDataCommand(v.ID)
	}
}

func (a *Agent) UpdateTextDataCommand(ID int) {
	var err error
	textData := storage.TextData{ID: ID}
	fmt.Println("Enter text data")
	_, err = fmt.Scan(&textData.TextData)
	if err != nil {
		log.Println(err)
		return
	}

	if textData.MetaInfo, err = scanMetaInfo(); err != nil {
		log.Println(err)
		return
	}

	err = a.grpcAgent.UpdateTextDataRecord(textData)
	if err != nil {
		log.Println(err)
		return
	}
	a.state.set(textData)
}

func (a *Agent) UpdateBankDataCommand(ID int) {
	var err error
	bankData := storage.BankData{ID: ID}
	fmt.Println("Enter bank data separated by spaces: CardNumber(without spaces), CardExpiry, CVV")
	_, err = fmt.Scan(&bankData.CardNumber, &bankData.CardExpire, &bankData.CVV)
	if err != nil {
		log.Println(err)
		return
	}

	if bankData.MetaInfo, err = scanMetaInfo(); err != nil {
		log.Println(err)
		return
	}

	err = a.grpcAgent.UpdateBankDataRecord(bankData)
	if err != nil {
		log.Println(err)
		return
	}
	a.state.set(bankData)
}

func (a *Agent) UpdateBinaryDataCommand(ID int) {
	var err error
	var binaryFP string // todo: для вывода оставить? или в метаинфу
	binaryData := storage.BinaryData{ID: ID}
	fmt.Println("Enter binary data filepath")
	_, err = fmt.Scan(&binaryFP)
	if err != nil {
		log.Println(err)
		return
	}

	f, err := os.Open(binaryFP)
	if err != nil {
		log.Println(err)
		return
	}
	binaryData.BinaryData, err = io.ReadAll(f)
	if err != nil {
		log.Println(err)
		return
	}

	if binaryData.MetaInfo, err = scanMetaInfo(); err != nil {
		log.Println(err)
		return
	}

	err = a.grpcAgent.UpdateBinaryDataRecord(binaryData)
	if err != nil {
		log.Println(err)
		return
	}
	a.state.set(binaryData)
}

// other commands

func (a *Agent) DeleteRecordCommand() {
	var err error
	var recordID int
	fmt.Println("Choose record ID")
	_, err = fmt.Scan(&recordID)
	if err != nil {
		log.Println(err)
		return
	}
	record, err := a.state.get(recordID)
	if err != nil {
		log.Println(err)
		return
	}

	switch v := record.(type) {
	case storage.TextData:
		err = a.grpcAgent.DeleteTextDataRecord(storage.TextData{ID: v.ID})
	case storage.BankData:
		err = a.grpcAgent.DeleteBankDataRecord(storage.BankData{ID: v.ID})
	case storage.BinaryData:
		err = a.grpcAgent.DeleteBinaryDataRecord(storage.BinaryData{ID: v.ID})
	}
	if err = a.state.delete(recordID); err != nil {
		log.Println("element was not found")
		return
	}
}

func (a *Agent) ShowAllRecordsCommand() {
	currentState, err := a.grpcAgent.ShowAllRecords()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Text data records:")
	fmt.Printf("ID MetaInfo")
	for _, d := range currentState.TextDataList {
		fmt.Printf("%d %s\n", d.ID, d.MetaInfo)
	}

	fmt.Println("Bank data records:")
	fmt.Printf("ID MetaInfo")
	for _, d := range currentState.BankDataList {
		fmt.Printf("%d %s\n", d.ID, d.MetaInfo)
	}

	fmt.Println("Binary data records:")
	fmt.Printf("ID MetaInfo")
	for _, d := range currentState.BinaryDataList {
		fmt.Printf("%d %s\n", d.ID, d.MetaInfo)
	}
}

func (a *Agent) HelpCommand() {
	fmt.Print(`
	Commands:
	register_user, login_user
	create_record, open_record, update_record, delete_record
	show_all_records
	help
`)
}
