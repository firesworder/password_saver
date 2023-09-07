package repositories

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/storage"
)

type BinaryData struct {
	Conn *sql.DB
}

func (br *BinaryData) AddBinaryData(ctx context.Context, bd storage.BinaryData, u *storage.User) (int, error) {
	var id int

	err := br.Conn.QueryRowContext(ctx,
		"INSERT INTO binarydata(binary_data, meta_info, user_id) VALUES ($1, $2, $3) RETURNING id",
		bd.BinaryData, bd.MetaInfo, u.ID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (br *BinaryData) UpdateBinaryData(ctx context.Context, bd storage.BinaryData, u *storage.User) error {
	result, err := br.Conn.ExecContext(ctx,
		`UPDATE binarydata SET binary_data = $1, meta_info = $2 WHERE id = $3 AND user_id = $4`,
		bd.BinaryData, bd.MetaInfo, bd.ID, u.ID)
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

func (br *BinaryData) DeleteBinaryData(ctx context.Context, bd storage.BinaryData, u *storage.User) error {
	result, err := br.Conn.ExecContext(ctx,
		"DELETE FROM binarydata WHERE id = $1 AND user_id = $2", bd.ID, u.ID)
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

func (br *BinaryData) GetAllRecords(ctx context.Context, u *storage.User) ([]storage.BinaryData, error) {
	result := make([]storage.BinaryData, 0)
	rows, err := br.Conn.QueryContext(ctx,
		"SELECT id, binary_data, meta_info, user_id FROM binarydata WHERE user_id = $1", u.ID)
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
