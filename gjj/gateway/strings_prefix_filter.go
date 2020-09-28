package gateway

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	RegisteFilters("StripPrefix", NewStringsPrefixFilter())
}

type StringsPrefixFilter struct {
}

func (s *StringsPrefixFilter) Apply(config interface{}) GateWayFilter {
	return func(exchange *ServerWebExchange) {
		request := exchange.request

		path := request.URL.Path
		// /v1/course => /course
		defIndex := 1
		valueConf := ValueConfig(config.(string))
		temp, err := strconv.Atoi(valueConf.Get()[0])
		if nil != err {
			return
		}

		defIndex = temp

		fmt.Printf("def index = %+v", temp)

		pathList := strings.Split(path, "/")
		newPath := strings.Join(pathList[defIndex+1:], "/")
		request.URL.Path = newPath
		fmt.Printf("new path = %+v\n", newPath)
	}
}

func NewStringsPrefixFilter() *StringsPrefixFilter {
	return &StringsPrefixFilter{}
}
