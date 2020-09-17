package endpoint

import (
	"context"
	"net/http"
)

type EndPoint func(ctx context.Context, requestParam interface{})(response interface{}, err error)

type EndRequestFunc func(ctx context.Context, r *http.Request, params interface{}) error
