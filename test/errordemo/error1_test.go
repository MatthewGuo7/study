package errordemo

import (
	//"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"os"
	//"test/errordemo/dao"
	"test/errordemo/queryError"
	"testing"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if nil != err {
		return nil, errors.Wrapf(err, "open failed")
	}
	defer f.Close()

	return nil, nil
}

func ReadConfig() ([]byte, error) {
	config, err := ReadFile("test")
	return config, errors.WithMessage(err, "cound not read config")
}

func TestError1(t *testing.T) {
	_, err := ReadConfig()
	if nil != err {
		fmt.Printf("original error:%T, %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trance:\n %+v \n", err)
		os.Exit(-1)
	}
}

func GetData() error {

	err := queryError.Get()
	if nil != err {
		return errors.WithMessage(err, "decompress: %w")
	}

	return nil
}

var ErrPermission = errors.New("Denied")

func WrapErrTest() error {
	return fmt.Errorf("access denied :%w", ErrPermission)
}

func TestRootCause(t *testing.T) {
	//err := dao.GetUserByID()
	//if errors.Cause(err) == sql.ErrNoRows {
	//	fmt.Println("not found")
	//}
	//
	//if errors.Is(err, sql.ErrNoRows) {
	//	fmt.Println("not found")
	//}

	queryErr := GetData()
	queryRootError := errors.Cause(queryErr)
	fmt.Printf("error = %+v\n", queryRootError)

	switch tmpErr := queryRootError.(type) {
	case nil:
		fmt.Println("no error")
	case *queryError.QueryError:
		fmt.Printf("query error = %+v\n", tmpErr)
	default:
		fmt.Println("undefined")
	}

	var e *queryError.QueryError
	if errors.As(queryErr, &e) {
		fmt.Printf("%+v\n", queryErr)
		fmt.Printf("equal")
	}

	err := WrapErrTest()
	if errors.Is(err, ErrPermission) {
		fmt.Println("permission denied")
	}

}
