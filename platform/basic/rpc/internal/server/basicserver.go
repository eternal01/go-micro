// Code generated by goctl. DO NOT EDIT.
// Source: basic.proto

package server

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/logic"
	"go-micro/platform/basic/rpc/internal/svc"
)

type BasicServer struct {
	svcCtx *svc.ServiceContext
	basic.UnimplementedBasicServer
}

func NewBasicServer(svcCtx *svc.ServiceContext) *BasicServer {
	return &BasicServer{
		svcCtx: svcCtx,
	}
}

func (s *BasicServer) Ping(ctx context.Context, in *basic.Request) (*basic.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

// 获取地区信息
func (s *BasicServer) GetRegion(ctx context.Context, in *basic.GetRegionRequest) (*basic.GetRegionResponse, error) {
	l := logic.NewGetRegionLogic(ctx, s.svcCtx)
	return l.GetRegion(in)
}

// 根据父级id获取地区信息
func (s *BasicServer) GetRegions(ctx context.Context, in *basic.GetRegionsRequest) (*basic.GetRegionsResponse, error) {
	l := logic.NewGetRegionsLogic(ctx, s.svcCtx)
	return l.GetRegions(in)
}

// 获取职业信息
func (s *BasicServer) GetIndustry(ctx context.Context, in *basic.GetIndustryRequest) (*basic.GetIndustryResponse, error) {
	l := logic.NewGetIndustryLogic(ctx, s.svcCtx)
	return l.GetIndustry(in)
}

// 根据父级id获取职业信息
func (s *BasicServer) GetIndustries(ctx context.Context, in *basic.GetIndustriesRequest) (*basic.GetIndustriesResponse, error) {
	l := logic.NewGetIndustriesLogic(ctx, s.svcCtx)
	return l.GetIndustries(in)
}

func (s *BasicServer) AddRegion(ctx context.Context, in *basic.AddRegionRequest) (*basic.AddRegionResponse, error) {
	l := logic.NewAddRegionLogic(ctx, s.svcCtx)
	return l.AddRegion(in)
}
