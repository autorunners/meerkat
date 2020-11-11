package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func (validates Validates) Check(resp *http.Response, t string) error {
	if t == "" {
		t = "json"
	}
	switch t {
	case "json":
		return validates.jsonCheck(resp)
	}
	return errors.New("not json")

}

func (validates Validates) jsonCheck(resp *http.Response) error {
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var mapBody map[string]interface{}
	err = json.Unmarshal(respBody, mapBody)
	if err != nil {
		return err
	}
	for _, validate := range validates {
		log.Println(validate)
		op := validate.Op
		log.Println(op)
		if op[0] == "http" {
			if err := checkHttpOp(op, resp); err != nil {
				return err
			}
		} else if op[0] == "body" {
			if err := checkBodyOp(op, mapBody); err != nil {
				return err
			}
		} else {
			return errors.New("unknown")
		}
	}
	return nil
}

func checkHttpOp(op []string, resp *http.Response) error {
	if op[1] == "eq" {
		if op[2] == "status" {
			if resp.Status == op[3] {
				return nil
			} else {
				return errors.New("not eq")
			}
		}
	}
	return errors.New("not eq")
}
func checkBodyOp(op []string, mapBody map[string]interface{}) error {
	log.Println("body:", op, mapBody)
	return nil
}
