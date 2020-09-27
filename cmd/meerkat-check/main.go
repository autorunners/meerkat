package main

import (
	"log"

	"github.com/autorunners/meerkat/cmd/meerkat-check/app"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {

	command := app.NewCommand()

	if err := command.Execute(); err != nil {
		log.Fatalf("%v\n", err)
	}

}
