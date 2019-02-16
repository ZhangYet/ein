package server

import "github.com/ZhangYet/ein/common"

func (s *EinServer) SyncLastQuotes() error {
	quotes, err := common.GetLastQuotes()
	if err != nil {
		return err
	}
	for _, quote := range quotes {
		common.QuoteData[quote.Symbol] = quote
	}
	return nil
}
