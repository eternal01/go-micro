package logic

import (
	"context"

	"go-micro/platform/account/rpc/account"
	"go-micro/platform/account/rpc/internal/svc"
	"go-micro/platform/basic/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByMobileLogic {
	return &GetUserByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByMobileLogic) GetUserByMobile(in *account.GetUserByMobileRequest) (*account.GetUserResponse, error) {
	// todo: add your logic here and delete this line
	user, error := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if error != nil && error != model.ErrNotFound {
		return nil, error
	}
	return &account.GetUserResponse{
		Id:       user.Id,
		UserName: user.UserName,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Gender:   user.Gender,
		Mobile:   user.Mobile,
		Email:    user.Email,
		Status:   user.Status,
	}, nil
}
