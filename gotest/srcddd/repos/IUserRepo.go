package repos

import "srcddd/domain/models"

type IUserRepo interface {
	FindByName(name string) *models.UserModel
	SaveUser(model *models.UserModel) error
	UpdateUser(model *models.UserModel) error
	DeleteUser(model *models.UserModel) error
}

type IUserLogRepo interface {
	FindByName(name string) *models.UserLogModel
	SaveLog(model *models.UserLogModel) error
}
