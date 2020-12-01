package result

type Result struct {
	UUID        string   `json:"uuid,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Status      string   `json:"status,omitempty"`
	Stage       string   `json:"stage,omitempty"`
	Steps       []step   `json:"steps,omitempty"`
	Start       int64    `json:"start,omitempty"`
	Stop        int64    `json:"stop,omitempty"`
	Children    []string `json:"children,omitempty"`
	FullName    string   `json:"fullName,omitempty"`
	Labels      []label  `json:"labels,omitempty"`
	Test        func()   `json:"-"`
}

type label struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type parameter struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type step struct {
	Name          string       `json:"name,omitempty"`
	Status        string       `json:"status,omitempty"`
	Stage         string       `json:"stage"`
	ChildrenSteps []step       `json:"steps"`
	Attachments   []attachment `json:"attachments"`
	Parameters    []parameter  `json:"parameters"`
	Start         int64        `json:"start"`
	Stop          int64        `json:"stop"`
	Action        func()       `json:"-"`
}

type attachment struct {
	uuid    string
	Name    string   `json:"name"`
	Source  string   `json:"source"`
	Type    MimeType `json:"type"`
	content []byte
}

type MimeType string
