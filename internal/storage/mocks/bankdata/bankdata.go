package bankdata

import (
	"context"
	"errors"
	"fmt"
	"log"
	"password_saver/internal/storage"
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
	bankData   map[int]storage.BankData
	lastUsedID int
}

func (m *MockBankData) AddBankData(ctx context.Context, bd storage.BankData) error {
	if err := checkBankData(bd); err != nil {
		return err
	}

	m.lastUsedID++
	bd.ID = m.lastUsedID
	m.bankData[bd.ID] = bd
	return nil
}

func (m *MockBankData) UpdateBankData(ctx context.Context, bd storage.BankData) error {
	if err := checkBankData(bd); err != nil {
		return err
	}
	if _, ok := m.bankData[bd.ID]; !ok {
		return ErrNotFound
	}
	m.bankData[bd.ID] = bd
	return nil
}

func (m *MockBankData) DeleteBankData(ctx context.Context, bd storage.BankData) error {
	if _, ok := m.bankData[bd.ID]; !ok {
		return ErrNotFound
	}
	delete(m.bankData, bd.ID)
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
