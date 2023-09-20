package sqlstorage

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/crypt"
	"github.com/firesworder/password_saver/internal/storage"
	"strings"
)

// BankData репозиторий банковских данных в БД.
type BankData struct {
	Conn *sql.DB

	Encoder *crypt.Encoder
	Decoder *crypt.Decoder
}

// AddBankData добавляет банковскую запись пользователя.
func (br *BankData) AddBankData(ctx context.Context, bd storage.BankData, u *storage.User) (int, error) {
	var id int
	var err error
	content, err := br.Encoder.Encode([]byte(strings.Join([]string{bd.CardNumber, bd.CardExpire, bd.CVV}, ",")))
	if err != nil {
		return 0, err
	}
	err = br.Conn.QueryRowContext(ctx,
		`INSERT INTO bankdata(bank_data, meta_info, user_id) VALUES ($1, $2, $3) RETURNING id`,
		content, bd.MetaInfo, u.ID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateBankData обновляет банковскую запись пользователя.
func (br *BankData) UpdateBankData(ctx context.Context, bd storage.BankData, u *storage.User) error {
	content, err := br.Encoder.Encode([]byte(strings.Join([]string{bd.CardNumber, bd.CardExpire, bd.CVV}, ",")))
	if err != nil {
		return err
	}
	result, err := br.Conn.ExecContext(ctx,
		`UPDATE bankdata SET bank_data = $1, meta_info = $2 WHERE id = $3 AND user_id = $4`,
		content, bd.MetaInfo, bd.ID, u.ID)
	if err != nil {
		return err
	}
	if rAff, err := result.RowsAffected(); err != nil || rAff == 0 {
		if err != nil {
			return err
		} else {
			return storage.ErrElementNotFound
		}
	}
	return nil
}

// DeleteBankData удаляет банковскую запись пользователя.
func (br *BankData) DeleteBankData(ctx context.Context, bd storage.BankData, u *storage.User) error {
	result, err := br.Conn.ExecContext(ctx,
		"DELETE FROM bankdata WHERE id = $1 AND user_id=$2", bd.ID, u.ID)
	if err != nil {
		return err
	}
	if rAff, err := result.RowsAffected(); err != nil || rAff == 0 {
		if err != nil {
			return err
		} else {
			return storage.ErrElementNotFound
		}
	}
	return nil
}

// GetAllRecords возвращает все банковские данные пользователя.
func (br *BankData) GetAllRecords(ctx context.Context, u *storage.User) ([]storage.BankData, error) {
	result := make([]storage.BankData, 0)
	rows, err := br.Conn.QueryContext(ctx,
		"SELECT id, bank_data, meta_info, user_id FROM bankdata WHERE user_id=$1", u.ID)
	if err != nil {
		return nil, nil
	}

	for rows.Next() {
		element := storage.BankData{}
		var content, contentDec []byte
		if err = rows.Scan(&element.ID, &content, &element.MetaInfo, &element.UserID); err != nil {
			return nil, err
		}
		if contentDec, err = br.Decoder.Decode(content); err != nil {
			return nil, err
		}
		strDec := strings.Split(string(contentDec), ",")
		element.CardNumber, element.CardExpire, element.CVV = strDec[0], strDec[1], strDec[2]
		result = append(result, element)
	}

	// проверяем на ошибки
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
