package category

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoriesLogic {
	return &GetCategoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoriesLogic) GetCategories(req *types.GatewayGetCategoriesRequest) (resp *types.GatewayGetCategoriesReply, err error) {
	if req.ParentId < 0 {
		return nil, errorx.ParamsError
	}

	categories, err := l.svcCtx.BasicCategoryRpc.GetCategories(l.ctx, &basic.GetCategoriesRequest{
		ParentId: req.ParentId,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.GatewayGetCategoriesReply)
	if len(categories.List) > 0 {
		for _, v := range categories.List {
			resp.List = append(resp.List, types.GatewayGetCategoriesChild{
				Id:          v.Id,
				ParentId:    v.ParentId,
				Name:        v.Name,
				Description: v.Description,
			})
		}
	}

	return
}
