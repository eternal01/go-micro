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

// 添加地区信息
func (s *BasicServer) AddRegion(ctx context.Context, in *basic.AddRegionRequest) (*basic.AddRegionResponse, error) {
	l := logic.NewAddRegionLogic(ctx, s.svcCtx)
	return l.AddRegion(in)
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

// 添加职业信息
func (s *BasicServer) AddIndustry(ctx context.Context, in *basic.AddIndustryRequest) (*basic.AddIndustryResponse, error) {
	l := logic.NewAddIndustryLogic(ctx, s.svcCtx)
	return l.AddIndustry(in)
}

// 获取分类信息
func (s *BasicServer) GetClassify(ctx context.Context, in *basic.GetClassifyRequest) (*basic.GetClassifyResponse, error) {
	l := logic.NewGetClassifyLogic(ctx, s.svcCtx)
	return l.GetClassify(in)
}

// 根据父级id获取分类信息
func (s *BasicServer) GetClassifies(ctx context.Context, in *basic.GetClassifiesRequest) (*basic.GetClassifiesResponse, error) {
	l := logic.NewGetClassifiesLogic(ctx, s.svcCtx)
	return l.GetClassifies(in)
}

// 添加分类信息
func (s *BasicServer) AddClassify(ctx context.Context, in *basic.AddClassifyRequest) (*basic.AddClassifyResponse, error) {
	l := logic.NewAddClassifyLogic(ctx, s.svcCtx)
	return l.AddClassify(in)
}

// 获取类别信息
func (s *BasicServer) GetCategory(ctx context.Context, in *basic.GetCategoryRequest) (*basic.GetCategoryResponse, error) {
	l := logic.NewGetCategoryLogic(ctx, s.svcCtx)
	return l.GetCategory(in)
}

// 根据父级id获取类别信息
func (s *BasicServer) GetCategories(ctx context.Context, in *basic.GetCategoriesRequest) (*basic.GetCategoriesResponse, error) {
	l := logic.NewGetCategoriesLogic(ctx, s.svcCtx)
	return l.GetCategories(in)
}

// 添加类别信息
func (s *BasicServer) AddCategory(ctx context.Context, in *basic.AddCategoryRequest) (*basic.AddCategoryResponse, error) {
	l := logic.NewAddCategoryLogic(ctx, s.svcCtx)
	return l.AddCategory(in)
}

// 获取阶段信息
func (s *BasicServer) GetStage(ctx context.Context, in *basic.GetStageRequest) (*basic.GetStageResponse, error) {
	l := logic.NewGetStageLogic(ctx, s.svcCtx)
	return l.GetStage(in)
}

// 根据父级id获取阶段信息
func (s *BasicServer) GetStages(ctx context.Context, in *basic.GetStagesRequest) (*basic.GetStagesResponse, error) {
	l := logic.NewGetStagesLogic(ctx, s.svcCtx)
	return l.GetStages(in)
}

// 添加阶段信息
func (s *BasicServer) AddStage(ctx context.Context, in *basic.AddStageRequest) (*basic.AddStageResponse, error) {
	l := logic.NewAddStageLogic(ctx, s.svcCtx)
	return l.AddStage(in)
}
