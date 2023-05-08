package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRegionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRegionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegionsLogic {
	return &GetRegionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRegionsLogic) GetRegions(req *types.GatewayGetRegionsRequest) (resp *types.GatewayGetRegionsReply, err error) {
	// todo: add your logic here and delete this line

	if req.ParentId <= 0 {
		return nil, errorx.ParamsError
	}

	regions, err := l.svcCtx.BasicRpc.GetRegions(l.ctx, &basic.GetRegionsRequest{
		ParentId: req.ParentId,
	})
	if err != nil {
		return nil, err
	}

	if len(regions.List) > 0 {
		for _, v := range regions.List {
			resp.List = append(resp.List, types.GatewayGetRegionsChild{
				Id:       v.Id,
				ParentId: v.ParentId,
				Name:     v.Name,
			})
		}
	}

	return
}
