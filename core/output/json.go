package output

import (
	"log"
	"sync"
)

type (
	JsonResult struct {
		filename string
		result   *Result
	}
)

func (jr JsonResult) Receiving(ch <-chan SuiteResult, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	for data := range ch {
		log.Println(data)
	}
}
