package logic

import (
	"context"

	"go-micro/platform/account/rpc/account"
	"go-micro/platform/account/rpc/internal/svc"
	"go-micro/platform/basic/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *account.GetUserRequest) (*account.GetUserResponse, error) {
	// todo: add your logic here and delete this line

	user, error := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if error != nil {
		if error == model.ErrNotFound {
			return &account.GetUserResponse{}, nil
		}
		return nil, error
	}
	return &account.GetUserResponse{
		Id:       user.Id,
		UserName: user.UserName,
		NickName: user.NickName,
		Gender:   user.Gender,
		Mobile:   user.Mobile,
		Email:    user.Email,
		Status:   user.Status,
	}, nil
}
