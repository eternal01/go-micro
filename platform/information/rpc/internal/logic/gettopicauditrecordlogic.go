package logic

import (
	"context"

	"go-micro/platform/information/rpc/information"
	"go-micro/platform/information/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicAuditRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopicAuditRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicAuditRecordLogic {
	return &GetTopicAuditRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取文章审核记录
func (l *GetTopicAuditRecordLogic) GetTopicAuditRecord(in *information.GetTopicAuditRecordRequest) (*information.GetTopicAuditRecordResponse, error) {
	// todo: add your logic here and delete this line

	return &information.GetTopicAuditRecordResponse{}, nil
}
