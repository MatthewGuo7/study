package dao

import (
	"gorm.io/gorm"
	"srcddd/domain/models"
)

type UserRepo struct {
	db *gorm.DB
}

func (u *UserRepo) FindByName(name string) *models.UserModel {
	panic("implement me")
}

func (u *UserRepo) SaveUser(model *models.UserModel) error {
	panic("implement me")
}

func (u *UserRepo) UpdateUser(model *models.UserModel) error {
	panic("implement me")
}

func (u *UserRepo) DeleteUser(model *models.UserModel) error {
	panic("implement me")
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}
