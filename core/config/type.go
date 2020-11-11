package config

type (
	Config struct {
		Global Global `yaml:"global"`
		Suite  Suite  `yaml:"suite"`
	}

	Global struct {
		Name      string    `yaml:"name"`
		Verify    bool      `yaml:"verify"`
		Request   Request   `yaml:"request"`
		Validates Validates `yaml:"validates"`
	}

	Suite []Scene
	Scene struct {
		Name  string `yaml:"name"`
		Steps []Step `yaml:"name"`
	}
	Step struct {
		Name      string    `yaml:"name"`
		Request   Request   `yaml:"request"`
		Validates Validates `yaml:"validates"`
		Response  Response  `yaml:"response"`
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
		Host    string  `yaml:"host"`
		Timeout int64   `yaml:"timeout"`
	}

	Response struct {
		Name   string   `yaml:"name"`
		Type   string   `yaml:"type"`
		Values []string `yaml:"value"`
	}
)
