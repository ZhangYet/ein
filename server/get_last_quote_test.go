package server

import (
	"context"
	"testing"
	"time"

	"github.com/ZhangYet/ein"
	"github.com/ZhangYet/ein/common"

	"github.com/stretchr/testify/assert"
)

func init() {
	common.QuoteData = map[string]*ein.LastQuote{
		"IBM": &ein.LastQuote{
			Symbol: "IBM",
			Time:   time.Now().Unix(),
			Price:  42.0,
			Size:   100,
		},
	}
}

func TestGetLastQuote(t *testing.T) {
	ctx := context.Background()
	req := &ein.QuoteRequest{
		Key: "IBM",
	}

	res, err := cli.GetLastQuote(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
