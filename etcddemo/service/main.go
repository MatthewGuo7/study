package main

import (
	"context"
	"etcddemo/register"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"github.com/satori/go.uuid"
)

func main() {
	name := flag.String("name", "", "服务名称")
	port := flag.Int("p", 0, "服务端口")
	flag.Parse()
	if *name == "" {
		log.Fatal("请指定服务名")
	}
	if *port == 0 {
		log.Fatal("请指定端口")
	}

	route := mux.NewRouter()
	route.HandleFunc("/prod/{id:\\d+}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		str := "produc id = " + vars["id"]
		writer.Write([]byte(str))
	})

	serviceName := *name
	serviceAddr := "127.0.0.1"
	servicePort := *port
	serviceID := uuid.NewV4().String()

	rc := register.NewRegisterService([]string{"127.0.0.1:2379"})
	err := rc.RegService(serviceID, serviceName, serviceAddr+":"+strconv.Itoa(servicePort))
	if nil != err {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println(rc.UnRegister(serviceID))
	}()

	httpServer := &http.Server{Addr: ":" + strconv.Itoa(servicePort), Handler: route}

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
	case <-sig:
	}
}
