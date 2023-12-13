package tfa

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/mpcsdk/mpcCommon/mpccode"

	"tfaserver/api/risk/nrpc"
	v1 "tfaserver/api/tfa/v1"
	"tfaserver/internal/service"
)

func (c *ControllerV1) VerifyCode(ctx context.Context, req *v1.VerifyCodeReq) (res *v1.VerifyCodeRes, err error) {
	// ///
	userInfo, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}

	///
	_, err = service.NrpcClient().RpcVerifyCode(ctx, &nrpc.VerifyCodeReq{
		UserId:     userInfo.UserId,
		RiskSerial: req.RiskSerial,
		MailCode:   req.MailCode,
		PhoneCode:  req.PhoneCode,
	})
	if err != nil {
		return nil, err
	}
	return nil, err
}
