package server

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/ZhangYet/ein"
	"github.com/ZhangYet/ein/common"

	"github.com/stretchr/testify/assert"
)

func init() {
	common.UpdateQuotaInfo = &ein.UpdateInfo{
		Timestamp: time.Now().Unix(),
		UpdateNum: 4,
	}
}

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
		assert.Equal(t, int64(4), res.UpdateNum)
	}
}
