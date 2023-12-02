package main

import (
	"context"
	"time"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	pb "Demo-client/proto"
)

var (
	serviceName = "helloworld" //定义某个服务的服务名，这个一般都是跟服务端的保持一致，否则就会出现： "level=fatal {"id":"go.micro.client","code":500,"detail":"service helloworld1: not found","status":"Internal Server Error"}" 
	// 一般在用于client 来连接。
	version = "latest"

)

// 
func main() {
	// Create service, 主要是用于获取 client句柄。
	srv := micro.NewService()
	srv.Init() 

	// Create client;  需要创建访问某中服务的客户端句柄。
	hellworldClient := pb.NewHelloworldService(serviceName, srv.Client())
	shenzhenClient := pb.NewShenzhenService(serviceName, srv.Client())

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
}
