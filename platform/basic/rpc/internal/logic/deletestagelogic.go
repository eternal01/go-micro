package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStageLogic {
	return &DeleteStageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除阶段信息
func (l *DeleteStageLogic) DeleteStage(in *basic.DeleteStageRequest) (*basic.DeleteStageResponse, error) {
	err := l.svcCtx.StageModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &basic.DeleteStageResponse{}, nil
}
