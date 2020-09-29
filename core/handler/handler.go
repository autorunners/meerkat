package handler

import (
	"context"
	"log"

	"github.com/autorunners/meerkat/core/config"
)

func Handler(ctx context.Context, conf config.C) error {

	steps := conf.TestSteps
	global := conf.Global
	//log.Println(steps, global)

	for _, step := range steps {
		log.Println(step)
		req := step.Request
		name := step.Name
		log.Printf("[req] %v is begin", name)
		if req.FullUri == "" {
			req.FullUri = global.Host + req.Uri
		}
		err := req.Handle()
		if err != nil {
			return err
		}
	}
	return nil

}
