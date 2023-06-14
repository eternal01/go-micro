package basicindustryservicelogic

import (
	"context"

	"go-micro/platform/basic/rpc/basic"
	"go-micro/platform/basic/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteIndustryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteIndustryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteIndustryLogic {
	return &DeleteIndustryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除职业信息
func (l *DeleteIndustryLogic) DeleteIndustry(in *basic.DeleteIndustryRequest) (*basic.DeleteIndustryResponse, error) {
	err := l.svcCtx.IndustryModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &basic.DeleteIndustryResponse{}, nil
}
