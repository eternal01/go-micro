package topic

import (
	"context"

	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
	"go-micro/platform/information/rpc/information"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopicListLogic {
	return &GetTopicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopicListLogic) GetTopicList(req *types.GatewayGetTopicListRequest) (resp *types.GatewayGetTopicListReply, err error) {
	list, err := l.svcCtx.InformationTopicRpc.GetTopicList(l.ctx, &information.GetTopicsRequest{
		Page:      req.Page,
		Size:      req.Size,
		LastId:    req.LastId,
		Keyword:   req.Keyword,
		SortField: req.SortField,
		SortType:  req.SortType,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.GatewayGetTopicListReply)
	if len(list.List) > 0 {
		for _, item := range list.List {
			resp.List = append(resp.List, types.GatewayListTopic{
				Id:         item.Id,
				Img:        item.Img,
				Content:    item.Content,
				Status:     item.Status,
				StageId:    item.StageId,
				CategoryId: item.CategoryId,
				UserId:     item.UserId,
				AuditorId:  item.AuditorId,
				CreatedAt:  item.CreatedAt,
				UpdatedAt:  item.UpdatedAt,
				AuditedAt:  item.AuditedAt,
				DeletedAt:  item.DeletedAt,
			})
		}
	}
	resp.Page = list.Page
	resp.Size = list.Size
	resp.Total = list.Total

	return
}
