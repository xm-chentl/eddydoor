package user

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/xm-chentl/eddydoor/internal/config"
	"github.com/xm-chentl/eddydoor/internal/model/enum/formats"
	"github.com/xm-chentl/eddydoor/internal/model/global"
	"github.com/xm-chentl/eddydoor/internal/response"
	"github.com/xm-chentl/eddydoor/internal/service/usersvc"
	"github.com/xm-chentl/eddydoor/utils/redisex"

	"github.com/golang-jwt/jwt/v5"
	"github.com/xm-chentl/goresource"
)

type LoginByPhoneAPI struct {
	MySqlDb  goresource.IResource `inject:"mysql"`
	RedisImp redisex.IRedis       `inject:""`

	Phone   string `json:"phone" validate:"required"`
	SMSCode string `json:"sms_code" validate:"required"`
}

func (s LoginByPhoneAPI) Call(ctx context.Context) (res interface{}, err error) {
	code, err := s.RedisImp.Get(ctx, fmt.Sprintf(formats.SMSLogin.String(), s.Phone))
	if err != nil {
		err = response.ErrSMSCodeValid
		return
	}

	var entry global.User
	db := s.MySqlDb.Db(ctx)
	if err = db.Query().Where("phone = ?", s.Phone).First(&entry); err != nil {
		err = response.ErrDataException
		return
	}
	if entry.ID == "" {
		err = response.ErrAccountNotExists
		return
	}
	if !strings.EqualFold(code, s.SMSCode) {
		err = response.ErrSMSCodeIncorrect
		return
	}

	loginTime := time.Now()
	expiredAt := loginTime.Add(config.Cfg.Secret.Expired * time.Minute)
	// 生成token
	user := usersvc.LoginUser{
		ID:        entry.ID,
		Phone:     entry.Phone,
		Nickname:  entry.Nickname,
		LoginTime: loginTime.Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expiredAt,
			},
		},
	}
	tokenImp := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	token, err := tokenImp.SignedString([]byte(config.Cfg.Secret.User))
	if err != nil {
		log.Fatal(err)
		err = response.ErrTokenGenerateFailed
		return
	}
	res = token

	return
}
