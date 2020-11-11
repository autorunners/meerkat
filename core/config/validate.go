package config

import (
	"net/http"
)

func (validates Validates) Check(resp *http.Response, t string) error {
	if t == "" {
		t = "json"
	}
	switch t {
	case "json":
		jsonCheck()
	}
	return nil
}

func jsonCheck() {

}
