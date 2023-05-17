package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/account/rpc/accountclient"
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
	if req.Id == 0 {
		return nil, errorx.ParamsError
	}
	var user *accountclient.GetUserResponse
	user, err = l.svcCtx.AccountRpc.GetUser(l.ctx, &accountclient.GetUserRequest{
		Id: req.Id,
	})
	if err != nil {
		l.Logger.Errorf("RPC-ACCOUNT GetUser Error - %+v", err)
		return nil, err
	}
	if user.Id == 0 {
		return nil, errorx.UserNotFound
	}
	return &types.GatewayGetUserResponse{
		Id:       user.Id,
		Avatar:   user.Avatar,
		Mobile:   user.Mobile,
		Email:    user.Email,
		UserName: user.UserName,
		NickName: user.NickName,
		Gender:   user.Gender,
		Status:   user.Status,
	}, nil
}
