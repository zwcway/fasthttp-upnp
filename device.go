package upnp

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"mime"
	"net"
	"net/url"
	"strings"

	"github.com/valyala/fasthttp"
	"github.com/zwcway/fasthttp-upnp/scpd"
	"github.com/zwcway/fasthttp-upnp/ssdp"
)

type DeviceServer struct {
	AuthName     string
	FriendlyName string

	DeviceType string

	SpecVersion        scpd.SpecVersion
	Manufacturer       string
	ServerName         string
	RootDescNamespaces map[string]string
	RootDescUrl        string

	ServiceList []*Controller

	BeforeRequestHandle func(ctx *fasthttp.RequestCtx) bool
	AfterRequestHandle  func(ctx *fasthttp.RequestCtx) bool

	ListenInterface *net.Interface
	ListenPort      uint16

	AllowIps []*net.IPNet
	DenyIps  []*net.IPNet

	ErrorChan chan error

	uuid         string
	ssdpList     []ssdp.SSDPServer
	ctx          context.Context
	conn         net.Listener
	listenedAddr *net.TCPAddr
	rootDescXML  []byte
	ifaces       []*net.Interface
}

func (s *DeviceServer) makeXMLStart() xml.StartElement {
	var attrs []xml.Attr

	for s, v := range s.RootDescNamespaces {
		attrs = append(attrs, xml.Attr{Name: xml.Name{Space: "", Local: s}, Value: v})
	}
	return xml.StartElement{
		Name: xml.Name{Space: s.DeviceNamespace(), Local: "root"},
		Attr: attrs,
	}
}

func (s *DeviceServer) Init() (err error) {
	if s.ctx == nil {
		return fmt.Errorf("context can not nil")
	}
	if s.FriendlyName == "" {
		return fmt.Errorf("FriendlyName con not empty")
	}
	if s.AuthName == "" {
		s.AuthName = AuthName
	}
	if s.SpecVersion.Major == 0 {
		s.SpecVersion.Major = 1
	}

	if s.RootDescUrl == "" {
		s.RootDescUrl = "/rootDesc.xml"
	} else {
		s.RootDescUrl = "/" + strings.Trim(s.RootDescUrl, "/")
	}
	if s.ListenInterface != nil && s.ListenPort == 0 {
		return fmt.Errorf("must specify a listen port")
	}

	for _, ss := range s.ServiceList {
		ss.Init(s.ctx)
	}

	address, err := getListenAddress(s.ListenInterface, s.ListenPort)
	if err != nil {
		return err
	}

	s.conn, err = net.Listen("tcp", address)
	if err != nil {
		return
	}

	s.ifaces = make([]*net.Interface, 0)
	if s.ListenInterface != nil {
		s.ifaces = append(s.ifaces, s.ListenInterface)
	} else {
		ifs, err := net.Interfaces()
		if err != nil {
			return err
		}

		for i := range ifs {
			ifi := ifs[i]
			s.ifaces = append(s.ifaces, &ifi)
		}
	}

	s.listenedAddr = s.conn.Addr().(*net.TCPAddr)
	s.ErrorChan = make(chan error, 10)

	bufio := &bytes.Buffer{}
	err = xml.NewEncoder(bufio).EncodeElement(s.makeDevice(), s.makeXMLStart())
	s.rootDescXML = bufio.Bytes()
	return
}

func (s *DeviceServer) makeServices() (srvs []scpd.Service) {
	for i := range s.ServiceList {
		c := s.ServiceList[i]
		srv := scpd.Service{
			ServiceType: c.ServiceURN(s.AuthName),
			ServiceId:   c.ServiceId(s.AuthName),
			SCPDURL:     c.SCPDHttpPath(),
			ControlURL:  c.ControlHttpPath(),
			EventSubURL: c.EventHttpPath(),
		}

		srvs = append(srvs, srv)
	}

	return
}

func (s *DeviceServer) DeviceURN() string {
	return fmt.Sprintf("urn:%s:device:%s:%d", s.AuthName, s.DeviceType, s.SpecVersion.Major)
}
func (s *DeviceServer) DeviceNamespace() string {
	return fmt.Sprintf("urn:%s:device-%d-%d", s.AuthName, s.SpecVersion.Major, s.SpecVersion.Minor)
}

func (s *DeviceServer) makeDevice() scpd.DeviceDesc {
	return scpd.DeviceDesc{
		SpecVersion: s.SpecVersion,
		Device: scpd.Device{
			DeviceType:   s.DeviceURN(),
			FriendlyName: s.FriendlyName,
			Manufacturer: s.Manufacturer,
			ModelName:    s.ServerName,
			UDN:          s.uuid,
			ServiceList:  s.makeServices(),
		},
	}
}

func (s *DeviceServer) Connection() net.Listener {
	return s.conn
}

func NewDeviceServer(ctx context.Context, name string) (s *DeviceServer, err error) {
	s = &DeviceServer{
		ctx:  ctx,
		uuid: MakeUUID(name),

		FriendlyName: name,
	}

	err = s.Init()

	return
}

func (s *DeviceServer) Close() {
	if s.conn == nil {
		return
	}
	for _, ss := range s.ssdpList {
		ss.Close()
	}
	for _, ss := range s.ServiceList {
		ss.DeInit()
	}
	close(s.ErrorChan)
	s.conn.Close()
	s.conn = nil
}

func (s *DeviceServer) Serve() {
	if s.conn == nil {
		return
	}

	s.startSSDP()

	server := fasthttp.Server{Handler: s.httpHandler}
	server.Serve(s.conn)
}

func (s *DeviceServer) httpHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetServer(s.ServerName)

	remoteIp := ctx.RemoteIP()
	for _, ipnet := range s.DenyIps {
		if ipnet.Contains(remoteIp) {
			ctx.SetStatusCode(fasthttp.StatusForbidden)
			s.notifyError(ssdp.NewIPDenyError(remoteIp))
			return
		}
	}
	for _, ipnet := range s.AllowIps {
		if !ipnet.Contains(remoteIp) {
			ctx.SetStatusCode(fasthttp.StatusForbidden)
			s.notifyError(ssdp.NewIPNotAllowError(remoteIp))
			return
		}
	}

	if s.BeforeRequestHandle != nil && !s.BeforeRequestHandle(ctx) {
		return
	}
	if s.AfterRequestHandle != nil {
		defer s.AfterRequestHandle(ctx)
	}

	uri := string(ctx.Path())

	for i := range s.ServiceList {
		c := s.ServiceList[i]
		switch uri {
		case c.SCPDHttpPath():
			if !ctx.IsGet() {
				ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
				return
			}
			if c.SCPDHandler != nil {
				c.SCPDHandler(c, ctx)
			} else {
				ctx.SetStatusCode(fasthttp.StatusNotFound)
			}
			return
		case c.ControlHttpPath():
			if !ctx.IsPost() {
				ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
				return
			}
			if !checkRequestIsXML(ctx) {
				ctx.SetStatusCode(fasthttp.StatusBadRequest)
				return
			}
			if c.ControlHandler != nil {
				c.ControlHandler(c, ctx)
			} else {
				ctx.SetStatusCode(fasthttp.StatusNotFound)
			}
			return
		case c.EventHttpPath():
			if c.EventHandler != nil {
				c.EventHandler(c, ctx)
			} else {
				ctx.SetStatusCode(fasthttp.StatusNotFound)
			}
			return
		}
	}

	if uri == s.RootDescUrl {
		ctx.Response.Header.SetContentType(`text/xml; charset="utf-8"`)
		ctx.Write(s.rootDescXML)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusNotFound)
}

func checkRequestIsXML(ctx *fasthttp.RequestCtx) bool {
	ct := string(ctx.Request.Header.ContentType())
	t, params, err := mime.ParseMediaType(ct)
	if err != nil {
		return false
	}
	if t != "text/xml" {
		return false
	}
	charset, ok := params["charset"]
	if !ok {
		return false
	}
	charset = strings.ToLower(charset)
	if charset != "utf-8" && charset != "utf8" {
		return false
	}

	return true
}

func (s *DeviceServer) makeSSDPLocation(ip net.IP) string {
	var host net.IP
	if s.listenedAddr.IP.IsUnspecified() {
		host = ip
	} else {
		host = s.listenedAddr.IP
	}
	url := url.URL{
		Scheme: "http",
		Host: (&net.TCPAddr{
			IP:   host,
			Port: s.listenedAddr.Port,
		}).String(),
		Path: s.RootDescUrl,
	}
	return url.String()
}

func (s *DeviceServer) notifyError(err error) {
	if s.ErrorChan == nil || len(s.ErrorChan) == cap(s.ErrorChan) {
		return
	}
	s.ErrorChan <- err
}

func (s *DeviceServer) ssdpRoutine(ifi *net.Interface, devices []string, services []string) {
	ss, err := ssdp.NewSSDPServer(s.ctx, ifi, s.uuid)
	if err != nil {
		s.notifyError(err)
		return
	}
	s.ssdpList = append(s.ssdpList, ss)

	ss.Location = s.makeSSDPLocation
	ss.ServerDesc = fmt.Sprintf("UPnP/1.0 %s", s.ServerName)
	ss.Devices = devices
	ss.Services = services
	ss.ErrChan = s.ErrorChan
	ss.InterfaceAddrsFilter = InterfaceAddrsFilter

	err = ss.Init()
	if err != nil {
		s.notifyError(err)
		return
	}

	err = ss.ListenAndServe()
	if err != nil {
		s.notifyError(err)
		return
	}
}

func (s *DeviceServer) startSSDP() {
	devices := []string{}

	services := []string{}
	for _, srv := range s.makeServices() {
		services = append(services, srv.ServiceType)
	}

	for _, iface := range s.ifaces {
		if iface != nil && iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagMulticast != 0 {
			go s.ssdpRoutine(iface, devices, services)
		}
	}
}
