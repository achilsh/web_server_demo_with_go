package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "Helloworld/proto"
)

type Helloworld struct{}

func (e *Helloworld) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Helloworld.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
