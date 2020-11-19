package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func (validates Validates) Check(resp *http.Response, t string) (error, []byte) {
	log.Println(t)
	//if t == "" {
	//	t = "json"
	//}
	switch t {
	case "json":
		return validates.jsonCheck(resp) // @todo 未来使用单独的类处理，不同的类处理不同的类型
	default:
		log.Println("type not support for now")
		return errors.New("type not support for now"), nil
	}

}

func (validates Validates) jsonCheck(resp *http.Response) (error, []byte) {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err, nil
	}
	log.Println(string(respBody))
	var mapBody map[string]interface{}
	err = json.Unmarshal(respBody, &mapBody)
	if err != nil {
		log.Println(err, respBody)
		return err, nil
	}
	for _, validate := range validates {
		log.Println(validate)
		op := validate.Op
		log.Println(op)
		if op[0] == "http" {
			if err := checkHttpOp(op, resp); err != nil {
				return err, nil
			}
		} else if op[0] == "body" {
			if err := checkBodyOp(op, mapBody); err != nil {
				return err, nil
			}
		} else {
			return errors.New("unknown"), nil
		}
	}
	return nil, respBody
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
