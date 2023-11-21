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
	res = &tfav1.TFARes{
		UserId: tfaInfo.UserId,
		Phone:  tfaInfo.Phone,
		UpPhoneTime: func() string {
			if tfaInfo.PhoneUpdatedAt == nil {
				return ""
			}

			return tfaInfo.PhoneUpdatedAt.String()
		}(),
		Mail: tfaInfo.Mail,
		UpMailTime: func() string {
			if tfaInfo.MailUpdatedAt == nil {
				return ""
			}
			return tfaInfo.MailUpdatedAt.String()
		}(),
	}
	g.Log().Info(ctx, "RpcTfaInfo:", res)
	return res, nil
}

func (s *sNrpcServer) RpcSendSmsCode(ctx context.Context, req *tfav1.SmsCodeReq) (res *tfav1.SmsCodeRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcSendSmsCode")
	defer span.End()
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	///
	risk := service.TFA().GetRiskVerify(ctx, info.UserId, req.RiskSerial)
	if risk == nil {
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	v := risk.Verifier(model.VerifierKind_Phone)
	if v == nil {
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	if err := s.limitSendPhone(ctx, req.Token, v.Destination()); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeLimitSendPhoneCode)
	}
	///
	_, err = service.TFA().SendPhoneCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
	}
	return nil, err
}

func (s *sNrpcServer) RpcSendMailCode(ctx context.Context, req *tfav1.MailCodekReq) (res *tfav1.MailCodekRes, err error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcSendMailCode")
	defer span.End()
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	/// limit send cnt
	risk := service.TFA().GetRiskVerify(ctx, info.UserId, req.RiskSerial)
	if risk == nil {
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	v := risk.Verifier(model.VerifierKind_Mail)
	if v == nil {
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	if err := s.limitSendMail(ctx, req.Token, v.Destination()); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeLimitSendMailCode)
	}
	///
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
		g.Log().Warningf(ctx, "%+v", err)
		return nil, mpccode.CodeTokenInvalid.Error()
	}
	code := &model.VerifyCode{
		MailCode:  req.MailCode,
		PhoneCode: req.PhoneCode,
	}
	err = service.TFA().VerifyCode(ctx, info.UserId, req.RiskSerial, code)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	return nil, err
}
