package logic

import (
	"context"
	"time"

	"go-micro/platform/account/rpc/account"
	"go-micro/platform/account/rpc/internal/svc"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *account.GenerateTokenRequest) (*account.GenerateTokenResponse, error) {
	// todo: add your logic here and delete this line

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId)
	if err != nil {
		return nil, err
	}

	return &account.GenerateTokenResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["user_id"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
