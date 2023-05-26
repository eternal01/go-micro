package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRegionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRegionLogic {
	return &AddRegionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRegionLogic) AddRegion(req *types.GatewayAddRegionRequest) (resp *types.GatewayAddRegionReply, err error) {
	if req.ParentId < 0 || len(req.Name) <= 0 {
		return nil, errorx.ParamsError
	}
	result, err := l.svcCtx.BasicRpc.AddRegion(l.ctx, &basic.AddRegionRequest{
		Name:     req.Name,
		ParentId: req.ParentId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GatewayAddRegionReply{
		Id: result.Id,
	}, nil
}
