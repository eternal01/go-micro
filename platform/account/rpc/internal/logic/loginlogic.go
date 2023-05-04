package logic

import (
	"context"

	"go-micro/common/errorx"
	"go-micro/platform/account/rpc/account"
	"go-micro/platform/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *account.LoginRequest) (*account.LoginResponse, error) {
	// todo: add your logic here and delete this line
	var user *account.GetUserResponse
	var err error
	if len(in.Mobile) == 0 && len(in.Email) == 0 {
		return nil, errorx.NewDefaultError("请填写账户")
	}
	if len(in.Mobile) > 0 {
		if user, err = NewGetUserByMobileLogic(l.ctx, l.svcCtx).GetUserByMobile(&account.GetUserByMobileRequest{
			Mobile: in.Mobile,
		}); err != nil {
			return nil, err
		}
	} else {
		if user, err = NewGetUserByEmailLogic(l.ctx, l.svcCtx).GetUserByEmail(&account.GetUserByEmailRequest{
			Email: in.Email,
		}); err != nil {
			return nil, err
		}
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := generateTokenLogic.GenerateToken(&account.GenerateTokenRequest{
		UserId: user.Id,
	})
	if err != nil {
		return nil, err
	}

	return &account.LoginResponse{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
		RefreshAfter: token.RefreshAfter,
	}, nil
}
