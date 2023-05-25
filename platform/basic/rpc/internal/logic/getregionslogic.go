package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRegionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRegionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegionsLogic {
	return &GetRegionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRegionsLogic) GetRegions(in *basic.GetRegionsRequest) (*basic.GetRegionsResponse, error) {
	whereBuilder := l.svcCtx.RegionModel.RowBuilder().Where(squirrel.Eq{"parent_id": in.ParentId})
	regions, err := l.svcCtx.RegionModel.FindChildrenList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	l.Logger.Slowf("result is %+v", regions)
	var list []*basic.GetRegionChild
	if len(regions) > 0 {
		for _, region := range regions {
			list = append(list, &basic.GetRegionChild{
				Id:       region.Id,
				ParentId: region.ParentId,
				Name:     region.Name,
			})
		}
	}

	return &basic.GetRegionsResponse{
		List: list,
	}, nil
}
