package fake

import (
	"encoding/json"
	"fmt"
	"github.com/antchfx/xmlquery"
	"io/ioutil"
	"net/http"

	"github.com/autorunners/meerkat/core/request"
)

func groupPostNormal() []testModel {
	return []testModel{
		{
			"http post 200 text",
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("hi"))
			},
			check(tHttpTest(200), tPostTest(`hi`, "hi")),
		},
		{
			"http post 200 text json",
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`{"key":"Gordon"}`))
			},
			check(tHttpTest(200), tPostJsonTest(`{"name":"Gordon"}`, "Gordon")),
		},
	}
}

func tPostTest(body, want string) checkFunc {
	return func(req request.Request) error {
		req.Body = body
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

func tPostJsonTest(body, want string) checkFunc {
	return func(req request.Request) error {
		req.Body = body
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
func tPostXmlTest(want string) checkFunc {
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
