package logic

import (
	"context"

	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GatewayGetUserRequest) (resp *types.GatewayGetUserResponse, err error) {
	// todo: add your logic here and delete this line

	// if req.Id == 0 {
	// 	return nil, errorx.NewDefaultError("参数错误")
	// }
	// var user *accountclient.GetUserResponse
	// user, err = l.svcCtx.AccountRpc.GetUser(l.ctx, &accountclient.GetUserRequest{
	// 	Id: req.Id,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// if user == nil {
	// 	return nil, errorx.NewDefaultError("用户不存在")
	// }
	// return &types.GatewayGetUserResponse{
	// 	Id:     user.Id,
	// 	Mobile: user.Mobile,
	// 	Email:  user.Email,
	// }, nil

	return &types.GatewayGetUserResponse{
		Id:     1,
		Mobile: "1223333",
	}, nil
}
