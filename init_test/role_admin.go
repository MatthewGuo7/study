package initTest

import "fmt"

func init() {
	RegisterRoleMange(2, &RoleAdmin{})
}

type RoleAdmin struct {
}

func (r *RoleAdmin) DoAction() error {
	fmt.Println("role admin")
	return nil
}
