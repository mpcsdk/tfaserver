package tfa

import (
	"context"
	"errors"
	"sync"
	"tfaserver/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
)

const (
	VerifierKind_Nil   = "nil"
	VerifierKind_Phone = "Phone"
	VerifierKind_Mail  = "Mail"
)

type riskPenddingContainer struct {
	lock                     sync.RWMutex
	riskPendding             map[UserRiskId]*riskVerifyPendding
	ctx                      context.Context
	verificationCodeDuration int
}

func newRiskPenddingContainer(times int) *riskPenddingContainer {
	s := &riskPenddingContainer{
		riskPendding:             make(map[UserRiskId]*riskVerifyPendding),
		ctx:                      context.Background(),
		verificationCodeDuration: times,
	}
	//
	gtimer.Add(s.ctx, time.Second*time.Duration(times), func(ctx context.Context) {
		s.lock.Lock()
		defer s.lock.Unlock()
		n := gtime.Now()
		for key, risk := range s.riskPendding {
			if risk.dealline.Before(n) {
				g.Log().Info(s.ctx, "RiskPenddingContainer dealline:", key, risk.dealline.Local().String())
				delete(s.riskPendding, key)
			}
		}
	})
	//
	return s
}

// /
// //
func (s *riskPenddingContainer) GetRiskVerify(userId, riskSerial string) *riskVerifyPendding {
	key := keyUserRiskId(userId, riskSerial)
	s.lock.RLock()
	defer s.lock.RUnlock()
	if risk, ok := s.riskPendding[key]; ok {
		return risk
	}
	return nil
}

func (s *riskPenddingContainer) NewRiskPendding(
	userId, riskSerial string,
	riskKind model.RiskKind,
) *riskVerifyPendding {
	risk := s.GetRiskVerify(userId, riskSerial)
	if risk == nil {
		risk = &riskVerifyPendding{
			RiskKind:   riskKind,
			UserId:     userId,
			RiskSerial: riskSerial,
			verifier:   map[VerifyKind]IVerifier{},
			dealline:   gtime.Now().Add(time.Duration(s.verificationCodeDuration) * time.Second),
			// dealline: gtime.Now(),
		}
		key := keyUserRiskId(userId, riskSerial)
		s.lock.Lock()
		s.riskPendding[key] = risk
		s.lock.Unlock()
	}
	return risk
}
func (s *riskPenddingContainer) Del(userId, riskSerial string) {
	key := keyUserRiskId(userId, riskSerial)
	delete(s.riskPendding, key)
}

var errRiskNotExist error = errors.New("risk not exist")
var errRiskNotDone error = errors.New("risk not done")
