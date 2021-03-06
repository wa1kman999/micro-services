// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/howAreYou.proto

package howAreYou

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

// Api Endpoints for HowAreYou service

func NewHowAreYouEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for HowAreYou service

type HowAreYouService interface {
	GetUserInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type howAreYouService struct {
	c    client.Client
	name string
}

func NewHowAreYouService(name string, c client.Client) HowAreYouService {
	return &howAreYouService{
		c:    c,
		name: name,
	}
}

func (c *howAreYouService) GetUserInfo(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "HowAreYou.GetUserInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HowAreYou service

type HowAreYouHandler interface {
	GetUserInfo(context.Context, *Request, *Response) error
}

func RegisterHowAreYouHandler(s server.Server, hdlr HowAreYouHandler, opts ...server.HandlerOption) error {
	type howAreYou interface {
		GetUserInfo(ctx context.Context, in *Request, out *Response) error
	}
	type HowAreYou struct {
		howAreYou
	}
	h := &howAreYouHandler{hdlr}
	return s.Handle(s.NewHandler(&HowAreYou{h}, opts...))
}

type howAreYouHandler struct {
	HowAreYouHandler
}

func (h *howAreYouHandler) GetUserInfo(ctx context.Context, in *Request, out *Response) error {
	return h.HowAreYouHandler.GetUserInfo(ctx, in, out)
}
