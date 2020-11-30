package UserModel

type UserModelAttrFunc func(model *UserModel)
type UserModelAttrs []UserModelAttrFunc

func (u UserModelAttrs) Apply(model *UserModel) {
	for _, f := range u {
		f(model)
	}
}

func WithUserID(id int) UserModelAttrFunc {
	return func(model *UserModel) {
		model.UserID = id
	}
}

func WithUserName(name string) UserModelAttrFunc {
	return func(model *UserModel) {
		model.UserName = name
	}
}
