package app

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"

	"github.com/autorunners/meerkat/core/config"
)

func readYaml(path string) (config.Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}
	var obj config.Config
	//log.Println(string(data))
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		log.Panic(err)
	}

	suites := obj.Suites
	// 把global中的配置合并到suites中
	for _, suite := range suites {
		log.Printf("scene name %s begin working", suite.Name)
		for _, step := range suite.Steps {
			log.Printf("step name %s begin", step.Name)
			req := step.Request
			headers := req.Headers
			method := req.Method
			uri := req.Uri
			cookie := req.Cookies
			log.Println(uri, method, headers, cookie)
		}
	}

	// just for debug
	log.Println(obj)
	return obj, nil

}
