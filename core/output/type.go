package output

type (
	Result struct {
		Id            string       `json:"id"`
		Name          string       `json:"name"`
		StartTime     int64        `json:"startTime"`
		EndTime       int64        `json:"endTime"`
		Time          int64        `json:"time"`
		Success       bool         `json:"success"`
		Number        int          `json:"number"`
		NumberFail    int          `json:"numberFail"`
		NumberSuccess int          `json:"numberSuccess"`
		SuiteResults  SuiteResults `json:"result"`
	}

	SuiteResults []SuiteResult
	SuiteResult  struct {
		Id            string      `json:"id"`
		Name          string      `json:"name"`
		StartTime     int64       `json:"startTime"`
		EndTime       int64       `json:"endTime"`
		Time          int64       `json:"time"`
		Success       bool        `json:"success"`
		Number        int         `json:"number"`
		NumberFail    int         `json:"numberFail"`
		NumberSuccess int         `json:"numberSuccess"`
		StepsResult   StepsResult `json:"result"`
	}
	StepsResult []StepResult
	StepResult  struct {
		Id              string          `json:"id"`
		Name            string          `json:"name"`
		StartTime       int64           `json:"startTime"`
		EndTime         int64           `json:"endTime"`
		Time            int64           `json:"time"`
		Success         bool            `json:"success"`
		Body            string          `json:"body"`
		ValidateResults ValidateResults `json:"result"`
	}
	ValidateResults []ValidateResult
	ValidateResult  struct {
		Id      string `json:"id"`
		Success bool   `json:"success"`
		Op      string `json:"op"`
		Detail  string `json:"detail"`
	}
)
