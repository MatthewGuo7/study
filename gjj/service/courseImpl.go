package service

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"gjj/Course"
	"gjj/def"
)

type CourseImpl struct {
	db *gorm.DB
}

func (c *CourseImpl) GetCourseDetail(ctx context.Context, req *Course.CourseDetailReq, resp *Course.CourseDetailResp) error {
	fmt.Printf("req = %+v", req)
	resCourese := &Course.CourseModel{}
	err := c.db.Table(def.TableName_CourseMain).Where("course_id = ?", req.CourseId).Find(resCourese).Error
	resp.CourseDetail = resCourese
	resp.CourseCounts = []*Course.CourseCount{
		&Course.CourseCount{
			CountId:    1,
			CourseId:   1,
			CountKey:   "click",
			CountValue: 20,
		},
	}
	return err
}

func NewCourseImpl(db *gorm.DB) *CourseImpl {
	return &CourseImpl{db: db.New()}
}

func (c *CourseImpl) ListForTop(ctx context.Context, req *Course.CourseListReq, resp *Course.CourseListResp) error {
	fmt.Println("req = ", req)
	courses := make([]*Course.CourseModel, 0)
	db := c.db.Table(def.TableName_CourseMain).Order("course_id desc")

	err := db.Offset((req.Page-1) * req.Size_).Limit(req.Size_).Find(&courses).Error
	if nil != err {
		return err
	}

	totalCount := 0
	err = db.Count(&totalCount).Error
	if nil != err {
		return err
	}

	resp.PageInfo = BuildPageInfo(req.Page, req.Size_, int32(totalCount))
	resp.Courses = courses
	return nil
}

func BuildPageInfo(page int32,size int32,count int32 ) *Course.PageInfo {
	pageInfo:=&Course.PageInfo{PageNo:page,PageSize:size,TotalCount:count}
	var pageNum int32=1
	if count>size{
		pageNum=count/size
		if count%size>0{
			pageNum++
		}
	}
	pageInfo.TotalPage=pageNum
	return pageInfo
}
