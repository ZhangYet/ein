package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/ZhangYet/ein"
	"github.com/lileio/lile"
)

var s = EinServer{}
var cli ein.EinClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		ein.RegisterEinServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = ein.NewEinClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
