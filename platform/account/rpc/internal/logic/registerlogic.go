package logic

import (
	"context"
	"database/sql"

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
	userResult, error := l.svcCtx.UserModel.Insert(l.ctx, &model.Users{
		Mobile:    in.Mobile,
		Email:     in.Email,
		Status:    model.Enable, //默认启用
		CreatedAt: sql.NullTime{},
		UpdatedAt: sql.NullTime{},
	})
	if error != nil {
		return nil, error
	}
	userId, error := userResult.LastInsertId()
	if error != nil {
		return nil, error
	}

	_, error = l.svcCtx.UserAuthModel.Insert(l.ctx, &model.UsersAuths{
		UserId: userId,
		// AuthType:       in.AuthType,
		AuthKey:        in.AuthKey,
		AuthCredential: in.AuthCredential,
		CreatedAt:      sql.NullTime{},
		UpdatedAt:      sql.NullTime{},
	})
	if error != nil {
		return nil, error
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, error := generateTokenLogic.GenerateToken(&account.GenerateTokenRequest{
		UserId: userId,
	})
	if error != nil {
		return nil, error
	}

	return &account.RegisterResponse{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
		RefreshAfter: token.RefreshAfter,
	}, nil
}
