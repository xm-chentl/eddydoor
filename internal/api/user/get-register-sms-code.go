package user

import (
	"context"
	"fmt"
	"time"

	"github.com/xm-chentl/eddydoor/internal/model/enum/formats"
	"github.com/xm-chentl/eddydoor/internal/response"
	"github.com/xm-chentl/eddydoor/utils/redisex"
)

type GetRegisterSMSCodeAPI struct {
	// todo: 未校验参数
	Phone string `json:"phone" validate:"required"`

	RedisImp redisex.IRedis `inject:""`
}

func (s GetRegisterSMSCodeAPI) Call(ctx context.Context) (res interface{}, err error) {
	// todo: 设计防刷的策略
	// 保存至redis用于验证，有效时间为60秒 key: login_sms_code_{phone} value: sms_code
	code := GenValidateCode(6)
	if err = s.RedisImp.Set(ctx, fmt.Sprintf(formats.SMSRegister.String(), s.Phone), code, time.Minute); err != nil {
		err = response.ErrSMSCodeValid
		return
	}
	// todo: 发送验证码至手机
	// todo: 为了方便暂时下放至响应数据中
	res = map[string]interface{}{
		"code":  code,
		"phone": s.Phone,
	}
	return
}
