package renderingcontrol1

import (
	"github.com/zwcway/fasthttp-upnp/service"
)

func ServiceController() *service.Controller {
	return &service.Controller{
		ServiceName: NAME,
		Actions: []*service.Action{
			GetBlueVideoBlackLevel(service.DefaultActionHandler),
			GetBlueVideoGain(service.DefaultActionHandler),
			GetBrightness(service.DefaultActionHandler),
			GetColorTemperature(service.DefaultActionHandler),
			GetContrast(service.DefaultActionHandler),
			GetGreenVideoBlackLevel(service.DefaultActionHandler),
			GetGreenVideoGain(service.DefaultActionHandler),
			GetHorizontalKeystone(service.DefaultActionHandler),
			GetLoudness(service.DefaultActionHandler),
			GetMute(service.DefaultActionHandler),
			GetRedVideoBlackLevel(service.DefaultActionHandler),
			GetRedVideoGain(service.DefaultActionHandler),
			GetSharpness(service.DefaultActionHandler),
			GetVerticalKeystone(service.DefaultActionHandler),
			GetVolume(service.DefaultActionHandler),
			GetVolumeDB(service.DefaultActionHandler),
			GetVolumeDBRange(service.DefaultActionHandler),
			ListPresets(service.DefaultActionHandler),
			SelectPreset(service.DefaultActionHandler),
			SetBlueVideoBlackLevel(service.DefaultActionHandler),
			SetBlueVideoGain(service.DefaultActionHandler),
			SetBrightness(service.DefaultActionHandler),
			SetColorTemperature(service.DefaultActionHandler),
			SetContrast(service.DefaultActionHandler),
			SetGreenVideoBlackLevel(service.DefaultActionHandler),
			SetGreenVideoGain(service.DefaultActionHandler),
			SetHorizontalKeystone(service.DefaultActionHandler),
			SetLoudness(service.DefaultActionHandler),
			SetMute(service.DefaultActionHandler),
			SetRedVideoBlackLevel(service.DefaultActionHandler),
			SetRedVideoGain(service.DefaultActionHandler),
			SetSharpness(service.DefaultActionHandler),
			SetVerticalKeystone(service.DefaultActionHandler),
			SetVolume(service.DefaultActionHandler),
			SetVolumeDB(service.DefaultActionHandler),
		},
	}
}
