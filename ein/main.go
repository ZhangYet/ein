package main

import (
	"context"
	"os"
	"time"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"

	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/ZhangYet/ein"
	"github.com/ZhangYet/ein/common"
	"github.com/ZhangYet/ein/ein/cmd"
	"github.com/ZhangYet/ein/server"

	_ "net/http/pprof"

	opentracing "github.com/opentracing/opentracing-go"
	zipkintracer "github.com/openzipkin/zipkin-go-opentracing"
)

var logrusLogger = logrus.New()

func main() {
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
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logFile, err := os.Create("ein.log")
	if err != nil {
		panic(err)
	}
	logrusLogger.SetOutput(logFile)
	logrusEntry := logrus.NewEntry(logrusLogger)
	lile.AddUnaryInterceptor(grpc_logrus.UnaryServerInterceptor(logrusEntry))

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	go func() {
		span := tracer.StartSpan("SyncLastQuotes")
		ctx := opentracing.ContextWithSpan(context.Background(), span)
		for range ticker.C {
			err := s.SyncLastQuotes()
			if err != nil {
				logrusLogger.Error(err)
			} else {
				logr.WithCtx(ctx).Info(common.UpdateQuotaInfo)
				logrusLogger.Info(common.UpdateQuotaInfo)
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
