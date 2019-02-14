package server

import (
	"context"
	"testing"

	"github.com/ZhangYet/ein"

	"github.com/stretchr/testify/assert"
)

func TestGetLastQuote(t *testing.T) {
	ctx := context.Background()
	req := &ein.QuoteRequest{}

	res, err := cli.GetLastQuote(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
