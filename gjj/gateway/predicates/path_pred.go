package predicates

import (
	"fmt"
	"path/filepath"
	"strings"
)

type PathPredicate string

func (p PathPredicate) IsMatch(path string) bool {
	if strings.Trim(string(p), " ") == "" {
		return true
	}
	matched, err := filepath.Match(string(p), path)
	if nil != err || !matched {
		return false
	}

	fmt.Printf("matched url path = %+v", path)

	return true
}
