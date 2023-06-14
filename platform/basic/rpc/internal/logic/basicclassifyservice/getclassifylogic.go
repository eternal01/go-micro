package basicclassifyservicelogic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassifyLogic {
	return &GetClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取分类信息
func (l *GetClassifyLogic) GetClassify(in *basic.GetClassifyRequest) (*basic.GetClassifyResponse, error) {
	classify, err := l.svcCtx.ClassifyModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &basic.GetClassifyResponse{
		Id:          classify.Id,
		Name:        classify.Name,
		ParentId:    classify.ParentId,
		Description: classify.Description,
	}, nil
}
