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

type AddIndustryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddIndustryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIndustryLogic {
	return &AddIndustryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddIndustryLogic) AddIndustry(req *types.GatewayAddIndustryRequest) (resp *types.GatewayAddIndustryReply, err error) {
	if len(req.ParentId) <= 0 || len(req.Name) <= 0 {
		return nil, errorx.ParamsError
	}
	result, err := l.svcCtx.BasicRpc.AddIndustry(l.ctx, &basic.AddIndustryRequest{
		IndustryId:  req.IndustryId,
		Name:        req.Name,
		ParentId:    req.ParentId,
		LevelType:   req.LevelType,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	cache.GetSingleton(l.svcCtx.Config.CacheRedis[0].RedisConf).DelRedisCache(model.IndustryTreeCacheKey)

	return &types.GatewayAddIndustryReply{
		Id: result.Id,
	}, nil
}
