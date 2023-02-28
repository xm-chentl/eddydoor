package formats

type Value string

func (v Value) String() string {
	return string(v)
}

const (
	SMSLogin    Value = "login_sms_code_%s"
	SMSRegister Value = "register_sms_code_%s"
)
