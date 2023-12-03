package handler

import (
	"context"
	// pb "backend/proto"

	"go-micro.dev/v4/logger"

	pb "etc_registry/proto"
)

type Backend struct{}

func (e *Backend) RetName(ctx context.Context, req *pb.BkReq, rsp *pb.BkResp) error {
	logger.Infof("Received Backend.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

