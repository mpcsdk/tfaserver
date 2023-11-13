package email

import (
	"context"
	"testing"
	"tfaserver/internal/service"
)

func Test_SendMail(t *testing.T) {
	service.RegisterMailCode(new())
	_, err := service.MailCode().SendVerificationCode(context.Background(), "xinwei.li@mixmarvel.com")
	if err != nil {
		t.Error(err)
	}
}
