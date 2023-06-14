package basicclassifyservicelogic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassifiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClassifiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassifiesLogic {
	return &GetClassifiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据父级id获取分类信息
func (l *GetClassifiesLogic) GetClassifies(in *basic.GetClassifiesRequest) (*basic.GetClassifiesResponse, error) {
	whereBuilder := l.svcCtx.ClassifyModel.RowBuilder().Where(squirrel.Eq{"parent_id": in.ParentId})
	classifies, err := l.svcCtx.ClassifyModel.FindChildrenList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	l.Logger.Slowf("result is %+v", classifies)
	var list []*basic.GetClassifyChild
	if len(classifies) > 0 {
		for _, classify := range classifies {
			list = append(list, &basic.GetClassifyChild{
				Id:          classify.Id,
				ParentId:    classify.ParentId,
				Name:        classify.Name,
				Description: classify.Description,
			})
		}
	}

	return &basic.GetClassifiesResponse{
		List: list,
	}, nil
}
