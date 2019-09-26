package solonn

import (
	"soloos/common/solofsapitypes"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

func (p *SolonnDriver) SolonnHeartBeat(heartbeat solofsapitypes.SolonnHeartBeat) error {
	var err error
	var peerID = snettypes.StrToPeerID(heartbeat.SRPCPeerID)
	var iptr, exists = p.solonnTable.Load(peerID)
	var solonnInfo = soloboattypes.SolonnInfo{PeerID: peerID}
	if exists {
		solonnInfo = iptr.(soloboattypes.SolonnInfo)
	}

	solonnInfo.LastHeatBeatAt = time.Now()
	solonnInfo.SolonnHeartBeat = heartbeat
	err = p.FormatSolonnInfo(&solonnInfo)
	if err != nil {
		return err
	}

	p.solonnTable.Store(peerID, solonnInfo)

	return nil
}