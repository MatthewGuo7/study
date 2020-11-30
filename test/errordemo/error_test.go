package errordemo

import (
	"errors"
	"fmt"
	"testing"
)

var NotFoundErr = errors.New("Not Found")
var ErrStructType = errors.New("EOF")

func NewError() error {
	return fmt.Errorf("new error:%+v", NotFoundErr)
}

type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%+v:%+v:%+v", e.File, e.Line, e.Msg)
}

func test() error {
	return &MyError{"some thing happened", "server.go", 42}
}

func TestError(t *testing.T) {
	err := test()
	switch err := err.(type) {
	case nil:
		fmt.Println("succ")
	case *MyError:
		fmt.Println("myerror", err)
	default:
		fmt.Println("unknown error")
	}

	//log.Println(auth.AuthenticateRequest())

}
