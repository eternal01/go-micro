// Code generated by goctl. DO NOT EDIT.
// Source: basic.proto

package basicclient

import (
	"context"

	"go-micro/platform/basic/rpc/basic"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddRegionRequest   = basic.AddRegionRequest
	AddRegionResponse  = basic.AddRegionResponse
	GetRegionChild     = basic.GetRegionChild
	GetRegionRequest   = basic.GetRegionRequest
	GetRegionResponse  = basic.GetRegionResponse
	GetRegionsRequest  = basic.GetRegionsRequest
	GetRegionsResponse = basic.GetRegionsResponse
	Request            = basic.Request
	Response           = basic.Response

	Basic interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetRegionResponse, error)
		GetRegions(ctx context.Context, in *GetRegionsRequest, opts ...grpc.CallOption) (*GetRegionsResponse, error)
		AddRegion(ctx context.Context, in *AddRegionRequest, opts ...grpc.CallOption) (*AddRegionResponse, error)
	}

	defaultBasic struct {
		cli zrpc.Client
	}
)

func NewBasic(cli zrpc.Client) Basic {
	return &defaultBasic{
		cli: cli,
	}
}

func (m *defaultBasic) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := basic.NewBasicClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultBasic) GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*GetRegionResponse, error) {
	client := basic.NewBasicClient(m.cli.Conn())
	return client.GetRegion(ctx, in, opts...)
}

func (m *defaultBasic) GetRegions(ctx context.Context, in *GetRegionsRequest, opts ...grpc.CallOption) (*GetRegionsResponse, error) {
	client := basic.NewBasicClient(m.cli.Conn())
	return client.GetRegions(ctx, in, opts...)
}

func (m *defaultBasic) AddRegion(ctx context.Context, in *AddRegionRequest, opts ...grpc.CallOption) (*AddRegionResponse, error) {
	client := basic.NewBasicClient(m.cli.Conn())
	return client.AddRegion(ctx, in, opts...)
}
