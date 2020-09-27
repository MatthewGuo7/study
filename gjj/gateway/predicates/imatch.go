package predicates

import "net/http"

type PredicateMatcher interface {
	IsMatch(req *http.Request ) bool
}
