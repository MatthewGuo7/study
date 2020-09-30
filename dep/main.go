package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()
	r.Handle("GET", "/", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "this is my web"})
	})
	r.Run(":80")
}
