package category

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

type GetCategoryTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryTreeLogic {
	return &GetCategoryTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryTreeLogic) GetCategoryTree(req *types.GatewayGetCategoryTreeRequest) (resp *types.GatewayGetCategoryTreeReply, err error) {
	resp = new(types.GatewayGetCategoryTreeReply)

	r := cache.GetSingleton(l.svcCtx.Config.CacheRedis[0].RedisConf)
	v, err := r.GetRedisCache(model.CategoryTreeCacheKey)
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

	node := types.GatewayCategoryTreeNode{}

	var generateTree func(node *types.GatewayCategoryTreeNode)
	generateTree = func(node *types.GatewayCategoryTreeNode) {
		nodes, err := l.svcCtx.BasicCategoryRpc.GetCategories(l.ctx, &basic.GetCategoriesRequest{
			ParentId: node.Id,
		})
		if err != nil {
			return
		}
		if len(nodes.List) > 0 {
			for _, v := range nodes.List {
				treeNode := types.GatewayCategoryTreeNode{
					Id:          v.Id,
					ParentId:    v.ParentId,
					Name:        v.Name,
					Description: v.Description,
				}
				generateTree(&treeNode)
				node.Children = append(node.Children, treeNode)
			}
		}
	}

	generateTree(&node)
	resp.Tree = node.Children

	cacheTree, err := json.Marshal(*resp)
	r.SetRedisCache(model.CategoryTreeCacheKey, string(cacheTree))

	return
}
