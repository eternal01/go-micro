package category

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(req *types.GatewayUpdateCategoryRequest) (resp *types.GatewayUpdateCategoryReply, err error) {
	category, err := l.svcCtx.BasicCategoryRpc.UpdateCategory(l.ctx, &basic.UpdateCategoryRequest{
		Id:          req.Id,
		ParentId:    req.ParentId,
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return nil, err
	}

	return &types.GatewayUpdateCategoryReply{
		Id: category.Id,
	}, nil
}
