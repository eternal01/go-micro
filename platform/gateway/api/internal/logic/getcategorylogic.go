package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryLogic) GetCategory(req *types.GatewayGetCategoryRequest) (resp *types.GatewayGetCategoryReply, err error) {
	if req.Id <= 0 {
		return nil, errorx.ParamsError
	}
	category, err := l.svcCtx.BasicRpc.GetCategory(l.ctx, &basic.GetCategoryRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("RPC-BASIC GetCategory Error - %+v", err)
		return nil, err
	}

	return &types.GatewayGetCategoryReply{
		Id:          category.Id,
		Name:        category.Name,
		ParentId:    category.ParentId,
		Description: category.Description,
	}, nil
}
