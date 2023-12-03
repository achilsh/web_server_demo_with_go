package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/etcd"
	httpServer "github.com/go-micro/plugins/v4/server/http"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"

	pb "etc_registry/proto"
)

const (
	HTTP_SERVER_NMAE = "go-micro_http_server"
)
func main() {
	ginR := gin.Default()
	vGroup := ginR.Group("v1")
	{
		vGroup.Handle("POST", "/demo1", func(c *gin.Context){
			innReq := &pb.BkReq{
				Name :"api: " + "aaaaaa",
			}

			rsp, e := bkClient.RetName(context.Background(), innReq)
			var rdata string =  ""
			if e != nil {
				rdata = "is error"
			} else {
				rdata = rsp.GetMsg()
			}
			//add logic....
			c.JSON(200, gin.H {
				"data": rdata,
			})
		})
	}
	
	httpSrv := httpServer.NewServer(
		server.Name(HTTP_SERVER_NMAE),
	)
	hd := httpSrv.NewHandler(ginR)
	if e := httpSrv.Handle(hd); e != nil {
		fmt.Println("fail, e: ", e)
		return 
	}

	mSrv := micro.NewService(
		micro.Server(httpSrv),	
		micro.Registry(etcd.NewRegistry(func(option *registry.Options){
				option.Addrs = []string{"127.0.0.1:2379"} //这是 etcd的服务地址
				option.Timeout = 2*time.Second //注册超时时间
		},
		//  etcd.Auth("", "") ,  //比如增加etcd的认证信息。
		 )),
		 micro.Address(":8090"),
	)

	InitClients(mSrv)

	mSrv.Init()
	mSrv.Run()
}


var bkClient  pb.BackendService
var backendServiceName = "backend"
func InitClients(srv micro.Service) {
	bkClient = pb.NewBackendService(backendServiceName, srv.Client())

}