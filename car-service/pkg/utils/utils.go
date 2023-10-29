package utils

import (
	"car-service/internal/storage/postgresql"
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrEmptyQuery = errors.New("ErrEmptyQuery")

type Response struct {
	Data interface{} `json:"data"`
	Err  *string     `json:"err"`
}

func ResponseFunc(c echo.Context, err error, i interface{}) {
	var status int
	if err != nil {
		errF := err.Error()
		if errors.Is(err, postgresql.ErrQueryRows) {
			status = http.StatusInternalServerError
		}
		if errors.Is(err, postgresql.ErrScanRows) {
			status = http.StatusInternalServerError
		}
		if errors.Is(err, sql.ErrNoRows) {
			status = http.StatusNotFound
		}
		if errors.Is(err, ErrEmptyQuery) {
			status = http.StatusUnprocessableEntity
		}
		c.JSON(status, Response{Data: nil, Err: &errF})
		return
	} else {
		status = http.StatusOK
	}

	c.JSON(status, Response{Data: i, Err: nil})
}
