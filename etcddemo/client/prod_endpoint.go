package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type Product struct {
	ID int
}

func ProdEndPoint(ctx context.Context, r *http.Request, params interface{}) error {

	req, ok := params.(Product)
	if !ok {
		return errors.New("params error")
	}

	r.URL.Path = "/prod/" + strconv.Itoa(req.ID)
	return nil
}
