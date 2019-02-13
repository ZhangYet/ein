package main

import (
	_ "net/http/pprof"

	"github.com/ZhangYet/ein"
	"github.com/ZhangYet/ein/ein/cmd"
	"github.com/ZhangYet/ein/server"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.EinServer{}

	lile.Name("ein")
	lile.Server(func(g *grpc.Server) {
		ein.RegisterEinServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
