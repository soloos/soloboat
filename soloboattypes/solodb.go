package soloboattypes

import (
	"soloos/common/snet"
	"soloos/common/solodbapitypes"
	"time"
	"unsafe"
)

const (
	SolodbInfoStructSize = unsafe.Sizeof(SolodbInfo{})
)

type SolodbInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSolodbInfoJSON(solodbInfoJSON SolodbInfoJSON) SolodbInfo {
	var ret SolodbInfo
	ret.PeerID.SetStr(solodbInfoJSON.PeerID)
	ret.LastHeatBeatAt = time.Unix(solodbInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSolodbInfoJSON(solodbInfo SolodbInfo) SolodbInfoJSON {
	var ret SolodbInfoJSON
	ret.PeerID = string(solodbInfo.PeerID.Str())
	ret.LastHeatBeatAt = solodbInfo.LastHeatBeatAt.Unix()
	return ret
}

type SolodbInfo struct {
	snet.PeerID
	LastHeatBeatAt    time.Time
	LastHeatBeatAtStr string
	SrpcServerAddr    string
	WebServerAddr     string
	solodbapitypes.SolodbHeartBeat
}
