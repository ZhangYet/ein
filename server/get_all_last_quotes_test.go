package server

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ZhangYet/ein"
	"github.com/ZhangYet/ein/common"
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

func TestGetAllLastQuotes(t *testing.T) {
	ctx := context.Background()
	req := &ein.QuoteRequest{}

	res, err := cli.GetAllLastQuotes(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 1, len(res.Data))
}
