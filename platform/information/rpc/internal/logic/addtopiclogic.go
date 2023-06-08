package logic

import (
	"context"
	"database/sql"

	"go-micro/platform/information/model"
	"go-micro/platform/information/rpc/information"
	"go-micro/platform/information/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicLogic {
	return &AddTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加文章
func (l *AddTopicLogic) AddTopic(in *information.AddTopicRequest) (*information.AddTopicResponse, error) {
	result, err := l.svcCtx.TopicModel.Insert(l.ctx, &model.Topic{
		Title:      in.Title,
		Img:        in.Img,
		Content:    sql.NullString{String: in.Content, Valid: true},
		Status:     model.TopicStatusNotAudited,
		StageId:    in.StageId,
		CategoryId: in.CategoryId,
		UserId:     0,
	})
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &information.AddTopicResponse{
		Id: id,
	}, nil
}
