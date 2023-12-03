package main

import (
	"fmt"
	"net/http"
	_ "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
)
func main() {
	//和gin 结合起来使用： 
	r := gin.Default() 
	//
	r.POST("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, "is test2 response.")
	})

	// 注册中心信息。
	webRegistry :=etcd.NewRegistry(func(option *registry.Options){
				option.Addrs = []string{"127.0.0.1:2379"} //这是 etcd的服务地址
				option.Timeout = 2*time.Second //注册超时时间
		},
		//  etcd.Auth("", "") ,  //比如增加etcd的认证信息。
	)

	webs := web.NewService(
		web.Name("http test2"), //当前的服务名。
		web.Address(":8070"),
		// web.Server(&http.Server{}), //定制 http 服务
		web.Handler(r), // 使用gin的 路由信息。
		web.Registry(webRegistry,
		//  etcd.Auth("", "") ,  //比如增加etcd的认证信息。
		 ),
		//  web.Context() //可以将业务的context存入传给 服务的context.
		// other options, you can web. get it.
		// ....
	)

	rserverList, _ := webRegistry.GetService("backend") //根据服务名获取 对应的服务节列表。
	for i := 0; i < len(rserverList); i++ {
		fmt.Printf("%+v\n", rserverList[i].Nodes[0])	
	}
	
	webs.Init() 

	webs.Run() 
}