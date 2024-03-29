package soloboattypes

import (
	"soloos/common/snet"
	"soloos/common/solofsapitypes"
	"time"
	"unsafe"
)

const (
	SolodnInfoStructSize = unsafe.Sizeof(SolodnInfo{})
)

type SolodnInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSolodnInfoJSON(solodnInfoJSON SolodnInfoJSON) SolodnInfo {
	var ret SolodnInfo
	ret.PeerID.SetStr(solodnInfoJSON.PeerID)
	ret.LastHeatBeatAt = time.Unix(solodnInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSolodnInfoJSON(solodnInfo SolodnInfo) SolodnInfoJSON {
	var ret SolodnInfoJSON
	ret.PeerID = string(solodnInfo.PeerID.Str())
	ret.LastHeatBeatAt = solodnInfo.LastHeatBeatAt.Unix()
	return ret
}

type SolodnInfo struct {
	snet.PeerID
	LastHeatBeatAt    time.Time
	LastHeatBeatAtStr string
	SrpcServerAddr    string
	WebServerAddr     string
	solofsapitypes.SolodnHeartBeat
}
