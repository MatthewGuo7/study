package main

import (
	"fmt"
	"net/http"
)

type Proxy struct {

}

func (p *Proxy) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	troundTrpip := http.DefaultTransport
	outReq := &http.Request{}
	*outReq = *request
}

func main() {

	fmt.Println("servier start")
	http.Handle("/", &Proxy{})
	http.ListenAndServe(":8080", nil)
}
