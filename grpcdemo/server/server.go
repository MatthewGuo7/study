package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpcdemo/services"
	"log"
	"net/http"
)

func main() {
	srv := grpc.NewServer()
	services.RegisterProductServiceServer(srv, &services.ProductService{})

	//l, err := net.Listen("tcp", ":8082")
	//if nil != err {
	//	log.Fatal(err)
	//}

	//srv.Serve(l)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request)
		srv.ServeHTTP(writer, request)
	})

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	err := httpServer.ListenAndServe()
	if nil != err {
		log.Fatal(err)
	}

}
