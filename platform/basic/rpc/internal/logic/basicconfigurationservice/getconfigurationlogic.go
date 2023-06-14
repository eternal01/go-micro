package basicconfigurationservicelogic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationLogic {
	return &GetConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取配置信息
func (l *GetConfigurationLogic) GetConfiguration(in *basic.GetConfigurationRequest) (*basic.GetConfigurationResponse, error) {
	configuration, err := l.svcCtx.ConfigurationModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &basic.GetConfigurationResponse{
		Id:          configuration.Id,
		Name:        configuration.Name,
		Description: configuration.Description,
		Content:     configuration.Content,
	}, nil
}
