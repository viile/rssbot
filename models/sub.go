package models

type Sub struct {
	ID int32 `gorm:"primary_key"`
	UserID string `gorm:"column:user_id"`
	ContentID int32 `gorm:"column:content_id"`
}
