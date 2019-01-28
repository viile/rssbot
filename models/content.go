package models

type Content struct {
	ID int32 `gorm:"primary_key"`
	Link string `gorm:"column:link"`
	Type int32 `gorm:"column:type"`  // 1 rss 2 weibo 3 youku
}
