package tfa

import (
	"testing"
	"tfaserver/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
)

func Test_Verifier(t *testing.T) {
	verifier := newEmptyVerifier(model.RiskKind_Tx)
	_, err := verifier.SendVerificationCode()
	if !gerror.Equal(err, errRiskKindTx) {
		t.Error(err)
	}
}
