package logic

import (
	"context"

	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
	"go-micro/platform/information/rpc/information"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicLogic {
	return &UpdateTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTopicLogic) UpdateTopic(req *types.GatewayUpdateTopicRequest) (resp *types.GatewayUpdateTopicReply, err error) {
	topic, err := l.svcCtx.InformationRpc.UpdateTopic(l.ctx, &information.UpdateTopicRequest{
		Id:         req.Id,
		Title:      req.Title,
		Img:        req.Img,
		Content:    req.Content,
		Status:     req.Status,
		StageId:    req.StageId,
		CategoryId: req.CategoryId,
	})
	if err != nil {
		return nil, err
	}
	return &types.GatewayUpdateTopicReply{
		Id: topic.Id,
	}, nil
}
