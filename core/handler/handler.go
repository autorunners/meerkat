package handler

import (
	"context"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/autorunners/meerkat/core/config"
	"github.com/autorunners/meerkat/core/output"
	"github.com/autorunners/meerkat/utils"
)

func Handler(ctx context.Context, conf config.Config) {
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
		//go handlerSuite(suite, gReq, ch)
		handlerSuite(suite, gReq, ch, &wgSuite)
	}

	wgSuite.Wait()
	time.Sleep(time.Millisecond * 100)
	close(ch)

	wg.Wait()

}

func handlerSuite(suite config.Suite, gReq config.Request, ch chan output.SuiteResult, wg *sync.WaitGroup) {
	defer wg.Done()
	suiteResult := output.SuiteResult{
		Id:        utils.GenerateUUID(),
		Name:      suite.Name,
		StartTime: utils.GetTimestampMilli(),
	}
	success := handlerSteps(suite.Steps, gReq, &suiteResult)
	suiteResult.EndTime = utils.GetTimestampMilli()
	suiteResult.Success = success

	ch <- suiteResult
}

func handlerSteps(steps config.Steps, gReq config.Request, suiteResult *output.SuiteResult) (success bool) {
	success = true
	for _, step := range steps {
		log.Println(step)
		stepResult := output.StepResult{
			Id:        utils.GenerateUUID(),
			Name:      step.Name,
			StartTime: utils.GetTimestampMilli(),
		}

		err, body := handlerStep(step, gReq)
		if err != nil {
			stepResult.Success = false
			stepResult.Body = err.Error()
			stepResult.EndTime = utils.GetTimestampMilli()
			success = false
		} else {
			stepResult.Success = true
			stepResult.Body = string(body)
			stepResult.EndTime = utils.GetTimestampMilli()
			// @todo Validate验证判断success是否为true
		}
		suiteResult.StepsResult = append(suiteResult.StepsResult, stepResult)
	}
	return
}

func handlerStep(step config.Step, gReq config.Request) (error, []byte) {
	req := step.Request
	name := step.Name
	log.Printf("[req] %v is begin", name)
	newReq := mergeGlobal(req, gReq)
	resp, err := newReq.Handle()
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	validates := step.Validates
	validates.Check(resp, step.Response.Type)
	return nil, body
}

// 如果没有相关配置，则使用global中的相关配置
func mergeGlobal(req config.Request, gReq config.Request) config.Request {
	if req.Host == "" {
		req.Host = gReq.Host
	}
	if req.Timeout == 0 {
		req.Timeout = gReq.Timeout
	}
	for hk, hv := range gReq.Headers {
		if req.Headers[hk] == "" {
			req.Headers[hk] = hv
		}
	}
	for ck, cv := range gReq.Cookies {
		if req.Cookies[ck] == "" {
			req.Cookies[ck] = cv
		}
	}
	return req
}
