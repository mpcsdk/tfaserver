// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	tfav1 "tfaserver/api/tfa/nrpc/v1"

	"github.com/golang/protobuf/ptypes/empty"
)

type (
	INrpcServer interface {
		RpcAlive(ctx context.Context, in *empty.Empty) (*empty.Empty, error)
		RpcTfaTx(ctx context.Context, req *tfav1.TfaTxReq) (*tfav1.TfaTxRes, error)
		RpcTfaInfo(ctx context.Context, req *tfav1.TFAReq) (res *tfav1.TFARes, err error)
		RpcSendSmsCode(ctx context.Context, req *tfav1.SmsCodeReq) (res *tfav1.SmsCodeRes, err error)
		RpcSendMailCode(ctx context.Context, req *tfav1.MailCodekReq) (res *tfav1.MailCodekRes, err error)
		RpcSendVerifyCode(ctx context.Context, req *tfav1.VerifyCodekReq) (res *tfav1.VerifyCodeRes, err error)
	}
)

var (
	localNrpcServer INrpcServer
)

func NrpcServer() INrpcServer {
	if localNrpcServer == nil {
		panic("implement not found for interface INrpcServer, forgot register?")
	}
	return localNrpcServer
}

func RegisterNrpcServer(i INrpcServer) {
	localNrpcServer = i
}
