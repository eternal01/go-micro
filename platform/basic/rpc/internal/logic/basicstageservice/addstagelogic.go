package basicstageservicelogic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStageLogic {
	return &AddStageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加阶段信息
func (l *AddStageLogic) AddStage(in *basic.AddStageRequest) (*basic.AddStageResponse, error) {
	stage := new(model.SystemStage)
	stage.ParentId = in.ParentId
	stage.Name = in.Name
	stage.Description = in.Description

	result, err := l.svcCtx.StageModel.Insert(l.ctx, stage)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &basic.AddStageResponse{
		Id: id,
	}, nil
}
