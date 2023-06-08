package logic

import (
	"context"
	"database/sql"

	"go-micro/platform/information/rpc/information"
	"go-micro/platform/information/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicLogic {
	return &UpdateTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 编辑文章
func (l *UpdateTopicLogic) UpdateTopic(in *information.UpdateTopicRequest) (*information.UpdateTopicResponse, error) {
	topic, err := l.svcCtx.TopicModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if in.Title != "" {
		topic.Title = in.Title
	}
	if in.Img != "" {
		topic.Img = in.Img
	}
	if in.Content != "" {
		topic.Content = sql.NullString{String: in.Content, Valid: true}
	}
	if in.Status != 0 {
		topic.Status = in.Status
	}
	if in.StageId != 0 {
		topic.StageId = in.StageId
	}
	if in.CategoryId != 0 {
		topic.CategoryId = in.CategoryId
	}

	err = l.svcCtx.TopicModel.Update(l.ctx, topic)
	if err != nil {
		return nil, err
	}

	return &information.UpdateTopicResponse{
		Id: in.Id,
	}, nil
}
