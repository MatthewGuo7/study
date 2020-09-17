package main

import (
	"context"
	"etcddemo/register"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/prod/{id:\\d+}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		str := "produc id = " + vars["id"]
		writer.Write([]byte(str))
	})

	serviceID := "p1"
	rc := register.NewRegisterService([]string{"127.0.0.1:2379"})
	err := rc.RegService(serviceID, "productservice", "127.0.0.1:8081")
	if nil != err {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println(rc.UnRegister(serviceID))
	}()

	httpServer := &http.Server{Addr: ":8081", Handler: route}

	go func() {
		err = httpServer.ListenAndServe()
		if nil != err {
			log.Fatal(err)
		}
		defer httpServer.Shutdown(context.Background())
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <- sig:
	}
}
