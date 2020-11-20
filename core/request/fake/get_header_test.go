package fake

import (
	"fmt"
	"net/http"

	"github.com/autorunners/meerkat/core/request"
)

func groupGetHeader() []testModel {
	return []testModel{
		{
			"http get 200 headers",
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("<html>"))
			},
			check(tHttpGetHeaderTest("Content-Type", "text/html; charset=utf-8")),
		}, {
			"http get 200 headers special",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "some/type")
				w.Write([]byte("<html>"))
			},
			check(tHttpGetHeaderTest("Content-Type", "some/type"), tHttpGetHeaderTest("Transfer-Encoding", "")),
		}, {
			"http get 200 headers special2 ",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Transfer-Encoding", "chunked")
				w.Write([]byte("<html>"))
			},
			check(tHttpGetHeaderTest("Content-Type", ""), tHttpGetHeaderTest("Transfer-Encoding", "")),
		}, {
			"http get 200 headers special2 ",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Key", "value")
				w.Write([]byte("<html>"))
			},
			check(tHttpGetHeaderTest("Content-Type", "text/html; charset=utf-8"), tHttpGetHeaderTest("Key", "value")),
		},
	}
}
func tHttpGetHeaderTest(key string, want string) checkFunc {
	return func(req request.Request) error {
		res, err := req.Handle()
		if err != nil {
			return err
		}
		contentType := res.Header.Get(key)
		if contentType != want {
			return fmt.Errorf("got = %s; want %s", contentType, want)
		}
		return nil
	}
}
