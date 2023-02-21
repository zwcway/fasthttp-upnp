package renderingcontrol1

import (
	"github.com/zwcway/fasthttp-upnp/service"
)

func GetBlueVideoBlackLevel(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetBlueVideoBlackLevel",
		Handler: handle,
		ArgIn:   &ArgInGetBlueVideoBlackLevel{},
		ArgOut:  &ArgOutGetBlueVideoBlackLevel{},
	}
}
func GetBlueVideoGain(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetBlueVideoGain",
		Handler: handle,
		ArgIn:   &ArgInGetBlueVideoGain{},
		ArgOut:  &ArgOutGetBlueVideoGain{},
	}
}
func GetBrightness(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetBrightness",
		Handler: handle,
		ArgIn:   &ArgInGetBrightness{},
		ArgOut:  &ArgOutGetBrightness{},
	}
}
func GetColorTemperature(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetColorTemperature",
		Handler: handle,
		ArgIn:   &ArgInGetColorTemperature{},
		ArgOut:  &ArgOutGetColorTemperature{},
	}
}
func GetContrast(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetContrast",
		Handler: handle,
		ArgIn:   &ArgInGetContrast{},
		ArgOut:  &ArgOutGetContrast{},
	}
}
func GetGreenVideoBlackLevel(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetGreenVideoBlackLevel",
		Handler: handle,
		ArgIn:   &ArgInGetGreenVideoBlackLevel{},
		ArgOut:  &ArgOutGetGreenVideoBlackLevel{},
	}
}
func GetGreenVideoGain(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetGreenVideoGain",
		Handler: handle,
		ArgIn:   &ArgInGetGreenVideoGain{},
		ArgOut:  &ArgOutGetGreenVideoGain{},
	}
}
func GetHorizontalKeystone(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetHorizontalKeystone",
		Handler: handle,
		ArgIn:   &ArgInGetHorizontalKeystone{},
		ArgOut:  &ArgOutGetHorizontalKeystone{},
	}
}
func GetLoudness(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetLoudness",
		Handler: handle,
		ArgIn:   &ArgInGetLoudness{},
		ArgOut:  &ArgOutGetLoudness{},
	}
}
func GetMute(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetMute",
		Handler: handle,
		ArgIn:   &ArgInGetMute{},
		ArgOut:  &ArgOutGetMute{},
	}
}
func GetRedVideoBlackLevel(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetRedVideoBlackLevel",
		Handler: handle,
		ArgIn:   &ArgInGetRedVideoBlackLevel{},
		ArgOut:  &ArgOutGetRedVideoBlackLevel{},
	}
}
func GetRedVideoGain(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetRedVideoGain",
		Handler: handle,
		ArgIn:   &ArgInGetRedVideoGain{},
		ArgOut:  &ArgOutGetRedVideoGain{},
	}
}
func GetSharpness(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetSharpness",
		Handler: handle,
		ArgIn:   &ArgInGetSharpness{},
		ArgOut:  &ArgOutGetSharpness{},
	}
}
func GetVerticalKeystone(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetVerticalKeystone",
		Handler: handle,
		ArgIn:   &ArgInGetVerticalKeystone{},
		ArgOut:  &ArgOutGetVerticalKeystone{},
	}
}
func GetVolume(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetVolume",
		Handler: handle,
		ArgIn:   &ArgInGetVolume{},
		ArgOut:  &ArgOutGetVolume{},
	}
}
func GetVolumeDB(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetVolumeDB",
		Handler: handle,
		ArgIn:   &ArgInGetVolumeDB{},
		ArgOut:  &ArgOutGetVolumeDB{},
	}
}
func GetVolumeDBRange(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "GetVolumeDBRange",
		Handler: handle,
		ArgIn:   &ArgInGetVolumeDBRange{},
		ArgOut:  &ArgOutGetVolumeDBRange{},
	}
}
func ListPresets(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "ListPresets",
		Handler: handle,
		ArgIn:   &ArgInListPresets{},
		ArgOut:  &ArgOutListPresets{},
	}
}
func SelectPreset(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SelectPreset",
		Handler: handle,
		ArgIn:   &ArgInSelectPreset{},
		ArgOut:  &ArgOutSelectPreset{},
	}
}
func SetBlueVideoBlackLevel(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetBlueVideoBlackLevel",
		Handler: handle,
		ArgIn:   &ArgInSetBlueVideoBlackLevel{},
		ArgOut:  &ArgOutSetBlueVideoBlackLevel{},
	}
}
func SetBlueVideoGain(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetBlueVideoGain",
		Handler: handle,
		ArgIn:   &ArgInSetBlueVideoGain{},
		ArgOut:  &ArgOutSetBlueVideoGain{},
	}
}
func SetBrightness(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetBrightness",
		Handler: handle,
		ArgIn:   &ArgInSetBrightness{},
		ArgOut:  &ArgOutSetBrightness{},
	}
}
func SetColorTemperature(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetColorTemperature",
		Handler: handle,
		ArgIn:   &ArgInSetColorTemperature{},
		ArgOut:  &ArgOutSetColorTemperature{},
	}
}
func SetContrast(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetContrast",
		Handler: handle,
		ArgIn:   &ArgInSetContrast{},
		ArgOut:  &ArgOutSetContrast{},
	}
}
func SetGreenVideoBlackLevel(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetGreenVideoBlackLevel",
		Handler: handle,
		ArgIn:   &ArgInSetGreenVideoBlackLevel{},
		ArgOut:  &ArgOutSetGreenVideoBlackLevel{},
	}
}
func SetGreenVideoGain(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetGreenVideoGain",
		Handler: handle,
		ArgIn:   &ArgInSetGreenVideoGain{},
		ArgOut:  &ArgOutSetGreenVideoGain{},
	}
}
func SetHorizontalKeystone(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetHorizontalKeystone",
		Handler: handle,
		ArgIn:   &ArgInSetHorizontalKeystone{},
		ArgOut:  &ArgOutSetHorizontalKeystone{},
	}
}
func SetLoudness(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetLoudness",
		Handler: handle,
		ArgIn:   &ArgInSetLoudness{},
		ArgOut:  &ArgOutSetLoudness{},
	}
}
func SetMute(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetMute",
		Handler: handle,
		ArgIn:   &ArgInSetMute{},
		ArgOut:  &ArgOutSetMute{},
	}
}
func SetRedVideoBlackLevel(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetRedVideoBlackLevel",
		Handler: handle,
		ArgIn:   &ArgInSetRedVideoBlackLevel{},
		ArgOut:  &ArgOutSetRedVideoBlackLevel{},
	}
}
func SetRedVideoGain(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetRedVideoGain",
		Handler: handle,
		ArgIn:   &ArgInSetRedVideoGain{},
		ArgOut:  &ArgOutSetRedVideoGain{},
	}
}
func SetSharpness(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetSharpness",
		Handler: handle,
		ArgIn:   &ArgInSetSharpness{},
		ArgOut:  &ArgOutSetSharpness{},
	}
}
func SetVerticalKeystone(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetVerticalKeystone",
		Handler: handle,
		ArgIn:   &ArgInSetVerticalKeystone{},
		ArgOut:  &ArgOutSetVerticalKeystone{},
	}
}
func SetVolume(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetVolume",
		Handler: handle,
		ArgIn:   &ArgInSetVolume{},
		ArgOut:  &ArgOutSetVolume{},
	}
}
func SetVolumeDB(handle service.ActionHandler) *service.Action {
	return &service.Action{
		Name:    "SetVolumeDB",
		Handler: handle,
		ArgIn:   &ArgInSetVolumeDB{},
		ArgOut:  &ArgOutSetVolumeDB{},
	}
}
