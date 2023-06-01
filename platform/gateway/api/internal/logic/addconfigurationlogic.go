package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigurationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigurationLogic {
	return &AddConfigurationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConfigurationLogic) AddConfiguration(req *types.GatewayAddConfigurationRequest) (resp *types.GatewayAddConfigurationReply, err error) {
	if len(req.Name) <= 0 {
		return nil, errorx.ParamsError
	}
	result, err := l.svcCtx.BasicRpc.AddConfiguration(l.ctx, &basic.AddConfigurationRequest{
		Name:        req.Name,
		Description: req.Description,
		Content:     req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &types.GatewayAddConfigurationReply{
		Id: result.Id,
	}, nil
}
