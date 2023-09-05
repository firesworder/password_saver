package bankdata

import (
	"context"
	"errors"
	"fmt"
	"github.com/firesworder/password_saver/internal/storage"
	"log"
	"regexp"
)

var cardNumberRE, cardExpiryRE, cvvRE *regexp.Regexp

func init() {
	var err error
	cardNumberRE, err = regexp.Compile(`^\d{4}\s*\d{4}\s*\d{4}\s*\d{4}$`)
	if err != nil {
		log.Fatal(err)
	}
	cardExpiryRE, err = regexp.Compile(`^\d{2}/\d{2}$`)
	if err != nil {
		log.Fatal(err)
	}
	cvvRE, err = regexp.Compile(`^\d{3}$`)
	if err != nil {
		log.Fatal(err)
	}
}

var ErrNotFound = errors.New("element not found")
var ErrDataInvalid = errors.New("bank data invalid")

type MockBankData struct {
	BankData   map[int]storage.BankData
	LastUsedID int
}

func (m *MockBankData) AddBankData(ctx context.Context, bd storage.BankData) (int, error) {
	if err := checkBankData(bd); err != nil {
		return 0, err
	}

	m.LastUsedID++
	bd.ID = m.LastUsedID
	m.BankData[bd.ID] = bd
	return bd.ID, nil
}

func (m *MockBankData) UpdateBankData(ctx context.Context, bd storage.BankData) error {
	if err := checkBankData(bd); err != nil {
		return err
	}
	if _, ok := m.BankData[bd.ID]; !ok {
		return ErrNotFound
	}
	m.BankData[bd.ID] = bd
	return nil
}

func (m *MockBankData) DeleteBankData(ctx context.Context, bd storage.BankData) error {
	if _, ok := m.BankData[bd.ID]; !ok {
		return ErrNotFound
	}
	delete(m.BankData, bd.ID)
	return nil
}

func checkBankData(bd storage.BankData) error {
	if bd.CardNumber == "" || bd.CardExpire == "" || bd.CVV == "" {
		return errors.Join(ErrDataInvalid, fmt.Errorf("none of bank data field can be empty"))
	}

	if !cardNumberRE.MatchString(bd.CardNumber) {
		return errors.Join(ErrDataInvalid, fmt.Errorf("card number format is not valid"))
	}
	if !cardExpiryRE.MatchString(bd.CardExpire) {
		return errors.Join(ErrDataInvalid, fmt.Errorf("card expiry format is not valid"))
	}
	if !cvvRE.MatchString(bd.CVV) {
		return errors.Join(ErrDataInvalid, fmt.Errorf("cvv format is not valid"))
	}

	return nil
}

func (m *MockBankData) GetAllRecords(ctx context.Context) ([]storage.BankData, error) {
	result := make([]storage.BankData, 0, len(m.BankData))
	for _, v := range m.BankData {
		result = append(result, v)
	}
	return result, nil
}
