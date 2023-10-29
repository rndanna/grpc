package postgresql

import (
	"car-service/internal/models"
	"database/sql"
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

func (db *DB) GetCar(carID string) (*models.Car, error) {
	var car models.Car

	if err := db.db.QueryRow(`
		SELECT id, name, price, description, brand, user_id, engine_id
		FROM cars
		WHERE id = $1
`, carID).Scan(&car.ID, &car.Name, &car.Price, &car.Description, &car.Brand, &car.UserID, &car.EngineID); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("err GetCar QueryRow: %w", ErrScanRows)
	}

	return &car, nil
}

func (db *DB) GetUserCars(userID string) ([]models.Car, error) {
	var cars []models.Car

	rows, err := db.db.Query(`
		SELECT id, name, price, description, brand, user_id, engine_id
		FROM cars
		WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("err getUserCars Query: %w", ErrQueryRows)
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car

		if err = rows.Scan(&car.ID, &car.Name, &car.Price, &car.Description, &car.Brand, &car.UserID, &car.EngineID); err != nil {
			return nil, fmt.Errorf("err getUserCars Scan: %w", ErrScanRows)
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (db *DB) GetCarByBrand(brand string) ([]models.Car, error) {
	var cars []models.Car

	rows, err := db.db.Query(`
		SELECT id, name, price, description, brand, user_id, engine_id
		FROM cars
		WHERE brand = $1
	`, brand)
	if err != nil {
		return nil, fmt.Errorf("err GetCarByBrand Query: %w", ErrQueryRows)
	}
	defer rows.Close()

	for rows.Next() {
		var car models.Car

		if err = rows.Scan(&car.ID, &car.Name, &car.Price, &car.Description, &car.Brand, &car.UserID, &car.EngineID); err != nil {
			return nil, fmt.Errorf("err GetCarByBrand Scan: %w", ErrScanRows)
		}
		cars = append(cars, car)
	}

	return cars, nil
}
