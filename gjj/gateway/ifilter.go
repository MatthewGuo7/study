package gateway

import "net/http"

type ServerWebExchange struct {
	request *http.Request
}

func NewServerWebExchange(request *http.Request) *ServerWebExchange {
	return &ServerWebExchange{request: request}
}

type IFilter interface {
	Apply() GateWayFilter
}

type GateWayFilter func(exchange *ServerWebExchange)
