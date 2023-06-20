package configuration

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigurationByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationByNameLogic {
	return &GetConfigurationByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigurationByNameLogic) GetConfigurationByName(req *types.GatewayGetConfigurationByNameRequest) (resp *types.GatewayGetConfigurationReply, err error) {
	if len(req.Name) <= 0 {
		return nil, errorx.ParamsError
	}
	configuration, err := l.svcCtx.BasicConfigurationRpc.GetConfigurationByName(l.ctx, &basic.GetConfigurationByNameRequest{
		Name: req.Name,
	})
	if err != nil {
		l.Logger.Errorf("RPC-BASIC GetConfigurationByName Error - %+v", err)
		return nil, err
	}

	return &types.GatewayGetConfigurationReply{
		Id:          configuration.Id,
		Name:        configuration.Name,
		Description: configuration.Description,
		Content:     configuration.Content,
	}, nil
}
