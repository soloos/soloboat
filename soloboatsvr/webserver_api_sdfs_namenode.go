package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/log"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboatsdk"
)

func (p *WebServer) apiSDFSNameNodeHeartBeat(ir *iron.Request) {
	var (
		req soloboatsdk.SDFSNameNodeHeartBeatReqJSON
		err error
	)

	err = ir.DecodeBodyJSONData(&req)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	log.Error("fuck heatbeat", req.PeerID)
	ir.ApiOutput(nil, snettypes.CODE_OK, "")
}