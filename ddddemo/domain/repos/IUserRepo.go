package repos

import "test/ddddemo/domain/models"

type IUserRepo interface {
	FindByName(name string) (*models.UserModel, error)
	SaveUser(*models.UserModel) error
}

type IUserLogRepo interface {
	FindByName(name string) *models.UserLogModel
	SaveLog(model models.UserLogModel) error
}
