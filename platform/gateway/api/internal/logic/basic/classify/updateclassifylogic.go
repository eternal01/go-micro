package classify

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClassifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClassifyLogic {
	return &UpdateClassifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateClassifyLogic) UpdateClassify(req *types.GatewayUpdateClassifyRequest) (resp *types.GatewayUpdateClassifyReply, err error) {
	classify, err := l.svcCtx.BasicClassifyRpc.UpdateClassify(l.ctx, &basic.UpdateClassifyRequest{
		Id:          req.Id,
		ParentId:    req.ParentId,
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		return nil, err
	}

	return &types.GatewayUpdateClassifyReply{
		Id: classify.Id,
	}, nil
}
