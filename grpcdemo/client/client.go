package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcdemo/services"
	"log"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if nil != err {
		log.Fatalf("conn error, error = %+v", err)
	}

	client := services.NewProductServiceClient(conn)
	req := services.ProductReq{
		Area:   services.ProductArea_B,
		ProdId: 1,
	}
	resp, err := client.GetProductStock(context.Background(), &req)

	fmt.Printf("resp = %+v, err = %+v", resp, err)

}
