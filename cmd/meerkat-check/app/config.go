package app

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"

	"github.com/autorunners/meerkat/core/config"
)

func readYaml() (config.C, error) {
	data, err := ioutil.ReadFile("../../config/config.yaml")
	if err != nil {
		log.Panic(err)
	}
	var obj config.C
	//log.Println(string(data))
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		log.Panic(err)
	}

	log.Println(obj)

	steps := obj.TestSteps
	for _, step := range steps {
		log.Println(step.Name)
		req := step.Request
		headers := req.Headers
		method := req.Method
		uri := req.Uri
		cookie := req.Cookie
		log.Println(uri, method, headers, cookie)
	}

	return obj, nil

}
