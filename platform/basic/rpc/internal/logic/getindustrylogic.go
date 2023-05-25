package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIndustryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIndustryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIndustryLogic {
	return &GetIndustryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取职业信息
func (l *GetIndustryLogic) GetIndustry(in *basic.GetIndustryRequest) (*basic.GetIndustryResponse, error) {
	industry, err := l.svcCtx.IndustryModel.FindOneByIndustryId(l.ctx, in.IndustryId)
	if err != nil {
		return nil, err
	}

	return &basic.GetIndustryResponse{
		Id:          industry.Id,
		IndustryId:  industry.IndustryId,
		Name:        industry.Name,
		ParentId:    industry.ParentId,
		LevelType:   industry.LevelType,
		Description: industry.Description,
	}, nil
}
