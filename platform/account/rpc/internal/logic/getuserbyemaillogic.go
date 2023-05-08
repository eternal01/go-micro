package logic

import (
	"context"

	"go-micro/platform/account/rpc/account"
	"go-micro/platform/account/rpc/internal/svc"
	"go-micro/platform/basic/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByEmailLogic {
	return &GetUserByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByEmailLogic) GetUserByEmail(in *account.GetUserByEmailRequest) (*account.GetUserResponse, error) {
	// todo: add your logic here and delete this line
	user, error := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
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
