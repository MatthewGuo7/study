package UserModel

type UserModel struct {
	UserID   int    `json:"id"`
	UserName string `json:"user_name" form:"name" binding:"min=4"`
}

func New(attrs ...UserModelAttrFunc) *UserModel {
	u := &UserModel{}
	UserModelAttrs(attrs).Apply(u)
	return u
}

func (u *UserModel) Modify(attrs ...UserModelAttrFunc) *UserModel {
	UserModelAttrs(attrs).Apply(u)
	return u
}
