package server

import (
	"io"
	"testing"

	"context"
	"github.com/ZhangYet/ein"
	"github.com/stretchr/testify/assert"
)

func TestStreamUpdateInfo(t *testing.T) {
	ctx := context.Background()
	req := &ein.StreamRequest{}

	stream, err := cli.StreamUpdateInfo(ctx, req)
	assert.Nil(t, err)

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			assert.Fail(t, err.Error())
			break
		}

		assert.Nil(t, err)
		assert.NotNil(t, res)
	}
}
