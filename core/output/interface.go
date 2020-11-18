package output

import (
	"fmt"
	"sync"
)

type CoreInterface interface {
	Receiving(ch <-chan SuiteResult, wg *sync.WaitGroup)
}

func Init(name string) CoreInterface {
	filename := fmt.Sprintf("/tmp/%s.json", name)
	return JsonResult{
		filename: filename,
	}
}
