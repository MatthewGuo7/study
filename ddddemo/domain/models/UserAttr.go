package models

type UserAttrFunc func(model *UserModel)

type UserAttrFuncs []UserAttrFunc

func (r UserAttrFuncs) apply(model *UserModel) {
	for _, f := range r {
		f(model)
	}
}

type UserExtraAttrFunc func(model *UserExtra)
type UserExtraAttrFuncs []UserExtraAttrFunc

func (r UserExtraAttrFuncs) apply(extra *UserExtra)  {
	for _, f := range r {
		f(extra)
	}
}

func WithUseID(userID int) UserAttrFunc {
	return func(model *UserModel) {
		model.UserID = userID
	}
}

func WithUserName(userName string) UserAttrFunc {
	return func(model *UserModel) {
		model.UserName = userName
	}
}

func WithUserPwd(pwd string) UserAttrFunc {
	return func(model *UserModel) {
		model.UserPwd = pwd
	}
}

func WithUserExtra(extra *UserExtra) UserAttrFunc {
	return func(model *UserModel) {
		model.UserExtra = extra
	}
}

func WithUserPhone(phone string) UserExtraAttrFunc {
	return func(model *UserExtra) {
		model.UserPhone = phone
	}
}

func WithUserCity(city string) UserExtraAttrFunc {
	return func(model *UserExtra) {
		model.UserCity = city
	}
}

func WithUserQQ(qq string) UserExtraAttrFunc {
	return func(model *UserExtra) {
		model.UserQQ = qq
	}
}
