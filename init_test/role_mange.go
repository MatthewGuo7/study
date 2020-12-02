package initTest

import "testing"

type RoleMangeInter interface {
	DoAction() error
}

var roleMange = map[int]RoleMangeInter{}

func RegisterRoleMange(role int, roleType RoleMangeInter) {
	roleMange[role] = roleType
}

func NewRoleMange(role int) RoleMangeInter {
	_, ok := roleMange[role]
	if !ok {
		panic("unsupport role type")
	}
	return roleMange[role]
}

func TestSliceCopy(t *testing.T)  {

}

