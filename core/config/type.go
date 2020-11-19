package config

import (
	"github.com/autorunners/meerkat/core/request"
)

type (
	Config struct {
		Global Global `yaml:"global"`
		Suites Suites `yaml:"suites"`
	}

	Global struct {
		Name      string    `yaml:"name"`
		Verify    bool      `yaml:"verify"`
		Request   Request   `yaml:"request"`
		Validates Validates `yaml:"validates"`
	}

	Suites []Suite
	Suite  struct {
		Name  string `yaml:"name"`
		Steps Steps  `yaml:"steps"`
	}
	Steps []Step
	Step  struct {
		Name      string          `yaml:"name"`
		Request   request.Request `yaml:"request"`
		Validates Validates       `yaml:"validates"`
		Response  Response        `yaml:"response"`
	}

	Validates []Validate
	Validate  struct {
		Op []string `yaml:"op"`
	}

	Response struct {
		Name   string   `yaml:"name"`
		Type   string   `yaml:"type"`
		Values []string `yaml:"value"`
	}
)
