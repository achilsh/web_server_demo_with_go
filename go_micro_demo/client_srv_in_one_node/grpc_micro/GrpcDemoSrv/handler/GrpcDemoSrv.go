package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "GrpcDemoSrv/proto"
)

type GrpcDemoSrv struct{}

func (e *GrpcDemoSrv) GrpcCall(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received GrpcDemoSrv.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
