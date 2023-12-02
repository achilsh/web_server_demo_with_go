package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "Helloworld/proto"
)

type ShenZhen struct{}

func (sz *ShenZhen) Call(ctx context.Context, req *pb.SZReq, rsp *pb.SZRsp) error {
	logger.Infof("Received Shenzhen.Call request: %v", req)
	rsp.Msg = "Shenzhen " + req.Name
	return nil
}

