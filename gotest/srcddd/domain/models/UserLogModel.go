package models

import "time"

type UserLogModel struct {
	ID         int       `gorm:"column:id;primary_key;auto_increment" json:"id"`
	UserName   string    `gorm:"column:user_name" json:"user_name"`
	LogType    uint8     `gorm:"column:log_type" json:"log_type"`
	LogComment string    `gorm:"column:log_comment" json:"log_comment"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}
