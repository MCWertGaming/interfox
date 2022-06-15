package database

type User struct {
	UserID     string `gorm:"primaryKey"`
	NameFormat string
	NameStatic string
	NamePretty string
	Email      string
	Password   string
	RegisterIP string
	CreatedAt  string
	Locked     bool
	VerifyCode []VerifyCode `gorm:"foreignKey:UserID"`
}

type VerifyCode struct {
	ID     uint `gorm:"primaryKey"`
	UserID string
	Code   string
}

type Profile struct {
}
