package interceptor

import (
	"context"
	"fmt"
	. "google.golang.org/grpc"
)

var ServerIntercept UnaryServerInterceptor = func(ctx context.Context, req interface{}, info *UnaryServerInfo,
	handler UnaryHandler) (resp interface{}, err error) {

	defer func() {
		if nil != err {
			resp, err = handler(ctx, req)
			return
		}
	}()

	err = check(ctx, req)
	if nil != err {
		return
	}

	return handler(ctx, req)
}

func check(ctx context.Context, req interface{}) error {
	fmt.Printf("this is req = %+v\n", req)
	return nil
}


