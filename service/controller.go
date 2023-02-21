package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/zwcway/fasthttp-upnp/event"
	"github.com/zwcway/fasthttp-upnp/scpd"
)

type httpHandler func(c *Controller, ctx *fasthttp.RequestCtx, uuid string) error

type Controller struct {
	SpecVersion scpd.SpecVersion
	ServiceName string
	PrefixPath  string

	event.Event

	Actions []*Action

	SCPDHandler    httpHandler
	ControlHandler httpHandler
	EventHandler   httpHandler

	scpdHttpPath    string
	controlHttpPath string
	eventHttpPath   string

	scpdXML []byte
}

func (c *Controller) Init(ctx context.Context) error {
	c.PrefixPath = strings.Trim(c.PrefixPath, "/")
	if c.PrefixPath != "" {
		c.PrefixPath = "/" + c.PrefixPath
	}
	if c.SpecVersion.Major == 0 {
		c.SpecVersion.Major = 1
	}

	c.scpdHttpPath = fmt.Sprintf("%s/%s%d.xml", c.PrefixPath, c.ServiceName, c.SpecVersion.Major)
	c.controlHttpPath = fmt.Sprintf("%s/%s/control", c.PrefixPath, c.ServiceName)
	c.eventHttpPath = fmt.Sprintf("%s/%s/event", c.PrefixPath, c.ServiceName)

	if c.SCPDHandler == nil {
		c.SCPDHandler = defaultSCPDHandler
	}
	if c.ControlHandler == nil {
		c.ControlHandler = defaultControlHandler
	}
	if c.EventHandler == nil {
		c.EventHandler = defaultEventHandler
	}
	if c.Actions == nil {
		return fmt.Errorf("actions can not empty")
	}

	err := c.init()

	return err
}

func (c *Controller) DeInit() {
}

func parseActionArgument(index int, a *Action, arv reflect.Value, isIn bool) (*scpd.Argument, *scpd.Variable) {
	arg := arv.Field(index)
	art := arv.Type().Field(index)
	tag := art.Tag

	var (
		aName string = art.Name
		vName string
		dir   string
		evt   string = "no"
		vdt   string
		def   string
	)

	if isIn {
		dir = "in"
	} else {
		dir = "out"
	}

	tags := strings.Split(tag.Get("soap"), ",")
	if len(tags) == 0 {
		vName = aName
	} else {
		vName = tags[0]
		if idx := strings.Index(vName, " "); idx > 0 {
			def = vName[idx+1:]
			vName = vName[:idx]
		}
		tags = tags[1:]
	}
	for _, t := range tags {
		switch strings.ToLower(t) {
		case "sendevent":
			evt = "yes"
		}
	}

	switch arg.Kind() {
	case reflect.Uint, reflect.Uint32:
		vdt = "ui4"
	case reflect.Uint8:
		vdt = "ui1"
	case reflect.Uint16:
		vdt = "ui2"
	case reflect.Int, reflect.Int32:
		vdt = "i4"
	case reflect.Int8:
		vdt = "i1"
	case reflect.Int16:
		vdt = "i2"
	case reflect.Bool:
		vdt = "boolean"
	case reflect.String:
		vdt = "string"
	default:
		// panic(fmt.Errorf("unsupported argument type '%s' for action '%s'", arg.Kind().String(), a.Name))
		return nil, nil
	}
	sa := &scpd.Argument{
		Name:            aName,
		Direction:       dir,
		RelatedStateVar: vName,
	}

	for _, v := range a.stateVariable {
		if v.Name == vName {
			return sa, v
		}
	}

	tav := strings.Split(tag.Get("allowed"), ",")
	tar := strings.Split(tag.Get("range"), ",")

	var av *[]string
	if len(tav) > 0 {
		av = &tav
	}

	var ar *scpd.AllowRange
	for i, r := range tar {
		si, err := strconv.ParseInt(r, 0, 32)
		if err != nil {
			continue
		}
		if ar == nil {
			ar = &scpd.AllowRange{}
		}
		if i == 0 {
			ar.Max = int(si)
		} else if i == 1 {
			ar.Min = int(si)
		} else if i == 2 {
			ar.Step = int(si)
		} else {
			break
		}
	}
	if ar != nil {
		if ar.Step == 0 {
			ar.Step = 1
		}
		if ar.Min != ar.Max && ((ar.Min-ar.Max > 0 && ar.Step > 0) || (ar.Max-ar.Min > 0 && ar.Step < 0)) {
			panic(fmt.Errorf("range invalid for field '%s' in struct '%s'", art.Name, arv.Type().Name()))
		}
	}
	sv := &scpd.Variable{
		Name:          vName,
		SendEvents:    evt,
		DataType:      vdt,
		Default:       def,
		AllowedValues: av,
		AllowedRange:  ar,
	}

	if isIn {
		a.inSoap[index] = soapArg{
			rv:   arg,
			name: aName,
			sv:   sv,
		}
	} else {
		a.outSoap[index] = soapArg{
			rv:   arg,
			name: aName,
			sv:   sv,
		}
	}

	a.stateVariable = append(a.stateVariable, sv)

	return sa, sv
}

func parseActionSCPD(a *Action) (sas []*scpd.Argument, svs []*scpd.Variable) {

	sas = make([]*scpd.Argument, 0)
	svs = make([]*scpd.Variable, 0)

	if reflect.ValueOf(a.ArgIn).Kind() != reflect.Pointer {
		panic(fmt.Errorf("ArgIn must be pointer of struct for action '%s'", a.Name))
	}
	argInReflect := reflect.ValueOf(&a.ArgIn).Elem().Elem().Elem()
	if !argInReflect.IsValid() || argInReflect.Kind() != reflect.Struct {
		panic(fmt.Errorf("ArgIn invalid for action '%s'", a.Name))
	}
	a.inSoap = make([]soapArg, argInReflect.NumField())

	for i := 0; i < argInReflect.NumField(); i++ {
		sa, sv := parseActionArgument(i, a, argInReflect, true)
		if sa == nil || sv == nil {
			continue
		}

		sas = append(sas, sa)
		svs = append(svs, sv)
	}

	if reflect.ValueOf(a.ArgOut).Kind() != reflect.Pointer {
		panic(fmt.Errorf("ArgOut must be pointer of struct for action '%s'", a.Name))
	}
	argOutReflect := reflect.ValueOf(&a.ArgOut).Elem().Elem().Elem()
	if !argOutReflect.IsValid() || argOutReflect.Kind() != reflect.Struct {
		panic(fmt.Errorf("ArgOut invalid for action '%s'", a.Name))
	}

	a.outSoap = make([]soapArg, argOutReflect.NumField())
	for i := 0; i < argOutReflect.NumField(); i++ {
		sa, sv := parseActionArgument(i, a, argOutReflect, false)
		if sa == nil || sv == nil {
			continue
		}

		sas = append(sas, sa)
		svs = append(svs, sv)
	}

	return
}

func (c *Controller) init() (err error) {
	var (
		actions   []scpd.Action
		variables []*scpd.Variable
	)
	for _, a := range c.Actions {
		if a.Name == "" {
			panic(fmt.Errorf("action name can not empty"))
		}

		args, sv := parseActionSCPD(a)

		variables = append(variables, sv...)

		actions = append(actions, scpd.Action{
			Name:      a.Name,
			Arguments: args,
		})
	}
	variables = sliceRemoveRepeatByLoop(variables)

	var s = scpd.SCPD{
		SpecVersion:       c.SpecVersion,
		ActionList:        actions,
		ServiceStateTable: variables,
	}
	c.scpdXML, err = xml.Marshal(s)
	return
}

func (c *Controller) ServiceURN(auth string) string {
	return fmt.Sprintf("%s:%d", c.ServiceId(auth), c.SpecVersion.Major)
}

func (c *Controller) ServiceId(auth string) string {
	return fmt.Sprintf("urn:%s:service:%s", auth, c.ServiceName)
}

func (c *Controller) SCPDHttpPath(uuid string) string    { return "/" + uuid + c.scpdHttpPath }
func (c *Controller) ControlHttpPath(uuid string) string { return "/" + uuid + c.controlHttpPath }
func (c *Controller) EventHttpPath(uuid string) string   { return "/" + uuid + c.eventHttpPath }

type ServiceController interface {
	Init(ctx context.Context) error
	Deinit()
}
