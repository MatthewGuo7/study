package service

import (
	"context"
	"gjj/Users"
)

type UserService struct {
}

func (u *UserService) Test(ctx context.Context, request *Users.UserRequest, resp *Users.UserResp) error {
	resp.Id = "users" + request.Id
	return nil
}

func NewUserService() *UserService {
	return &UserService{}
}
