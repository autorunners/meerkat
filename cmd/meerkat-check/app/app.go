package app

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/autorunners/meerkat/core/handler"
)

func NewCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:  "meerkat-check",
		Long: `主要用于在新api上线前、单元测试后进行功能测试，确保api接口能按设想的运行。`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runCommand(); err != nil {
				log.Fatalf("%v\n", err)
			}
		},
	}
	return cmd
}

func runCommand() error {
	log.Println("meerkat-check is running...")

	config, err := readYaml()
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		log.Println("canceling context")
		cancel()
	}()

	go handler.Handler(ctx, config)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	log.Println("Received termination, signaling shutdown")

	return err
}
