package register

import (
	"context"
	"etcddemo/endpoint"
	"fmt"
	etcdClient "go.etcd.io/etcd/clientv3"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

type Client struct {
	client   *etcdClient.Client
	ctx      context.Context
	Services []*ServiceInfo
}

type ServiceInfo struct {
	ServiceID   string
	ServiceName string
	ServiceAddr string
}

func NewClient(endPoints []string) *Client {
	config := etcdClient.Config{
		Endpoints:   endPoints,
		DialTimeout: 5 * time.Second,
	}
	client, err := etcdClient.New(config)
	if nil != err {
		log.Fatal(err)
	}

	return &Client{client: client, ctx: context.Background()}
}

func (c *Client) LoadService() error {
	kv := etcdClient.NewKV(c.client)
	res, err := kv.Get(c.ctx, "/services", etcdClient.WithPrefix())
	if nil != err {
		return err
	}

	for _, service := range res.Kvs {
		c.parseService(service.Key, service.Value)
	}

	return nil
}

func (this *Client) parseService(key []byte, value []byte) {
	reg := regexp.MustCompile("/services/(\\w+)/(\\w+)")
	if reg.Match(key) {
		idandname := reg.FindSubmatch(key)
		sid := idandname[1]
		sname := idandname[2]
		fmt.Println(string(sid), string(sname))
		this.Services = append(this.Services, &ServiceInfo{ServiceID: string(sid),
			ServiceName: string(sname), ServiceAddr: string(value)})
	}
}

func (c *Client) GetService(serviceName, method string, requestFunc endpoint.EndRequestFunc) endpoint.EndPoint {
	for _, service := range c.Services {
		if service.ServiceName == serviceName {
			return func(ctx context.Context, requestParam interface{}) (response interface{}, err error) {
				tp := http.DefaultTransport
				client := http.Client{Transport: tp}
				request, err := http.NewRequest(method, "http://"+service.ServiceAddr, nil)
				if nil != err {
					return nil, err
				}

				err = requestFunc(ctx, request, requestParam)
				if nil != err {
					return nil, err
				}

				resp, err := client.Do(request)
				if nil != err {
					return nil, err
				}
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				if nil != err {
					return nil, err
				}

				return string(body), nil
			}
		}
	}
	return nil
}
