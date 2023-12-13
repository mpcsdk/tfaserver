package tfa

import (
	"context"

	"tfaserver/api/risk/nrpc"
	v1 "tfaserver/api/tfa/v1"
	"tfaserver/internal/service"
)

func (c *ControllerV1) TfaRequest(ctx context.Context, req *v1.TfaRequestReq) (res *v1.TfaRequestRes, err error) {
	//
	///
	tres, err := service.NrpcClient().TfaRequest(ctx, &nrpc.TfaRequestReq{
		Token:    req.Token,
		CodeType: req.CodeType,
	})
	if err != nil {
		return nil, err
	}
	///
	res = &v1.TfaRequestRes{
		RiskSerial: tres.RiskSerial,
		VList:      tres.VList,
	}
	return res, nil
	///
}
