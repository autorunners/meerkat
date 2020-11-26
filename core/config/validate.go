package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/autorunners/meerkat/core/output"
	"github.com/autorunners/meerkat/utils"
)

func (validates Validates) Check(resp *http.Response) (body string, numberFail int, numberSuccess int, results output.ValidateResults) {
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		numberFail++
		return
	}
	body = string(respBody)
	log.Println(body)
	for _, validate := range validates {
		log.Println(validate)
		result := output.ValidateResult{
			Id: utils.GenerateUUID(),
			Op: strings.Join(validate.Op, "-"),
		}
		ops := validate.Op
		log.Println(ops)
		if err := checkOp(ops, respBody, resp); err != nil {
			numberFail++
			result.Success = false
			result.Detail = err.Error()
			results = append(results, result)
			continue
		}
		numberSuccess++
		result.Detail = "success"
		result.Success = true
		results = append(results, result)
	}
	return
}

func checkOp(ops []string, respBody []byte, resp *http.Response) error {
	if ops[0] == "http" {
		if err := checkHttpOp(ops, resp); err != nil {
			return err
		}
	} else if ops[0] == "body" {
		var mapBody map[string]interface{}
		if ops[1] == "json" {
			err := json.Unmarshal(respBody, &mapBody)
			if err != nil {
				log.Println(err)
				return err
			}
		} else {
			return errors.New("body type not eq json")
		}
		if err := checkBodyOp(ops, mapBody); err != nil {
			return err
		}
	} else {
		return errors.New("invalid OP type")
	}
	return nil
}

func checkHttpOp(ops []string, resp *http.Response) error {
	log.Println(ops[1], ops[2], ops[3], resp.StatusCode)
	if ops[1] == "eq" {
		if ops[2] == "status_code" {
			if resp.Status == ops[3] {
				return nil
			} else {
				return errors.New("status not eq")
			}
		}
	}
	return errors.New("invalid eq")
}
func checkBodyOp(ops []string, mapBody map[string]interface{}) error {
	log.Println("body:", ops, mapBody)
	return nil
}
