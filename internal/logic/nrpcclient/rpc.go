package nrpcclient

import (
	"context"
	"tfaserver/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"google.golang.org/protobuf/types/known/emptypb"

	riskv1 "tfaserver/api/risk/nrpc/v1"
)

func (s *sNrpcClient) RpcRiskTFA(ctx context.Context, userId string, riskData *model.RiskTfa) (string, int32) {

	rst, err := s.cli.RpcRiskTFA(&riskv1.TFARiskReq{
		UserId: userId,
		Type:   riskData.Type,
		Mail:   riskData.Mail,
		Phone:  riskData.Phone,
	})

	///
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RpcRiskTFA TimeOut:")
			s.Flush()
			return "", mpccode.RiskCodeError
		}
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("useid", userId),
			mpccode.ErrDetail("riskData", riskData),
		))
		g.Log().Errorf(ctx, "%+v", err)
		return "", mpccode.RiskCodeError
	}
	///
	g.Log().Notice(ctx, "RpcRiskTFA:", "rst:", rst)
	return rst.RiskSerial, rst.Ok
}
func (s *sNrpcClient) RpcAlive(ctx context.Context) error {

	_, err := s.cli.RpcAlive(&emptypb.Empty{})
	if err != nil {
		if err.Error() == mpccode.ErrNrpcTimeOut.Error() {
			g.Log().Warning(ctx, "RpcAlive TimeOut:")
			s.Flush()
			return nil
		}
	}
	return err
}
