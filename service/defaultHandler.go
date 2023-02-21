package service

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/zwcway/fasthttp-upnp/scpd"
	"github.com/zwcway/fasthttp-upnp/soap"
)

func ParseSOAPAction(ctx *fasthttp.RequestCtx) (*soap.SoapAction, error) {
	return soap.ParseSOAPAction(string(ctx.Request.Header.Peek("SOAPACTION")))
}

func defaultSCPDHandler(c *Controller, ctx *fasthttp.RequestCtx, uuid string) error {
	ctx.Response.Header.SetContentType(ResponseContentTypeXML)
	ctx.Write(c.scpdXML)
	return nil
}

func defaultControlHandler(c *Controller, ctx *fasthttp.RequestCtx, uuid string) error {
	soapAction, err := ParseSOAPAction(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return err
	}

	var action *Action
	for _, a := range c.Actions {
		if a.Name == soapAction.Action {
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
		return fmt.Errorf("the action '%s' handler is nil", action.Name)
	}

	action.Handler(action.ArgIn, action.ArgOut, ctx, uuid)

	var resp []byte
	resp, err = marshalActionResponse(action, c)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotImplemented)
		return err
	}

	ctx.SetContentType(ResponseContentTypeXML)
	ctx.Response.Header.Set("Ext", "")

	ctx.Write(resp)

	return nil
}

func defaultEventHandler(c *Controller, ctx *fasthttp.RequestCtx, uuid string) error {
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

func parseRequestArguments(a *Action, ctx *fasthttp.RequestCtx) (err error) {
	for _, arg := range a.inSoap {
		reqArg := string(ctx.Request.Header.Peek(arg.name))
		if reqArg == "" {
			reqArg = arg.sv.Default
		}
		al := arg.sv.AllowedValues
		ar := arg.sv.AllowedRange
		switch arg.rv.Kind() {
		case reflect.String:
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
				return soap.NewErrorf(fasthttp.StatusBadRequest, "string value %s=%s must in list ", arg.name, reqArg)
			}
			arg.rv.SetString(reqArg)
		case reflect.Bool:
			reqArg = strings.ToLower(reqArg)
			if reqArg == "true" {
				arg.rv.SetBool(true)
			} else if reqArg == "false" {
				arg.rv.SetBool(false)
			} else {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "bool value %s=%s must boolean ", arg.name, reqArg)
			}
		case reflect.Int32:
			var i uint64
			i, err = checkRequestInt(reqArg, al, ar, false, 32)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "int32 value %s=%s %s", arg.name, reqArg, err.Error())
			}
			arg.rv.SetInt(int64(i))
		case reflect.Uint32:
			var i uint64
			i, err = checkRequestInt(reqArg, al, ar, true, 32)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "uint32 value %s=%s %s", arg.name, reqArg, err.Error())
			}
			arg.rv.SetUint(i)
		case reflect.Int16:
			var i uint64
			i, err = checkRequestInt(reqArg, al, ar, true, 16)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "int16 value %s=%s %s", arg.name, reqArg, err.Error())
			}
			arg.rv.SetInt(int64(i))
		case reflect.Uint16:
			var i uint64
			i, err = checkRequestInt(reqArg, al, ar, true, 16)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "uint16 value %s=%s %s", arg.name, reqArg, err.Error())
			}
			arg.rv.SetUint(i)
		case reflect.Int8:
			var i uint64
			i, err = checkRequestInt(reqArg, al, ar, true, 8)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "int8 value %s=%s %s", arg.name, reqArg, err.Error())
			}
			arg.rv.SetInt(int64(i))
		case reflect.Uint8:
			var i uint64
			i, err = checkRequestInt(reqArg, al, ar, true, 8)
			if err != nil {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return soap.NewErrorf(fasthttp.StatusBadRequest, "uint8 value %s=%s %s", arg.name, reqArg, err.Error())
			}
			arg.rv.SetUint(i)
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

func marshalActionResponse(action *Action, c *Controller) (body []byte, err error) {
	body, err = xml.Marshal(action.ArgOut)
	if err != nil {
		return
	}

	resp := soap.EnvelopeResponse{
		XMLSpace:      ServiceNS(c.ServiceName, int(c.SpecVersion.Major)),
		EncodingStyle: soap.EncodingStyle,

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

func DefaultActionHandler(input any, output any, ctx *fasthttp.RequestCtx, uuid string) {

}
