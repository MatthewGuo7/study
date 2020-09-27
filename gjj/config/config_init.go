package config

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"os"
)

type DataConfig struct {
	MySql *struct {
		Dsn     string
		Maxidle int
		Maxopen int
	}
	Redis *struct {
		Ip   string
		Port uint64
	}
}

type LocalConfig struct {
	Address string
	Path    string
	Port    uint64
}

type Service struct {
	Namespace string
	Name      string
}

type GlobalConfig struct {
	LocalConfig *LocalConfig
	Service     *Service
	DataConfig  *DataConfig
}

func NewGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		&LocalConfig{},
		&Service{},
		&DataConfig{},
	}
}

var JConfig *GlobalConfig
var nacosClient config_client.IConfigClient

func InitConfig(configFile string) {
	//configFile := "app.yaml"
	err := config.LoadFile(configFile)
	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(configFile)

	JConfig = NewGlobalConfig()
	err = config.Get("jtthink").Scan(JConfig)
	if nil != err {
		log.Fatal(err)
	}

	fmt.Println(JConfig.DataConfig, JConfig.LocalConfig)
	serverConfigs := []constant.ServerConfig{
		{IpAddr: JConfig.LocalConfig.Address, ContextPath: JConfig.LocalConfig.Path, Port: JConfig.LocalConfig.Port},
	}
	nacosClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})

	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(JConfig.DataConfig)
	listenNacos("jt-dataconfig", "jtthink", JConfig.DataConfig)
}

func writeFile(group, dataid, data string, model interface{}) {

	cacheFile := fmt.Sprintf("runtime/configcache/%s-%s.yaml", group, dataid)
	file, err := os.OpenFile(cacheFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		log.Println(err)
		return
	}
	err = config.LoadFile(cacheFile)
	if err != nil {
		log.Println(err)
		return
	}
	err = config.Scan(model)
	if err != nil {
		log.Println(err)
		return
	}
}

func listenNacos(dataid string, group string, model interface{}) {
	param := vo.ConfigParam{
		DataId: dataid,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println(data)
			writeFile(group, dataid, data, model)
		},
	}
	c, err := nacosClient.GetConfig(param)
	if nil != err {
		log.Fatal(err)
	}

	fmt.Printf("get config = %+v", c)
	writeFile(group, dataid, c, JConfig.DataConfig)
	fmt.Println(JConfig.DataConfig, JConfig.LocalConfig, JConfig.Service)
	err = nacosClient.ListenConfig(param)
	if err != nil {
		log.Println("listen config error,dataid:", dataid, err)
	}
	fmt.Println(JConfig.DataConfig, JConfig.LocalConfig, JConfig.Service)
}
