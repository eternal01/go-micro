package logic

import (
	"context"
	"strings"

	"go-micro/common/errorx"
	"go-micro/common/tool"
	"go-micro/platform/account/model"
	"go-micro/platform/account/rpc/accountclient"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.GatewayRegisterRequest) (resp *types.GatewayRegisterReply, err error) {
	email := strings.TrimSpace(req.Email)
	mobile := strings.TrimSpace(req.Mobile)
	if len(email) == 0 && len(mobile) == 0 {
		return nil, errorx.ParamsError
	}
	var user *accountclient.GetUserResponse
	var authKey string
	if len(email) > 0 {
		authKey = email
		user, err = l.svcCtx.AccountRpc.GetUserByEmail(l.ctx, &accountclient.GetUserByEmailRequest{
			Email: email,
		})
	}
	if len(mobile) > 0 {
		authKey = mobile
		user, err = l.svcCtx.AccountRpc.GetUserByMobile(l.ctx, &accountclient.GetUserByMobileRequest{
			Mobile: mobile,
		})
	}

	if err != nil {
		return nil, err
	}

	if user.Id != 0 {
		return nil, errorx.UserExistError
	}

	hash, err := tool.Encryption(req.AuthCredential)
	if err != nil {
		return nil, err
	}
	respUser, err := l.svcCtx.AccountRpc.Register(l.ctx, &accountclient.RegisterRequest{
		Mobile:         mobile,
		Email:          email,
		AuthType:       model.AuthTypeSystem,
		AuthKey:        authKey,
		AuthCredential: string(hash),
	})
	if err != nil {
		logx.Errorf("RPC-ACCOUNT Register Error - %+v", err)
		return nil, err
	}

	return &types.GatewayRegisterReply{
		AccessToken:  respUser.AccessToken,
		AccessExpire: respUser.AccessExpire,
		RefreshAfter: respUser.RefreshAfter,
	}, nil
}
