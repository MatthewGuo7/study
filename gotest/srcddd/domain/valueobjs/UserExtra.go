package valueobjs

type UserExtra struct {
	UserPhone string `gorm:"column:user_phone" json:"user_phone"`
	UserQQ    string `gorm:"column:user_qq" json:"user_qq"`
	UserCity  string `gorm:"column:user_city" json:"user_city"`
}

func NewUserExtra(attr ... UserExtraAttrFunc) *UserExtra {
	extra := &UserExtra{}
	UserExtraAttrFuncs(attr).apply(extra)
	return extra
}

func (u *UserExtra) Equals(other *UserExtra) bool {
	return u.UserCity == other.UserCity && u.UserQQ == other.UserQQ &&
		u.UserPhone == other.UserPhone
}
