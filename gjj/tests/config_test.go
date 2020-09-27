package tests

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gjj/Boot"
	"gjj/config"
	"net/http"
	"testing"
	"time"
)

func CheckForReady() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.JConfig.DataConfig.MySql == nil {
			c.AbortWithStatusJSON(200, gin.H{"res": "server is starting"})
		} else {
			c.Next()
		}
	}
}

func waitForReady() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	for {

		select {
		case <-ctx.Done():
			return fmt.Errorf("init config error")
		default:
			if config.JConfig.DataConfig.MySql != nil {
				return nil
			}
		}
	}
}

func TestLoadConfig(t *testing.T) {

	Boot.BootInit()

	errChan := make(chan error)

	r := gin.New()
	r.Use(CheckForReady())
	r.Handle("GET", "/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": config.JConfig.DataConfig.MySql})
	})

	go func() {
		err := r.Run()
		if nil != err {
			errChan <- err
		}
	}()

	err := <-errChan
	fmt.Println(err)
}

func TestConfigWithGin(t *testing.T) {

}
