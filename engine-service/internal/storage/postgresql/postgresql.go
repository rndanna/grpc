package postgresql

import (
	"database/sql"
	"engine-service/internal/models"
	"errors"
	"fmt"
)

var ErrQueryRows = errors.New("ErrQueryRows")
var ErrScanRows = errors.New("ErrScanRows")

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (db *DB) GetEngine(id int) (*models.Engine, error) {
	var engine models.Engine

	if err := db.db.QueryRow(`
		SELECT id, name, description 
		FROM engines
		WHERE id = $1
    `, id).Scan(&engine.ID, &engine.Name, &engine.Description); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("err GetEngine QueryRow: %w", ErrScanRows)
	}

	return &engine, nil
}
