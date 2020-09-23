package main

import (
	"github.com/micro/go-micro/v2"
	"gjj/Course"
	"log"
)

func main() {

	server := micro.NewService(
		micro.Name("go.micro.api.snoopy.course"),
		)

	server.Init()

	err := Course.RegisterCourseServiceHandler(server.Server(), Course.NewCourseImpl())
	if nil != err {
		log.Fatal(err)
	}

	log.Fatal(server.Run())

}
