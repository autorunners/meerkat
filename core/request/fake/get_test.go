package fake

import (
	"encoding/json"
	"fmt"
	"github.com/antchfx/xmlquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

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

func groupGet(gs ...[]testModel) (newGs []testModel) {
	for _, g := range gs {
		newGs = append(newGs, g...)
	}
	return newGs
}

func groupGetNormal() []testModel {
	return []testModel{
		{
			"http get 200 text",
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hi"))
			},
			check(tHttpTest(200), tContentTest("hi")),
		}, {
			"http get 500",
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(500)
			},
			check(tHttpTest(500), tContentTest("")),
		}, {
			"http get 200 json",
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`{"key":"value"}`))
			},
			check(tHttpTest(200), tContentJsonTest("value")),
		}, {
			"http get 200 xml",
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><rss key="value"></rss>`))
			},
			check(tHttpTest(200), tContentXmlTest("value")),
		},
	}
}

func tHttpTest(wantCode int) checkFunc {
	return func(req request.Request) error {
		res, err := req.Handle()
		if err != nil {
			return err
		}
		if res.StatusCode != wantCode {
			return fmt.Errorf("status = %d; want %d", res.StatusCode, wantCode)
		}
		return nil
	}
}
func tContentTest(want string) checkFunc {
	return func(req request.Request) error {
		res, err := req.Handle()
		if err != nil {
			return err
		}
		got, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return err
		}
		if string(got) != want {
			return fmt.Errorf("want: %s; got: %s", want, string(got))
		}
		return nil
	}
}
func tContentJsonTest(want string) checkFunc {
	return func(req request.Request) error {
		res, err := req.Handle()
		if err != nil {
			return err
		}
		got, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return err
		}
		var m map[string]string
		json.Unmarshal(got, &m)
		if m["key"] != want {
			return fmt.Errorf("want: %s; got: %s", want, m["key"])
		}
		return nil
	}
}
func tContentXmlTest(want string) checkFunc {
	return func(req request.Request) error {
		res, err := req.Handle()
		if err != nil {
			return err
		}
		defer res.Body.Close()
		var m map[string]string
		doc, err := xmlquery.Parse(res.Body)
		if err != nil {
			return err
		}
		value := doc.LastChild.Attr[0].Value
		if value != want {
			return fmt.Errorf("want: %s; got: %s", want, m["key"])
		}
		return nil
	}
}

func groupGetHeader() []testModel {
	return []testModel{
		{
			"http get 200 headers",
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("<html>"))
			},
			check(tHttpGetHeaderTest("Content-Type", "text/html; charset=utf-8")),
		}, {
			"http get 200 headers special ",
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

func TestGet(t *testing.T) {
	tts := groupGet(groupGetNormal(), groupGetHeader())
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(tt.h)
			server := httptest.NewServer(handler)

			req := request.Request{
				Method:  "GET",
				Timeout: 50,
			}
			req.FullUri = server.URL
			defer server.Close()

			for _, check := range tt.checks {
				if err := check(req); err != nil {
					t.Error(err)
				}
			}
		})
	}

}
