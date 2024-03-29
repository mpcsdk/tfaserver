package tfa

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/mpcsdk/mpcCommon/mpccode"
)

func (c *ControllerV1) exit(ctx context.Context) {

}
func (c *ControllerV1) counter(ctx context.Context, tokenId string, method string) error {
	key := tokenId + method + "counter"
	if v, err := c.cache.Get(ctx, key); err != nil || !v.IsEmpty() {
		return gerror.NewCode(mpccode.CodeApiLimit)
	} else {
		c.cache.Set(ctx, key, 1, apiInterval)
		return nil
	}
}
func (c *ControllerV1) limitSendVerification(ctx context.Context, tokenId string, method string) error {
	key := tokenId + method + "limitSendVerification"
	if v, err := c.cache.Get(ctx, key); err != nil || !v.IsEmpty() {
		_, err = json.Marshal(func() {})
		err = gerror.Wrap(err,
			mpccode.ErrDetails(mpccode.ErrDetail("key", key),
				mpccode.ErrDetail("method", method)),
		)
		return err
	} else {
		c.cache.Set(ctx, key, 1, limitSendInterval)
		return nil
	}
}

func (c *ControllerV1) delTimeOut(s []*gtime.Time, limitDuration time.Duration) []*gtime.Time {
	i := 0
	n := gtime.Now()

	beforTime := n.Add(-limitDuration)
	for _, st := range s {
		if st.After(beforTime) {
			s[i] = st
			i++
		}
	}
	return s[:i]
}

func (c *ControllerV1) limitSendPhone(ctx context.Context, tokenId string, phone string) error {
	key := phone + "limitSendPhone"
	sendtimes := []*gtime.Time{}
	if v, err := c.cache.Get(ctx, key); err != nil {
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
			sendtimes = c.delTimeOut(sendtimes, limitSendPhoneDuration)
		}

		if len(sendtimes) >= limitSendPhoneDurationCnt {
			return mpccode.CodeLimitSendPhoneCode.Error()
		}
		sendtimes = append(sendtimes, gtime.Now())
		c.cache.Set(ctx, key, sendtimes, 0)
		return nil
	} else {
		sendtimes = append(sendtimes, gtime.Now())
		c.cache.Set(ctx, key, sendtimes, 0)
		return nil
	}
}
func (c *ControllerV1) limitSendMail(ctx context.Context, tokenId string, mail string) error {
	key := mail + "limitSendMail"
	sendtimes := []*gtime.Time{}
	if v, err := c.cache.Get(ctx, key); err != nil {
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
			sendtimes = c.delTimeOut(sendtimes, limitSendMailDuration)
		}

		if len(sendtimes) >= limitSendMailDurationCnt {
			return err
		}
		sendtimes = append(sendtimes, gtime.Now())
		c.cache.Set(ctx, key, sendtimes, 0)
		return nil
	} else {
		sendtimes = append(sendtimes, gtime.Now())
		c.cache.Set(ctx, key, sendtimes, 0)
		return nil
	}
}
