package common

import "github.com/ZhangYet/ein"

// 配置文件，公用变量

var QuoteData map[string]*ein.LastQuote

var UpdateQuotaInfo *ein.UpdateInfo

func init() {
	QuoteData = make(map[string]*ein.LastQuote)
	UpdateQuotaInfo = &ein.UpdateInfo{}
}
