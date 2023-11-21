package nats

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (s *sNrpcServer) delTimeOut(stlist []*gtime.Time, limitDuration time.Duration) []*gtime.Time {
	i := 0
	n := gtime.Now()
	beforTime := n.Add(-limitDuration)
	for _, st := range stlist {
		if st.After(beforTime) {
			stlist[i] = st
			i++
		}
	}
	return stlist[:i]
}

func (s *sNrpcServer) limitSendPhone(ctx context.Context, tokenId string, phone string) error {
	key := phone + "limitSendPhone"
	sendtimes := []*gtime.Time{}
	if v, err := s.cache.Get(ctx, key); err != nil {
		err = gerror.Wrap(err,
			mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
		)
		return err
	} else if !v.IsEmpty() {
		err := v.Structs(&sendtimes)
		if err != nil {
			return gerror.Wrap(err,
				mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
			)
		}
		////
		if len(sendtimes) >= limitSendPhoneDurationCnt {
			sendtimes = s.delTimeOut(sendtimes, limitSendPhoneDuration)
		}

		if len(sendtimes) >= limitSendPhoneDurationCnt {
			return mpccode.CodeLimitSendPhoneCode.Error()
		}
		sendtimes = append(sendtimes, gtime.Now())
		s.cache.Set(ctx, key, sendtimes, 0)
		return nil
	} else {
		sendtimes = append(sendtimes, gtime.Now())
		s.cache.Set(ctx, key, sendtimes, 0)
		return nil
	}
}
func (s *sNrpcServer) limitSendMail(ctx context.Context, tokenId string, mail string) error {
	key := mail + "limitSendMail"
	sendtimes := []*gtime.Time{}
	if v, err := s.cache.Get(ctx, key); err != nil {
		err = gerror.Wrap(err,
			mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
		)
		return err
	} else if !v.IsEmpty() {
		err := v.Structs(&sendtimes)
		if err != nil {
			return gerror.Wrap(err,
				mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
			)
		}
		////
		if len(sendtimes) >= limitSendMailDurationCnt {
			sendtimes = s.delTimeOut(sendtimes, limitSendMailDuration)
		}

		if len(sendtimes) >= limitSendMailDurationCnt {
			return mpccode.CodeLimitSendMailCode.Error()
		}
		sendtimes = append(sendtimes, gtime.Now())
		s.cache.Set(ctx, key, sendtimes, 0)
		return nil
	} else {
		sendtimes = append(sendtimes, gtime.Now())
		s.cache.Set(ctx, key, sendtimes, 0)
		return nil
	}
}
