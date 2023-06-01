package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIndustryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateIndustryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIndustryLogic {
	return &UpdateIndustryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateIndustryLogic) UpdateIndustry(req *types.GatewayUpdateIndustryRequest) (resp *types.GatewayUpdateIndustryReply, err error) {
	industry, err := l.svcCtx.BasicRpc.UpdateIndustry(l.ctx, &basic.UpdateIndustryRequest{
		Id:          req.Id,
		IndustryId:  req.IndustryId,
		Name:        req.Name,
		ParentId:    req.ParentId,
		LevelType:   req.LevelType,
		Description: req.Description,
	})

	if err != nil {
		return nil, err
	}

	return &types.GatewayUpdateIndustryReply{
		Id: industry.Id,
	}, nil
}
