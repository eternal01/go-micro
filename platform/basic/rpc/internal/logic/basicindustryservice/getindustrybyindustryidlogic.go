package basicindustryservicelogic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIndustryByIndustryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIndustryByIndustryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIndustryByIndustryIdLogic {
	return &GetIndustryByIndustryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据industry_id获取职业信息
func (l *GetIndustryByIndustryIdLogic) GetIndustryByIndustryId(in *basic.GetIndustryByIndustryIdRequest) (*basic.GetIndustryResponse, error) {
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
