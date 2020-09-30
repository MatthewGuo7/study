package gateway

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

var filters sync.Map

func RegisteFilters(name string, f IFilter) {
	filters.Store(name, f)
}

type ServerWebExchange struct {
	request *http.Request
}

func NewServerWebExchange(request *http.Request) *ServerWebExchange {
	return &ServerWebExchange{request: request}
}

type IFilter interface {
	Apply(config interface{}) GateWayFilter
}

type ResponseFilter func(response *http.Response)
type GateWayFilter func(exchange *ServerWebExchange)ResponseFilter

type ResponseFilters []ResponseFilter

func (r ResponseFilters) Filter(response *http.Response) {
	for _, filter := range r {
		filter(response)
	}
}

type SimpleFilter string

func (s SimpleFilter) Filter() GateWayFilter {
	str := string(s)
	vlist := strings.Split(str, "=")
	if len(vlist) != 2 {
		return nil
	}
	key := strings.Trim(vlist[0], " ")
	value := strings.Trim(vlist[1], " ")
	fmt.Printf("key = '%+v',value = '%+v'", key, value)
	f, ok := filters.Load(key)
	if ok {
		return f.(IFilter).Apply(value)
	}
	return nil
}

func (r Router) FilterRequest(exchange *ServerWebExchange) ResponseFilters {
	retFilters := make(ResponseFilters, 0)
	for _, f := range r.Filters {
		v := reflect.ValueOf(f)
		if v.Kind() == reflect.String {
			fmt.Println(v.String())
			gwFilter := SimpleFilter(strings.Trim(v.String(), " ")).Filter()
			if nil != gwFilter {
				responseFilter := gwFilter(exchange)
				if nil != responseFilter {
					retFilters = append(retFilters, responseFilter)
				}
			}
		}
	}

	return retFilters
}
