package output

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/autorunners/meerkat/utils"
)

type (
	JsonResult struct {
		filename string
		result   *Result
	}
)

func New(name string) *JsonResult {
	filename := fmt.Sprintf("/tmp/%s.json", name)
	return &JsonResult{
		filename: filename,
		result: &Result{
			Id:        utils.GenerateUUID(),
			Name:      name,
			StartTime: utils.GetTimestampMilli(),
		},
	}
}

func (jr *JsonResult) Receiving(ch <-chan SuiteResult, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	number := 0
	success := true
	for data := range ch {
		log.Println(data)
		jr.result.SuiteResults = append(jr.result.SuiteResults, data)
		number = number + data.Number
		if data.Success == false {
			success = false
		}
	}
	jr.result.EndTime = utils.GetTimestampMilli()
	jr.result.Time = jr.result.EndTime - jr.result.StartTime
	jr.result.Success = success
	jr.result.Number = number
	log.Println(jr.result)
	json, err := json.Marshal(jr.result)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(jr.filename, json, os.ModePerm)
}
