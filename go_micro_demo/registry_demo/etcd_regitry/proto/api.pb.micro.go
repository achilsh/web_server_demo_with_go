// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/api.proto

package backend

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
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

// Api Endpoints for Api service

func NewApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Api service

type ApiService interface {
	PName(ctx context.Context, in *PNReq, opts ...client.CallOption) (*PNRsp, error)
}

type apiService struct {
	c    client.Client
	name string
}

func NewApiService(name string, c client.Client) ApiService {
	return &apiService{
		c:    c,
		name: name,
	}
}

func (c *apiService) PName(ctx context.Context, in *PNReq, opts ...client.CallOption) (*PNRsp, error) {
	req := c.c.NewRequest(c.name, "Api.PName", in)
	out := new(PNRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiHandler interface {
	PName(context.Context, *PNReq, *PNRsp) error
}

func RegisterApiHandler(s server.Server, hdlr ApiHandler, opts ...server.HandlerOption) error {
	type api interface {
		PName(ctx context.Context, in *PNReq, out *PNRsp) error
	}
	type Api struct {
		api
	}
	h := &apiHandler{hdlr}
	return s.Handle(s.NewHandler(&Api{h}, opts...))
}

type apiHandler struct {
	ApiHandler
}

func (h *apiHandler) PName(ctx context.Context, in *PNReq, out *PNRsp) error {
	return h.ApiHandler.PName(ctx, in, out)
}
