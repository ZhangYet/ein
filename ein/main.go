package main

import (
	"time"

	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"google.golang.org/grpc"

	"github.com/ZhangYet/ein"
	"github.com/ZhangYet/ein/common"
	"github.com/ZhangYet/ein/ein/cmd"
	"github.com/ZhangYet/ein/server"

	_ "net/http/pprof"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
)

func main() {

	s := &server.EinServer{}

	lile.Name("ein")
	lile.Server(func(g *grpc.Server) {
		ein.RegisterEinServer(g, s)
	})

	lile.AddUnaryInterceptor(grpc_logrus.UnaryServerInterceptor(&common.LogrusEntry))
	lile.AddUnaryInterceptor(grpc_opentracing.UnaryServerInterceptor())

	if err := common.LoadQuoteData(); err != nil {
		common.LogrusLogger.Error(err)
	}

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			err := s.SyncLastQuotes()
			if err != nil {
				common.LogrusLogger.Error(err)
			} else {
				common.LogrusLogger.Info(common.UpdateQuotaInfo)
			}

		}
	}()

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
