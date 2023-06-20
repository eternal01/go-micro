package topic

import (
	"context"

	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
	"go-micro/platform/information/rpc/information"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicLogic {
	return &GetTopicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicLogic) GetTopic(req *types.GatewayGetTopicRequest) (resp *types.GatewayGetTopicReply, err error) {
	topic, err := l.svcCtx.InformationTopicRpc.GetTopic(l.ctx, &information.GetTopicRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.GatewayGetTopicReply{
		Id:         topic.Id,
		Title:      topic.Title,
		Img:        topic.Img,
		Content:    topic.Content,
		Status:     topic.Status,
		StageId:    topic.StageId,
		CategoryId: topic.CategoryId,
		UserId:     topic.UserId,
		AuditorId:  topic.AuditorId,
		CreatedAt:  topic.CreatedAt,
		UpdatedAt:  topic.UpdatedAt,
		AuditedAt:  topic.AuditedAt,
		DeletedAt:  topic.DeletedAt,
	}, nil
}
