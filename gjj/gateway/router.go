package gateway

import (
	"gjj/gateway/predicates"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type Predicates struct {
	Header predicates.HeaderPredicate
	Method []string
	Host   string
	Path   predicates.PathPredicate
}

type Filter struct {
}

type Router struct {
	Id         string
	Namespace  string
	Url        string
	Predicates Predicates
	Filters    []interface{}
}

func (r *Router) GetAddress() *url.URL {
	discovery := GetSysConfig().Discovery
	remote, err := url.Parse(r.Url)
	if nil != err {
		return nil
	}

	if remote.Scheme == discovery.Loadbalancer.SchemePrefix {
		serviceName := remote.Host
		if r.Namespace != "" {
			serviceName = r.Namespace + "." + remote.Host
		}
		node := discovery.GetServiceNode(serviceName)
		if nil == node {
			panic("error service")
		}
		r, _ := url.Parse("http://" + node.Address)
		return r
	} else {
		return remote
	}

}

func (r *Router) IsMatch(request *http.Request) bool {

	/*
		if !r.Predicates.Path.IsMatch(request.URL.Path) {
			return false
		}

		if !r.Predicates.Header.IsMatch(request.Header) {
			return false
		}
	*/

	v := reflect.ValueOf(r.Predicates)
	for i := 0; i < v.NumField(); i++ {
		if matcher, ok := v.Field(i).Interface().(predicates.PredicateMatcher); ok &&
			strings.Trim(v.Field(i).String(), " ") != "" {
			if !matcher.IsMatch(request) {
				return false
			}
		}
	}

	return true
}

type Routers []*Router

func (r Routers) IsMatch(request *http.Request) *Router {

	for _, router := range r {
		if router.IsMatch(request) {
			return router
		}
	}

	return nil
}
