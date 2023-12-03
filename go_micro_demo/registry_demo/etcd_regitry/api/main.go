package main

import (
	"time"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	"etc_registry/api/handler"
	pb "etc_registry/proto"
)

var (
	service = "api" 
	version = "latest"

	backendServiceName = "backend"
)

func main() {
	// Create service
	//对外也提供 http 请求，但是client的请求投应该携带类似：
    //  -H 'Content-Type: application/json' \
    //  -H 'Micro-Endpoint: Api.PName' \


	srv := micro.NewService(
		micro.RegisterInterval(1*time.Second), //重复向regitry 注册的时间间隔。
		micro.RegisterTTL(2*time.Second), //注册服务在registry中过期的时间
		micro.Registry(etcd.NewRegistry(func(option *registry.Options){
			option.Addrs = []string{"127.0.0.1:2379"} //这是 etcd的服务地址
			option.Timeout = 2*time.Second //注册超时时间
		},
		//  etcd.Auth("", "") ,  //比如增加etcd的认证信息。
		 )),
		
		micro.Address("127.0.0.1:8080"), //本地服务的ip和地址。
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	//
	handler.BackendClient = pb.NewBackendService(backendServiceName, srv.Client())

	// Register handler
	if err := pb.RegisterApiHandler(srv.Server(), new(handler.Api)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
