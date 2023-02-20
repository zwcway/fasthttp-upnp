package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
	upnp "github.com/zwcway/fasthttp-upnp"
)

var (
	playUrl string
	playDur int = 0
	ticker  *time.Ticker
)

func playing(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			playDur++
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	uuids := map[string]string{}
	for _, n := range []string{"我是DLNA", "备用DLNA"} {
		uuids[upnp.MakeUUID(n)] = n
	}
	upnpSrv, err := upnp.NewDeviceServers(ctx, uuids)
	if err != nil {
		return
	}
	upnpSrv.DeviceType = upnp.DeviceType_MediaRenderer
	upnpSrv.ServerName = "UPnPServer/1.0"
	upnpSrv.RootDescNamespaces = map[string]string{
		"xmlns:dlna": "urn:schemas-dlna-org:device-1-0",
	}
	upnpSrv.ServiceList = []*upnp.Controller{
		{
			ServiceName: upnp.ServiceName_AVTransport,
			Actions:     upnp.AVTransportV1,
		},
	}
	// upnpSrv.ListenPort = 1900
	upnpSrv.AllowIps = []*net.IPNet{
		{IP: net.ParseIP("10.2.2.24"), Mask: net.IPv4Mask(0xFF, 0xFF, 0xFF, 0xFF)},
		{IP: net.ParseIP("10.2.2.113"), Mask: net.IPv4Mask(0xFF, 0xFF, 0xFF, 0xFF)},
	}
	upnpSrv.ErrorChan = make(chan error, 10)
	upnpSrv.BeforeRequestHandle = func(ctx *fasthttp.RequestCtx) bool {
		fmt.Println("from", ctx.RemoteAddr(), ctx.Request.String())
		return true
	}

	upnp.AVT_SetAVTransportURI.Handler = func(input, output any, ctx *fasthttp.RequestCtx, uuid string) {
		in := input.(*upnp.AVTArgIn_SetAVTransportURI)

		playUrl = in.CurrentURI

		fmt.Println(in.InstanceID, in.CurrentURI)
	}
	upnp.AVT_GetPositionInfo.Handler = func(input, output any, ctx *fasthttp.RequestCtx, uuid string) {
		out := output.(*upnp.AVTArgOut_GetPositionInfo)

		playDur++

		out.TrackDuration = fmt.Sprintf("00:00:%2d", playDur)

		out.TrackURI = playUrl
	}
	upnp.AVT_Play.Handler = func(input, output any, ctx *fasthttp.RequestCtx, uuid string) {
		playDur = 0
		ticker.Reset(time.Second)
	}
	upnp.AVT_Pause.Handler = func(input, output any, ctx *fasthttp.RequestCtx, uuid string) {
		ticker.Stop()
	}
	upnp.AVT_Stop.Handler = func(input, output any, ctx *fasthttp.RequestCtx, uuid string) {
		playDur = 0
		ticker.Stop()
	}

	err = upnpSrv.Init()
	if err != nil {
		panic(err)
	}
	go upnpSrv.Serve()

	ticker = time.NewTicker(time.Second)
	go playing(ctx)
	ticker.Stop()

	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGABRT)

	for {
		select {
		case <-signalChan:
			close(signalChan)
			cancel()
			upnpSrv.Close()
			return
		case err := <-upnpSrv.ErrorChan:
			fmt.Println(err)
		}
	}

}
