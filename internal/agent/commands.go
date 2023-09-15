package agent

import (
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
	"io"
	"os"
	"strconv"
	"strings"
)

func (a *Agent) scanMetaInfo() (string, error) {
	a.writeString("Enter meta info")
	input, err := a.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func (a *Agent) registerUserCommand() {
	a.writeString("Enter login and password separated by space")
	input, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString("input error")
		return
	}
	fields := strings.Fields(input)
	if len(fields) != 2 {
		a.writeErrorString("input error")
		return
	}
	login, password := fields[0], fields[1]

	err = a.grpcAgent.RegisterUser(login, password)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	a.isAuth = true
}

func (a *Agent) loginUserCommand() {
	a.writeString("Enter login and password separated by space")
	input, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	fields := strings.Fields(input)
	if len(fields) != 2 {
		a.writeErrorString("input error")
		return
	}
	login, password := fields[0], fields[1]

	err = a.grpcAgent.LoginUser(login, password)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	a.isAuth = true
}

// create commands
func (a *Agent) createRecordCommand() {
	if !a.isAuth {
		a.writeErrorString("auth required")
		return
	}
	a.writeString("Choose data type(enter name type): text, bank or binary")
	dataType, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}

	switch strings.TrimSpace(dataType) {
	case "text":
		a.createTextDataCommand()
	case "bank":
		a.createBankDataCommand()
	case "binary":
		a.createBinaryDataCommand()
	default:
		a.writeErrorString("unknown data type")
		return
	}
}

func (a *Agent) createTextDataCommand() {
	var textData storage.TextData
	a.writeString("Enter text data")
	input, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	textData.TextData = strings.TrimSpace(input)

	if textData.MetaInfo, err = a.scanMetaInfo(); err != nil {
		a.writeErrorString(err.Error())
		return
	}

	textData.ID, err = a.grpcAgent.CreateTextDataRecord(textData)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	a.state.set(textData)
}

func (a *Agent) createBankDataCommand() {
	var bankData storage.BankData
	a.writeString("Enter bank data separated by spaces: CardNumber(without spaces), CardExpiry, CVV")
	input, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	fields := strings.Fields(input)
	if len(fields) != 3 {
		a.writeErrorString("input error")
		return
	}
	bankData.CardNumber, bankData.CardExpire, bankData.CVV = fields[0], fields[1], fields[2]

	if bankData.MetaInfo, err = a.scanMetaInfo(); err != nil {
		a.writeErrorString(err.Error())
		return
	}

	bankData.ID, err = a.grpcAgent.CreateBankDataRecord(bankData)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	a.state.set(bankData)
}

func (a *Agent) createBinaryDataCommand() {
	var binaryData storage.BinaryData
	a.writeString("Enter binary data filepath")
	binaryFP, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}

	f, err := os.Open(strings.TrimSpace(binaryFP))
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	binaryData.BinaryData, err = io.ReadAll(f)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}

	if binaryData.MetaInfo, err = a.scanMetaInfo(); err != nil {
		a.writeErrorString(err.Error())
		return
	}

	binaryData.ID, err = a.grpcAgent.CreateBinaryDataRecord(binaryData)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	a.state.set(binaryData)
}

// open commands

func (a *Agent) openRecordCommand() {
	if !a.isAuth {
		a.writeErrorString("auth required")
		return
	}

	a.writeString("Enter recordID and dataType")
	input, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	fields := strings.Fields(input)
	if len(fields) != 2 {
		a.writeErrorString("input error")
		return
	}
	recordID, err := strconv.Atoi(fields[0])
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	dataType := fields[1]

	record, err := a.state.get(recordID, dataType)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}

	switch v := record.(type) {
	case storage.TextData:
		output := fmt.Sprintf(`Text data record:
ID: %d
Content: %s
MetaInfo: %s
`, v.ID, v.TextData, v.MetaInfo)
		a.writeString(output)
	case storage.BankData:
		output := fmt.Sprintf(`Bank data record:
ID: %d
CardNumber: %s
CardExpiry: %s | CVV: %s
MetaInfo: %s
`, v.ID, v.CardNumber, v.CardExpire, v.CVV, v.MetaInfo)
		a.writeString(output)
	case storage.BinaryData:
		a.writeString("Enter filepath to save binary content")
		fp, err := a.reader.ReadString('\n')
		if err != nil {
			a.writeErrorString(err.Error())
			return
		}

		f, err := os.OpenFile(strings.TrimSpace(fp), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			a.writeErrorString(err.Error())
			return
		}
		if _, err = f.Write(v.BinaryData); err != nil {
			a.writeErrorString(err.Error())
			return
		}
		a.writeString("writing complete")
	}
}

// update commands

func (a *Agent) updateRecordCommand() {
	if !a.isAuth {
		a.writeErrorString("auth required")
		return
	}

	a.writeString("Enter recordID and dataType")
	input, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	fields := strings.Fields(input)
	if len(fields) != 2 {
		a.writeErrorString("input error")
		return
	}
	recordID, err := strconv.Atoi(fields[0])
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	dataType := fields[1]

	_, err = fmt.Scan(&recordID, &dataType)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	record, err := a.state.get(recordID, dataType)
	if err != nil {
		a.writeErrorString(err.Error())
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
	a.writeString("Enter text data")
	input, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	textData.TextData = strings.TrimSpace(input)

	if textData.MetaInfo, err = a.scanMetaInfo(); err != nil {
		a.writeErrorString(err.Error())
		return
	}

	err = a.grpcAgent.UpdateTextDataRecord(textData)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	a.state.set(textData)
}

func (a *Agent) updateBankDataCommand(ID int) {
	var err error
	bankData := storage.BankData{ID: ID}
	a.writeString("Enter bank data separated by spaces: CardNumber(without spaces), CardExpiry, CVV")
	input, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	fields := strings.Fields(input)
	if len(fields) != 3 {
		a.writeErrorString("input error")
		return
	}
	bankData.CardNumber, bankData.CardExpire, bankData.CVV = fields[0], fields[1], fields[2]

	if bankData.MetaInfo, err = a.scanMetaInfo(); err != nil {
		a.writeErrorString(err.Error())
		return
	}

	err = a.grpcAgent.UpdateBankDataRecord(bankData)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	a.state.set(bankData)
}

func (a *Agent) updateBinaryDataCommand(ID int) {
	var err error
	binaryData := storage.BinaryData{ID: ID}
	a.writeString("Enter binary data filepath")
	binaryFP, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}

	f, err := os.Open(strings.TrimSpace(binaryFP))
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	binaryData.BinaryData, err = io.ReadAll(f)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}

	if binaryData.MetaInfo, err = a.scanMetaInfo(); err != nil {
		a.writeErrorString(err.Error())
		return
	}

	err = a.grpcAgent.UpdateBinaryDataRecord(binaryData)
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	a.state.set(binaryData)
}

// other commands

func (a *Agent) deleteRecordCommand() {
	if !a.isAuth {
		a.writeErrorString("auth required")
		return
	}

	a.writeString("Enter recordID and dataType")
	input, err := a.reader.ReadString('\n')
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	fields := strings.Fields(input)
	if len(fields) != 2 {
		a.writeErrorString("input error")
		return
	}
	recordID, err := strconv.Atoi(fields[0])
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}
	dataType := fields[1]

	record, err := a.state.get(recordID, dataType)
	if err != nil {
		a.writeErrorString(err.Error())
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
		a.writeErrorString("element was not found")
		return
	}
}

func (a *Agent) showAllRecordsCommand() {
	if !a.isAuth {
		a.writeErrorString("auth required")
		return
	}
	currentState, err := a.grpcAgent.ShowAllRecords()
	if err != nil {
		a.writeErrorString(err.Error())
		return
	}

	a.writeString("Text data records:")
	a.writeString("ID MetaInfo")
	for _, d := range currentState.TextDataList {
		a.state.set(d)
		a.writeString(fmt.Sprintf("%d %s", d.ID, d.MetaInfo))
	}

	a.writeString("Bank data records:")
	a.writeString("ID MetaInfo")
	for _, d := range currentState.BankDataList {
		a.state.set(d)
		a.writeString(fmt.Sprintf("%d %s", d.ID, d.MetaInfo))
	}

	a.writeString("Binary data records:")
	a.writeString("ID MetaInfo")
	for _, d := range currentState.BinaryDataList {
		a.state.set(d)
		a.writeString(fmt.Sprintf("%d %s", d.ID, d.MetaInfo))
	}
}

func (a *Agent) helpCommand() {
	a.writeString(`Commands:
Auth methods:
- register_user, login_user

User data methods(required auth!):
- create_record, open_record, update_record, delete_record
- show_all_records
`)
}
