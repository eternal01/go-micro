package basicconfigurationservicelogic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigurationByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationByNameLogic {
	return &GetConfigurationByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据配置名称获取配置信息
func (l *GetConfigurationByNameLogic) GetConfigurationByName(in *basic.GetConfigurationByNameRequest) (*basic.GetConfigurationResponse, error) {
	configuration, err := l.svcCtx.ConfigurationModel.FindOneByName(l.ctx, in.Name)
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
