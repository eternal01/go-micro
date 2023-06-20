package classify

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

type DeleteClassifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClassifyLogic {
	return &DeleteClassifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteClassifyLogic) DeleteClassify(req *types.GatewayDeleteClassifyRequest) (resp *types.GatewayDeleteClassifyReply, err error) {
	classifies, err := l.svcCtx.BasicClassifyRpc.GetClassifies(l.ctx, &basic.GetClassifiesRequest{
		ParentId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if len(classifies.List) > 0 {
		return nil, errorx.ExistChildrenError
	}
	_, err = l.svcCtx.BasicClassifyRpc.DeleteClassify(l.ctx, &basic.DeleteClassifyRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	cache.GetSingleton(l.svcCtx.Config.CacheRedis[0].RedisConf).DelRedisCache(model.ClassifyTreeCacheKey)
	return
}
