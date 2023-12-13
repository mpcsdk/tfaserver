package nrpcclient

import (
	"context"
	"tfaserver/api/risk/nrpc"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sNrpcClient) TfaRequest(ctx context.Context, req *nrpc.TfaRequestReq) (*nrpc.TfaRequestRes, error) {
	rst, err := s.cli.RpcTfaRequest(req)
	///
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RpcTfaRequest TimeOut:")
			s.Flush()
			return nil, nil
			//, mpccode.RiskCodeError
		}
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("useid", req.Token),
			mpccode.ErrDetail("riskData", req),
		))
		g.Log().Errorf(ctx, "%+v", err)
		return nil, err
	}
	///
	g.Log().Notice(ctx, "RpcTfaRequest:", "rst:", rst)
	return rst, nil
}
