package admin

import (
	"context"
	"log"
	"time"

	"github.com/xm-chentl/eddydoor/internal/config"
	"github.com/xm-chentl/eddydoor/internal/model/global"
	"github.com/xm-chentl/eddydoor/internal/response"
	"github.com/xm-chentl/eddydoor/internal/service/adminsvc"
	"github.com/xm-chentl/goresource"

	"github.com/golang-jwt/jwt/v5"
)

type LoginAPI struct {
	MySql goresource.IResource `inject:"mysql"`

	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s LoginAPI) Call(ctx context.Context) (res interface{}, err error) {
	db := s.MySql.Db(ctx)
	var entry global.Administrator
	if err = db.Query().Where("account = ? and password = ?", s.UserName, s.Password).First(&entry); err != nil {
		log.Println(err.Error())
		err = response.ErrDataException
		return
	}
	if entry.ID == "" {
		err = response.ErrAccountOrPasswordFailed
		return
	}

	loginTime := time.Now()
	expiredAt := loginTime.Add(config.Cfg.Secret.Expired * time.Minute)
	// 生成token
	admin := adminsvc.LoginAdmin{
		ID:        entry.ID,
		Account:   entry.Account,
		Nickname:  entry.Nickname,
		LoginTime: loginTime.Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expiredAt,
			},
		},
	}
	tokenImp := jwt.NewWithClaims(jwt.SigningMethodHS256, admin)
	token, err := tokenImp.SignedString([]byte(config.Cfg.Secret.Admin))
	if err != nil {
		log.Fatal(err)
		err = response.ErrTokenGenerateFailed
		return
	}
	res = map[string]string{
		"accessToken": token,
	}

	// todo: 需要校验token的过期，可以考虑把token放至redis加入过期时间

	return
}
