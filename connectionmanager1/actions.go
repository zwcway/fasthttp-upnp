package connectionmanager1

import (
	"github.com/zwcway/fasthttp-upnp/service"
)

func ConnectionComplete(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "ConnectionComplete",
		Handler: handle,
		ArgIn:   &ArgInConnectionComplete{},
		ArgOut:  &ArgOutConnectionComplete{},
	}
}
func GetCurrentConnectionIDs(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetCurrentConnectionIDs",
		Handler: handle,
		ArgIn:   &ArgInGetCurrentConnectionIDs{},
		ArgOut:  &ArgOutGetCurrentConnectionIDs{},
	}
}
func GetCurrentConnectionInfo(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetCurrentConnectionInfo",
		Handler: handle,
		ArgIn:   &ArgInGetCurrentConnectionInfo{},
		ArgOut:  &ArgOutGetCurrentConnectionInfo{},
	}
}
func GetProtocolInfo(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetProtocolInfo",
		Handler: handle,
		ArgIn:   &ArgInGetProtocolInfo{},
		ArgOut:  &ArgOutGetProtocolInfo{},
	}
}
func PrepareForConnection(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "PrepareForConnection",
		Handler: handle,
		ArgIn:   &ArgInPrepareForConnection{},
		ArgOut:  &ArgOutPrepareForConnection{},
	}
}
