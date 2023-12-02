package main

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"Helloworld/handler"
	pb "Helloworld/proto"
)

var (
	service = "helloworld"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld)); err != nil {
		logger.Fatal(err)
	}
	
	//注册 handler， 新增的业务handler...
	if err := pb.RegisterShenzhenHandler(srv.Server(), new(handler.ShenZhen)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
