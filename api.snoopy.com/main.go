package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"topic.snoopy.com/src/dao/topicDao"
	"topic.snoopy.com/src/model"
	"topic.snoopy.com/src/proto"
)

func main() {

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("topurl", model.TopicUrl)
	}

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("api/v1/topics")
	v1.GET("", func(context *gin.Context) {
		getTopic := &proto.GetTopicReq{}
		err := context.Bind(getTopic)
		if nil != err {
			context.String(http.StatusOK, "params error, error = %+v", err)
		} else {
			context.String(http.StatusOK, "get topis list, user name = %+v", getTopic.UserName)
		}
	})

	v1.GET("/:topic_id", topicDao.GetTopicDetail)

	v1.Use(topicDao.LoginMid())
	{
		v1.POST("/newTopic", topicDao.NewTopic)
	}

	router.Run()

}
