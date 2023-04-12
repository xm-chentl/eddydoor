package global

type Administrator struct {
	ID           string `gorm:"column:id"`
	Account      string `gorm:"column:account"`
	Password     string `gorm:"column:account"`
	Nickname     string `gorm:"column:nickname"`
	Avatar       string `gorm:"column:avatar"`
	Introduction string `gorm:"introduction"`
	EMail        string `gorm:"email"`
	CreatedAt    int64  `gorm:"created_at"`
	LastLoginAt  int64  `gorm:"last_login_at"`
}

func (m Administrator) GetID() interface{} {
	return m.ID
}

func (m *Administrator) SetID(v interface{}) {
	m.ID = v.(string)
}

func (m Administrator) Table() string {
	return "administrators"
}
