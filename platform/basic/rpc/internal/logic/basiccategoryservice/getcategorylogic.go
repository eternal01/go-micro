package basiccategoryservicelogic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取类别信息
func (l *GetCategoryLogic) GetCategory(in *basic.GetCategoryRequest) (*basic.GetCategoryResponse, error) {
	category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &basic.GetCategoryResponse{
		Id:          category.Id,
		Name:        category.Name,
		ParentId:    category.ParentId,
		Description: category.Description,
	}, nil
}
