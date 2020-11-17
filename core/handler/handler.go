package handler

import (
	"context"
	"log"
	"sync"

	"github.com/autorunners/meerkat/core/config"
	"github.com/autorunners/meerkat/core/output"
)

func Handler(ctx context.Context, conf config.Config) {
	log.Println("Handler start===============================")
	var (
		wg sync.WaitGroup
		ch chan map[string]interface{}
	)
	ch = make(chan map[string]interface{}, 100)
	wg.Add(1)
	suites := conf.Suites
	global := conf.Global
	gReq := global.Request

	result := output.Init(global.Name)
	go result.Receiving(wg, ch)

	startSuites := map[string]interface{}{
		"key":  "start-suites",
		"name": global.Name,
	}
	ch <- startSuites

	for _, suite := range suites {
		log.Printf("suite name %s begin working", suite.Name)
		startSuite := map[string]interface{}{
			"key":  "start-suite",
			"name": suite.Name,
		}
		ch <- startSuite
		handlerSteps(suite.Steps, gReq, ch)

		endSuite := map[string]interface{}{
			"key":  "end-suite",
			"name": suite.Name,
		}
		ch <- endSuite
	}

	endSuites := map[string]interface{}{
		"key":  "end-suites",
		"name": global.Name,
	}
	ch <- endSuites

	wg.Wait()

}

func handlerSteps(steps config.Steps, gReq config.Request, ch chan map[string]interface{}) {
	for _, step := range steps {
		log.Println(step)
		startStep := map[string]interface{}{
			"key":  "start-step",
			"name": step.Name,
		}
		ch <- startStep

		endStep := map[string]interface{}{
			"key":  "end-step",
			"name": step.Name,
		}
		if err := handlerStep(step, gReq); err != nil {
			endStep["success"] = false
			endStep["body"] = err.Error()
			ch <- endStep
		} else {
			endStep["success"] = true
			endStep["body"] = "success"
			ch <- endStep
		}
	}
}

func handlerStep(step config.Step, gReq config.Request) error {
	req := step.Request
	name := step.Name
	log.Printf("[req] %v is begin", name)
	newReq := mergeGlobal(req, gReq)
	resp, err := newReq.Handle()
	if err != nil {
		return err
	}
	validates := step.Validates
	validates.Check(resp, step.Response.Type)
	return nil
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
