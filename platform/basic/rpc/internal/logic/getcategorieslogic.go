package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoriesLogic {
	return &GetCategoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据父级id获取类别信息
func (l *GetCategoriesLogic) GetCategories(in *basic.GetCategoriesRequest) (*basic.GetCategoriesResponse, error) {
	whereBuilder := l.svcCtx.CategoryModel.RowBuilder().Where(squirrel.Eq{"parent_id": in.ParentId})
	categories, err := l.svcCtx.CategoryModel.FindChildrenList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	l.Logger.Slowf("result is %+v", categories)
	var list []*basic.GetCategoryChild
	if len(categories) > 0 {
		for _, classify := range categories {
			list = append(list, &basic.GetCategoryChild{
				Id:          classify.Id,
				ParentId:    classify.ParentId,
				Name:        classify.Name,
				Description: classify.Description,
			})
		}
	}

	return &basic.GetCategoriesResponse{
		List: list,
	}, nil
}
