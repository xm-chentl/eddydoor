package user

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/xm-chentl/eddydoor/internal/model/enum/formats"
	"github.com/xm-chentl/eddydoor/internal/model/global"
	"github.com/xm-chentl/eddydoor/internal/response"
	"github.com/xm-chentl/eddydoor/utils/redisex"
	"github.com/xm-chentl/goresource"
)

type GetLoginSMSCodeAPI struct {
	RedisImp redisex.IRedis       `inject:""`
	MySqlDb  goresource.IResource `inject:"mysql"`

	Phone string `json:"phone" validate:"required"`
}

func (s GetLoginSMSCodeAPI) Call(ctx context.Context) (res interface{}, err error) {
	// todo: 放至redis key: login_sms_code_{phone} value: sms_code
	code := GenValidateCode(6)
	err = s.RedisImp.Set(ctx, fmt.Sprintf(formats.SMSLogin.String(), s.Phone), code, time.Minute)
	if err != nil {
		err = response.ErrSMSCodeValid
		return
	}

	var entry global.User
	if err = s.MySqlDb.Db(ctx).Query().Where("phone = ?", s.Phone).First(&entry); err != nil {
		err = response.ErrDataException
		return
	}
	if entry.ID == "" {
		err = response.ErrAccountNotExists
		return
	}
	// 发送验证码到手机
	res = code

	return
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
