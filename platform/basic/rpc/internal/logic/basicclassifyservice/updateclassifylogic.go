package basicclassifyservicelogic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClassifyLogic {
	return &UpdateClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新分类信息
func (l *UpdateClassifyLogic) UpdateClassify(in *basic.UpdateClassifyRequest) (*basic.UpdateClassifyResponse, error) {
	err := l.svcCtx.ClassifyModel.Update(l.ctx, &model.SystemClassify{
		Id:          in.Id,
		Name:        in.Name,
		ParentId:    in.ParentId,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}

	return &basic.UpdateClassifyResponse{
		Id: in.Id,
	}, nil
}
