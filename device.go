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
	service "github.com/zwcway/fasthttp-upnp/service"
	"github.com/zwcway/fasthttp-upnp/soap"
	"github.com/zwcway/fasthttp-upnp/ssdp"
	"github.com/zwcway/fasthttp-upnp/utils"
)

const (
	replaced_FriendlyName = "**FRIENDLYNAME**"
	replaced_UUID         = "**UUID**"
)

type multiServer struct {
	name string
	uuid string
	root string
}
type DeviceServer struct {
	AuthName     string
	MultiDevices []*multiServer

	DeviceType string

	SpecVersion        scpd.SpecVersion
	Manufacturer       string
	ServerName         string
	RootDescNamespaces map[string]string
	UrlPrefix          string

	ServiceList []*service.Controller

	BeforeRequestHandle func(ctx *fasthttp.RequestCtx) bool
	AfterRequestHandle  func(ctx *fasthttp.RequestCtx) bool

	ListenInterface *net.Interface
	ListenPort      uint16

	AllowIps []*net.IPNet
	DenyIps  []*net.IPNet

	ErrorHandler ssdp.ErrorHandler
	InfoHandler  ssdp.InfoHandler

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

func (s *DeviceServer) AddServer(name, uuid, root string) string {
	if uuid == "" {
		uuid = utils.MakeUUID(name)
	}
	for _, ms := range s.MultiDevices {
		if ms.uuid == uuid {
			ms.name = name
			if root != "" {
				ms.root = root
			}
			s.initMultiServer()
			return ms.uuid
		}
	}
	s.MultiDevices = append(s.MultiDevices, &multiServer{name, uuid, root})
	s.initMultiServer()

	return uuid
}

func (s *DeviceServer) DelServer(uuid string) {
	for i, ms := range s.MultiDevices {
		if ms.uuid == uuid {
			if i == len(s.MultiDevices)-1 {
				s.MultiDevices = s.MultiDevices[:i]
				return
			}
			s.MultiDevices = append(s.MultiDevices[:i], s.MultiDevices[i+1:]...)
			return
		}
	}
	s.initMultiServer()
}

func (s *DeviceServer) initMultiServer() error {
	for _, md := range s.MultiDevices {
		if md.name == "" {
			return fmt.Errorf("FriendlyName con not empty")
		}
		if md.uuid == "" {
			md.uuid = utils.MakeUUID(md.name)
		}
		if md.root == "" {
			md.root = fmt.Sprintf("%s/%s/rootDesc.xml", s.UrlPrefix, md.uuid)
		} else {
			md.root = fmt.Sprintf("%s/%s", s.UrlPrefix, strings.Trim(md.root, "/"))
		}
	}
	return nil
}

func (s *DeviceServer) Init() (err error) {
	if s.ctx == nil {
		return fmt.Errorf("context can not nil")
	}
	s.UrlPrefix = strings.Trim(s.UrlPrefix, "/")
	if s.UrlPrefix != "" {
		s.UrlPrefix = "/" + s.UrlPrefix
	}

	if err := s.initMultiServer(); err != nil {
		return err
	}

	if s.AuthName == "" {
		s.AuthName = soap.AuthName
	}
	if s.SpecVersion.Major == 0 {
		s.SpecVersion.Major = 1
	}

	if s.ListenInterface != nil && s.ListenPort == 0 {
		return fmt.Errorf("must specify a listen port")
	}

	for _, ss := range s.ServiceList {
		ss.Init(s.ctx)
	}

	address, err := utils.GetListenAddress(s.ListenInterface, s.ListenPort)
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
			if ifi.Flags&net.FlagLoopback == 0 && ifi.Flags&net.FlagMulticast != 0 {
				s.ifaces = append(s.ifaces, &ifi)
			}
		}
	}

	s.listenedAddr = s.conn.Addr().(*net.TCPAddr)

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
			SCPDURL:     c.SCPDHttpPath(replaced_UUID),
			ControlURL:  c.ControlHttpPath(replaced_UUID),
			EventSubURL: c.EventHttpPath(replaced_UUID),
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
			FriendlyName: replaced_FriendlyName,
			Manufacturer: s.Manufacturer,
			ModelName:    s.ServerName,
			UDN:          replaced_UUID,
			ServiceList:  s.makeServices(),
		},
	}
}

func (s *DeviceServer) Connection() net.Listener {
	return s.conn
}

func NewDeviceServer(ctx context.Context, friendlyName string) (s *DeviceServer, uuid string, err error) {
	uuid = utils.MakeUUID(friendlyName)
	s = &DeviceServer{
		ctx:          ctx,
		MultiDevices: []*multiServer{{friendlyName, uuid, ""}},
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
	s.conn.Close()
	s.conn = nil
}

func (s *DeviceServer) Serve() error {
	if s.conn == nil {
		return nil
	}

	err := s.startSSDP()
	if err != nil {
		return err
	}

	s.notifyInfo("UPnP: listen on " + s.conn.Addr().String())

	server := fasthttp.Server{Handler: s.httpHandler}
	server.Serve(s.conn)

	return nil
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
	allowed := s.AllowIps == nil
	for _, ipnet := range s.AllowIps {
		if ipnet.Contains(remoteIp) {
			allowed = true
			break
		}
	}
	if !allowed {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		s.notifyError(ssdp.NewIPNotAllowError(remoteIp))
		return
	}

	if s.BeforeRequestHandle != nil && !s.BeforeRequestHandle(ctx) {
		return
	}

	var err error

	defer func() {
		if s.AfterRequestHandle != nil {
			s.AfterRequestHandle(ctx)
		}
		if err != nil {
			s.notifyError(err)
		}
	}()

	uri := string(ctx.Path())

	for i := range s.ServiceList {
		c := s.ServiceList[i]
		for _, md := range s.MultiDevices {
			switch uri {
			case c.SCPDHttpPath(md.uuid):
				if !ctx.IsGet() {
					ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
					return
				}
				if c.SCPDHandler != nil {
					err = c.SCPDHandler(c, ctx, md.uuid)
				} else {
					ctx.SetStatusCode(fasthttp.StatusNotFound)
				}
				return
			case c.ControlHttpPath(md.uuid):
				if !ctx.IsPost() {
					ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
					return
				}
				if !checkRequestIsXML(ctx) {
					ctx.SetStatusCode(fasthttp.StatusBadRequest)
					return
				}
				if c.ControlHandler != nil {
					err = c.ControlHandler(c, ctx, md.uuid)
				} else {
					ctx.SetStatusCode(fasthttp.StatusNotFound)
				}
				return
			case c.EventHttpPath(md.uuid):
				if c.EventHandler != nil {
					err = c.EventHandler(c, ctx, md.uuid)
				} else {
					ctx.SetStatusCode(fasthttp.StatusNotFound)
				}
				return
			}
		}
	}

	for _, md := range s.MultiDevices {
		if uri == md.root {
			ctx.Response.Header.SetContentType(`text/xml; charset="utf-8"`)

			xml := bytes.Replace(s.rootDescXML, []byte(replaced_FriendlyName), []byte(md.name), 1)
			xml = bytes.ReplaceAll(xml, []byte(replaced_UUID), []byte(md.uuid))

			ctx.Write(xml)
			return
		}
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

func (s *DeviceServer) makeSSDPLocation(uuid string, ip net.IP) string {
	var host net.IP
	if s.listenedAddr.IP.IsUnspecified() {
		host = ip
	} else {
		host = s.listenedAddr.IP
	}
	path := ""
	for _, md := range s.MultiDevices {
		if uuid == md.uuid {
			path = md.root
			break
		}
	}

	if path == "" {
		path = s.MultiDevices[0].root
	}

	url := url.URL{
		Scheme: "http",
		Host: (&net.TCPAddr{
			IP:   host,
			Port: s.listenedAddr.Port,
		}).String(),
		Path: path,
	}
	return url.String()
}

func (s *DeviceServer) notifyError(err error) {
	if s.ErrorHandler == nil {
		return
	}
	s.ErrorHandler(err)
}

func (s *DeviceServer) notifyInfo(err string) {
	if s.InfoHandler == nil {
		return
	}
	s.InfoHandler(err)
}

func (s *DeviceServer) startSSDP() error {
	services := []string{}
	for _, srv := range s.makeServices() {
		services = append(services, srv.ServiceType)
	}

	uuids := []string{}
	for _, md := range s.MultiDevices {
		uuids = append(uuids, md.uuid)
	}
	ss, err := ssdp.NewSSDPServer(s.ctx, s.ifaces, uuids)
	if err != nil {
		s.notifyError(err)
		return err
	}
	s.ssdpList = append(s.ssdpList, ss)

	ss.Location = s.makeSSDPLocation
	ss.ServerDesc = fmt.Sprintf("UPnP/1.0 %s", s.ServerName)
	ss.Services = services
	ss.ErrorHandler = s.ErrorHandler
	ss.InfoHandler = s.InfoHandler
	ss.InterfaceAddrsFilter = utils.InterfaceAddrsFilter
	ss.AllowIps = s.AllowIps
	ss.DenyIps = s.DenyIps

	err = ss.ListenAndServe()
	if err != nil {
		s.notifyError(err)
		return err
	}

	return nil
}
