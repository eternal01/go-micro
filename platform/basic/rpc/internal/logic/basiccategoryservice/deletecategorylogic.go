package basiccategoryservicelogic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除类别信息
func (l *DeleteCategoryLogic) DeleteCategory(in *basic.DeleteCategoryRequest) (*basic.DeleteCategoryResponse, error) {
	err := l.svcCtx.CategoryModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &basic.DeleteCategoryResponse{}, nil
}
