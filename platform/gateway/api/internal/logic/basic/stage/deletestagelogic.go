package stage

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStageLogic {
	return &DeleteStageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStageLogic) DeleteStage(req *types.GatewayDeleteStageRequest) (resp *types.GatewayDeleteStageReply, err error) {
	_, err = l.svcCtx.BasicStageRpc.DeleteStage(l.ctx, &basic.DeleteStageRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}
