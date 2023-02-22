package service

import (
	"fmt"
	"reflect"

	"github.com/valyala/fasthttp"
	"github.com/zwcway/fasthttp-upnp/scpd"
	"github.com/zwcway/fasthttp-upnp/soap"
)

const (
	DeviceType_MediaServer   = "MediaServer"
	DeviceType_MediaRenderer = "MediaRenderer"

	ServiceName_AVTransport       = "AVTransport"
	ServiceName_ConnectionManager = "ConnectionManager"
	ServiceName_RenderingControl  = "RenderingControl"
)

const ResponseContentTypeXML = `text/xml; charset="utf-8"`

type ActionHandler func(input any, output any, ctx *fasthttp.RequestCtx, uuid string) error

type Argument struct {
	Name            string
	Direction       string
	RelatedStateVar *scpd.Variable
}

type soapArg struct {
	name string
	rv   reflect.Value
	sv   *scpd.Variable
}
type Action struct {
	Handler ActionHandler
	ArgIn   any
	ArgOut  any

	Name          string
	stateVariable []*scpd.Variable
	inSoap        []soapArg
	outSoap       []soapArg
}

func (a *Action) SetStateVariables(sv []*scpd.Variable) *Action {
	a.stateVariable = make([]*scpd.Variable, len(sv))
	for i := range sv {
		a.stateVariable[i] = sv[i]
	}

	return a
}

func ServiceNS(service string, ver int) string {
	return fmt.Sprintf("urn:%s:service:%s:%d", soap.AuthName, service, ver)
}


type ActionArgumentIn interface {
	
}

type ActionArgumentOut interface {
	
}
