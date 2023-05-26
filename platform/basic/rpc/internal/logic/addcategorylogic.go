package logic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加类别信息
func (l *AddCategoryLogic) AddCategory(in *basic.AddCategoryRequest) (*basic.AddCategoryResponse, error) {
	category := new(model.SystemCategory)
	category.ParentId = in.ParentId
	category.Name = in.Name
	category.Description = in.Description

	result, err := l.svcCtx.CategoryModel.Insert(l.ctx, category)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &basic.AddCategoryResponse{
		Id: id,
	}, nil
}
