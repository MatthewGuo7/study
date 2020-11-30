package initTest

import "fmt"

func init() {
	RegisterRoleMange(1, &RoleLeader{})
}

type RoleLeader struct {
}

func (r *RoleLeader) DoAction() error {
	fmt.Println("role leader")
	return nil
}
