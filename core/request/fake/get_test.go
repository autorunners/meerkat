package fake

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/autorunners/meerkat/core/request"
)

func TestGet(t *testing.T) {
	req := getRequest()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	server := httptest.NewServer(handler)
	defer server.Close()
	req.FullUri = server.URL
	req.Method = "GET"
	req.Timeout = 50

	res, err := req.Handle()
	if err != nil {
		t.Fatal(err)
	}
	got, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "hello" {
		t.Errorf("got %q, want hello", string(got))
	}

}

func getRequest() request.Request {
	return request.Request{}
}
