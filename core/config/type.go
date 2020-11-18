package config

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
		Name      string    `yaml:"name"`
		Request   Request   `yaml:"request"`
		Validates Validates `yaml:"validates"`
		Response  Response  `yaml:"response"`
	}

	Validates []Validate
	Validate  struct {
		Op []string `yaml:"op"`
	}

	Cookie  map[string]string
	Headers map[string]string
	Request struct {
		Cookies Cookie  `yaml:"cookies"`
		Headers Headers `yaml:"headers"`
		Method  string  `yaml:"method"`
		Uri     string  `yaml:"uri"`
		Host    string  `yaml:"host"`
		Timeout int64   `yaml:"timeout"`
	}

	Response struct {
		Name   string   `yaml:"name"`
		Type   string   `yaml:"type"`
		Values []string `yaml:"value"`
	}
)
