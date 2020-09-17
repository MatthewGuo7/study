package main

import (
	"context"
	"fmt"
	etcdClient "go.etcd.io/etcd/clientv3"
	"log"
	"testing"
	"time"
)

func TestFirst(t *testing.T)  {
	config := etcdClient.Config{
		Endpoints:            []string{"127.0.0.1:2379"},
		DialTimeout:          2 * time.Second,
	}

	client, err := etcdClient.New(config)
	if nil != err {
		log.Fatal(err)
	}
	defer client.Close()

	kv := etcdClient.NewKV(client)
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	resp, err := kv.Put(ctx, "/service/key1", "value1")
	cancel()
	if nil != err {
		log.Fatal(err)
	}

	fmt.Println(resp)

}
