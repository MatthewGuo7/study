package topicDao

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"topic.snoopy.com/src/model"
)

func LoginMid() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, ok := context.GetQuery("token")
		if !ok {
			context.String(http.StatusUnauthorized, "need token")
			context.Abort()
			return
		}

		context.Next()
	}
}

func GetTopicDetail(context *gin.Context) {
	topicID := context.Query("topic_id")
	context.String(http.StatusOK, "id = %+v", topicID)
}

func NewTopic(context *gin.Context) {
	m := &model.Topic{}
	if err := context.Bind(m); nil != err {
		context.JSONP(http.StatusOK, gin.H{
			"code": http.StatusPaymentRequired,
			"msg":  err.Error(),
		})
	} else {
		context.String(http.StatusOK, "new topic")
	}
}
