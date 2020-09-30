package main

import (
	"fmt"
	"gjj/gateway"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {

	/*
		reg := etcd.NewRegistry(registry.Addrs("localhost:2379"))
		serviceList, err := reg.GetService("go.micro.api.snoopy.http.course")
		if nil != err {
			log.Fatal(err)
		}

		for _, service := range serviceList {
			fmt.Println(service)
		}

		next := selector.Random(serviceList)
		node, err := next()
		if nil != err {
			log.Fatal(err)
		}

		fmt.Println(node)


		return
	*/

	configs := gateway.GetSysConfig()

	fmt.Printf("configs = %+v\n", configs)

	routers := configs.Routers
	if nil == routers {
		log.Fatal("routers is nil")
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if r := routers.IsMatch(request); nil != r {
			//remote, _ := url.Parse(r.Url)

			remote := r.GetAddress()

			//request.URL.Path = strings.Replace(request.URL.Path, "/v1", "", -1)

			exchange := gateway.NewServerWebExchange(request)

			//gateway.NewStringsPrefixFilter().Apply()(exchange)
			respFilters := r.FilterRequest(exchange)

			fmt.Printf("replaced url path = %+v", request.URL.Path)

			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.ModifyResponse = func(response *http.Response) error {
				respFilters.Filter(response)
				//response.Header.Add("name", "abc")
				return nil
			}

			proxy.ServeHTTP(writer, request)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}

	})
	fmt.Println("the gw is starting")
	http.ListenAndServe(":8080", nil)
}
