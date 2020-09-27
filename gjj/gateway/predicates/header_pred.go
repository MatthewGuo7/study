package predicates

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type MethodPredicate string
func(this MethodPredicate) IsMatch(req *http.Request ) bool{
	param:=req.Method
	s:=string(this)
	slist:=strings.Split(s,",")
	if len(slist)==0{
		return true
	}
	for _,item :=range slist{
		if item==param{
			return true
		}
	}
	return false
}

type HeaderPredicate string

func (this HeaderPredicate) IsMatch(param http.Header) bool {
	s := string(this)
	if strings.Trim(s, " ") == "" {
		return true
	}

	slist := strings.Split(s, ",")
	if len(slist) < 2 || len(slist)%2 != 0 { // header 规则写错，就当没有
		return true
	}

	for i := 0; i < len(slist); i = i + 2 {
		key := slist[i]
		pattern := slist[i+1]
		if value, ok := param[key]; !ok { //没取到头
			return false
		} else {
			//正则判断
			reg, err := regexp.Compile(pattern)
			if err != nil {
				log.Println(err)
				return false
			}
			fmt.Printf("header = %+v ,reg = %+v\n", value, reg)
			if !reg.MatchString(value[0]) {
				return false
			}
		}
	}

	return true
}
