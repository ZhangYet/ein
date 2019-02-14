package server

import (
	"errors"
	"github.com/ZhangYet/ein/common"

	"context"
	"github.com/ZhangYet/ein"
)

func (s EinServer) GetLastQuote(ctx context.Context, r *ein.QuoteRequest) (*ein.LastQuoteResponse, error) {
	if r.Key == "" {
		return nil, errors.New("invalid param")
	}
	lastQuote, ok := common.QuoteData[r.Key]
	if !ok {
		return nil, nil
	}
	ret := &ein.LastQuoteResponse{
		Data: lastQuote,
	}
	return ret, nil
}
