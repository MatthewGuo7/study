package gateway

import (
	"github.com/micro/go-micro/v2/config"
	"log"
)

func InitConfig() Routers {
	configFile := "gateway.yaml"
	err := config.LoadFile(configFile)
	if nil != err {
		log.Fatal(err)
	}

	routers := make(Routers, 0)
	err = config.Get("routers").Scan(&routers)
	if nil != err {
		log.Fatal(err)
	}
	return routers
}
