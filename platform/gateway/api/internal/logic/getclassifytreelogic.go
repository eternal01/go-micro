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

type GetClassifyTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassifyTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassifyTreeLogic {
	return &GetClassifyTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassifyTreeLogic) GetClassifyTree(req *types.GatewayGetClassifyTreeRequest) (resp *types.GatewayGetClassifyTreeReply, err error) {
	resp = new(types.GatewayGetClassifyTreeReply)

	r := cache.GetSingleton(l.svcCtx.Config.CacheRedis[0].RedisConf)
	v, err := r.GetRedisCache(model.ClassifyTreeCacheKey)
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

	node := types.GatewayClassifyTreeNode{}

	var generateTree func(node *types.GatewayClassifyTreeNode)
	generateTree = func(node *types.GatewayClassifyTreeNode) {
		nodes, err := l.svcCtx.BasicRpc.GetClassifies(l.ctx, &basic.GetClassifiesRequest{
			ParentId: node.Id,
		})
		if err != nil {
			logx.Errorf("RPC-BASIC GetClassifies Error - %+v", err)
			return
		}
		if len(nodes.List) > 0 {
			for _, v := range nodes.List {
				treeNode := types.GatewayClassifyTreeNode{
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
	r.SetRedisCache(l.ctx, model.ClassifyTreeCacheKey, string(cacheTree))

	return
}
