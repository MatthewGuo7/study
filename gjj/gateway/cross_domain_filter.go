package gateway

import "net/http"

func init() {
	RegisteFilters("CrossDomain", NewCrossDomainFilter())
}

type CrossDomainFilter struct {
}

func (c *CrossDomainFilter) Apply(config interface{}) GateWayFilter {
	return func(exchange *ServerWebExchange) ResponseFilter {
		return func(response *http.Response) {
			response.Header.Add("Access-Control-Allow-Origin", "*")
			response.Header.Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			response.Header.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			response.Header.Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			response.Header.Add("Access-Control-Allow-Credentials", "true")
		}
	}
}

func NewCrossDomainFilter() *CrossDomainFilter {
	return &CrossDomainFilter{}
}
