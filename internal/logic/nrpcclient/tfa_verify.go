package nrpcclient

import (
	"context"
	"tfaserver/api/risk/nrpc"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sNrpcClient) RpcVerifyCode(ctx context.Context, req *nrpc.VerifyCodeReq) (*nrpc.VerifyCodeRes, error) {
	rst, err := s.cli.RpcVerifyCode(req)
	///
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RpcRiskTFA TimeOut:")
			s.Flush()
			return nil, nil
			//, mpccode.RiskCodeError
		}
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("useid", req.UserId),
			mpccode.ErrDetail("riskData", req),
		))
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	///
	g.Log().Notice(ctx, "RpcRiskTFA:", "rst:", rst)
	return &nrpc.VerifyCodeRes{
		Ok: 0,
	}, nil
	// return rst.RiskSerial, .niOk
}
