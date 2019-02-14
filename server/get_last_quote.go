package server

import (
	"errors"

	"context"
	"github.com/ZhangYet/ein"
)

func (s EinServer) GetLastQuote(ctx context.Context, r *ein.QuoteRequest) (*ein.LastQuoteResponse, error) {
	return nil, errors.New("not yet implemented")
}
