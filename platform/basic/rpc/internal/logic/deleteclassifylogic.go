package logic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClassifyLogic {
	return &DeleteClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除分类信息
func (l *DeleteClassifyLogic) DeleteClassify(in *basic.DeleteClassifyRequest) (*basic.DeleteClassifyResponse, error) {
	err := l.svcCtx.ClassifyModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &basic.DeleteClassifyResponse{}, nil
}
