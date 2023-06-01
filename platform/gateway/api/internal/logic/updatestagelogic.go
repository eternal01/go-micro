package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStageLogic {
	return &UpdateStageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStageLogic) UpdateStage(req *types.GatewayUpdateStageRequest) (resp *types.GatewayUpdateStageReply, err error) {
	stage, err := l.svcCtx.BasicRpc.UpdateStage(l.ctx, &basic.UpdateStageRequest{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return nil, err
	}

	return &types.GatewayUpdateStageReply{
		Id: stage.Id,
	}, nil
}
