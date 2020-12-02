package models

type UserModel struct {
	UserID   int    `gorm:"column:user_id;primary_key;auto_increment"`
	UserName string `gorm:"column:user_name"`
	UserPwd  string `gorm:"column:user_pwd"`
	//UserAddTime int64  `gorm:"column:user_addtime"`
	UserExtra *UserExtra
}

func NewUserModel(funcs ...UserAttrFunc) *UserModel {
	u := &UserModel{}
	UserAttrFuncs(funcs).apply(u)
	return u
}

type UserExtra struct {
	UserPhone string `gorm:"column:user_phone"`
	UserCity  string `gorm:"column:user_city"`
	UserQQ    string `gorm:"column:user_qq"`
}

func NewUserExtra(funcs ...UserExtraAttrFunc) *UserExtra {
	u := &UserExtra{}
	UserExtraAttrFuncs(funcs).apply(u)
	return u
}

func (r *UserExtra) Equals(other *UserExtra) bool {
	return r.UserCity == other.UserCity && r.UserPhone == other.UserPhone && r.UserQQ == other.UserQQ
}
