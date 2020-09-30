package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"jaegerdemo/utils"
	"time"
)

func main() {
	r := gin.New()
	utils.InitJaeger()
	r.Handle("GET","/", func(ginContext *gin.Context) {
		orderSpan := opentracing.StartSpan("创建订单API")
		defer orderSpan.Finish()

		spanCtx := opentracing.ContextWithSpan(context.Background(), orderSpan)

		orderNo := time.Now().Nanosecond()
		orderSpan.SetTag("orderno", orderNo)
		fmt.Println(orderNo)
		time.Sleep(time.Millisecond * 100)
		if GetOrder(spanCtx, "获取订单") {
			ginContext.JSON(200, gin.H{"msg": "success"})
		} else {
			orderSpan.SetTag("status", "error")

			ginContext.JSON(200, gin.H{"msg": "failed"})
		}
	})

	r.Run(":8080")
}

func GetOrder(ctx context.Context, name string)  bool {

	span, ctx := opentracing.StartSpanFromContext(ctx, name)
	defer span.Finish()

	time.Sleep(time.Millisecond * 200)

	if GetStock(ctx,"检查库存") > 0 {
		return true
	}

	return false
}

func GetStock(ctx context.Context, name string) int  {
	span, ctx := opentracing.StartSpanFromContext(ctx, name)
	defer span.Finish()

	time.Sleep(time.Millisecond * 300)

	return 0
}
