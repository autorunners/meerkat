package app

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"

	"github.com/autorunners/meerkat/core/config"
)

func readYaml() (config.Config, error) {
	data, err := ioutil.ReadFile("../../config/config.yaml")
	if err != nil {
		log.Panic(err)
	}
	var obj config.Config
	//log.Println(string(data))
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		log.Panic(err)
	}

	// just for debug
	log.Println(obj)
	suite := obj.Suite

	// 把global中的配置合并到suites中
	for _, scene := range suite {
		log.Printf("scene name %s begin working", scene.Name)
		for _, step := range scene.Steps {
			log.Printf("step name %s begin", step.Name)
			req := step.Request
			headers := req.Headers
			method := req.Method
			uri := req.Uri
			cookie := req.Cookie
			log.Println(uri, method, headers, cookie)
		}
	}

	return obj, nil

}
