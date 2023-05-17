package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRegionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegionLogic {
	return &GetRegionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRegionLogic) GetRegion(req *types.GatewayGetRegionRequest) (resp *types.GatewayGetRegionReply, err error) {
	// todo: add your logic here and delete this line
	if req.Id <= 0 {
		return nil, errorx.ParamsError
	}
	region, err := l.svcCtx.BasicRpc.GetRegion(l.ctx, &basic.GetRegionRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("RPC-BASIC GetRegion Error - %+v", err)
		return nil, err
	}

	return &types.GatewayGetRegionReply{
		Id:       region.Id,
		ParentId: region.ParentId,
		Name:     region.Name,
	}, nil
}
