package config

import (
	"log"
	"net/http"
	"time"
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
	for key, value := range req.Cookie {
		cookie := &http.Cookie{
			Name:  key,
			Value: value,
		}
		request.AddCookie(cookie)
	}

	client := &http.Client{
		Timeout: time.Second, // @todo
	}

	return client.Do(request)
}
