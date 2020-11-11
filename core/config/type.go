package config

type (
	Config struct {
		Global Global `yaml:"global"`
		Suite  Suite  `yaml:"suite"`
	}

	Global struct {
		Name   string `yaml:"name"`
		Host   string `yaml:"host"`
		Verify bool   `yaml:"verify"`
	}

	Suite []Scene
	Scene struct {
		Name  string `yaml:"name"`
		Steps []Step `yaml:"name"`
	}
	Step struct {
		Name     string    `yaml:"name"`
		Request  Request   `yaml:"request"`
		Validate Validates `yaml:"validate"`
		Response Response  `yaml:"response"`
	}

	Validates []Validate
	Validate  struct {
		Eq []string `yaml:"eq"`
	}

	Cookie  map[string]string
	Headers map[string]string
	Request struct {
		Cookie  Cookie  `yaml:"cookie"`
		Headers Headers `yaml:"headers"`
		Method  string  `yaml:"method"`
		Uri     string  `yaml:"uri"`
		FullUri string  `yaml:"fullUri"` // 带host, 和uri只需要一个
	}

	Response struct {
		Name   string   `yaml:"name"`
		Type   string   `yaml:"type"`
		Values []string `yaml:"value"`
	}
)
