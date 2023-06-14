package basicregionservicelogic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRegionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegionLogic {
	return &GetRegionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取地区信息
func (l *GetRegionLogic) GetRegion(in *basic.GetRegionRequest) (*basic.GetRegionResponse, error) {
	region, err := l.svcCtx.RegionModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &basic.GetRegionResponse{
		Id:       region.Id,
		ParentId: region.ParentId,
		Name:     region.Name,
	}, nil
}
