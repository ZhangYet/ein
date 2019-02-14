package server

import (
	"errors"

	"context"
	"github.com/ZhangYet/ein"
)

func (s EinServer) GetAllLastQuotes(ctx context.Context, r *ein.QuoteRequest) (*ein.LastQuotesResponse, error) {
	return nil, errors.New("not yet implemented")
}
