package config

import (
	"github.com/autorunners/meerkat/core/http"
)

type (
	C struct {
		Global    Config    `yaml:"global"`
		TestSteps TestSteps `yaml:"teststeps"`
	}

	Config struct {
		Name   string `yaml:"name"`
		Host   string `yaml:"host"`
		Verify bool   `yaml:"verify"`
	}

	TestSteps []TestStep
	TestStep  struct {
		Name     string       `yaml:"name"`
		Request  http.Request `yaml:"request"`
		Validate Validates    `yaml:"validate"`
	}

	Validates []Validate
	Validate  struct {
		Eq []string `yaml:"eq"`
	}
)
