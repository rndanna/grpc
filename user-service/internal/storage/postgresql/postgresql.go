package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"user-service/internal/models"
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

func (db *DB) GetUser(userID int) (*models.User, error) {
	var user models.User

	if err := db.db.QueryRow(`
		SELECT id, username, email
		FROM users
		WHERE id = $1
	`, userID).Scan(&user.ID, &user.Username, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("err GetUser QueryRow: %w", ErrScanRows)

	}

	return &user, nil
}
