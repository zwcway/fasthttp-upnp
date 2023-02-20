package upnp

import (
	"time"

	"github.com/valyala/fasthttp"
	"github.com/zwcway/fasthttp-upnp/soap"
)

func defaultSubscribeHandler(e *Event, ctx *fasthttp.RequestCtx) (ret *soap.Error) {
	var (
		sid string
	)

	if nt := string(ctx.Request.Header.Peek("NT")); nt != NTEvent {
		return soap.NewErrorf(fasthttp.StatusBadRequest, "request invalid: NT=%s", nt)
	}

	reqSid := string(ctx.Request.Header.Peek("SID"))
	if reqSid == "" {
		// 第一次订阅
		sid = MakeUUID(time.Now().String())
	} else {
		// 续订
		sid = string(reqSid)
		if !e.HasSubscribed(sid) {
			return soap.NewErrorf(fasthttp.StatusBadRequest, "request invalid: sid=%s", sid)
		}
	}

	ctx.Request.Header.Set("SID", sid)

	defer func() {
		if ret == nil {
			ctx.Request.Header.Set("TIMEOUT", e.MakeTimeout(sid))
		} else {
			ResponseError(ctx, ret)
		}
	}()

	urls, err := ParseCallback(string(ctx.Request.Header.Peek("CALLBACK")))
	if err != nil {
		return soap.NewError(fasthttp.StatusBadRequest, err)
	}
	t, err := ParseTimeout(string(ctx.Request.Header.Peek("TIMEOUT")))
	if err != nil {
		return soap.NewError(fasthttp.StatusBadRequest, err)
	}

	e.Subscribe(sid, urls, t)

	return nil
}
