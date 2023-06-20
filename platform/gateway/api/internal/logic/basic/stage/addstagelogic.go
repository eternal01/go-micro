package stage

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStageLogic {
	return &AddStageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStageLogic) AddStage(req *types.GatewayAddStageRequest) (resp *types.GatewayAddStageReply, err error) {
	if len(req.Name) <= 0 {
		return nil, errorx.ParamsError
	}
	result, err := l.svcCtx.BasicStageRpc.AddStage(l.ctx, &basic.AddStageRequest{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &types.GatewayAddStageReply{
		Id: result.Id,
	}, nil
}
