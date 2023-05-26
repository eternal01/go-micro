package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassifyLogic {
	return &AddClassifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClassifyLogic) AddClassify(req *types.GatewayAddClassifyRequest) (resp *types.GatewayAddClassifyReply, err error) {
	if req.ParentId < 0 || len(req.Name) <= 0 {
		return nil, errorx.ParamsError
	}
	result, err := l.svcCtx.BasicRpc.AddClassify(l.ctx, &basic.AddClassifyRequest{
		Name:        req.Name,
		ParentId:    req.ParentId,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &types.GatewayAddClassifyReply{
		Id: result.Id,
	}, nil
}
