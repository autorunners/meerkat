package fake

import (
	"encoding/json"
	"fmt"
	"github.com/antchfx/xmlquery"
	"io/ioutil"
	"net/http"

	"github.com/autorunners/meerkat/core/request"
)

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
