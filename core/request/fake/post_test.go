package fake

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/autorunners/meerkat/core/request"
)

func groupPost(gs ...[]testModel) (newGs []testModel) {
	for _, g := range gs {
		newGs = append(newGs, g...)
	}
	return newGs
}

func TestPost(t *testing.T) {
	tts := groupPost(groupPostNormal())
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(tt.h)
			server := httptest.NewServer(handler)

			req := request.Request{
				Method:  "POST",
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
