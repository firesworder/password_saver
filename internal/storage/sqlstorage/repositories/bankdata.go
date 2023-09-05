package repositories

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/storage"
)

type BankData struct {
	conn *sql.DB
}

func (br *BankData) AddBankData(ctx context.Context, bd storage.BankData) (int, error) {
	var id int

	err := br.conn.QueryRowContext(ctx,
		`INSERT INTO bankdata(card_number, card_expiry, cvv, meta_info, user_id) VALUES ($1, $2, $3, $4, $5) 
                                                                        RETURNING id`,
		bd.CardNumber, bd.CardExpire, bd.CVV, bd.MetaInfo, bd.UserID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (br *BankData) UpdateBankData(ctx context.Context, bd storage.BankData) error {
	result, err := br.conn.ExecContext(ctx,
		`UPDATE bankdata SET card_number = $1, card_expiry = $2, cvv = $3, meta_info = $4 WHERE id = $5`,
		bd.CardNumber, bd.CardExpire, bd.CVV, bd.MetaInfo, bd.ID)
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

func (br *BankData) DeleteBankData(ctx context.Context, bd storage.BankData) error {
	result, err := br.conn.ExecContext(ctx,
		"DELETE FROM bankdata WHERE id = $1", bd.ID)
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

func (br *BankData) GetAllRecords(ctx context.Context) ([]storage.BankData, error) {
	result := make([]storage.BankData, 0)
	rows, err := br.conn.QueryContext(ctx,
		"SELECT id, card_number, card_expiry, cvv, meta_info, user_id FROM bankdata")
	if err != nil {
		return nil, nil
	}

	for rows.Next() {
		element := storage.BankData{}
		if err = rows.Scan(&element.ID, &element.CardNumber, &element.CardExpire, &element.CVV, &element.MetaInfo,
			&element.UserID); err != nil {
			return nil, err
		}
		result = append(result, element)
	}

	// проверяем на ошибки
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
