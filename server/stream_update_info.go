package server

import (
	"errors"
	"github.com/ZhangYet/ein"
)

func (s EinServer) StreamUpdateInfo(r *ein.StreamRequest, stream ein.Ein_StreamUpdateInfoServer) error {
	res := &ein.UpdateInfo{}
	err := stream.Send(res)
	if err != nil {
		return err
	}

	return errors.New("not yet implemented")
}
