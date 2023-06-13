package basiccategoryservicelogic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新类别信息
func (l *UpdateCategoryLogic) UpdateCategory(in *basic.UpdateCategoryRequest) (*basic.UpdateCategoryResponse, error) {
	err := l.svcCtx.CategoryModel.Update(l.ctx, &model.SystemCategory{
		Id:          in.Id,
		Name:        in.Name,
		ParentId:    in.ParentId,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}

	return &basic.UpdateCategoryResponse{
		Id: in.Id,
	}, nil
}
