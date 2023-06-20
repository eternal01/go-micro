package configuration

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigurationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationsLogic {
	return &GetConfigurationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigurationsLogic) GetConfigurations(req *types.GatewayGetConfigurationsRequest) (resp *types.GatewayGetConfigurationsReply, err error) {
	configurations, err := l.svcCtx.BasicConfigurationRpc.GetConfigurations(l.ctx, &basic.GetConfigurationsRequest{})
	if err != nil {
		return nil, err
	}
	resp = new(types.GatewayGetConfigurationsReply)
	if len(configurations.List) > 0 {
		for _, v := range configurations.List {
			resp.List = append(resp.List, types.GatewayConfiguration{
				Id:          v.Id,
				Name:        v.Name,
				Description: v.Description,
				Content:     v.Content,
			})
		}
	}

	return
}
