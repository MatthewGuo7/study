package Course

import "context"

func NewCourseModel(courseId int32, courseName string) *CourseModel {
	return &CourseModel{CourseId: courseId, CourseName: courseName}
}

type CourseImpl struct {
}

func (c *CourseImpl) ListForTop(ctx context.Context, req *CourseReq, resp *CourseResp) error {

	courses := []*CourseModel{
		&CourseModel{CourseId: 2, CourseName: "java"},
		&CourseModel{CourseId: 4, CourseName: "go"},
	}
	resp.Courses = courses
	return nil
}

func NewCourseImpl() *CourseImpl {
	return &CourseImpl{}
}
