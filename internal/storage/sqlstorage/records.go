package sqlstorage

import (
	"context"
	"database/sql"

	"github.com/firesworder/password_saver/internal/storage"
)

type RecordRepository struct {
	conn *sql.DB
}

func (rr *RecordRepository) AddRecord(ctx context.Context, r storage.Record, uid int) (int, error) {
	var recordID int

	err := rr.conn.QueryRowContext(ctx,
		"INSERT INTO records(record_type, content, meta_info, user_id) VALUES($1, $2, $3, $4) RETURNING id",
		r.RecordType, r.Content, r.MetaInfo, uid).Scan(&recordID)
	if err != nil {
		return 0, err
	}
	return recordID, nil
}

func (rr *RecordRepository) UpdateRecord(ctx context.Context, r storage.Record, uid int) error {
	result, err := rr.conn.ExecContext(ctx,
		"UPDATE records SET record_type = $1, content = $2, meta_info = $3 WHERE id = $4 AND user_id = $5",
		r.RecordType, r.Content, r.MetaInfo, r.ID, uid)
	if err != nil {
		return err
	}
	if rAff, err := result.RowsAffected(); err != nil || rAff == 0 {
		if rAff == 0 {
			return storage.ErrElementNotFound
		}
		return err
	}
	return nil
}

func (rr *RecordRepository) DeleteRecord(ctx context.Context, r storage.Record, uid int) error {
	result, err := rr.conn.ExecContext(ctx,
		"DELETE FROM records WHERE id = $1 and user_id = $2", r.ID, uid)
	if err != nil {
		return err
	}
	if rAff, err := result.RowsAffected(); err != nil || rAff == 0 {
		if rAff == 0 {
			return storage.ErrElementNotFound
		}
		return err
	}
	return nil
}

func (rr *RecordRepository) GetAll(ctx context.Context, uid int) ([]storage.Record, error) {
	result := make([]storage.Record, 0)
	rows, err := rr.conn.QueryContext(ctx,
		"SELECT id, record_type, content, meta_info FROM records WHERE user_id = $1", uid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r storage.Record
		if err = rows.Scan(&r.ID, &r.RecordType, &r.Content, &r.MetaInfo); err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
