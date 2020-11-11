package handler

import (
	"context"
	"log"

	"github.com/autorunners/meerkat/core/config"
)

func Handler(ctx context.Context, conf config.Config) error {

	suite := conf.Suite
	global := conf.Global
	gReq := global.Request
	for _, scene := range suite {
		log.Printf("scene name %s begin working", scene.Name)
		for _, step := range scene.Steps {
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
			for ck, cv := range gReq.Cookie {
				if req.Cookie[ck] == "" {
					req.Cookie[ck] = cv
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
