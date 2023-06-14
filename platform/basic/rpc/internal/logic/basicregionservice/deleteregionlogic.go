package basicregionservicelogic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRegionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRegionLogic {
	return &DeleteRegionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除地区信息
func (l *DeleteRegionLogic) DeleteRegion(in *basic.DeleteRegionRequest) (*basic.DeleteRegionResponse, error) {
	err := l.svcCtx.RegionModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &basic.DeleteRegionResponse{}, nil
}
