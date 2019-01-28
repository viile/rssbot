package models

type User struct {
	ID int32 `gorm:"primary_key"`
	OpenID string `gorm:"column:open_id"`
}
