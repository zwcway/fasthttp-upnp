package connectionmanager1

import (
	"github.com/zwcway/fasthttp-upnp/service"
)

func ServiceController() *service.Controller {
	return &service.Controller{
		ServiceName: NAME,
		Actions: []*service.Action{
			ConnectionComplete(service.DefaultActionHandler),
			GetCurrentConnectionIDs(service.DefaultActionHandler),
			GetCurrentConnectionInfo(service.DefaultActionHandler),
			GetProtocolInfo(service.DefaultActionHandler),
			PrepareForConnection(service.DefaultActionHandler),
		},
	}
}
