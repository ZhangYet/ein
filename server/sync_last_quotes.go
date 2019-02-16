package server

import (
	"time"

	"github.com/ZhangYet/ein/common"
)

func (s *EinServer) SyncLastQuotes() error {
	quotes, err := common.GetLastQuotes()
	if err != nil {
		return err
	}
	for _, quote := range quotes {
		common.QuoteData[quote.Symbol] = quote
	}
	common.UpdateQuotaInfo.UpdateNum = int64(len(quotes))
	common.UpdateQuotaInfo.Timestamp = time.Now().Unix()
	return nil
}
