package models

import "time"

type UserLogModel struct {
	ID         int       `gorm:"column:id;primary_key;auto_increment"`
	UserName   string    `gorm:"column:user_name"`
	LogType    int       `gorm:"column:log_type"`
	LogComment string    `gorm:"column:log_comment"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

func NewUserLogModel(userName string, logType int, logComment string) *UserLogModel {
	return &UserLogModel{UserName: userName, LogType: logType, LogComment: logComment}
}

