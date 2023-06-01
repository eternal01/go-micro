package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConfigurationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationsLogic {
	return &GetConfigurationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取配置列表
func (l *GetConfigurationsLogic) GetConfigurations(in *basic.GetConfigurationsRequest) (*basic.GetConfigurationsResponse, error) {
	whereBuilder := l.svcCtx.ConfigurationModel.RowBuilder()
	configurations, err := l.svcCtx.ConfigurationModel.FindList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	l.Logger.Slowf("result is %+v", configurations)
	var list []*basic.Configuration
	if len(configurations) > 0 {
		for _, configuration := range configurations {
			list = append(list, &basic.Configuration{
				Id:          configuration.Id,
				Name:        configuration.Name,
				Description: configuration.Description,
				Content:     configuration.Content,
			})
		}
	}

	return &basic.GetConfigurationsResponse{
		List: list,
	}, nil
}
