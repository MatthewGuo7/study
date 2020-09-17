package main

import (
	"context"
	"etcddemo/register"
	"fmt"
	"log"
)

func main() {
	rc := register.NewClient([]string{"127.0.0.1:2379"})
	err := rc.LoadService()
	if nil != err {
		log.Fatal(err)
	}

	endFunc := rc.GetService("productservice", "GET", ProdEndPoint)
	resp, err := endFunc(context.Background(), Product{
		ID: 100,
	})
	if nil != err {
		log.Fatal(err)
	}

	fmt.Println(resp)

}
