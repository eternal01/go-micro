package informationtopicservicelogic

import (
	"context"
	"database/sql"
	"time"

	"go-micro/platform/information/model"
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
	topic := model.Topic{
		Id:         in.Id,
		Title:      in.Title,
		Img:        in.Img,
		Content:    sql.NullString{String: in.Content, Valid: true},
		Status:     in.Status,
		StageId:    in.StageId,
		CategoryId: in.CategoryId,
		UserId:     in.UserId,
	}
	if in.AuditorId != 0 {
		topic.AuditorId = in.AuditorId
		topic.AuditedAt = sql.NullTime{Time: time.Now(), Valid: true}
	}
	if len(in.DeletedAt) != 0 {
		deleteAt, err := time.ParseInLocation(time.DateTime, in.DeletedAt, time.FixedZone("CST", 8*3600))
		if err != nil {
			return nil, err
		}
		topic.DeletedAt = sql.NullTime{Time: deleteAt, Valid: true}
	}
	err := l.svcCtx.TopicModel.Update(l.ctx, &topic)
	if err != nil {
		return nil, err
	}

	return &information.UpdateTopicResponse{
		Id: in.Id,
	}, nil
}
