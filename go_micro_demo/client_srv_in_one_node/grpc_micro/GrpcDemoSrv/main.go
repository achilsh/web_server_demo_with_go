package main

import (
	"context"
	"time"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"GrpcDemoSrv/handler"
	pb "GrpcDemoSrv/proto"
)

var (
	serviceName = "grpcdemosrv"
	version = "latest"
)


func InitClients(srv micro.Service) {
	// 定义 访问服务的客户端： Create client; 
	hellworldClient := pb.NewGrpcDemoSrvService(serviceName, srv.Client())
	shenzhenClient := pb.NewGrpcDemo2SrvService(serviceName, srv.Client())

	go func() {
		time.Sleep(4*time.Second)
		
		// 通过client的 rpc 接口调用 服务端。
		for {
		// Call service; 其中Call()就是 proto 定义的rpc 接口
		rsp, err := hellworldClient.GrpcCall(context.Background(), &pb.CallRequest{Name: "John"})
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info(rsp)

		time.Sleep(1 * time.Second)

		r1, e := shenzhenClient.Grpc2Call(context.Background(), &pb.Demo2Req{Name: "Tim"})
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
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()), //主要是为了客户端使用的。
	)
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
	)

	// Register handler;服务端上， 注册业务定义和实现的rpc 接口对象。
	if err := pb.RegisterGrpcDemoSrvHandler(srv.Server(), new(handler.GrpcDemoSrv)); err != nil {
		logger.Fatal(err)
	}

	if e := pb.RegisterGrpcDemo2SrvHandler(srv.Server(), new(handler.GrpcDemoSrv2)); e != nil {
		logger.Fatal(e)
		//
	}

	//初始化访问grpc 服务的client.
	InitClients(srv)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
