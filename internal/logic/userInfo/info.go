package userInfo

import (
	"context"
	"tfaserver/internal/config"
	"tfaserver/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/mpcsdk/mpcCommon/mpccode"
	"github.com/mpcsdk/mpcCommon/userInfoGeter"
)

type sUserInfo struct {
	url string
	///
	userGeter *userInfoGeter.UserTokenInfoGeter
	c         *gcache.Cache
}

func (s *sUserInfo) GetUserInfo(ctx context.Context, userToken string) (userInfo *userInfoGeter.UserInfo, err error) {
	if userToken == "" {
		return nil, mpccode.ErrArg
	}
	///
	// 用户信息示例
	// "id": 10,
	// "appPubKey": "038c90b87d77f2cc3d26132e1ea26e14646d663e3f43f17180345df3d54b8b5c70",
	// "email": "sunwenhao0421@163.com",
	// "loginType": "tkey-auth0-twitter-cyan",
	// "address": "0xe73E35d8Ecc3972481138D01799ED3934cc57853",
	// "keyHash": "U2FsdGVkX1/O6j9czaWzdjjDo/XPjk1hI8pIoaxSuS52zIxVuStK/nS07ucgiM5si8NjN97rAux3aH7Ld2i5oO8UuL6tpNZmLMG9ZpwVTxvGkCa3H14vTxWNz+yBoWG8",
	// "create_time": 1691118876
	if v, ok := s.c.Get(ctx, userToken); ok == nil && !v.IsEmpty() {
		info := &userInfoGeter.UserInfo{}
		err = v.Struct(info)
		if err != nil {
			return nil, err
		}
		return info, nil
	}
	///
	info, err := s.userGeter.GetUserInfo(ctx, userToken)
	if err != nil {
		err = gerror.Wrap(err, mpccode.ErrDetails(
			mpccode.ErrDetail("userToken", userToken),
		))
		return info, err
	}
	s.c.Set(ctx, userToken, info, 0)
	return info, err
}

func new() *sUserInfo {
	///
	// url, err := gcfg.Instance().Get(context.Background(), "userTokenUrl")
	// if err != nil {
	// 	panic(err)
	// }
	url := config.Config.UserTokenUrl
	///
	userGeter := userInfoGeter.NewUserInfoGeter(url)
	_, err := userGeter.GetUserInfo(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBQdWJLZXkiOiIwMjI1YmI1MmU5NTcyMDUwZmZjMGM4MGRjZDBhYTBmNjQyNDFjMDk5ZDAzZjFlYTFjODEzMmZkMzViY2Q3MDBiMWMiLCJpYXQiOjE2OTQ0Mjk5OTEsImV4cCI6MTcyNTk2NTk5MX0.8YaF5spnD1SjI-NNbBCIBj9H5pspXMMkPJrKk23LdnM")
	if err != nil {
		panic(err)
	}
	//
	s := &sUserInfo{
		userGeter: userGeter,
		c:         gcache.New(),
	}
	///

	return s
}

func init() {
	service.RegisterUserInfo(new())
}
