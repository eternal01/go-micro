package logic

import (
	"context"

	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRegionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRegionLogic {
	return &AddRegionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRegionLogic) AddRegion(req *types.GatewayAddRegionRequest) (resp *types.GatewayAddRegionReply, err error) {
	// todo: add your logic here and delete this line

	return
}
