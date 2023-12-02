package main

import (
	"context"
	"time"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"Helloworld/handler"
	pb "Helloworld/proto"
)

var (
	serviceName = "helloworld"
	version = "latest"
)

func initClients(srv micro.Service) {
	// 定义 访问服务的客户端： Create client; 
	hellworldClient := pb.NewHelloworldService(serviceName, srv.Client())
	shenzhenClient := pb.NewShenzhenService(serviceName, srv.Client())

	go func() {
		time.Sleep(4*time.Second)
		
		// 通过client的 rpc 接口调用 服务端。
		for {
		// Call service; 其中Call()就是 proto 定义的rpc 接口
		rsp, err := hellworldClient.Call(context.Background(), &pb.CallRequest{Name: "John"})
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info(rsp)

		time.Sleep(1 * time.Second)

		r1, e := shenzhenClient.Call(context.Background(), &pb.SZReq{Name: "Tim"})
		if e != nil {
			logger.Fatal("e: ", e)
			return 
		}
		logger.Info(r1)
	}
	}()
}


func main() {
	// Create service
	srv := micro.NewService(
	)
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld)); err != nil {
		logger.Fatal(err)
	}
	
	//注册 handler; 新增的业务ShenZhen, ShenZhen 实现了grpc的 service定义接口。
	if err := pb.RegisterShenzhenHandler(srv.Server(), new(handler.ShenZhen)); err != nil {
		logger.Fatal(err)
	}

	//有时候 在服务上要访问其他服务，那么就需要初始化访问其它服务的client. 
	initClients(srv)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
