package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCategoryLogic) AddCategory(req *types.GatewayAddCategoryRequest) (resp *types.GatewayAddCategoryReply, err error) {
	if req.ParentId < 0 || len(req.Name) <= 0 {
		return nil, errorx.ParamsError
	}
	result, err := l.svcCtx.BasicRpc.AddCategory(l.ctx, &basic.AddCategoryRequest{
		Name:        req.Name,
		ParentId:    req.ParentId,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &types.GatewayAddCategoryReply{
		Id: result.Id,
	}, nil
}
