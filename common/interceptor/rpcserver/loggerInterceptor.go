package rpcserver

import (
	"context"
	"go-micro/common/errorx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logx.WithContext(ctx).Slowf("RPC Request - %+v", req)
	resp, err = handler(ctx, req)
	if err != nil {
		if err == sqlx.ErrNotFound {
			err = errorx.NotFoundError
		}
		causeErr := errors.Cause(err)                  // err类型
		if e, ok := causeErr.(*errorx.CodeError); ok { //自定义错误类型
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)

			//转成grpc err
			err = status.Error(codes.Code(e.Code), e.Msg)
		} else {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
		}
	}
	logx.WithContext(ctx).Slowf("RPC Response - %+v", resp)
	logx.CollectSysLog()
	return resp, err
}
