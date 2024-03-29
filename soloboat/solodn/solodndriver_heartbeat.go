package solodn

import (
	"soloos/common/snet"
	"soloos/common/solofsapitypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

func (p *SolodnDriver) SolodnHeartBeat(heartbeat solofsapitypes.SolodnHeartBeat) error {
	var err error
	var peerID = snet.StrToPeerID(heartbeat.SrpcPeerID)
	var iptr, exists = p.solodnTable.Load(peerID)
	var solodnInfo = soloboattypes.SolodnInfo{PeerID: peerID}
	if exists {
		solodnInfo = iptr.(soloboattypes.SolodnInfo)
	}

	solodnInfo.LastHeatBeatAt = time.Now()
	solodnInfo.SolodnHeartBeat = heartbeat
	err = p.FormatSolodnInfo(&solodnInfo)
	if err != nil {
		return err
	}

	p.solodnTable.Store(peerID, solodnInfo)

	return nil
}
