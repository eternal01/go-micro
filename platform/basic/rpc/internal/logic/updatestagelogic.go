package logic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStageLogic {
	return &UpdateStageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新阶段信息
func (l *UpdateStageLogic) UpdateStage(in *basic.UpdateStageRequest) (*basic.UpdateStageResponse, error) {
	err := l.svcCtx.StageModel.Update(l.ctx, &model.SystemStage{
		Id:          in.Id,
		Name:        in.Name,
		ParentId:    in.ParentId,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}

	return &basic.UpdateStageResponse{
		Id: in.Id,
	}, nil
}
