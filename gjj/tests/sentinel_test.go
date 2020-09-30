package tests

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/gin-gonic/gin"
	"log"
	"testing"
	"time"
)

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
)

func printTest(i int) {
	fmt.Println(i)
}

func TestSentinel(t *testing.T) {
	err := sentinel.InitWithConfigFile("sentinel.yaml")
	if nil != err {
		log.Fatal(err)
	}

	_, err = flow.LoadRules([]*flow.Rule{
		{
			ID:                0,
			Resource:          "abc",
			MetricType:        flow.QPS,
			Count:             2,
			ControlBehavior:   flow.Throttling,
			MaxQueueingTimeMs: uint32(time.Second * 2),
		},
	})
	if nil != err {
		log.Fatal(err)
	}

	/*
		for i := 0; i < 100; i++ {
			e, b := sentinel.Entry("abc")
			if nil != b {
				log.Println("out of num = ", b)
				time.Sleep(time.Second * 1)
			} else {
				printTest(i)
				e.Exit()
			}
		}
	*/

	r := gin.New()
	/*
	r.Use(func(context *gin.Context) {
		fmt.Println("request /")
		entry, blockErr := sentinel.Entry("abc")
		if nil != blockErr {
			context.AbortWithStatusJSON(400, gin.H{"error": blockErr.Error() + ", limit has reached"})
		} else {
			entry.Exit()
			context.Next()
		}
	})
	 */

	r.Handle("GET", "/prodlist", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg":"success"})
	})

	r.Run(":8080")

}

func TestCircuitbreaker(t *testing.T)  {
	err := sentinel.InitDefault()
	if nil != err {
		log.Fatal(err)
	}	
	
	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		{
			Id:               "prods",
			Resource:         "",
			Strategy:         circuitbreaker.SlowRequestRatio,
			RetryTimeoutMs:   3000,
			MinRequestAmount: 0,
			StatIntervalMs:   0,
			MaxAllowedRtMs:   0,
			Threshold:        0,
		},
	})
	

}
