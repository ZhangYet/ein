package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLastQuotes(t *testing.T) {
	res, err := GetLastQuotes()
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
