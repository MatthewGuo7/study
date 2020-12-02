package auth

import "errors"

type Request struct {
	User string
}

func AuthenticateRequest() error  {
	return auth()
}

func auth() error {
	return errors.New("auth error")
}
