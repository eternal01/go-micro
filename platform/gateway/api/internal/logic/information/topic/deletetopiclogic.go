package topic

import (
	"context"
	"time"

	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
	"go-micro/platform/information/rpc/information"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicLogic {
	return &DeleteTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTopicLogic) DeleteTopic(req *types.GatewayDeleteTopicRequest) (resp *types.GatewayDeleteTopicReply, err error) {
	topic, err := l.svcCtx.InformationTopicRpc.GetTopic(l.ctx, &information.GetTopicRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.InformationTopicRpc.UpdateTopic(l.ctx, &information.UpdateTopicRequest{
		Id:         topic.Id,
		Title:      topic.Title,
		Img:        topic.Img,
		Content:    topic.Content,
		Status:     topic.Status,
		StageId:    topic.StageId,
		CategoryId: topic.CategoryId,
		UserId:     topic.UserId,
		AuditorId:  topic.AuditorId,
		DeletedAt:  time.Now().Format(time.DateTime),
	})
	if err != nil {
		return nil, err
	}
	return &types.GatewayDeleteTopicReply{}, nil
}
