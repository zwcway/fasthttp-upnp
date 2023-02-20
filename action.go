package upnp

import (
	"reflect"

	"github.com/valyala/fasthttp"
	"github.com/zwcway/fasthttp-upnp/scpd"
)

type Argument struct {
	Name            string
	Direction       string
	RelatedStateVar *scpd.Variable
}

type Action struct {
	name         string
	Handler      func(ctx *fasthttp.RequestCtx)
	ArgIn        any
	ArgOut       any
	arguments    []Argument
	argInReflect reflect.Value
}

func (a *Action) Name() string { return a.name }

type ActionMap map[string]*Action
