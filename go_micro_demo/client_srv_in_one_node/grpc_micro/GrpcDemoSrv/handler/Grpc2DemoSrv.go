package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "GrpcDemoSrv/proto"
)

type GrpcDemoSrv2 struct{}


func (e *GrpcDemoSrv2)Grpc2Call(ctx context.Context, req *pb.Demo2Req, rsp *pb.Demo2Resp) error {
	logger.Infof("Received GrpcDemoSrv.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}