package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpcdemo/services"
	"log"
	"net"
)

func main() {

	//creds, err := credentials.NewServerTLSFromFile("pem/server.crt", "pem/server.key")
	//if nil != err {
	//	log.Fatal(err)
	//}

	//srv := grpc.NewServer(grpc.Creds(creds))
	srv := grpc.NewServer()
	services.RegisterProductServiceServer(srv, &services.ProductService{})

	reflection.Register(srv)

	l, err := net.Listen("tcp", ":8081")
	if nil != err {
		log.Fatal(err)
	}

	srv.Serve(l)

	/*
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("req = %+v", request)
		srv.ServeHTTP(writer, request)
	})

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	//err = httpServer.ListenAndServeTLS("pem/server.crt", "pem/server.key")
	err := httpServer.ListenAndServe()
	if nil != err {
		log.Fatal(err)
	}
	 */
}
