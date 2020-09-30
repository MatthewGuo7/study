package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"gjj/Boot"
	"gjj/Course"
	"gjj/service"
	"log"
	"time"
)

func main() {

	Boot.BootInit()

	server := micro.NewService(
		micro.Name("go.micro.api.snoopy.course"),
	)

	server.Init()

	time.Sleep(time.Second * 5)

	fmt.Println(Boot.GetDB())

	err := Course.RegisterCourseServiceHandler(server.Server(), service.NewCourseImpl(Boot.GetDB()))
	if nil != err {
		log.Fatal(err)
	}

	log.Fatal(server.Run())

}
