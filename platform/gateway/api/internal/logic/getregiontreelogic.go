package logic

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

type GetRegionTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRegionTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRegionTreeLogic {
	return &GetRegionTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRegionTreeLogic) GetRegionTree(req *types.GatewayGetRegionTreeRequest) (resp *types.GatewayGetRegionTreeReply, err error) {

	resp = new(types.GatewayGetRegionTreeReply)
	r := cache.GetSingleton(l.svcCtx.Config.CacheRedis[0].RedisConf)
	v, err := r.GetRedisCache(model.RegionTreeCacheKey)
	if err != nil {
		l.Logger.Errorf("REDIS CACHE ERR - %+v", err)
		return resp, errorx.CacheError
	}
	if len(v) > 0 {
		err = json.Unmarshal([]byte(v), resp)
		if err != nil {
			l.Logger.Errorf("REDIS CACHE ERR - %+v", err)
			return resp, errorx.CacheError
		}

		return
	}

	region := types.GatewayRegionTreeNode{}
	var regionTree func(node *types.GatewayRegionTreeNode)

	regionTree = func(node *types.GatewayRegionTreeNode) {
		regions, err := l.svcCtx.BasicRpc.GetRegions(l.ctx, &basic.GetRegionsRequest{
			ParentId: node.Id,
		})
		if err != nil {
			logx.Errorf("RPC-BASIC GetRegions Error - %+v", err)
			return
		}
		if len(regions.List) > 0 {
			for _, v := range regions.List {
				region := types.GatewayRegionTreeNode{
					Id:       v.Id,
					ParentId: v.ParentId,
					Name:     v.Name,
				}
				regionTree(&region)
				node.Children = append(node.Children, region)
			}
		}
	}
	regionTree(&region)
	resp.Tree = region.Children

	cacheTree, err := json.Marshal(*resp)
	r.SetRedisCache(model.RegionTreeCacheKey, string(cacheTree))

	return
}
