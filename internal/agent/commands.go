package agent

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
	"io"
	"log"
	"os"
)

func (a *Agent) registerUserCommand() {
	var err error
	var login, password string
	fmt.Println("Enter login and password separated by space")
	_, err = fmt.Scan(&login, &password)
	if err != nil {
		log.Println(err)
		return
	}

	err = a.grpcAgent.RegisterUser(login, password)
	if err != nil {
		log.Println(err)
		return
	}
	a.isAuth = true
}

func (a *Agent) loginUserCommand() {
	var err error
	var login, password string
	fmt.Println("Enter login and password separated by space")
	_, err = fmt.Scan(&login, &password)
	if err != nil {
		log.Println(err)
		return
	}

	err = a.grpcAgent.LoginUser(login, password)
	if err != nil {
		log.Println(err)
		return
	}
	a.isAuth = true
}

// create commands

func (a *Agent) createRecordCommand() {
	if !a.isAuth {
		log.Println("auth required")
		return
	}

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
		a.createTextDataCommand()
	case "bank":
		a.createBankDataCommand()
	case "binary":
		a.createBinaryDataCommand()
	default:
		fmt.Println("unknown data type")
		return
	}
}

func (a *Agent) createTextDataCommand() {
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

func (a *Agent) createBankDataCommand() {
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

func (a *Agent) createBinaryDataCommand() {
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

func (a *Agent) openRecordCommand() {
	if !a.isAuth {
		log.Println("auth required")
		return
	}
	var err error
	var recordID int
	var dataType string
	fmt.Println("Enter recordID and dataType")
	_, err = fmt.Scan(&recordID, &dataType)
	if err != nil {
		log.Println(err)
		return
	}
	record, err := a.state.get(recordID, dataType)
	if err != nil {
		log.Println(err)
		return
	}

	switch v := record.(type) {
	case storage.TextData:
		fmt.Printf("Text data record:")
		fmt.Printf("ID: %d\n", v.ID)
		fmt.Printf("Content: %s\n", v.TextData)
		fmt.Printf("MetaInfo: %s\n", v.MetaInfo)
	case storage.BankData:
		fmt.Printf("Bank data record:")
		fmt.Printf("ID: %d\n", v.ID)
		fmt.Printf("CardNumber: %s\n", v.CardNumber)
		fmt.Printf("CardExpiry: %s | CVV: %sn", v.CardExpire, v.CVV)
		fmt.Printf("MetaInfo: %s\n", v.MetaInfo)
	case storage.BinaryData:
		fmt.Printf("Binary data record:")
		fmt.Printf("ID: %d\n", v.ID)
		fmt.Printf("Content: %s\n", v.BinaryData)
		fmt.Printf("MetaInfo: %s\n", v.MetaInfo)
	}
}

// update commands

func (a *Agent) updateRecordCommand() {
	if !a.isAuth {
		log.Println("auth required")
		return
	}
	var err error
	var recordID int
	var dataType string
	fmt.Println("Enter recordID and dataType")
	_, err = fmt.Scan(&recordID, &dataType)
	if err != nil {
		log.Println(err)
		return
	}
	record, err := a.state.get(recordID, dataType)
	if err != nil {
		log.Println(err)
		return
	}

	switch v := record.(type) {
	case storage.TextData:
		a.updateTextDataCommand(v.ID)
	case storage.BankData:
		a.updateBankDataCommand(v.ID)
	case storage.BinaryData:
		a.updateBinaryDataCommand(v.ID)
	}
}

func (a *Agent) updateTextDataCommand(ID int) {
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

func (a *Agent) updateBankDataCommand(ID int) {
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

func (a *Agent) updateBinaryDataCommand(ID int) {
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

func (a *Agent) deleteRecordCommand() {
	if !a.isAuth {
		log.Println("auth required")
		return
	}
	var err error
	var recordID int
	var dataType string
	fmt.Println("Enter recordID and dataType")
	_, err = fmt.Scan(&recordID, &dataType)
	if err != nil {
		log.Println(err)
		return
	}
	record, err := a.state.get(recordID, dataType)
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

func (a *Agent) showAllRecordsCommand() {
	if !a.isAuth {
		log.Println("auth required")
		return
	}
	currentState, err := a.grpcAgent.ShowAllRecords()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Text data records:")
	fmt.Println("ID MetaInfo")
	for _, d := range currentState.TextDataList {
		a.state.set(d)
		fmt.Printf("%d %s\n", d.ID, d.MetaInfo)
	}

	fmt.Println("Bank data records:")
	fmt.Println("ID MetaInfo")
	for _, d := range currentState.BankDataList {
		a.state.set(d)
		fmt.Printf("%d %s\n", d.ID, d.MetaInfo)
	}

	fmt.Println("Binary data records:")
	fmt.Println("ID MetaInfo")
	for _, d := range currentState.BinaryDataList {
		a.state.set(d)
		fmt.Printf("%d %s\n", d.ID, d.MetaInfo)
	}
}

func (a *Agent) helpCommand() {
	fmt.Print(`Commands:
Auth methods:
- register_user, login_user

User data methods(required auth!):
create_record, open_record, update_record, delete_record
show_all_records
`)
}
