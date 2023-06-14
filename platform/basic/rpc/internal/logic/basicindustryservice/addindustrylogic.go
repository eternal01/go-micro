package basicindustryservicelogic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIndustryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddIndustryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIndustryLogic {
	return &AddIndustryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加职业信息
func (l *AddIndustryLogic) AddIndustry(in *basic.AddIndustryRequest) (*basic.AddIndustryResponse, error) {
	industry := new(model.SystemIndustry)
	industry.IndustryId = in.IndustryId
	industry.Name = in.Name
	industry.ParentId = in.ParentId
	industry.LevelType = in.LevelType
	industry.Description = in.Description

	result, err := l.svcCtx.IndustryModel.Insert(l.ctx, industry)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &basic.AddIndustryResponse{
		Id: id,
	}, nil
}
