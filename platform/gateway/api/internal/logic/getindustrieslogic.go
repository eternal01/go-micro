package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIndustriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIndustriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIndustriesLogic {
	return &GetIndustriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIndustriesLogic) GetIndustries(req *types.GatewayGetIndustriesRequest) (resp *types.GatewayGetIndustriesReply, err error) {
	if len(req.ParentId) <= 0 {
		return nil, errorx.ParamsError
	}

	industries, err := l.svcCtx.BasicRpc.GetIndustries(l.ctx, &basic.GetIndustriesRequest{
		ParentId: req.ParentId,
	})
	if err != nil {
		logx.Errorf("RPC-BASIC GetIndustries Error - %+v", err)
		return nil, err
	}
	resp = new(types.GatewayGetIndustriesReply)
	if len(industries.List) > 0 {
		for _, v := range industries.List {
			resp.List = append(resp.List, types.GatewayGetIndustriesChild{
				Id:          v.Id,
				IndustryId:  v.IndustryId,
				Name:        v.Name,
				ParentId:    v.ParentId,
				LevelType:   v.LevelType,
				Description: v.Description,
			})
		}
	}

	return
}
