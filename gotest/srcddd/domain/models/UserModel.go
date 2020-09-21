package models

import (
	"crypto/md5"
	"fmt"
	. "srcddd/domain/valueobjs"
)

type UserModel struct {
	UserID   int    `gorm:"column:user_id;primary_key;auto_increment" json:"user_id"`
	UserPwd  string `gorm:"column:user_pwd" json:"user_pwd"`
	UserName string `gorm:"column:user_name" json:"user_name"`
	Extra    *UserExtra
}

func NewUserModel(attrs ... UserAttrFunc) *UserModel {
	u := &UserModel{}
	UserAttrFuncs(attrs).apply(u)
	return u
}

func (u *UserModel) BeforeSave() {
	u.UserPwd = fmt.Sprintf("%x", md5.Sum([]byte(u.UserPwd)))
}
