package handler

import (
	"context"
	"fmt"

	"go-micro.dev/v4/logger"

	pb "etc_registry/proto"
)

//定义 访问其他服务的client.
var BackendClient pb.BackendService

type Api struct{}



func (e *Api) PName(ctx context.Context, req *pb.PNReq, rsp *pb.PNRsp) error {
	logger.Infof("Received Api.Call request: %v", req)

	innReq := &pb.BkReq{
		Name :"api: " + req.GetName(),
	}

	//调用其他服务的rpc接口
	v, err := BackendClient.RetName(ctx, innReq)
	if err != nil {
		fmt.Println("call fail, e: ", e)
		return nil
	}

	rsp.Msg = "Hello "  + v.GetMsg()
	return nil
}