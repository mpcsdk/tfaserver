package nats

import (
	"context"
	tfav1 "tfaserver/api/tfa/nrpc/v1"
	"tfaserver/internal/model"
	"tfaserver/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sNrpcServer) RpcTfaTx(ctx context.Context, req *tfav1.TfaTxReq) (*tfav1.TfaTxRes, error) {

	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
	if err != nil || tfaInfo == nil {
		g.Log().Warning(ctx, "TFATx:", "req:", req)
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeInternalError)
	}
	kinds, err := service.TFA().TFATx(ctx, tfaInfo, req.RiskSerial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodePerformRiskError)
	}
	//

	return &tfav1.TfaTxRes{
		Kinds: kinds,
	}, nil
}

func (*sNrpcServer) RpcTfaInfo(ctx context.Context, req *tfav1.TFAReq) (res *tfav1.TFARes, err error) {

	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcSendSmsCode")
	defer span.End()
	g.Log().Info(ctx, "RpcTfaInfo:", req)
	// info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	// if err != nil {
	// 	g.Log().Warning(ctx, "TFAInfo no userId:", "req:", req, "userInfo:", info)
	// 	return nil, gerror.NewCode(mpccode.CodeTokenInvalid)
	// }
	if req.UserId == "" {
		g.Log().Warning(ctx, "TFAInfo no userId:", "req:", req, "userInfo:", req)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, req.UserId)
	if err != nil {
		g.Log().Warning(ctx, "TFAInfo no info:", "req:", req, "tfaInfo:", tfaInfo)
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	if tfaInfo == nil {
		return nil, nil
	}
	g.Log().Info(ctx, "RpcTfaInfo:", res)
	res = &tfav1.TFARes{
		UserId: tfaInfo.UserId,
		Phone:  tfaInfo.Phone,
		UpPhoneTime: func() string {
			if tfaInfo.PhoneUpdatedAt == nil {
				return ""
			}
			return tfaInfo.PhoneUpdatedAt.Local().String()
		}(),
		Mail: tfaInfo.Mail,
		UpMailTime: func() string {
			if tfaInfo.MailUpdatedAt == nil {
				return ""
			}
			return tfaInfo.MailUpdatedAt.Local().String()
		}(),
	}
	return res, nil
}

func (*sNrpcServer) RpcSendSmsCode(ctx context.Context, req *tfav1.SmsCodeReq) (res *tfav1.SmsCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcSendSmsCode")
	defer span.End()
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	_, err = service.TFA().SendPhoneCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
	}
	return nil, err
}

func (*sNrpcServer) RpcSendMailCode(ctx context.Context, req *tfav1.MailCodekReq) (res *tfav1.MailCodekRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcSendMailCode")
	defer span.End()
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	// err = service.Risk().RiskMailCode(ctx, req.RiskSerial)
	_, err = service.TFA().SendMailCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
	}
	return nil, err
}

func (*sNrpcServer) RpcSendVerifyCode(ctx context.Context, req *tfav1.VerifyCodekReq) (res *tfav1.VerifyCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcSendVerifyCode")
	defer span.End()
	//
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	// err = service.Risk().VerifyCode(ctx, req.RiskSerial, req.Code)
	code := &model.VerifyCode{
		MailCode:  req.MailCode,
		PhoneCode: req.PhoneCode,
	}
	err = service.TFA().VerifyCode(ctx, info.UserId, req.RiskSerial, code)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
	}
	return nil, err
}
