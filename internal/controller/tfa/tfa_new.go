// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package tfa

import (
	"tfaserver/internal/config"
	"tfaserver/api/tfa"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

type ControllerV1 struct{
	cache *gcache.Cache
}


var apiInterval = time.Second * 1
var limitSendInterval = time.Second * 60

var limitSendPhoneDurationCnt = 50
var limitSendPhoneDuration  = time.Hour
var limitSendMailDurationCnt = 10
var limitSendMailDuration  = time.Hour

func init() {
	apiInterval = time.Duration(config.Config.Cache.ApiInterval) *time.Second
	limitSendInterval  = time.Duration(config.Config.Cache.LimitSendInterval) *time.Second	

	limitSendPhoneDurationCnt = config.Config.Cache.LimitSendPhoneCount
	limitSendPhoneDuration  = time.Duration(config.Config.Cache.LimitSendPhoneDuration) *time.Second
	limitSendMailDurationCnt =  config.Config.Cache.LimitSendMailCount
	limitSendMailDuration  = time.Duration(config.Config.Cache.LimitSendMailDuration)*time.Second
}
func NewV1() tfa.ITfaV1 {
	redisCache := gcache.NewAdapterRedis(g.Redis())
	c := gcache.New()
	c.SetAdapter(redisCache)

	return &ControllerV1{
		cache : c,
	}
}

