package global

import "github.com/xm-chentl/eddydoor/internal/model/enum/datastatus"

type User struct {
	ID        string           `gorm:"column:id;primaryKey"`
	Phone     string           `gorm:"column:phone"`
	Password  string           `gorm:"column:password"`
	Nickname  string           `gorm:"column:nickname"`
	CreatedAt int64            `gorm:"column:created_at"`
	UpdatedAt int64            `gorm:"column:updated_at"`
	SMSCode   string           `gorm:"column:sms_code"`
	Status    datastatus.Value `gorm:"column:status"`
}

func (m User) GetID() interface{} {
	return nil
}

func (m *User) SetID(v interface{}) {
	m.ID = v.(string)
}

func (m User) Table() string {
	return "users"
}
