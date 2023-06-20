package industry

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

type DeleteIndustryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteIndustryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteIndustryLogic {
	return &DeleteIndustryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteIndustryLogic) DeleteIndustry(req *types.GatewayDeleteIndustryRequest) (resp *types.GatewayDeleteIndustryReply, err error) {
	industry, err := l.svcCtx.BasicIndustryRpc.GetIndustry(l.ctx, &basic.GetIndustryRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	industries, err := l.svcCtx.BasicIndustryRpc.GetIndustries(l.ctx, &basic.GetIndustriesRequest{
		ParentId: industry.IndustryId,
	})
	if err != nil {
		return nil, err
	}
	if len(industries.List) > 0 {
		return nil, errorx.ExistChildrenError
	}
	_, err = l.svcCtx.BasicIndustryRpc.DeleteIndustry(l.ctx, &basic.DeleteIndustryRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	cache.GetSingleton(l.svcCtx.Config.CacheRedis[0].RedisConf).DelRedisCache(model.IndustryTreeCacheKey)
	return
}
