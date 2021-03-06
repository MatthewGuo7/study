package Mapper

import (
	"github.com/jinzhu/gorm"
	"gjj/Boot"
	"gjj/def"
)

func GetCourseList() *gorm.DB {
	return Boot.GetDB().Table(def.TableName_CourseMain).
		Order("course_id desc").Limit(3)
}

const course_list="select * from course_main order by course_id desc limit ?"
func GetCourseListBySql(args ...interface{})  *gorm.DB{
	return Boot.GetDB().Raw(course_list,args...)
}