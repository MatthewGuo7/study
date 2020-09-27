package gateway

import (
	"fmt"
	"strings"
)

type StringsPrefixFilter struct {
}

func (s *StringsPrefixFilter) Apply() GateWayFilter {
	return func(exchange *ServerWebExchange) {
		request := exchange.request

		path := request.URL.Path
		// /v1/course => /course
		pathList := strings.Split(path, "/")
		newPath := strings.Join(pathList[2:], "/")
		request.URL.Path = newPath
		fmt.Printf("new path = %+v\n", newPath)
	}
}

func NewStringsPrefixFilter() *StringsPrefixFilter {
	return &StringsPrefixFilter{}
}
