package main

import (
	"context"
	_ "net/http/pprof"
	"time"

	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"google.golang.org/grpc"

	"github.com/ZhangYet/ein"
	"github.com/ZhangYet/ein/common"
	"github.com/ZhangYet/ein/ein/cmd"
	"github.com/ZhangYet/ein/server"

	opentracing "github.com/opentracing/opentracing-go"
	zipkintracer "github.com/openzipkin/zipkin-go-opentracing"
)

func main() {
	logr.SetLevelFromEnv()
	recorder := zipkintracer.NewInMemoryRecorder()
	tracer, err := zipkintracer.NewTracer(recorder)
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)

	s := &server.EinServer{}

	lile.Name("ein")
	lile.Server(func(g *grpc.Server) {
		ein.RegisterEinServer(g, s)
	})

	ticker := time.NewTicker(time.Second * 60)
	defer ticker.Stop()
	go func() {
		span := tracer.StartSpan("SyncLastQuotes")
		ctx := opentracing.ContextWithSpan(context.Background(), span)
		for range ticker.C {
			err := s.SyncLastQuotes()
			if err != nil {
				logr.WithCtx(ctx).Error(err)
			} else {
				logr.WithCtx(ctx).Info(common.UpdateQuotaInfo)
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
