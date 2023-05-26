package logic

import (
	"context"

	"go-micro/platform/basic/model"
	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassifyLogic {
	return &AddClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加分类信息
func (l *AddClassifyLogic) AddClassify(in *basic.AddClassifyRequest) (*basic.AddClassifyResponse, error) {
	classify := new(model.SystemClassify)
	classify.ParentId = in.ParentId
	classify.Name = in.Name
	classify.Description = in.Description

	result, err := l.svcCtx.ClassifyModel.Insert(l.ctx, classify)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &basic.AddClassifyResponse{
		Id: id,
	}, nil
}
