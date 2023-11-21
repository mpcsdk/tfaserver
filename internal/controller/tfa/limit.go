package tfa

import (
	"context"
	"encoding/json"

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

func (c *ControllerV1) delTimeOut(s []*gtime.Time) []*gtime.Time {
	i := 0
	beforTime := gtime.Now().Add(-limitSendPhoneDuration)
	for _, st := range s {
		if st.Before(beforTime) {
			s[i] = st
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
		err := v.Structs(sendtimes)
		if err != nil {
			return gerror.Wrap(err,
				mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
			)
		}
		////
		if len(sendtimes) >= limitSendPhoneDurationCnt {
			sendtimes = c.delTimeOut(sendtimes)
		}

		if len(sendtimes) >= limitSendPhoneDurationCnt {
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
func (c *ControllerV1) limitSendMail(ctx context.Context, tokenId string, mail string) error {
	key := mail + "limitSendMail"
	sendtimes := []*gtime.Time{}
	if v, err := c.cache.Get(ctx, key); err != nil {
		err = gerror.Wrap(err,
			mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
		)
		return err
	} else if !v.IsEmpty() {
		err := v.Structs(sendtimes)
		if err != nil {
			return gerror.Wrap(err,
				mpccode.ErrDetails(mpccode.ErrDetail("key", key)),
			)
		}
		////
		if len(sendtimes) >= limitSendPhoneDurationCnt {
			sendtimes = c.delTimeOut(sendtimes)
		}

		if len(sendtimes) >= limitSendPhoneDurationCnt {
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
