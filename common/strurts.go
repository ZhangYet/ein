package common

import (
	"os"

	"github.com/ZhangYet/ein"

	"github.com/sirupsen/logrus"
)

// 配置文件，公用变量

var (
	QuoteData       map[string]*ein.LastQuote
	UpdateQuotaInfo *ein.UpdateInfo

	LogrusLogger = logrus.New()
	LogrusEntry  = logrus.Entry{}
)

func init() {
	QuoteData = make(map[string]*ein.LastQuote)
	UpdateQuotaInfo = &ein.UpdateInfo{}

	LogrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logFile, err := os.Create("ein.log")
	if err != nil {
		panic(err)
	}
	LogrusLogger.SetOutput(logFile)
	LogrusEntry = *logrus.NewEntry(LogrusLogger)
}
