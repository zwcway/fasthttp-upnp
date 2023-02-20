package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/valyala/fasthttp"
	upnp "github.com/zwcway/fasthttp-upnp"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	upnpSrv, err := upnp.NewDeviceServer(ctx, "dlna server")
	if err != nil {
		return
	}
	upnpSrv.DeviceType = upnp.DeviceType_MediaRenderer
	upnpSrv.ServerName = "UPnPServer/1.0"
	upnpSrv.RootDescNamespaces = map[string]string{
		"xmlns:dlna": "urn:schemas-dlna-org:device-1-0",
	}
	upnpSrv.ServiceList = []*upnp.Controller{
		&upnp.Controller{
			ServiceName: upnp.ServiceName_AVTransport,
			Actions:     upnp.AVTransportV1,
		},
	}
	// upnpSrv.ListenPort = 1900
	upnpSrv.AllowIps = []*net.IPNet{
		&net.IPNet{IP: net.ParseIP("10.2.2.24"), Mask: net.IPv4Mask(0xFF, 0xFF, 0xFF, 0xFF)},
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
