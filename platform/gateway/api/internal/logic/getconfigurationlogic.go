package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationLogic {
	return &GetConfigurationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigurationLogic) GetConfiguration(req *types.GatewayGetConfigurationRequest) (resp *types.GatewayGetConfigurationReply, err error) {
	if req.Id <= 0 {
		return nil, errorx.ParamsError
	}
	configuration, err := l.svcCtx.BasicRpc.GetConfiguration(l.ctx, &basic.GetConfigurationRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("RPC-BASIC GetConfiguration Error - %+v", err)
		return nil, err
	}

	return &types.GatewayGetConfigurationReply{
		Id:          configuration.Id,
		Name:        configuration.Name,
		Description: configuration.Description,
		Content:     configuration.Content,
	}, nil
}
