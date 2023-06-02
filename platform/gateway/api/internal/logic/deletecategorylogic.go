package logic

import (
	"context"

	"go-micro/common/cache"
	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
	"go-micro/platform/gateway/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(req *types.GatewayDeleteCategoryRequest) (resp *types.GatewayDeleteCategoryReply, err error) {
	categories, err := l.svcCtx.BasicRpc.GetCategories(l.ctx, &basic.GetCategoriesRequest{
		ParentId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if len(categories.List) > 0 {
		return nil, errorx.ExistChildrenError
	}

	_, err = l.svcCtx.BasicRpc.DeleteCategory(l.ctx, &basic.DeleteCategoryRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	cache.GetSingleton(l.svcCtx.Config.CacheRedis[0].RedisConf).DelRedisCache(model.CategoryTreeCacheKey)
	return
}
