package httphandle

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"gjj/Boot"
	"gjj/Course"
	"gjj/service"
)

type CourseHttpService struct {
}

func (c *CourseHttpService) GetCourse(ginContext *gin.Context) goft.Json {
	/*
	size := ginContext.Param("size")
	sizeInt, _ := strconv.ParseInt(size, 10, 32)
	req := Course.CourseListReq{Size_: int32(sizeInt)}
	courseImpl := service.NewCourseImpl(Boot.GetDB())
	resp := &Course.CourseListResp{}
	err := courseImpl.ListForTop(context.Background(), &req, resp)
	if nil != err {
		fmt.Println(err)
		return gin.H{"code":0, "msg": "get data error"}
	}
	 */


	req := &Course.CourseListReq{}
	courseImpl := service.NewCourseImpl(Boot.GetDB())
	resp := &Course.CourseListResp{}

	err := ginContext.BindJSON(req)
	if nil != err {
		fmt.Println(err)
		return gin.H{"code":0, "msg": err.Error()}
	}

	err = courseImpl.ListForTop(context.Background(), req, resp)
	if nil != err {
		fmt.Println(err)
		return gin.H{"code":0, "msg": "get data error"}
	}

	return resp
}

func (c *CourseHttpService)GetCourseDetail(ginContext *gin.Context) goft.Json  {
	/*
	course_id := ginContext.Param("course_id")
	courseIdInt, _ := strconv.ParseInt(course_id, 10,32)
	fmt.Println(course_id)
	req := &Course.CourseDetailReq{CourseId: int32(courseIdInt)}
	courseImpl := service.NewCourseImpl(Boot.GetDB())
	resp := &Course.CourseDetailResp{}
	err := courseImpl.GetCourseDetail(context.Background(), req, resp)
	if nil != err {
		fmt.Println(err)
		return gin.H{"code":0, "msg": "get data error"}
	}
	 */

	req := &Course.CourseDetailReq{}
	err := ginContext.BindUri(req)
	if nil != err {
		fmt.Println(err)
		return gin.H{"code":0, "msg": "get data error"}
	}

	err = ginContext.BindHeader(req)
	if nil != err {
		fmt.Println(err)
		return gin.H{"code":0, "msg": "get data error"}
	}

	resp := &Course.CourseDetailResp{}
	courseImpl := service.NewCourseImpl(Boot.GetDB())
	err = courseImpl.GetCourseDetail(context.Background(), req, resp)
	if nil != err {
		fmt.Println(err)
		return gin.H{"code":0, "msg": "get data error"}
	}

	return resp
}

func (c *CourseHttpService) Build(goft *goft.Goft) {
	goft.Handle("GET", "/courses", c.GetCourse)
	goft.Handle("GET", "/course/:course_id", c.GetCourseDetail)
}

func (c *CourseHttpService) Name() string {
	return "CourseHttpService"
}

func NewCourseHttpService() *CourseHttpService {
	return &CourseHttpService{}
}
