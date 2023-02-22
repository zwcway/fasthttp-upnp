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
	"github.com/zwcway/fasthttp-upnp/avtransport1"
	"github.com/zwcway/fasthttp-upnp/service"
	"github.com/zwcway/fasthttp-upnp/ssdp"
	"github.com/zwcway/fasthttp-upnp/utils"
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
		uuids[utils.MakeUUID(n)] = n
	}
	upnpSrv, err := upnp.NewDeviceServers(ctx, uuids)
	if err != nil {
		return
	}
	upnpSrv.DeviceType = service.DeviceType_MediaRenderer
	upnpSrv.ServerName = "UPnPServer/1.0"
	upnpSrv.RootDescNamespaces = map[string]string{
		"xmlns:dlna": "urn:schemas-dlna-org:device-1-0",
	}
	upnpSrv.ServiceList = []*service.Controller{
		{
			ServiceName: avtransport1.NAME,
			Actions: []*service.Action{
				avtransport1.GetPositionInfo(func(input, output any, ctx *fasthttp.RequestCtx, uuid string) error {
					out := output.(*avtransport1.ArgOutGetPositionInfo)

					out.TrackDuration = fmt.Sprintf("00:00:%2d", playDur)
					out.TrackURI = playUrl
					return nil
				}),
				avtransport1.Pause(func(input, output any, ctx *fasthttp.RequestCtx, uuid string) error {
					ticker.Stop()
					return nil
				}),
				avtransport1.Play(func(input, output any, ctx *fasthttp.RequestCtx, uuid string) error {
					playDur = 0
					ticker.Reset(time.Second)
					return nil
				}),
				avtransport1.SetAVTransportURI(func(input, output any, ctx *fasthttp.RequestCtx, uuid string) error {
					in := input.(*avtransport1.ArgInSetAVTransportURI)

					playUrl = in.CurrentURI
					fmt.Println(in.InstanceID, in.CurrentURI)
					return nil
				}),
				avtransport1.Stop(func(input, output any, ctx *fasthttp.RequestCtx, uuid string) error {
					playDur = 0
					ticker.Stop()
					return nil
				}),
			},
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
			if ssdp.IsIPDenyError(err) {
				break
			}
			fmt.Println(err)
		case err := <-upnpSrv.InfoChan:
			fmt.Println(err)
		}
	}

}
