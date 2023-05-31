package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStageLogic {
	return &GetStageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取阶段信息
func (l *GetStageLogic) GetStage(in *basic.GetStageRequest) (*basic.GetStageResponse, error) {
	stage, err := l.svcCtx.StageModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &basic.GetStageResponse{
		Id:          stage.Id,
		Name:        stage.Name,
		ParentId:    stage.ParentId,
		Description: stage.Description,
	}, nil
}
