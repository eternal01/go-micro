// Code generated by goctl. DO NOT EDIT.
// Source: account.proto

package server

import (
	"context"

	"go-micro/platform/account/rpc/account"
	"go-micro/platform/account/rpc/internal/logic"
	"go-micro/platform/account/rpc/internal/svc"
)

type AccountServer struct {
	svcCtx *svc.ServiceContext
	account.UnimplementedAccountServer
}

func NewAccountServer(svcCtx *svc.ServiceContext) *AccountServer {
	return &AccountServer{
		svcCtx: svcCtx,
	}
}

func (s *AccountServer) Register(ctx context.Context, in *account.RegisterRequest) (*account.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *AccountServer) Login(ctx context.Context, in *account.LoginRequest) (*account.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *AccountServer) GetUser(ctx context.Context, in *account.GetUserRequest) (*account.GetUserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *AccountServer) GetUserByEmail(ctx context.Context, in *account.GetUserByEmailRequest) (*account.GetUserResponse, error) {
	l := logic.NewGetUserByEmailLogic(ctx, s.svcCtx)
	return l.GetUserByEmail(in)
}

func (s *AccountServer) GetUserByMobile(ctx context.Context, in *account.GetUserByMobileRequest) (*account.GetUserResponse, error) {
	l := logic.NewGetUserByMobileLogic(ctx, s.svcCtx)
	return l.GetUserByMobile(in)
}

func (s *AccountServer) GetUserAuthByAuthKey(ctx context.Context, in *account.GetUserAuthByAuthKeyRequest) (*account.GetUserAuthResponse, error) {
	l := logic.NewGetUserAuthByAuthKeyLogic(ctx, s.svcCtx)
	return l.GetUserAuthByAuthKey(in)
}

func (s *AccountServer) GetUserAuthByUserId(ctx context.Context, in *account.GetUserAuthByUserIdRequest) (*account.GetUserAuthResponse, error) {
	l := logic.NewGetUserAuthByUserIdLogic(ctx, s.svcCtx)
	return l.GetUserAuthByUserId(in)
}

func (s *AccountServer) GenerateToken(ctx context.Context, in *account.GenerateTokenRequest) (*account.GenerateTokenResponse, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}