package register

import (
	"context"
	"fmt"
	etcdClient "go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type Service struct {
	client *etcdClient.Client
	ctx context.Context
}

func NewRegisterService(endPoints []string) *Service {
	config := etcdClient.Config{
		Endpoints:   endPoints,
		DialTimeout: 5 * time.Second,
	}
	client, err := etcdClient.New(config)
	if nil != err {
		log.Fatal(err)
	}

	return &Service{client: client, ctx: context.Background()}
}

func (r *Service) RegService(id string, name string, address string) error {
	kv := etcdClient.NewKV(r.client)
	keyPrefix := "/services"

	lease := etcdClient.NewLease(r.client)

	leaseRes, err := lease.Grant(r.ctx, 20)
	if nil != err {
		return err
	}

	_, err = kv.Put(r.ctx, keyPrefix+"/"+id+"/"+name, address, etcdClient.WithLease(leaseRes.ID))
	if nil != err {
		return err
	}

	keepAlive, err := lease.KeepAlive(r.ctx, leaseRes.ID)
	if nil != err {
		return err
	}

	go r.KeepAlive(keepAlive)


	return err
}

func (r *Service) KeepAlive(keepAliveChan <-chan *etcdClient.LeaseKeepAliveResponse)  {
	for {
		select {
		case ret := <-keepAliveChan:
			if nil != ret {
				fmt.Println("get lease success", time.Now(), ret)
			} else {
				fmt.Println("get lease failed", time.Now(), ret)
			}
		}
	}
}


func (r *Service)UnRegister(id string) error {
	kv := etcdClient.NewKV(r.client)
	keyPrefix := "/services"
	_, err := kv.Delete(context.Background(), keyPrefix + "/" + id, etcdClient.WithPrefix())
	return err
}
