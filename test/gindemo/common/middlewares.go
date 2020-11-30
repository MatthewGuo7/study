package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("err = %+v", e)
				context.JSON(400, gin.H{"message": e})
			}
		}()

		context.Next()
	}
}
