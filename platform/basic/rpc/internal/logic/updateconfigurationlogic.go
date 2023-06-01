package logic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConfigurationLogic {
	return &UpdateConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新配置信息
func (l *UpdateConfigurationLogic) UpdateConfiguration(in *basic.UpdateConfigurationRequest) (*basic.UpdateConfigurationResponse, error) {
	err := l.svcCtx.ConfigurationModel.Update(l.ctx, &model.SystemConfiguration{
		Id:          in.Id,
		Name:        in.Name,
		Description: in.Description,
		Content:     in.Content,
	})
	if err != nil {
		return nil, err
	}

	return &basic.UpdateConfigurationResponse{
		Id: in.Id,
	}, nil
}
