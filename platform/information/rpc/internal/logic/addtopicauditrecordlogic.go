package logic

import (
	"context"

	"go-micro/platform/information/model"
	"go-micro/platform/information/rpc/information"
	"go-micro/platform/information/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTopicAuditRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTopicAuditRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicAuditRecordLogic {
	return &AddTopicAuditRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加文章审核记录
func (l *AddTopicAuditRecordLogic) AddTopicAuditRecord(in *information.AddTopicAuditRecordRequest) (*information.AddTopicAuditRecordResponse, error) {
	record, err := l.svcCtx.TopicAuditRecordModel.Insert(l.ctx, &model.TopicAuditRecord{
		TopicId: in.TopicId,
		Status:  in.Status,
		Remark:  in.Remark,
	})
	if err != nil {
		return nil, err
	}
	id, err := record.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &information.AddTopicAuditRecordResponse{
		Id: id,
	}, nil
}
