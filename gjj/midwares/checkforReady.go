package midwares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gjj/config"
)

type CheckForReadyMid struct {
}

func NewCheckForReadyMid() *CheckForReadyMid {
	return &CheckForReadyMid{}
}

func (c *CheckForReadyMid) OnRequest(context *gin.Context) error {
	if config.JConfig.DataConfig.MySql == nil {
		return fmt.Errorf("server is starting")
	}
	return nil
}

func (c *CheckForReadyMid) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}
