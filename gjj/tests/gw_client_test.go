package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2/client/grpc"
	"log"
	"strings"
	"testing"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func CallService(reg registry.Registry, args ...string) ([]byte, error) {
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

func TestGwClient(t *testing.T) {
	serviceName := "go.micro.api.snoopy.course"
	reg := etcd.NewRegistry(registry.Addrs("localhost:2379"))
	if nil == reg {
		log.Fatal("registry etcd error")
	}

	serviceList, err := reg.GetService(serviceName)
	if nil != err {
		log.Fatal(err)
	}

	fmt.Println(serviceList[0].Endpoints)

	req:=map[string]interface{}{
		"size":3,
		"page":1,
	}
	breq,_:=json.Marshal(req)
	rsp,err:=CallService(reg,serviceName,"CourseService.ListForTop",string(breq))
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(rsp))

}
