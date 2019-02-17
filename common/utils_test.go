package common

import (
	"testing"
	"time"

	"github.com/ZhangYet/ein"

	"github.com/stretchr/testify/assert"
)

func TestGetLastQuotes(t *testing.T) {
	res, err := GetLastQuotes()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestCacheQuoteDataAndLoadQuoteData(t *testing.T) {
	QuoteData = map[string]*ein.LastQuote{
		"IBM": &ein.LastQuote{
			Symbol: "IBM",
			Time:   time.Now().Unix(),
			Price:  42.0,
			Size:   100,
		},
	}
	errCache := CacheQuoteData()
	assert.Nil(t, errCache)
	QuoteData = map[string]*ein.LastQuote{}
	errLoad := LoadQuoteData()
	assert.Nil(t, errLoad)
	assert.NotNil(t, QuoteData)
}
