package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
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
	fmt.Println("meerkat-check is running...")
	return nil
}
