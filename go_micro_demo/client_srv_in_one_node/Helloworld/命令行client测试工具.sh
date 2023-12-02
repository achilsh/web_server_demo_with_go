#! /usr/bin/env bash 
##解释说明： go-micro 通用client 工具
#   helloworld 是定义在 main.go文件中 service的值，其实就是 服务名。
#   Helloworld.Call  是定义在proto 文件中的具体业务服务名和rpc 接口名 
#   '{"name": "John"}' 就是 定义在proto 文件中 req 的json 形式。因为go-micro默认使用json 和proto 格式进行编解码。

go-micro call  helloworld Helloworld.Call '{"name": "John"}'
go-micro call  helloworld Shenzhen.Call '{"name": "John"}'