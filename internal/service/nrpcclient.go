// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"tfaserver/api/risk/nrpc"
)

type (
	INrpcClient interface {
		Flush()
		TfaRequest(ctx context.Context, req *nrpc.TfaRequestReq) (*nrpc.TfaRequestRes, error)
		RpcSendPhoneCode(ctx context.Context, req *nrpc.SendPhoneCodeReq) (*nrpc.SendPhoneCodeRes, error)
		RpcSendMailCode(ctx context.Context, req *nrpc.SendMailCodeReq) (*nrpc.SendMailCodeRes, error)
		RpcVerifyCode(ctx context.Context, req *nrpc.VerifyCodeReq) (*nrpc.VerifyCodeRes, error)
	}
)

var (
	localNrpcClient INrpcClient
)

func NrpcClient() INrpcClient {
	if localNrpcClient == nil {
		panic("implement not found for interface INrpcClient, forgot register?")
	}
	return localNrpcClient
}

func RegisterNrpcClient(i INrpcClient) {
	localNrpcClient = i
}
