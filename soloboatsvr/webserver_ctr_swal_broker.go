package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/sdbone/offheap"
	"soloos/soloboat/soloboattypes"
)

func (p *WebServer) prepareCtrSWALBroker(ir *iron.Request) bool {
	var module = ViewModule{
		Name: "SWALBroker",
		CH:   "SWALBroker",
		URL:  "/SWAL/Broker",
	}
	ir.ViewData["Module"] = module
	return true
}

func (p *WebServer) ctrSWALBroker(ir *iron.Request) {
	var ret []soloboattypes.SWALBrokerInfo
	p.soloBoatSvr.swalBrokerDriver.swalBrokerTable.ListObject(func(uObj offheap.LKVTableObjectUPtrWithBytes64) bool {
		var uptr = soloboattypes.SWALBrokerInfoUintptr(uObj)
		var obj = *uptr.Ptr()
		ret = append(ret, obj)
		return true
	})
	ir.ViewData["BrokerArr"] = ret
	ir.Render("/SWAL/Broker/Index")
}
