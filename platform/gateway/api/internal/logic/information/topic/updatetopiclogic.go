package topic

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
	topic, err := l.svcCtx.InformationTopicRpc.GetTopic(l.ctx, &information.GetTopicRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	title := topic.Title
	if len(req.Title) != 0 {
		title = req.Title
	}
	img := topic.Img
	if len(req.Img) != 0 {
		img = req.Img
	}
	content := topic.Content
	if len(req.Content) != 0 {
		content = req.Content
	}
	status := topic.Status
	if req.Status != 0 {
		status = req.Status
	}
	stageId := topic.StageId
	if req.StageId != 0 {
		stageId = req.StageId
	}
	categoryId := topic.CategoryId
	if req.CategoryId != 0 {
		categoryId = req.CategoryId
	}

	result, err := l.svcCtx.InformationTopicRpc.UpdateTopic(l.ctx, &information.UpdateTopicRequest{
		Id:         req.Id,
		Title:      title,
		Img:        img,
		Content:    content,
		Status:     status,
		StageId:    stageId,
		CategoryId: categoryId,
	})
	if err != nil {
		return nil, err
	}
	return &types.GatewayUpdateTopicReply{
		Id: result.Id,
	}, nil
}
