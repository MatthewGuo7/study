package queryError

type QueryError struct {
	Query string
	//err   error
}

func (e *QueryError) Error() string {
	return e.Query
}

//func (e *QueryError) Unwrap() error {
//	return e.err
//}

func Get() error {
	return &QueryError{Query: "something", /*errors.New("query error")*/}
}
