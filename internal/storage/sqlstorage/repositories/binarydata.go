package repositories

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/storage"
)

type BinaryData struct {
	conn *sql.DB
}

func (br *BinaryData) AddBinaryData(ctx context.Context, bd storage.BinaryData) (int, error) {
	var id int

	err := br.conn.QueryRowContext(ctx,
		"INSERT INTO binarydata(binary_data, meta_info, user_id) VALUES ($1, $2, $3) RETURNING id",
		bd.BinaryData, bd.MetaInfo, bd.UserID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (br *BinaryData) UpdateBinaryData(ctx context.Context, bd storage.BinaryData) error {
	result, err := br.conn.ExecContext(ctx,
		`UPDATE binarydata SET binary_data = $1, meta_info = $2 WHERE id = $3`, bd.BinaryData, bd.MetaInfo, bd.ID)
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

func (br *BinaryData) DeleteBinaryData(ctx context.Context, bd storage.BinaryData) error {
	result, err := br.conn.ExecContext(ctx,
		"DELETE FROM binarydata WHERE id = $1", bd.ID)
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

// todo: добавить зависимость от пользователя(мб даже во всех методах!)
func (br *BinaryData) GetAllRecords(ctx context.Context) ([]storage.BinaryData, error) {
	result := make([]storage.BinaryData, 0)
	rows, err := br.conn.QueryContext(ctx, "SELECT id, binary_data, meta_info, user_id FROM binarydata")
	if err != nil {
		return nil, nil
	}

	for rows.Next() {
		element := storage.BinaryData{}
		if err = rows.Scan(&element.ID, &element.BinaryData, &element.MetaInfo, &element.UserID); err != nil {
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
