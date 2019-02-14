package server

import (
	"testing"

	"context"
	"github.com/ZhangYet/ein"
	"github.com/stretchr/testify/assert"
)

func TestGetAllLastQuotes(t *testing.T) {
	ctx := context.Background()
	req := &ein.QuoteRequest{}

	res, err := cli.GetAllLastQuotes(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
