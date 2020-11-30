package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

func GetUserByID() error {
	return errors.Wrapf(sql.ErrNoRows, "get user by id")
}
