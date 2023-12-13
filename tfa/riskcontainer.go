package tfa

import (
	"context"
	"tfaserver/internal/model"
)

func (s *sTFA) GetRiskVerify(ctx context.Context, userId, riskSerial string) (risk *model.RiskVerifyPendding) {
	risk = s.riskPenddingContainer.GetRiskVerify(userId, riskSerial)
	return risk
}
