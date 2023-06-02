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

type DeleteRegionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRegionLogic {
	return &DeleteRegionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRegionLogic) DeleteRegion(req *types.GatewayDeleteRegionRequest) (resp *types.GatewayDeleteRegionReply, err error) {
	regions, err := l.svcCtx.BasicRpc.GetRegions(l.ctx, &basic.GetRegionsRequest{
		ParentId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if len(regions.List) > 0 {
		return nil, errorx.ExistChildrenError
	}
	_, err = l.svcCtx.BasicRpc.DeleteRegion(l.ctx, &basic.DeleteRegionRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	cache.GetSingleton(l.svcCtx.Config.CacheRedis[0].RedisConf).DelRedisCache(model.RegionTreeCacheKey)
	return
}
