package avtransport1

import (
	"github.com/zwcway/fasthttp-upnp/service"
)

func ServiceController() *service.Controller {
	return &service.Controller{
		ServiceName: NAME,
		Actions: []*service.Action{
			GetCurrentTransportActions(service.DefaultActionHandler),
			GetDeviceCapabilities(service.DefaultActionHandler),
			GetMediaInfo(service.DefaultActionHandler),
			GetPositionInfo(service.DefaultActionHandler),
			GetTransportInfo(service.DefaultActionHandler),
			GetTransportSettings(service.DefaultActionHandler),
			Next(service.DefaultActionHandler),
			Pause(service.DefaultActionHandler),
			Play(service.DefaultActionHandler),
			Previous(service.DefaultActionHandler),
			Record(service.DefaultActionHandler),
			Seek(service.DefaultActionHandler),
			SetAVTransportURI(service.DefaultActionHandler),
			SetNextAVTransportURI(service.DefaultActionHandler),
			SetPlayMode(service.DefaultActionHandler),
			SetRecordQualityMode(service.DefaultActionHandler),
			Stop(service.DefaultActionHandler),
		},
	}
}
