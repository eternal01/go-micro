// Code generated by goctl. DO NOT EDIT.
// Source: basic.proto

package server

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/logic/basicregionservice"
	"go-micro/platform/basic/rpc/internal/svc"
)

type BasicRegionServiceServer struct {
	svcCtx *svc.ServiceContext
	basic.UnimplementedBasicRegionServiceServer
}

func NewBasicRegionServiceServer(svcCtx *svc.ServiceContext) *BasicRegionServiceServer {
	return &BasicRegionServiceServer{
		svcCtx: svcCtx,
	}
}

// 获取地区信息
func (s *BasicRegionServiceServer) GetRegion(ctx context.Context, in *basic.GetRegionRequest) (*basic.GetRegionResponse, error) {
	l := basicregionservicelogic.NewGetRegionLogic(ctx, s.svcCtx)
	return l.GetRegion(in)
}

// 根据父级id获取地区信息
func (s *BasicRegionServiceServer) GetRegions(ctx context.Context, in *basic.GetRegionsRequest) (*basic.GetRegionsResponse, error) {
	l := basicregionservicelogic.NewGetRegionsLogic(ctx, s.svcCtx)
	return l.GetRegions(in)
}

// 添加地区信息
func (s *BasicRegionServiceServer) AddRegion(ctx context.Context, in *basic.AddRegionRequest) (*basic.AddRegionResponse, error) {
	l := basicregionservicelogic.NewAddRegionLogic(ctx, s.svcCtx)
	return l.AddRegion(in)
}

// 更新地区信息
func (s *BasicRegionServiceServer) UpdateRegion(ctx context.Context, in *basic.UpdateRegionRequest) (*basic.UpdateRegionResponse, error) {
	l := basicregionservicelogic.NewUpdateRegionLogic(ctx, s.svcCtx)
	return l.UpdateRegion(in)
}

// 删除地区信息
func (s *BasicRegionServiceServer) DeleteRegion(ctx context.Context, in *basic.DeleteRegionRequest) (*basic.DeleteRegionResponse, error) {
	l := basicregionservicelogic.NewDeleteRegionLogic(ctx, s.svcCtx)
	return l.DeleteRegion(in)
}