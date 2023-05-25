package logic

import (
	"context"
	"encoding/json"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
	"go-micro/platform/gateway/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
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
	r, err := redis.NewRedis(redis.RedisConf{
		Host: "redis:6379",
		Type: redis.NodeType,
		Pass: "secret_redis",
	})

	if err != nil {
		return
	}
	v, err := r.Get(model.IndustryTreeCacheKey)
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

	industry := types.GatewayIndustryTreeNode{}
	var industryTree func(node *types.GatewayIndustryTreeNode)

	industryTree = func(node *types.GatewayIndustryTreeNode) {
		industries, err := l.svcCtx.BasicRpc.GetIndustries(l.ctx, &basic.GetIndustriesRequest{
			ParentId: node.IndustryId,
		})
		if err != nil {
			logx.Errorf("RPC-BASIC GetIndustries Error - %+v", err)
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
	r.Set(model.IndustryTreeCacheKey, string(cacheTree))

	return
}
