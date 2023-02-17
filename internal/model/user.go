package model

type User struct {
	Username string `gorm:"column:username;not null;unique"`
	Password string `gorm:"column:password"`
	Nickname string `gorm:"column:nickname"`
	Email    string `gorm:"column:email"`
	Mobile   string `gorm:"column:mobile;size:11"`
	Active   int    `gorm:"column:active;default:1"`
	BaseModel
}
