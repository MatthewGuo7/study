package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/gindemo/common"
	"test/gindemo/models/UserModel"
	"test/gindemo/result"
)

func main() {
	r := gin.New()
	r.Use(common.ErrorHandler())

	r.Handle("GET", "/test", func(context *gin.Context) {
		u := UserModel.New()
		result.Result(context.ShouldBind(u)).Unwrap()
		context.JSON(http.StatusOK, u)
	})

	r.Run()
}
