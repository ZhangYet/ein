package server

import (
	"github.com/ZhangYet/ein/common"

	"context"
	"github.com/ZhangYet/ein"
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
