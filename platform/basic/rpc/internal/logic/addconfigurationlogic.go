package logic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConfigurationLogic {
	return &AddConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加配置信息
func (l *AddConfigurationLogic) AddConfiguration(in *basic.AddConfigurationRequest) (*basic.AddConfigurationResponse, error) {
	configuration := new(model.SystemConfiguration)
	configuration.Name = in.Name
	configuration.Description = in.Description
	configuration.Content = in.Content

	result, err := l.svcCtx.ConfigurationModel.Insert(l.ctx, configuration)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &basic.AddConfigurationResponse{
		Id: id,
	}, nil
}
