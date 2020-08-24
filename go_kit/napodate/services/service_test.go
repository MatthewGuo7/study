package services

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func setUp() (Service, context.Context) {
	return newDateService(), context.TODO()
}

func Test_dateService_Get(t *testing.T) {
	Convey("test get", t, func() {

	})
}

func Test_dateService_Status(t *testing.T) {
	Convey("test get status", t, func() {
		s, context := setUp()
		ret, err := s.Status(context)
		fmt.Println(ret, err)
		So(err, ShouldBeNil)
	})
}

func Test_dateService_Validate(t *testing.T) {

}
