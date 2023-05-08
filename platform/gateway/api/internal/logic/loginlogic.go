package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/account/model"
	"go-micro/platform/account/rpc/accountclient"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.GatewayLoginRequest) (resp *types.GatewayLoginResponse, err error) {
	// todo: add your logic here and delete this line
	if len(req.Email) == 0 && len(req.Mobile) == 0 {
		return nil, errorx.ParamsError
	}
	var user *accountclient.GetUserResponse
	var authKey string
	if len(req.Email) > 0 {
		authKey = req.Email
		user, err = l.svcCtx.AccountRpc.GetUserByEmail(l.ctx, &accountclient.GetUserByEmailRequest{
			Email: req.Email,
		})
	}
	if len(req.Mobile) > 0 {
		authKey = req.Mobile
		user, err = l.svcCtx.AccountRpc.GetUserByMobile(l.ctx, &accountclient.GetUserByMobileRequest{
			Mobile: req.Mobile,
		})
	}
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, errorx.UserNotFound
	}

	userAuth, err := l.svcCtx.AccountRpc.GetUserAuthByUserIdAuthKeyAuthType(l.ctx, &accountclient.GetUserAuthByUserIdAuthKeyAuthTypeRequest{
		UserId:   user.Id,
		AuthKey:  authKey,
		AuthType: model.AuthTypeSystem,
	})
	if err != nil {
		return nil, err
	}

	if userAuth.Id == 0 {
		return nil, errorx.UserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(userAuth.AuthCredential), []byte(req.AuthCredential))
	if err != nil {
		return nil, errorx.PasswordError
	}

	jwtToken, err := l.svcCtx.AccountRpc.Login(l.ctx, &accountclient.LoginRequest{
		Mobile:         req.Mobile,
		Email:          req.Email,
		AuthCredential: req.AuthCredential,
	})
	if err != nil {
		return nil, err
	}

	return &types.GatewayLoginResponse{
		Id:     user.Id,
		Mobile: user.Mobile,
		Email:  user.Email,

		AccessToken:  jwtToken.AccessToken,
		AccessExpire: jwtToken.AccessExpire,
		RefreshAfter: jwtToken.RefreshAfter,
	}, nil
}
