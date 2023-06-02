package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigurationLogic {
	return &DeleteConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除配置信息
func (l *DeleteConfigurationLogic) DeleteConfiguration(in *basic.DeleteConfigurationRequest) (*basic.DeleteConfigurationResponse, error) {
	err := l.svcCtx.ConfigurationModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &basic.DeleteConfigurationResponse{}, nil
}
