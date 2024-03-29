package solodb

import (
	"soloos/common/snet"
	"soloos/common/solodbapitypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

func (p *SolodbDriver) SolodbHeartBeat(heartbeat solodbapitypes.SolodbHeartBeat) error {
	var err error
	var peerID = snet.StrToPeerID(heartbeat.SrpcPeerID)
	var iptr, exists = p.solodbTable.Load(peerID)
	var solodbInfo = soloboattypes.SolodbInfo{PeerID: peerID}
	if exists {
		solodbInfo = iptr.(soloboattypes.SolodbInfo)
	}

	solodbInfo.LastHeatBeatAt = time.Now()
	solodbInfo.SolodbHeartBeat = heartbeat
	err = p.FormatSolodbInfo(&solodbInfo)
	if err != nil {
		return err
	}

	p.solodbTable.Store(peerID, solodbInfo)

	return nil
}
