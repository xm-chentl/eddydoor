package user

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/xm-chentl/eddydoor/internal/model/enum/datastatus"
	"github.com/xm-chentl/eddydoor/internal/model/enum/formats"
	"github.com/xm-chentl/eddydoor/internal/model/global"
	"github.com/xm-chentl/eddydoor/internal/response"
	"github.com/xm-chentl/eddydoor/utils/guidex"
	"github.com/xm-chentl/eddydoor/utils/redisex"

	"github.com/xm-chentl/goresource"
)

type RegisterAPI struct {
	RedisImp redisex.IRedis       `inject:""`
	MySql    goresource.IResource `inject:""`
	GuidImp  guidex.IGenerate     `inject:""`

	Phone   string `json:"phone" validate:"required"`
	SMSCode string `json:"sms_code" validate:"required"`
}

func (s RegisterAPI) Call(ctx context.Context) (res interface{}, err error) {
	code, err := s.RedisImp.Get(ctx, fmt.Sprintf(formats.SMSRegister.String(), s.Phone))
	if err != nil {
		err = response.ErrSMSCodeValid
		return
	}
	if !strings.EqualFold(code, s.SMSCode) {
		err = response.ErrSMSCodeIncorrect
		return
	}

	db := s.MySql.Db(ctx)
	var entry global.User
	if err = db.Query().Where("phone = ?", s.Phone).First(&entry); err != nil {
		//todo: 打异常日志
		err = response.ErrDataException
		return
	}
	if entry.ID != "" {
		err = response.ErrAccountExists
		return
	}

	entry = global.User{
		ID:        s.GuidImp.String(),
		Phone:     s.Phone,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Status:    datastatus.Normal,
		Nickname:  s.Phone,
	}
	if err = db.Create(&entry); err != nil {
		err = response.ErrRegisterFailed
		return
	}

	return
}
