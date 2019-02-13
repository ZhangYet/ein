package cmd

import (
	"fmt"
	"os"

	"github.com/lileio/lile"
)

var cfgFile string

var RootCmd = lile.BaseCommand("ein", "A gRPC based service")

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
