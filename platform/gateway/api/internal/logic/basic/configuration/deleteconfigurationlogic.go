package configuration

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigurationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigurationLogic {
	return &DeleteConfigurationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteConfigurationLogic) DeleteConfiguration(req *types.GatewayDeleteConfigurationRequest) (resp *types.GatewayDeleteConfigurationReply, err error) {
	_, err = l.svcCtx.BasicConfigurationRpc.DeleteConfiguration(l.ctx, &basic.DeleteConfigurationRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}
