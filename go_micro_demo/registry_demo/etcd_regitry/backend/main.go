package main

import (
	"time"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	"etc_registry/backend/handler"
	pb "etc_registry/proto"
)

var (
	service = "backend"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.RegisterInterval(1*time.Second), //定义重复注册的时间间隔。 
		micro.RegisterTTL(2*time.Second), //注册服务在registry中的过期时间。 
		micro.Registry(etcd.NewRegistry(func(option *registry.Options){
				option.Addrs = []string{"127.0.0.1:2379"} //这是 etcd的服务地址
				option.Timeout = 2*time.Second //注册超时时间
		},
		//  etcd.Auth("", "") ,  //比如增加etcd的认证信息。
		 )),

		micro.Address("127.0.0.1:8081"), //本地ip和地址。
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterBackendHandler(srv.Server(), new(handler.Backend)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
