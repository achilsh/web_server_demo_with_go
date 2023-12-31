// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/GrpcDemoSrv.proto

package GrpcDemoSrv

import (
	fmt "fmt"
	math "math"

	proto "google.golang.org/protobuf/proto"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for GrpcDemoSrv service

func NewGrpcDemoSrvEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GrpcDemoSrv service

type GrpcDemoSrvService interface {
	GrpcCall(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
}

type grpcDemoSrvService struct {
	c    client.Client
	name string
}

func NewGrpcDemoSrvService(name string, c client.Client) GrpcDemoSrvService {
	return &grpcDemoSrvService{
		c:    c,
		name: name,
	}
}

func (c *grpcDemoSrvService) GrpcCall(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "GrpcDemoSrv.GrpcCall", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GrpcDemoSrv service

type GrpcDemoSrvHandler interface {
	GrpcCall(context.Context, *CallRequest, *CallResponse) error
}

func RegisterGrpcDemoSrvHandler(s server.Server, hdlr GrpcDemoSrvHandler, opts ...server.HandlerOption) error {
	type grpcDemoSrv interface {
		GrpcCall(ctx context.Context, in *CallRequest, out *CallResponse) error
	}
	type GrpcDemoSrv struct {
		grpcDemoSrv
	}
	h := &grpcDemoSrvHandler{hdlr}
	return s.Handle(s.NewHandler(&GrpcDemoSrv{h}, opts...))
}

type grpcDemoSrvHandler struct {
	GrpcDemoSrvHandler
}

func (h *grpcDemoSrvHandler) GrpcCall(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.GrpcDemoSrvHandler.GrpcCall(ctx, in, out)
}
