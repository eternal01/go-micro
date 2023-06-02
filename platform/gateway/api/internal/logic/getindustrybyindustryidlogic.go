package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIndustryByIndustryIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIndustryByIndustryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIndustryByIndustryIdLogic {
	return &GetIndustryByIndustryIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIndustryByIndustryIdLogic) GetIndustryByIndustryId(req *types.GatewayGetIndustryByIndustryIdRequest) (resp *types.GatewayGetIndustryReply, err error) {
	if len(req.IndustryId) <= 0 {
		return nil, errorx.ParamsError
	}
	industry, err := l.svcCtx.BasicRpc.GetIndustryByIndustryId(l.ctx, &basic.GetIndustryByIndustryIdRequest{
		IndustryId: req.IndustryId,
	})
	if err != nil {
		l.Logger.Errorf("RPC-BASIC GetIndustryByIndustryId Error - %+v", err)
		return nil, err
	}

	return &types.GatewayGetIndustryReply{
		Id:          industry.Id,
		IndustryId:  industry.IndustryId,
		Name:        industry.Name,
		LevelType:   industry.LevelType,
		ParentId:    industry.ParentId,
		Description: industry.Description,
	}, nil
}
