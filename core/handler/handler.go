package handler

import (
	"context"
	"log"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/autorunners/meerkat/core/config"
	"github.com/autorunners/meerkat/core/output"
	"github.com/autorunners/meerkat/core/request"
	"github.com/autorunners/meerkat/utils"
)

func Handler(ctx context.Context, conf config.Config, chSignal chan os.Signal) {
	log.Println("Handler start===============================")
	var (
		wg      sync.WaitGroup
		wgSuite sync.WaitGroup
		ch      chan output.SuiteResult
	)
	ch = make(chan output.SuiteResult, 100)
	wg.Add(1)
	suites := conf.Suites
	global := conf.Global
	gReq := global.Request

	// @todo 先只实现写json文件，后期可以扩展
	result := output.Init(global.Name)
	go result.Receiving(ch, &wg)

	for _, suite := range suites {
		wgSuite.Add(1)
		log.Printf("suite name %s begin working", suite.Name)
		//go handlerSuite(suite, gReq, ch, &wgSuite)
		handlerSuite(suite, gReq, ch, &wgSuite)
	}

	wgSuite.Wait()
	time.Sleep(time.Millisecond * 100)
	close(ch)

	wg.Wait()
	ctx.Done()
	chSignal <- syscall.SIGINT
}

func handlerSuite(suite config.Suite, gReq request.Request, ch chan output.SuiteResult, wg *sync.WaitGroup) {
	defer wg.Done()
	suiteResult := output.SuiteResult{
		Id:        utils.GenerateUUID(),
		Name:      suite.Name,
		StartTime: utils.GetTimestampMilli(),
	}
	success, number, numberFail, numberSuccess, stepResults := handlerSteps(suite.Steps, gReq)
	suiteResult.StepsResult = stepResults
	suiteResult.Success = success
	suiteResult.Number = number
	suiteResult.NumberFail = numberFail
	suiteResult.NumberSuccess = numberSuccess
	suiteResult.EndTime = utils.GetTimestampMilli()
	suiteResult.Time = suiteResult.EndTime - suiteResult.StartTime

	ch <- suiteResult
}

func handlerSteps(steps config.Steps, gReq request.Request) (success bool, number int, numberFail int, numberSuccess int, stepsResult output.StepsResult) {
	success = true
	number = 0
	numberFail = 0
	numberSuccess = 0
	stepsResult = output.StepsResult{}
	for _, step := range steps {
		log.Println(step)
		stepResult := output.StepResult{
			Id:        utils.GenerateUUID(),
			Name:      step.Name,
			StartTime: utils.GetTimestampMilli(),
		}

		body, numFail, numSuccess, results := handlerStep(step, gReq)
		if numberFail == 0 {
			stepResult.Success = true
		}
		stepResult.Body = body
		stepResult.Number = numSuccess + numFail
		stepResult.NumberSuccess = numSuccess
		stepResult.NumberFail = numFail
		stepResult.EndTime = utils.GetTimestampMilli()
		stepResult.Time = stepResult.EndTime - stepResult.StartTime
		stepResult.ValidateResults = results

		stepsResult = append(stepsResult, stepResult)

		numberFail = numberFail + numFail
		number = number + numSuccess + numFail
		numberSuccess = numberSuccess + numSuccess
	}
	return
}

func handlerStep(step config.Step, gReq request.Request) (body string, numberFail int, numberSuccess int, results output.ValidateResults) {
	req := step.Request
	name := step.Name
	log.Printf("[req] %v is begin", name)
	newReq := mergeGlobal(req, gReq)
	resp, err := newReq.Handle()
	if err != nil {
		numberFail++
		return
	}
	defer resp.Body.Close()
	validates := step.Validates
	body, numberFail, numberSuccess, results = validates.Check(resp)
	log.Println(body, numberFail, numberSuccess, results)
	return
}
