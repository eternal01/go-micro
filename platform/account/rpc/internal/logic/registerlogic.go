package logic

import (
	"context"

	"go-micro/platform/account/model"
	"go-micro/platform/account/rpc/account"
	"go-micro/platform/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *account.RegisterRequest) (*account.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	var user *model.Users
	var err error
	if len(in.Mobile) > 0 {
		user, err = l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	} else {
		user, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	}
	if err != nil {
		return nil, err
	}

	return &account.RegisterResponse{}, nil
}
