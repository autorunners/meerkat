package fake

import (
	"log"
	"net/http"

	"github.com/autorunners/meerkat/core/request"
)

type (
	checkFunc func(req request.Request) error
	testModel struct {
		name   string
		h      func(w http.ResponseWriter, r *http.Request)
		checks []checkFunc
	}
)

var check func(fns ...checkFunc) []checkFunc

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	check = func(fns ...checkFunc) []checkFunc { return fns }
}
