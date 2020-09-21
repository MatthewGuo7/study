package aggregates

import (
	"srcddd/domain/models"
	"srcddd/repos"
)

type Member struct {
	User *models.UserModel
	Log *models.UserLogModel
	userRepo repos.IUserRepo
	userLogRepo repos.IUserLogRepo
}

func NewMember(user *models.UserModel, log *models.UserLogModel, userRepo *repos.IUserRepo, userLogRepo *repos.IUserLogRepo) *Member {
	return &Member{User: user, Log: log, userRepo: userRepo, userLogRepo: userLogRepo}
}

func (m *Member) Create() error  {
	err := m.userRepo.SaveUser(m.User)
	if nil != err {
		return err
	}

	return nil
}


