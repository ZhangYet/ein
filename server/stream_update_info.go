package server

import (
	"github.com/ZhangYet/ein"
	"github.com/ZhangYet/ein/common"
)

func (s EinServer) StreamUpdateInfo(r *ein.StreamRequest, stream ein.Ein_StreamUpdateInfoServer) error {
	res := common.UpdateQuotaInfo
	err := stream.Send(res)
	if err != nil {
		return err
	}
	return nil
}
