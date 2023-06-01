package logic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIndustryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateIndustryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIndustryLogic {
	return &UpdateIndustryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新职业信息
func (l *UpdateIndustryLogic) UpdateIndustry(in *basic.UpdateIndustryRequest) (*basic.UpdateIndustryResponse, error) {
	err := l.svcCtx.IndustryModel.Update(l.ctx, &model.SystemIndustry{
		Id:          in.Id,
		IndustryId:  in.IndustryId,
		Name:        in.Name,
		ParentId:    in.ParentId,
		LevelType:   in.LevelType,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}

	return &basic.UpdateIndustryResponse{
		Id: in.Id,
	}, nil
}
