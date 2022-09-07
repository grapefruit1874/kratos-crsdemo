package server

import (
	etcdregitry "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	etcdclient "go.etcd.io/etcd/client/v3"
	"kratos-crsdemo/app/user/internal/conf"
)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	client, _ := etcdclient.New(etcdclient.Config{
		Endpoints:   []string{conf.Addr},
		DialTimeout: conf.Timeout.AsDuration(),
	})
	// 创建一个 registry 对象，就是对 ectd client 操作的一个包装
	r := etcdregitry.New(client)
	return r
}
