package server

import (
	"context"
	"errors"

	"github.com/ZhangYet/ein"
	"github.com/ZhangYet/ein/common"
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
