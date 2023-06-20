package industry

import (
	"context"
	"encoding/json"

	"go-micro/common/cache"
	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
	"go-micro/platform/gateway/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIndustryTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIndustryTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIndustryTreeLogic {
	return &GetIndustryTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIndustryTreeLogic) GetIndustryTree(req *types.GatewayGetIndustryTreeRequest) (resp *types.GatewayGetIndustryTreeReply, err error) {
	resp = new(types.GatewayGetIndustryTreeReply)
	r := cache.GetSingleton(l.svcCtx.Config.CacheRedis[0].RedisConf)
	v, err := r.GetRedisCache(model.IndustryTreeCacheKey)
	if err != nil {
		return resp, errorx.CacheError
	}
	if len(v) > 0 {
		err = json.Unmarshal([]byte(v), resp)
		if err != nil {
			return resp, errorx.CacheError
		}

		return
	}

	industry := types.GatewayIndustryTreeNode{
		IndustryId: "0",
	}
	var industryTree func(node *types.GatewayIndustryTreeNode)

	industryTree = func(node *types.GatewayIndustryTreeNode) {
		industries, err := l.svcCtx.BasicIndustryRpc.GetIndustries(l.ctx, &basic.GetIndustriesRequest{
			ParentId: node.IndustryId,
		})
		if err != nil {
			return
		}
		if len(industries.List) > 0 {
			for _, v := range industries.List {
				industry := types.GatewayIndustryTreeNode{
					Id:          v.Id,
					IndustryId:  v.IndustryId,
					Name:        v.Name,
					ParentId:    v.ParentId,
					LevelType:   v.LevelType,
					Description: v.Description,
				}
				industryTree(&industry)
				node.Children = append(node.Children, industry)
			}
		}
	}
	industryTree(&industry)
	resp.Tree = industry.Children

	cacheTree, err := json.Marshal(*resp)
	r.SetRedisCache(model.IndustryTreeCacheKey, string(cacheTree))

	return
}
