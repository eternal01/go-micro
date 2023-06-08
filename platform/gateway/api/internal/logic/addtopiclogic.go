package logic

import (
	"context"

	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
	"go-micro/platform/information/rpc/information"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicLogic {
	return &AddTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTopicLogic) AddTopic(req *types.GatewayAddTopicRequest) (resp *types.GatewayAddTopicReply, err error) {
	topic, err := l.svcCtx.InformationRpc.AddTopic(l.ctx, &information.AddTopicRequest{
		Title:      req.Title,
		Img:        req.Img,
		Content:    req.Content,
		StageId:    req.StageId,
		CategoryId: req.CategoryId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GatewayAddTopicReply{
		Id: topic.Id,
	}, nil
}
