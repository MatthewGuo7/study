package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"test/ddddemo/domain/models"
	"test/ddddemo/domain/repos"
)

var _ repos.IUserRepo = &UserRepo{}

type UserRepo struct {
	db *gorm.DB
}

func (u *UserRepo) FindByName(name string) (*models.UserModel, error) {
	ret := &models.UserModel{}
	d := u.db.Where("user_name = ?", name).Find(&ret)
	return ret, errors.WithStack(d.Error)
}

func (u *UserRepo) SaveUser(model *models.UserModel) error {
	return nil
}
