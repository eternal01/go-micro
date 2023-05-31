package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStagesLogic {
	return &GetStagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStagesLogic) GetStages(req *types.GatewayGetStagesRequest) (resp *types.GatewayGetStagesReply, err error) {
	categories, err := l.svcCtx.BasicRpc.GetStages(l.ctx, &basic.GetStagesRequest{
		ParentId: 0,
	})
	if err != nil {
		logx.Errorf("RPC-BASIC GetStages Error - %+v", err)
		return nil, err
	}
	resp = new(types.GatewayGetStagesReply)
	if len(categories.List) > 0 {
		for _, v := range categories.List {
			resp.List = append(resp.List, types.GatewayGetStagesChild{
				Id:          v.Id,
				Name:        v.Name,
				Description: v.Description,
			})
		}
	}

	return
}
