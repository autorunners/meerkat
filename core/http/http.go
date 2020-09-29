package http

import (
	"log"
	"net/http"
	"time"
)

type (
	Cookie  map[string]string
	Headers map[string]string
	Request struct {
		Cookie  Cookie  `yaml:"cookie"`
		Headers Headers `yaml:"headers"`
		Method  string  `yaml:"method"`
		Uri     string  `yaml:"uri"`
		FullUri string  `yaml:"fullUri"` // 带host, 和uri只需要一个
	}
)

func (req Request) Handle() error {
	log.Println(req)
	request, err := http.NewRequest(req.Method, req.FullUri, nil)
	if err != nil {
		log.Println(err)
		return err
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

	resp, err := client.Do(request)
	if err != nil {
		doError(err)
	}
	log.Println(resp)
	doCheck(resp)

	return nil
}

func doError(err error) {
	log.Println(err)
}

func doCheck(resp *http.Response) {

}
