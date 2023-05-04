package logic

import (
	"context"
	"strings"

	"go-micro/common/errorx"
	"go-micro/platform/account/rpc/accountclient"
	"go-micro/platform/gateway/api/internal/svc"
	"go-micro/platform/gateway/api/internal/types"
	"go-micro/platform/gateway/model"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
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

func (l *RegisterLogic) Register(req *types.GatewayRegisterRequest) (resp *types.GatewayRegisterResponse, err error) {
	// todo: add your logic here and delete this line
	email := strings.TrimSpace(req.Email)
	mobile := strings.TrimSpace(req.Mobile)
	if len(email) == 0 && len(mobile) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
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
	if user != nil || user.Id != 0 {
		return nil, errorx.NewDefaultError("用户已存在")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.AuthCredential), bcrypt.DefaultCost)
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
		return nil, err
	}

	return &types.GatewayRegisterResponse{
		AccessToken:  respUser.AccessToken,
		AccessExpire: respUser.AccessExpire,
		RefreshAfter: respUser.RefreshAfter,
	}, nil
}
