package handler

import (
	"context"
	"log"

	"github.com/autorunners/meerkat/core/config"
)

func Handler(ctx context.Context, conf config.Config) error {

	suite := conf.Suite
	global := conf.Global
	for _, scene := range suite {
		log.Printf("scene name %s begin working", scene.Name)
		for _, step := range scene.Steps {
			log.Println(step)
			req := step.Request
			name := step.Name
			log.Printf("[req] %v is begin", name)
			if req.FullUri == "" {
				req.FullUri = global.Host + req.Uri
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
