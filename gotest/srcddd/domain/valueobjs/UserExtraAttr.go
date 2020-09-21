package valueobjs

type UserExtraAttrFunc func(extra *UserExtra)
type UserExtraAttrFuncs []UserExtraAttrFunc

func (u UserExtraAttrFuncs)apply(extra *UserExtra)  {
	for _, f := range u {
		f(extra)
	}
}
