package main

import (
	"github.com/micro/go-micro/v2"
	"gjj/Users"
	"log"
)

func main() {
	server := micro.NewService(
		micro.Name("go.micro.api.snoopy.user"),
		)
	server.Init()

	err := Users.RegisterUserServiceHandler(server.Server(), Users.NewUserServiceImpl())
	if nil != err {
		log.Fatal(err)
	}

	log.Fatal(server.Run())


}
