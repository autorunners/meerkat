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

func (jr JsonResult) Receiving(wg sync.WaitGroup, ch <-chan map[string]interface{}) {
	for datas := range ch {
		key := datas["key"]
		switch key {
		case "start-suites":
			doStartSuites(datas)
		case "end-suites":
			doEndSuites(datas)
		default:
			doEndSuites(datas)

		}
	}
}

func doStartSuites(datas map[string]interface{}) {
	log.Println(datas)
}

func doEndSuites(datas map[string]interface{}) {
	log.Println(datas)
}
