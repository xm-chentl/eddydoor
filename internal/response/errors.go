package response

import "github.com/xm-chentl/eddydoor/internal/common"

// 500-599 服务内部逻辑错误
// 600- 授权中心

var (
	// ErrDataException 数据异常
	ErrDataException = common.ResponseErr(600, "Data exception, please be serious")
	// ErrAccountExists 此帐号已存在, 请直接登录
	ErrAccountExists = common.ResponseErr(601, "This account already exists, please log in directly")
	// ErrAccountNotExist 帐号未注册，请先注册
	ErrAccountNotExists = common.ResponseErr(602, "This account is not registered, please register first")
	// ErrRegisterFailed 注册失败
	ErrRegisterFailed = common.ResponseErr(603, "register failed")
	// ErrSMSCodeValid 短信验证已失效
	ErrSMSCodeValid = common.ResponseErr(605, "The SMS verification code is invalid")
	// ErrSMSCodeIncorrect 短信验证码不正确
	ErrSMSCodeIncorrect = common.ResponseErr(606, "The SMS verification code is incorrect")
	// ErrTokenGenerateFailed 令牌创建失败
	ErrTokenGenerateFailed = common.ResponseErr(607, "Failed to generate the access token")
	// ErrTokenValid 令牌无效
	ErrTokenValid = common.ResponseErr(608, "Invalid token")
)
