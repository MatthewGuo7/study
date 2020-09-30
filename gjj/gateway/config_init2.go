package gateway

import (
	"fmt"
	"log"
	"sync"

	"github.com/micro/go-micro/v2/config"
)

var sysConfig *SysConfig     //系统配置
var sysConfig_Once sync.Once //单例模式
type SysConfig struct {      //新增全局配置 目前包含了 路由和 服务发现配置
	Routers   Routers
	Discovery *Discovery
}

func NewSysConfig() *SysConfig {
	return &SysConfig{Routers: make(Routers, 0), Discovery: &Discovery{}}
}
func GetSysConfig() *SysConfig {
	sysConfig_Once.Do(func() {
		configFile := "gateway.yaml"
		err := config.LoadFile(configFile)
		if err != nil {
			log.Fatal(err)
		}
		sysConfig = NewSysConfig()
		err = config.Scan(sysConfig)
		if err != nil {
			log.Fatal(err)
		}
		//sortFilters(sysConfig.Routes) //排序
	})
	return sysConfig
}

func (s *SysConfig) String() string {
	return fmt.Sprintf("discovery = %+v, routers = %+v\n", s.Discovery, s.Routers)
}

/*
func sortFilters(routes Routers) {
	for _, route := range routes {
		route.sortFilter()
	}
}
*/
