package dao

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SQLiteDAO struct {
	DB *sqlx.DB
}

func NewSQLiteDAO(db *sqlx.DB) *SQLiteDAO {
	return &SQLiteDAO{DB: db}
}

func (d *SQLiteDAO) CreateMessage(message string) error {
	query := `INSERT INTO first_service_message (message) VALUES (?)`
	_, err := d.DB.Exec(query, message)
	if err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}
	return nil
}
