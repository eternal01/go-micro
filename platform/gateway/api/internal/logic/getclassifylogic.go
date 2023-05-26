package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassifyLogic {
	return &GetClassifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassifyLogic) GetClassify(req *types.GatewayGetClassifyRequest) (resp *types.GatewayGetClassifyReply, err error) {
	if req.Id <= 0 {
		return nil, errorx.ParamsError
	}
	classify, err := l.svcCtx.BasicRpc.GetClassify(l.ctx, &basic.GetClassifyRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("RPC-BASIC GetClassify Error - %+v", err)
		return nil, err
	}

	return &types.GatewayGetClassifyReply{
		Id:          classify.Id,
		Name:        classify.Name,
		ParentId:    classify.ParentId,
		Description: classify.Description,
	}, nil
}
