package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/web"
	"github.com/shenyisyn/goft-gin/goft"
	"gjj/Boot"
	"gjj/config"
	"gjj/httphandle"
	"gjj/midwares"
	"log"
	"net/http"
)

type CourseHttpHandle struct {
}

func (c *CourseHttpHandle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("http api test"))
}

func main() {

	Boot.BootInit()

	g := goft.Ignite().
		Attach(midwares.NewCheckForReadyMid()).
		Mount("", httphandle.NewCourseHttpService())

	/*
		r := gin.New()
		r.Handle("GET", "/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"data": gin.H{"size": 2}})
		})
	*/

	fmt.Println("server starting")
	server := web.NewService(
		web.Name(config.JConfig.Service.Namespace+"."+config.JConfig.Service.Name),
		web.Handler(g),
	)
	err := server.Init()
	if nil != err {
		log.Fatal("fatal:", err)
	}

	fmt.Println("server init")
	//server.Handle("/httptest", &CourseHttpHandle{})
	/*
		server.HandleFunc("/httptest", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("http api test"))
		})
	*/

	fmt.Println("server run")
	log.Fatal(server.Run())
	fmt.Println("server started")
}
