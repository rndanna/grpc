package handle

import (
	"fmt"
	"strconv"
	"user-service/internal/models"
	"user-service/internal/storage/postgresql"
	"user-service/pkg/utils"

	"github.com/labstack/echo/v4"
)

type handle struct {
	db *postgresql.DB
}

func New(db *postgresql.DB) *handle {
	return &handle{
		db: db,
	}
}

func (h *handle) GetUser(c echo.Context) (err error) {
	var user *models.User

	defer func() {
		utils.ResponseFunc(c, err, user)
	}()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return fmt.Errorf("failed Atoi: %w", err)
	}
	if id <= 0 {
		return fmt.Errorf("id < 0")
	}

	user, err = h.db.GetUser(id)
	if err != nil {
		err = fmt.Errorf("err GetUser : %w", err)
	}

	return
}
