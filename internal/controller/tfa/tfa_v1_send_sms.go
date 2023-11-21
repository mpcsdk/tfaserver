package tfa

import (
	"context"
	v1 "tfaserver/api/tfa/v1"
	"tfaserver/internal/model"
	"tfaserver/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (res *v1.SendSmsCodeRes, err error) {

	//limit
	if err := c.counter(ctx, req.Token, "SendSmsCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	if err := c.limitSendVerification(ctx, req.Token, "SendSmsCode"); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeLimitSendMailCode)
	}

	//trace
	ctx, span := gtrace.NewSpan(ctx, "SendSmsCode")
	defer span.End()
	///
	//
	info, err := service.UserInfo().GetUserInfo(ctx, req.Token)
	if err != nil || info == nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTokenInvalid)
	}
	tfaInfo, err := service.DB().FetchTfaInfo(ctx, info.UserId)
	if err != nil || tfaInfo == nil {
		g.Log().Warning(ctx, "SendSmsCode:", req, err)
		return nil, gerror.NewCode(mpccode.CodeTFANotExist)
	}
	////
	riskKind, err := service.TFA().TfaRiskKind(ctx, tfaInfo, req.RiskSerial)
	if err != nil {
		g.Log().Warningf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
	}
	///
	////
	switch riskKind {
	case model.RiskKind_BindPhone, model.RiskKind_UpPhone:
		if req.Phone == "" {
			return nil, gerror.NewCode(mpccode.CodeParamInvalid)
		}
		err = service.DB().TfaPhoneNotExists(ctx, req.Phone)
		if err != nil {
			g.Log().Warningf(ctx, "%+v", err)
			return nil, gerror.NewCode(mpccode.CodeTFAPhoneExists)
		}
		////
		service.TFA().TfaSetPhone(ctx, tfaInfo, req.Phone, req.RiskSerial, riskKind)
		///
	case model.RiskKind_BindMail, model.RiskKind_UpMail:
	default:
		return nil, gerror.NewCode(mpccode.CodeRiskSerialNotExist)
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
	if err := c.limitSendPhone(ctx, req.Token, v.Destination()); err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeLimitSendPhoneCode)
	}
	///
	_, err = service.TFA().SendPhoneCode(ctx, info.UserId, req.RiskSerial)
	if err != nil {
		g.Log().Errorf(ctx, "%+v", err)
		return nil, gerror.NewCode(mpccode.CodeTFASendSmsFailed)
	}
	return nil, nil
}
