package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRegionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRegionLogic {
	return &UpdateRegionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRegionLogic) UpdateRegion(req *types.GatewayUpdateRegionRequest) (resp *types.GatewayUpdateRegionReply, err error) {
	region, err := l.svcCtx.BasicRpc.UpdateRegion(l.ctx, &basic.UpdateRegionRequest{
		Id:       req.Id,
		ParentId: req.ParentId,
		Name:     req.Name,
	})

	if err != nil {
		return nil, err
	}

	return &types.GatewayUpdateRegionReply{
		Id: region.Id,
	}, nil
}
