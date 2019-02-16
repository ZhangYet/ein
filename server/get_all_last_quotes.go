package server

import (
	"context"

	"github.com/ZhangYet/ein"
	"github.com/ZhangYet/ein/common"
)

func (s EinServer) GetAllLastQuotes(ctx context.Context, r *ein.QuoteRequest) (*ein.LastQuotesResponse, error) {
	ret := &ein.LastQuotesResponse{
		Data: []*ein.LastQuote{},
	}
	for _, lastQuote := range common.QuoteData {
		ret.Data = append(ret.Data, lastQuote)
	}
	return ret, nil
}
