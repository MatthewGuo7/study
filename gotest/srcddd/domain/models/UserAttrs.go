package models

type UserAttrFunc func(model *UserModel)
type UserAttrFuncs []UserAttrFunc

func (funcs UserAttrFuncs)apply(user *UserModel)  {
	for _, f := range funcs {
		f(user)
	}
}

func WithUserID(id int) UserAttrFunc {
	return func(u *UserModel) {
		u.UserID = id
	}
}

func WithUserName(name string) UserAttrFunc {
	return func(u *UserModel) {
		u.UserName = name
	}
}

func WithUserPwd(pwd string) UserAttrFunc {
	return func(u *UserModel) {
		u.UserPwd = pwd
	}
}
