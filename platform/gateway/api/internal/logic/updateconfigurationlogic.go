package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigurationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigurationLogic {
	return &UpdateConfigurationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateConfigurationLogic) UpdateConfiguration(req *types.GatewayUpdateConfigurationRequest) (resp *types.GatewayUpdateConfigurationReply, err error) {
	configuration, err := l.svcCtx.BasicRpc.UpdateConfiguration(l.ctx, &basic.UpdateConfigurationRequest{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Content:     req.Content,
	})

	if err != nil {
		return nil, err
	}

	return &types.GatewayUpdateConfigurationReply{
		Id: configuration.Id,
	}, nil
}
