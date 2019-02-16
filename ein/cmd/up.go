package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ZhangYet/ein/subscribers"
	"github.com/lileio/lile"
	"github.com/lileio/pubsub"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up runs both RPC service",
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		go func() {
			lile.Run()
		}()
		go func() {
			pubsub.Subscribe(&subscribers.EinServiceSubscriber{})
		}()

		<-c // TODO graceful exit
		lile.Shutdown()
		pubsub.Shutdown()
	},
}

func init() {
	RootCmd.AddCommand(upCmd)
}
