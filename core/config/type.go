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

	Result struct {
		Id           string       `json:"id"`
		Name         string       `json:"name"`
		Time         uint64       `json:"time"`
		Success      bool         `json:"success"`
		Number       uint16       `json:"number"`
		SuiteResults SuiteResults `json:"result"`
	}

	SuiteResults []SuiteResult
	SuiteResult  struct {
		Id          string      `json:"id"`
		Name        string      `json:"name"`
		Time        uint64      `json:"time"`
		Success     bool        `json:"success"`
		StepsResult StepsResult `json:"result"`
	}
	StepsResult []StepResult
	StepResult  struct {
		Id              string          `json:"id"`
		Name            string          `json:"name"`
		Time            uint64          `json:"time"`
		Success         bool            `json:"success"`
		Body            []byte          `json:"body"`
		ValidateResults ValidateResults `json:"result"`
	}
	ValidateResults []ValidateResult
	ValidateResult  struct {
		Id      string `json:"id"`
		Success bool   `json:"success"`
		Detail  string `json:"detail"`
	}
)
