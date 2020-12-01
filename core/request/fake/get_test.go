package fake

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/autorunners/meerkat/core/request"
)

func groupGet(gs ...[]testModel) (newGs []testModel) {
	for _, g := range gs {
		newGs = append(newGs, g...)
	}
	return newGs
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
