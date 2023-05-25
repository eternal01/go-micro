package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetIndustriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIndustriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIndustriesLogic {
	return &GetIndustriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据父级id获取职业信息
func (l *GetIndustriesLogic) GetIndustries(in *basic.GetIndustriesRequest) (*basic.GetIndustriesResponse, error) {
	whereBuilder := l.svcCtx.IndustryModel.RowBuilder().Where(squirrel.Eq{"parent_id": in.ParentId})
	industries, err := l.svcCtx.IndustryModel.FindChildrenList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	l.Logger.Slowf("result is %+v", industries)
	var list []*basic.GetIndustryChild
	if len(industries) > 0 {
		for _, industry := range industries {
			list = append(list, &basic.GetIndustryChild{
				Id:          industry.Id,
				IndustryId:  industry.IndustryId,
				Name:        industry.Name,
				ParentId:    industry.ParentId,
				LevelType:   industry.LevelType,
				Description: industry.Description,
			})
		}
	}

	return &basic.GetIndustriesResponse{
		List: list,
	}, nil
}
