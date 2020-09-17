package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcdemo/services"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if nil != err {
		log.Fatalf("conn error, error = %+v", err)
	}

	client := services.NewProductServiceClient(conn)
	req := services.ProductsReq{
		QuerySize: 1,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.GetProducts(ctx, &req, grpc.EmptyCallOption{})

	fmt.Printf("resp = %+v, err = %+v", resp, err)
}
