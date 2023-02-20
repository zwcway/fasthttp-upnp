package upnp

import (
	"context"
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/zwcway/fasthttp-upnp/scpd"
	"github.com/zwcway/fasthttp-upnp/soap"
)

type httpHandler func(c *Controller, ctx *fasthttp.RequestCtx) error

type Controller struct {
	SpecVersion scpd.SpecVersion
	ServiceName string
	PrefixPath  string

	Event

	Actions ActionMap

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

func (c *Controller) init() (err error) {
	var (
		actions   []scpd.Action
		variables []*scpd.Variable
	)
	for n, a := range c.Actions {
		a.name = n
		a.argInReflect = reflect.ValueOf(&a.ArgIn).Elem().Elem().Elem()
		args := []scpd.Argument{}
		for _, g := range a.arguments {
			args = append(args, scpd.Argument{
				Name:            g.Name,
				Direction:       g.Direction,
				RelatedStateVar: g.RelatedStateVar.Name,
			})
			variables = append(variables, g.RelatedStateVar)
		}
		actions = append(actions, scpd.Action{
			Name:      n,
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

func (c *Controller) SCPDHttpPath() string    { return c.scpdHttpPath }
func (c *Controller) ControlHttpPath() string { return c.controlHttpPath }
func (c *Controller) EventHttpPath() string   { return c.eventHttpPath }

type ServiceController interface {
	Init(ctx context.Context) error
	Deinit()
}

func ParseSOAPAction(ctx *fasthttp.RequestCtx) (*soap.SoapAction, error) {
	return soap.ParseSOAPAction(string(ctx.Request.Header.Peek("SOAPACTION")))
}

func defaultSCPDHandler(c *Controller, ctx *fasthttp.RequestCtx) error {
	ctx.Response.Header.SetContentType(ResponseContentTypeXML)
	ctx.Write(c.scpdXML)
	return nil
}

func defaultControlHandler(c *Controller, ctx *fasthttp.RequestCtx) error {
	soapAction, err := ParseSOAPAction(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return err
	}

	var action *Action
	for _, a := range c.Actions {
		if a.Name() == soapAction.Action {
			action = a
			break
		}
	}
	if action == nil {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		return fmt.Errorf("unknown action '%s'", soapAction.Action)
	}

	err = unmarshalActionRequest(action.ArgIn, ctx.Request.Body())
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return err
	}

	if action.Handler == nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return fmt.Errorf("the action '%s' handler is nil", action.name)
	}

	action.Handler(action.ArgIn, action.ArgOut, ctx)

	var resp []byte
	resp, err = marshalActionResponse(&action.ArgOut)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotImplemented)
		return err
	}

	ctx.SetContentType(ResponseContentTypeXML)
	ctx.Response.Header.Set("Ext", "")

	ctx.Write(resp)

	return nil
}

func defaultEventHandler(c *Controller, ctx *fasthttp.RequestCtx) error {
	method := string(ctx.Method())

	switch method {
	case "SUBSCRIBE":
		err := defaultSubscribeHandler(&c.Event, ctx)
		if err != nil {
			ResponseError(ctx, err)
		}
		return err
	case "UNSUBSCRIBE":

	default:
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
	}

	return nil
}

func ResponseError(ctx *fasthttp.RequestCtx, e *soap.Error) error {
	if e == nil {
		return nil
	}
	body, err := xml.Marshal(e)
	if err != nil {
		return err
	}
	ctx.SetBody(body)
	return nil
}

func parseRequestArguments(action *Action, ctx *fasthttp.RequestCtx) (err error) {

	for _, arg := range action.arguments {
		if arg.Direction != DirIn {
			continue
		}
		arv := action.argInReflect.FieldByName(arg.Name)

		if !arv.IsValid() {
			return soap.NewErrorf(fasthttp.StatusBadRequest, "argument '%s' not allowed", arg.Name)
		}

		reqArg := string(ctx.Request.Header.Peek(arg.Name))
		if reqArg == "" {
			reqArg = arg.RelatedStateVar.Default
		}
		al := arg.RelatedStateVar.AllowedValues
		ar := arg.RelatedStateVar.AllowedRange
		switch arg.RelatedStateVar.DataType {
		case DataTypeStr:
			allow := true
			if al != nil {
				allow = false
				for _, v := range *al {
					if v == reqArg {
						allow = true
						break
					}
				}
			}
			if !allow {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "string value %s=%s must in list ", arg.Name, reqArg)
			}
			arv.SetString(reqArg)
		case DataTypeBool:
			reqArg = strings.ToLower(reqArg)
			if reqArg == "true" {
				arv.SetBool(true)
			} else if reqArg == "false" {
				arv.SetBool(false)
			} else {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "bool value %s=%s must boolean ", arg.Name, reqArg)
			}
		case DataTypeInt32:
			var i uint64
			i, err = checkRequestInt(reqArg, al, ar, false, 32)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "int32 value %s=%s %s", arg.Name, reqArg, err.Error())
			}
			arv.SetInt(int64(i))
		case DataTypeUint32:
			var i uint64
			i, err = checkRequestInt(reqArg, al, ar, true, 32)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "uint32 value %s=%s %s", arg.Name, reqArg, err.Error())
			}
			arv.SetUint(i)
		case DataTypeInt16:
			var i uint64
			i, err = checkRequestInt(reqArg, al, ar, true, 16)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "int16 value %s=%s %s", arg.Name, reqArg, err.Error())
			}
			arv.SetInt(int64(i))
		case DataTypeUInt16:
			var i uint64
			i, err = checkRequestInt(reqArg, al, ar, true, 16)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "uint16 value %s=%s %s", arg.Name, reqArg, err.Error())
			}
			arv.SetUint(i)
		}
	}

	return nil
}

func unmarshalActionRequest(args any, body []byte) error {
	var env soap.Envelope
	err := xml.Unmarshal(body, &env)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(env.Body.Action, args)
	if err != nil {
		return err
	}
	return nil
}

func marshalActionResponse(args *any) (body []byte, err error) {
	body, err = xml.Marshal(args)
	if err != nil {
		return
	}

	resp := soap.EnvelopeResponse{
		Body: soap.EnvelopeBody{
			Action: body,
		},
	}

	body, err = xml.Marshal(resp)

	return
}

func sliceRemoveRepeatByLoop(slc []*scpd.Variable) []*scpd.Variable {
	result := []*scpd.Variable{}
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}

func checkRequestInt(reqArg string, al *[]string, ar *scpd.AllowRange, unsigned bool, bits int) (uint64, error) {
	if reqArg == "" {
		reqArg = "0"
	}
	if al != nil {
		for _, v := range *al {
			if v == reqArg {
				if unsigned {
					return strconv.ParseUint(reqArg, 0, bits)
				} else {
					i, err := strconv.ParseUint(reqArg, 0, bits)
					return uint64(i), err
				}
			}
		}
	} else {
		var (
			intv uint64
			err  error
		)
		if unsigned {
			var i uint64
			i, err = strconv.ParseUint(reqArg, 0, bits)
			intv = i
		} else {
			var i int64
			i, err = strconv.ParseInt(reqArg, 0, bits)
			intv = uint64(i)
		}
		if err != nil {
			return 0, err
		}
		if ar == nil {
			return intv, nil
		}
		step := ar.Step
		if step < 1 {
			step = 1
		}

		for i := ar.Min; i <= ar.Max; i += ar.Step {
			if i == int(intv) {
				return intv, nil
			}
		}
	}

	return 0, fmt.Errorf("invalid value %s", reqArg)
}
