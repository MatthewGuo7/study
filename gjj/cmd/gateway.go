package main

import (
	"fmt"
	"gjj/gateway"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	routers := gateway.InitConfig()
	if nil == routers {
		log.Fatal("routers is nil")
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if r := routers.IsMatch(request); nil != r {
			remote, _ := url.Parse(r.Url)

			//request.URL.Path = strings.Replace(request.URL.Path, "/v1", "", -1)

			exchange := gateway.NewServerWebExchange(request)

			gateway.NewStringsPrefixFilter().Apply()(exchange)

			fmt.Printf("replaced url path = %+v", request.URL.Path)

			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.ServeHTTP(writer, request)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}

	})
	fmt.Println("the gw is starting")
	http.ListenAndServe(":8080", nil)
}
