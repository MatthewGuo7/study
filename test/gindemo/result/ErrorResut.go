package result

import "fmt"

type ErrorResult struct {
	err error
}

func (r *ErrorResult) Unwrap() interface{} {
	if r.err != nil {
		fmt.Println(r.err)
		panic(r.err)
	}

	return nil
}

func Result(err error) *ErrorResult {
	return &ErrorResult{
		err: err,
	}
}
