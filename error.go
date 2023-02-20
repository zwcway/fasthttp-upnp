package upnp

import (
	"fmt"

	"github.com/zwcway/fasthttp-upnp/ssdp"
)

func IsIPDenyError(err error) bool {
	_, ok := err.(*ssdp.IPDenyError)
	return ok
}

type UPnPError struct {
	Err error
}

func (e *UPnPError) Error() string {
	return fmt.Sprintf("UPnP: %s", e.Err.Error())
}
