package handler

import (
	"context"
	"log"

	"github.com/autorunners/meerkat/core/config"
)

func Handler(ctx context.Context, conf config.Config) error {
	suites := conf.Suites
	global := conf.Global
	gReq := global.Request
	for _, suite := range suites {
		log.Printf("scene name %s begin working", suite.Name)
		if err := handlerSteps(suite.Steps, gReq); err != nil {
			return err
		}
	}
	return nil

}

func handlerSteps(steps config.Steps, gReq config.Request) error {
	for _, step := range steps {
		log.Println(step)
		if err := handlerStep(step, gReq); err != nil {
			return err
		}
	}
	return nil
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
