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
		for _, step := range suite.Steps {
			log.Println(step)
			req := step.Request
			name := step.Name
			log.Printf("[req] %v is begin", name)
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

			resp, err := req.Handle()
			if err != nil {
				return err
			}
			validates := step.Validates
			validates.Check(resp, step.Response.Type)
		}
	}

	return nil

}
