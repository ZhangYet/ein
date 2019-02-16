package server

import (
	"testing"

	"github.com/ZhangYet/ein/common"

	"github.com/stretchr/testify/assert"
)

func TestSyncLastQuotes(t *testing.T) {
	server := EinServer{}
	err := server.SyncLastQuotes()
	assert.Nil(t, err)
	assert.NotNil(t, common.QuoteData)
}
