package informationtopicservicelogic

import (
	"context"
	"time"

	"go-micro/platform/information/rpc/information"
	"go-micro/platform/information/rpc/internal/svc"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicListLogic {
	return &GetTopicListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取文章列表
func (l *GetTopicListLogic) GetTopicList(in *information.GetTopicsRequest) (*information.GetTopicsResponse, error) {
	countBuilder := l.svcCtx.TopicModel.CountBuilder()
	whereBuilder := l.svcCtx.TopicModel.RowBuilder()
	countBuilder = countBuilder.Where(squirrel.Eq{"deleted_at": nil})
	whereBuilder = whereBuilder.Where(squirrel.Eq{"deleted_at": nil})
	count, err := l.svcCtx.TopicModel.Count(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}
	topics, err := l.svcCtx.TopicModel.FindList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	var list []*information.Topic
	var auditedAt string
	var deletedAt string
	if len(topics) > 0 {
		for _, topic := range topics {
			if topic.AuditedAt.Valid {
				auditedAt = topic.AuditedAt.Time.Format(time.DateTime)
			}
			if topic.DeletedAt.Valid {
				deletedAt = topic.DeletedAt.Time.Format(time.DateTime)
			}

			list = append(list, &information.Topic{
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
			})
		}
	}

	return &information.GetTopicsResponse{
		List:  list,
		Page:  in.Page,
		Size:  in.Size,
		Total: count,
	}, nil
}
