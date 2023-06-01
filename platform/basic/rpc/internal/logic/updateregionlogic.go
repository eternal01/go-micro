package logic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRegionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRegionLogic {
	return &UpdateRegionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新地区信息
func (l *UpdateRegionLogic) UpdateRegion(in *basic.UpdateRegionRequest) (*basic.UpdateRegionResponse, error) {
	err := l.svcCtx.RegionModel.Update(l.ctx, &model.SystemRegion{
		Id:       in.Id,
		ParentId: in.ParentId,
		Name:     in.Name,
	})
	if err != nil {
		return nil, err
	}

	return &basic.UpdateRegionResponse{
		Id: in.Id,
	}, nil
}
