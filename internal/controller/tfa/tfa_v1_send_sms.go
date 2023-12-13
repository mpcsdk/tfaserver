package tfa

import (
	"context"
	"tfaserver/api/risk/nrpc"
	v1 "tfaserver/api/tfa/v1"
	"tfaserver/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error) {

	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil || info == nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTokenInvalid)
	}
	////
	_, err = service.NrpcClient().RpcSendPhoneCode(ctx, &nrpc.SendPhoneCodeReq{
		Phone:      req.Phone,
		RiskSerial: req.RiskSerial,
		UserId:     info.UserId,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
