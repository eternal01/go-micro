package informationtopicservicelogic

import (
	"context"

	"go-micro/platform/information/rpc/information"
	"go-micro/platform/information/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTopicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTopicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicLogic {
	return &DeleteTopicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除文章
func (l *DeleteTopicLogic) DeleteTopic(in *information.DeleteTopicRequest) (*information.DeleteTopicResponse, error) {
	err := l.svcCtx.TopicModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &information.DeleteTopicResponse{}, nil
}
