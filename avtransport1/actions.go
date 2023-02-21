package avtransport1

import (
	"github.com/zwcway/fasthttp-upnp/service"
)

func GetCurrentTransportActions(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetCurrentTransportActions",
		Handler: handle,
		ArgIn:   &ArgInGetCurrentTransportActions{},
		ArgOut:  &ArgOutGetCurrentTransportActions{},
	}
}
func GetDeviceCapabilities(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetDeviceCapabilities",
		Handler: handle,
		ArgIn:   &ArgInGetDeviceCapabilities{},
		ArgOut:  &ArgOutGetDeviceCapabilities{},
	}
}
func GetMediaInfo(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetMediaInfo",
		Handler: handle,
		ArgIn:   &ArgInGetMediaInfo{},
		ArgOut:  &ArgOutGetMediaInfo{},
	}
}
func GetPositionInfo(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetPositionInfo",
		Handler: handle,
		ArgIn:   &ArgInGetPositionInfo{},
		ArgOut:  &ArgOutGetPositionInfo{},
	}
}
func GetTransportInfo(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetTransportInfo",
		Handler: handle,
		ArgIn:   &ArgInGetTransportInfo{},
		ArgOut:  &ArgOutGetTransportInfo{},
	}
}
func GetTransportSettings(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetTransportSettings",
		Handler: handle,
		ArgIn:   &ArgInGetTransportSettings{},
		ArgOut:  &ArgOutGetTransportSettings{},
	}
}
func Next(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "Next",
		Handler: handle,
		ArgIn:   &ArgInNext{},
		ArgOut:  &ArgOutNext{},
	}
}
func Pause(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "Pause",
		Handler: handle,
		ArgIn:   &ArgInPause{},
		ArgOut:  &ArgOutPause{},
	}
}
func Play(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "Play",
		Handler: handle,
		ArgIn:   &ArgInPlay{},
		ArgOut:  &ArgOutPlay{},
	}
}
func Previous(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "Previous",
		Handler: handle,
		ArgIn:   &ArgInPrevious{},
		ArgOut:  &ArgOutPrevious{},
	}
}
func Record(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "Record",
		Handler: handle,
		ArgIn:   &ArgInRecord{},
		ArgOut:  &ArgOutRecord{},
	}
}
func Seek(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "Seek",
		Handler: handle,
		ArgIn:   &ArgInSeek{},
		ArgOut:  &ArgOutSeek{},
	}
}
func SetAVTransportURI(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetAVTransportURI",
		Handler: handle,
		ArgIn:   &ArgInSetAVTransportURI{},
		ArgOut:  &ArgOutSetAVTransportURI{},
	}
}
func SetNextAVTransportURI(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetNextAVTransportURI",
		Handler: handle,
		ArgIn:   &ArgInSetNextAVTransportURI{},
		ArgOut:  &ArgOutSetNextAVTransportURI{},
	}
}
func SetPlayMode(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetPlayMode",
		Handler: handle,
		ArgIn:   &ArgInSetPlayMode{},
		ArgOut:  &ArgOutSetPlayMode{},
	}
}
func SetRecordQualityMode(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetRecordQualityMode",
		Handler: handle,
		ArgIn:   &ArgInSetRecordQualityMode{},
		ArgOut:  &ArgOutSetRecordQualityMode{},
	}
}
func Stop(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "Stop",
		Handler: handle,
		ArgIn:   &ArgInStop{},
		ArgOut:  &ArgOutStop{},
	}
}
