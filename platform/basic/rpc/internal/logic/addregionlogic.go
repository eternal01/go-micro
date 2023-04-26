package logic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRegionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRegionLogic {
	return &AddRegionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRegionLogic) AddRegion(in *basic.AddRegionRequest) (*basic.AddRegionResponse, error) {
	// todo: add your logic here and delete this line

	region := new(model.SystemRegion)
	region.Id = 1
	region.ParentId = in.ParentId
	region.Name = in.Name

	_, err := l.svcCtx.RegionModel.Insert(l.ctx, region)
	if err != nil {
		return nil, err
	}

	return &basic.AddRegionResponse{}, nil
}
