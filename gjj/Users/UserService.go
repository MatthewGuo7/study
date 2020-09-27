package Users

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"gjj/Course"
)

type UserServiceImpl struct {
	client client.Client
}

func (u *UserServiceImpl) Test(ctx context.Context, request *UserRequest, resp *UserResp) error {
	resp.Id = "users" + request.Id
	c := Course.NewCourseService("go.micro.api.snoopy.course",  u.client)
	courseResp, err := c.ListForTop(ctx, &Course.CourseReq{Size: 1})
	if nil != err {
		return err
	}

	fmt.Println(courseResp)

	return nil
}

func NewUserServiceImpl(client client.Client) *UserServiceImpl {
	return &UserServiceImpl{client: client}
}
