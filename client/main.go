package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/ZhangYet/ein"

	opentracing "github.com/opentracing/opentracing-go"
	zipkintracer "github.com/openzipkin/zipkin-go-opentracing"
	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	einClient := ein.GetEinClient("127.0.0.1:8000") // TODO 通过公共参数配置

	recorder := zipkintracer.NewInMemoryRecorder()
	tracer, err := zipkintracer.NewTracer(recorder)
	if err != nil {
		panic(err)
	}
	commands := []*cli.Command{
		{
			Name:  "getAllLastQuotes",
			Usage: "get all quotes",
			Action: func(ctx *cli.Context) error {
				span := tracer.StartSpan(ctx.Command.Name)
				einCtx := opentracing.ContextWithSpan(context.Background(), span)
				req := &ein.QuoteRequest{}
				res, err := einClient.GetAllLastQuotes(einCtx, req)
				if err == nil {
					for _, quote := range res.Data {
						fmt.Printf("Symbol: %s, price: %f \n", quote.Symbol, quote.Price) // TODO opti output
					}
				}
				return err
			},
		},
		{
			Name:  "getLastQuote",
			Usage: "get one quote",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "symbol",
					Aliases: []string{"s"},
				},
			},
			Action: func(ctx *cli.Context) error {
				symbol := ctx.Args().Get(1) // TODO 公共的参数解析参数
				span := tracer.StartSpan(ctx.Command.Name)
				einCtx := opentracing.ContextWithSpan(context.Background(), span)
				req := &ein.QuoteRequest{
					Key: symbol,
				}
				res, err := einClient.GetLastQuote(einCtx, req)
				if err == nil {
					fmt.Printf("Symbol: %s, price: %f \n", res.Data.Symbol, res.Data.Price)
				}
				return err
			},
		},
		{
			Name:  "daemon",
			Usage: "run in daemon and waiting for update info",
			Action: func(ctx *cli.Context) error {
				span := tracer.StartSpan(ctx.Command.Name)
				einCtx := opentracing.ContextWithSpan(context.Background(), span)
				req := &ein.StreamRequest{}
				resChan := make(chan ein.Ein_StreamUpdateInfoClient)
				go func() {
					for {
						recv, err := einClient.StreamUpdateInfo(einCtx, req)
						if err != nil {
							panic(err)
						}
						resChan <- recv
						time.Sleep(3 * time.Second)
					}
				}()

				for {
					recv := <-resChan
					res, err := recv.Recv()
					if err != nil && err != io.EOF {
						break
					}
					if res != nil {
						fmt.Printf("number: %d, timestamp: %d\n", res.UpdateNum, res.Timestamp)
					}
					time.Sleep(time.Second)
				}
				return nil
			},
		},
	}
	app := cli.App{}
	app.Name = "ein-client"
	app.Commands = commands
	app.Run(os.Args)
}
