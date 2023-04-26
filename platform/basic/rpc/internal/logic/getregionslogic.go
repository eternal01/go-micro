package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

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
	// todo: add your logic here and delete this line
	return &basic.GetRegionsResponse{
		List: []*basic.GetRegionChild{},
	}, nil
}
