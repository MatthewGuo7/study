package gateway

import "strings"

type ValueConfig string

func (v ValueConfig) Get() []string  {
	ret := make([]string, 0)
	values := strings.Split(string(v), ",")
	for _, v := range values {
		ret = append(ret,strings.Trim(v, " "))
	}

	return ret
}
