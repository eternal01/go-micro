package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStageLogic {
	return &GetStageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStageLogic) GetStage(req *types.GatewayGetStageRequest) (resp *types.GatewayGetStageReply, err error) {
	if req.Id <= 0 {
		return nil, errorx.ParamsError
	}
	classify, err := l.svcCtx.BasicRpc.GetStage(l.ctx, &basic.GetStageRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("RPC-BASIC GetStage Error - %+v", err)
		return nil, err
	}

	return &types.GatewayGetStageReply{
		Id:          classify.Id,
		Name:        classify.Name,
		Description: classify.Description,
	}, nil
}
