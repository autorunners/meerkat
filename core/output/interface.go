package output

import (
	"fmt"
	"sync"
)

type CoreInterface interface {
	Receiving(wg sync.WaitGroup, ch <-chan map[string]interface{})
}

func Init(name string) CoreInterface {
	filename := fmt.Sprintf("/tmp/%s.json", name)
	return JsonResult{
		filename: filename,
	}
}
