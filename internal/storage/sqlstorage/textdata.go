package sqlstorage

import (
	"context"
	"database/sql"
	"github.com/firesworder/password_saver/internal/crypt"
	"github.com/firesworder/password_saver/internal/storage"
)

type TextData struct {
	Conn *sql.DB

	Encoder *crypt.Encoder
	Decoder *crypt.Decoder
}

func (tr *TextData) AddTextData(ctx context.Context, td storage.TextData, u *storage.User) (int, error) {
	var id int
	var err error
	content, err := tr.Encoder.Encode([]byte(td.TextData))
	if err != nil {
		return 0, err
	}
	err = tr.Conn.QueryRowContext(ctx,
		"INSERT INTO textdata(text_data, meta_info, user_id) VALUES ($1, $2, $3) RETURNING id",
		content, td.MetaInfo, u.ID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tr *TextData) UpdateTextData(ctx context.Context, td storage.TextData, u *storage.User) error {
	content, err := tr.Encoder.Encode([]byte(td.TextData))
	if err != nil {
		return err
	}
	result, err := tr.Conn.ExecContext(ctx,
		`UPDATE textdata SET text_data = $1, meta_info = $2 WHERE id = $3 AND user_id = $4`,
		content, td.MetaInfo, td.ID, u.ID,
	)
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

func (tr *TextData) DeleteTextData(ctx context.Context, td storage.TextData, u *storage.User) error {
	result, err := tr.Conn.ExecContext(ctx,
		"DELETE FROM textdata WHERE id = $1 AND user_id = $2", td.ID, u.ID)
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

func (tr *TextData) GetAllRecords(ctx context.Context, u *storage.User) ([]storage.TextData, error) {
	result := make([]storage.TextData, 0)
	rows, err := tr.Conn.QueryContext(ctx,
		"SELECT id, text_data, meta_info, user_id FROM textdata WHERE user_id = $1", u.ID)
	if err != nil {
		return nil, nil
	}

	for rows.Next() {
		element := storage.TextData{}
		var content, contentDec []byte
		if err = rows.Scan(&element.ID, &content, &element.MetaInfo, &element.UserID); err != nil {
			return nil, err
		}
		if contentDec, err = tr.Decoder.Decode(content); err != nil {
			return nil, err
		}
		element.TextData = string(contentDec)
		result = append(result, element)
	}

	// проверяем на ошибки
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
