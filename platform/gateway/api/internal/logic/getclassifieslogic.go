package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassifiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassifiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassifiesLogic {
	return &GetClassifiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassifiesLogic) GetClassifies(req *types.GatewayGetClassifiesRequest) (resp *types.GatewayGetClassifiesReply, err error) {
	if req.ParentId < 0 {
		return nil, errorx.ParamsError
	}

	classifies, err := l.svcCtx.BasicRpc.GetClassifies(l.ctx, &basic.GetClassifiesRequest{
		ParentId: req.ParentId,
	})
	if err != nil {
		logx.Errorf("RPC-BASIC GetClassifies Error - %+v", err)
		return nil, err
	}
	resp = new(types.GatewayGetClassifiesReply)
	if len(classifies.List) > 0 {
		for _, v := range classifies.List {
			resp.List = append(resp.List, types.GatewayGetClassifiesChild{
				Id:          v.Id,
				ParentId:    v.ParentId,
				Name:        v.Name,
				Description: v.Description,
			})
		}
	}

	return
}
