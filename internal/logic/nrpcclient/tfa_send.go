package nrpcclient

import (
	"context"
	"tfaserver/api/risk/nrpc"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sNrpcClient) RpcSendPhoneCode(ctx context.Context, req *nrpc.SendPhoneCodeReq) (*nrpc.SendPhoneCodeRes, error) {
	rst, err := s.cli.RpcSendPhoneCode(req)
	///
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RpcSendPhoneCode TimeOut:")
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
	g.Log().Notice(ctx, "RpcSendPhoneCode:", "rst:", rst)
	return rst, nil
}
func (s *sNrpcClient) RpcSendMailCode(ctx context.Context, req *nrpc.SendMailCodeReq) (*nrpc.SendMailCodeRes, error) {
	rst, err := s.cli.RpcSendMailCode(req)
	///
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RpcSendMailCode TimeOut:")
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
	g.Log().Notice(ctx, "RpcSendMailCode:", "rst:", rst)
	return rst, nil
}
