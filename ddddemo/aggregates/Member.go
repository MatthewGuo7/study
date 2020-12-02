package aggregates

import (
	"test/ddddemo/domain/models"
	"test/ddddemo/domain/repos"
)

type Member struct {
	User    *models.UserModel
	UserLog *models.UserLogModel
	userRepo *repos.IUserRepo
	userLogRepo *repos.IUserLogRepo
}

func NewMember(user *models.UserModel, userRepo *repos.IUserRepo, userLogRepo *repos.IUserLogRepo) *Member {
	return &Member{User: user, userRepo: userRepo, userLogRepo: userLogRepo}
}






