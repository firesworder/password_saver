package repositories

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/storage"
)

type TextData struct {
	conn *sql.DB
}

func (tr *TextData) AddTextData(ctx context.Context, td storage.TextData) (int, error) {
	var id int

	err := tr.conn.QueryRowContext(ctx,
		"INSERT INTO textdata(text_data, meta_info, user_id) VALUES ($1, $2, $3) RETURNING id",
		td.TextData, td.MetaInfo, td.UserID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TextData) UpdateTextData(ctx context.Context, td storage.TextData) error {
	result, err := tr.conn.ExecContext(ctx,
		`UPDATE textdata SET text_data = $1, meta_info = $2 WHERE id = $3`, td.TextData, td.MetaInfo, td.ID)
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

func (tr *TextData) DeleteTextData(ctx context.Context, td storage.TextData) error {
	result, err := tr.conn.ExecContext(ctx,
		"DELETE FROM textdata WHERE id = $1", td.ID)
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
func (tr *TextData) GetAllRecords(ctx context.Context) ([]storage.TextData, error) {
	result := make([]storage.TextData, 0)
	rows, err := tr.conn.QueryContext(ctx, "SELECT id, text_data, meta_info, user_id FROM textdata")
	if err != nil {
		return nil, nil
	}

	for rows.Next() {
		element := storage.TextData{}
		if err = rows.Scan(&element.ID, &element.TextData, &element.MetaInfo, &element.UserID); err != nil {
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
