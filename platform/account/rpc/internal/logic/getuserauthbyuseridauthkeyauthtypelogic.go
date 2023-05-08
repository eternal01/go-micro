package logic

import (
	"context"

	"go-micro/platform/account/model"
	"go-micro/platform/account/rpc/account"
	"go-micro/platform/account/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByUserIdAuthKeyAuthTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByUserIdAuthKeyAuthTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByUserIdAuthKeyAuthTypeLogic {
	return &GetUserAuthByUserIdAuthKeyAuthTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByUserIdAuthKeyAuthTypeLogic) GetUserAuthByUserIdAuthKeyAuthType(in *account.GetUserAuthByUserIdAuthKeyAuthTypeRequest) (*account.GetUserAuthResponse, error) {
	// todo: add your logic here and delete this line
	userAuth, error := l.svcCtx.UserAuthModel.FindOneByUserIdAuthKeyAuthType(l.ctx, in.UserId, in.AuthKey, in.AuthType)
	if error != nil {
		if error == model.ErrNotFound {
			return &account.GetUserAuthResponse{}, nil
		}
		return nil, error
	}
	return &account.GetUserAuthResponse{
		Id:             userAuth.Id,
		UserId:         userAuth.UserId,
		AuthType:       userAuth.AuthType,
		AuthKey:        userAuth.AuthKey,
		AuthCredential: userAuth.AuthCredential,
	}, nil
}
