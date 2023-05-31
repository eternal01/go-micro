package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStagesLogic {
	return &GetStagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据父级id获取阶段信息
func (l *GetStagesLogic) GetStages(in *basic.GetStagesRequest) (*basic.GetStagesResponse, error) {
	whereBuilder := l.svcCtx.StageModel.RowBuilder().Where(squirrel.Eq{"parent_id": in.ParentId})
	stages, err := l.svcCtx.StageModel.FindChildrenList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	l.Logger.Slowf("result is %+v", stages)
	var list []*basic.GetStageChild
	if len(stages) > 0 {
		for _, stage := range stages {
			list = append(list, &basic.GetStageChild{
				Id:          stage.Id,
				ParentId:    stage.ParentId,
				Name:        stage.Name,
				Description: stage.Description,
			})
		}
	}

	return &basic.GetStagesResponse{
		List: list,
	}, nil
}
