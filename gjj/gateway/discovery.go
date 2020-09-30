package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
	"log"
	"strings"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

type Discovery struct {
	Loadbalancer *Loadbalancer
	Registry     *Registry
}

type Loadbalancer struct {
	SchemePrefix string
}

type Registry struct {
	Type    string
	Address string
}

//获取注册中心，目前只支持etcd
func (this *Discovery) getRegistry() registry.Registry {
	if this.Registry.Type == "etcd" { //注册中心支持类型，目前只支持etcd
		return etcd.NewRegistry(
			registry.Addrs(this.Registry.Address),
		)
	}
	panic("error registry")
}

//调用Grpc
func (this *Discovery) CallGrpcService(args ...string) ([]byte, error) {
	reg := this.getRegistry()
	if len(args) < 2 {
		return nil, errors.New(`require service and endpoint  `)
	}
	var req, service, endpoint string
	service = args[0]
	endpoint = args[1]
	if len(args) > 2 {
		req = strings.Join(args[2:], " ")
	}
	// empty request
	if len(req) == 0 {
		req = `{}`
	}
	var request map[string]interface{}
	var response []byte

	d := json.NewDecoder(strings.NewReader(req))
	d.UseNumber()
	if err := d.Decode(&request); err != nil {
		return nil, err
	}
	ctx := context.Background()
	se := selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.RoundRobin),
	)
	cli := grpc.NewClient(
		client.Registry(reg),
		client.Selector(se),
	)
	creq := cli.NewRequest(service, endpoint, request, client.WithContentType("application/json"))
	var opts []client.CallOption

	var err error
	var rsp json.RawMessage
	err = cli.Call(ctx, creq, &rsp, opts...)
	if err == nil {
		var out bytes.Buffer
		defer out.Reset()
		if err := json.Indent(&out, rsp, "", "\t"); err != nil {
			return nil, err
		}
		response = out.Bytes()
	}
	if err != nil {
		return nil, fmt.Errorf("error calling %s.%s: %v", service, endpoint, err)
	}
	return response, nil
}

func (d *Discovery) GetServiceNode(serviceName string) *registry.Node {
	reg := d.getRegistry()
	if nil == reg {
		return nil
	}
	s := selector.NewSelector(
		selector.Registry(reg),
		selector.SetStrategy(selector.RoundRobin),
	)

	next, err := s.Select(serviceName)
	if nil != err {
		log.Fatal(err)
	}

	node, err := next()
	if nil != err {
		log.Fatal(err)
	}
	return node
}
