package cmd

import (
	"os"

	"k8s.io/component-base/cli"
)

func Execute() {
	os.Exit(cli.Run(rootCmd))
}
