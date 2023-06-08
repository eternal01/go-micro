package logic

import (
	"context"
	"time"

	"go-micro/platform/information/rpc/information"
	"go-micro/platform/information/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicLogic {
	return &GetTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取文章
func (l *GetTopicLogic) GetTopic(in *information.GetTopicRequest) (*information.GetTopicResponse, error) {
	topic, err := l.svcCtx.TopicModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	var auditedAt string
	if topic.AuditedAt.Valid {
		auditedAt = topic.AuditedAt.Time.Format(time.DateTime)
	}

	var deletedAt string
	if topic.DeletedAt.Valid {
		deletedAt = topic.DeletedAt.Time.Format(time.DateTime)
	}

	return &information.GetTopicResponse{
		Id:         topic.Id,
		Title:      topic.Title,
		Img:        topic.Img,
		Content:    topic.Content.String,
		Status:     topic.Status,
		StageId:    topic.StageId,
		CategoryId: topic.CategoryId,
		UserId:     topic.UserId,
		AuditorId:  topic.AuditorId,
		CreatedAt:  topic.CreatedAt.Format(time.DateTime),
		UpdatedAt:  topic.UpdatedAt.Format(time.DateTime),
		AuditedAt:  auditedAt,
		DeletedAt:  deletedAt,
	}, nil
}
