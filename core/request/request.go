package request

import (
	"log"
	"net/http"
	"time"
)

type (
	Cookie  map[string]string
	Headers map[string]string
	Request struct {
		Cookies Cookie  `yaml:"cookies"`
		Headers Headers `yaml:"headers"`
		Method  string  `yaml:"method"`
		Uri     string  `yaml:"uri"`
		Host    string  `yaml:"host"`
		FullUri string  `yaml:"fullUri"`
		Timeout int64   `yaml:"timeout"`
	}
)

func (req Request) Handle() (*http.Response, error) {
	log.Println(req)
	request, err := http.NewRequest(req.Method, req.FullUri, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range req.Headers {
		request.Header.Add(key, value)
	}
	for key, value := range req.Cookies {
		cookie := &http.Cookie{
			Name:  key,
			Value: value,
		}
		request.AddCookie(cookie)
	}

	client := &http.Client{
		Timeout: time.Duration(req.Timeout) * time.Millisecond,
	}

	return client.Do(request)
}
