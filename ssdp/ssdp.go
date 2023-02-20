package ssdp

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/ipv4"
)

const (
	MulticastAddrPort string = "239.255.255.250:1900"
)

type SSDPServer interface {
	Init() error
	ListenAndServe() error
	Close()
}
type ssdpServer struct {
	ctx        context.Context
	signal     chan int
	conn       *net.UDPConn
	listenAddr *net.UDPAddr
	iface      *net.Interface
	addrs      []*net.IPNet

	ErrChan chan error

	UUID       string
	ServerDesc string

	Devices  []string
	Services []string

	AllowIps []*net.IPNet
	DenyIps  []*net.IPNet

	Location func(ip net.IP) string

	NotifyInterval uint

	InterfaceAddrsFilter func(iface *net.Interface, ip net.IP) bool
}

func (s *ssdpServer) interfaceAddrs(iface *net.Interface) []*net.IPNet {
	if iface == nil {
		return nil
	}
	addrs, err := iface.Addrs()
	if err != nil {
		return nil
	}
	addrStr := []*net.IPNet{}
	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && s.InterfaceAddrsFilter(iface, ip.IP) {
			addrStr = append(addrStr, ip)
		} else if ip, ok := addr.(*net.IPAddr); ok && s.InterfaceAddrsFilter(iface, ip.IP) {
			if len(ip.IP) == net.IPv4len {
				addrStr = append(addrStr, &net.IPNet{
					IP:   ip.IP,
					Mask: net.IPv4Mask(0, 0, 0, 0),
				})
			} else if len(ip.IP) == net.IPv6len {
				addrStr = append(addrStr, &net.IPNet{
					IP:   ip.IP,
					Mask: net.CIDRMask(0, 8*net.IPv6len),
				})
			}
		}
	}
	if len(addrStr) == 0 {
		return nil
	}

	return addrStr
}

func (s *ssdpServer) Init() error {
	if s.signal == nil {
		return fmt.Errorf("signal chan nil")
	}
	if s.iface == nil {
		return fmt.Errorf("interface nil")
	}
	if s.NotifyInterval == 0 {
		return fmt.Errorf("NotifyInterval zero")
	}
	if s.Location == nil {
		return fmt.Errorf("location function nil")
	}
	if s.InterfaceAddrsFilter == nil {
		return fmt.Errorf("interfaceAddrsFilter function nil")
	}
	if s.UUID == "" {
		return fmt.Errorf("uuid can not empty")
	}

	s.ServerDesc = "Linux/3.14.29 " + s.ServerDesc

	s.addrs = s.interfaceAddrs(s.iface)
	if s.addrs == nil {
		return &InterfaceError{s.iface}
	}

	return nil
}

func NewSSDPServer(ctx context.Context, iface *net.Interface, uuid string) (*ssdpServer, error) {
	return &ssdpServer{
		ctx:    ctx,
		signal: make(chan int, 1),
		iface:  iface,

		UUID: uuid,

		Location: func(ip net.IP) string {
			return ip.String()
		},
		NotifyInterval: 30,
	}, nil
}

func (s *ssdpServer) Close() {
	if s.signal == nil {
		return
	}
	s.signal <- 1 // 一切阻塞终将退出

	close(s.signal)
	s.signal = nil

	if s.conn != nil {
		s.sendByeBye()
		s.conn.Close()
		s.conn = nil
	}
}

func (s *ssdpServer) ListenAndServe() error {
	if s.signal == nil {
		return fmt.Errorf("please init ssdp server first")
	}
	var err error
	s.listenAddr, err = net.ResolveUDPAddr("udp4", MulticastAddrPort)
	if err != nil {
		panic(err)
	}
	s.conn, err = net.ListenMulticastUDP("udp", s.iface, s.listenAddr)
	if err != nil {
		return err
	}
	pack := ipv4.NewPacketConn(s.conn)
	err = pack.SetMulticastTTL(2)
	if err != nil {
		return err
	}

	go s.readUdpRoutine()

	go s.multicast()
	return nil
}

func (s *ssdpServer) multicast() {
	tick := time.NewTicker(time.Duration(s.NotifyInterval) * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-s.signal:
			return
		case <-tick.C:
		}
		for _, addr := range s.addrs {
			extHeads := map[string]string{
				"CACHE-CONTROL": fmt.Sprintf("max-age=%d", 5*s.NotifyInterval/2),
				"LOCATION":      s.Location(addr.IP),
			}
			s.sendAlive(extHeads)
		}
	}
}
func (s *ssdpServer) readUdpRoutine() {
	bs := int(math.Max(65535, float64(s.iface.MTU)))
	if bs <= 0 {
		bs = 65535
	}

	buf := make([]byte, bs)
	for {
		num, src, err := s.conn.ReadFromUDP(buf)

		if err != nil {
			// 全部退出
			s.signal <- 1

			s.notifyError(err) // 通知错误
			return
		}
		go s.readRequestRoutine(buf[:num], src)
	}
}
func (s *ssdpServer) checkRequest(req *http.Request) bool {
	man := req.Header.Get("man")
	if req.Method != "M-SEARCH" || man != `"ssdp:discover"` {
		return false
	}

	return true
}
func (s *ssdpServer) readMX(req *http.Request) int64 {
	if req.Header.Get("Host") == MulticastAddrPort {
		mxhd := req.Header.Get("mx")
		i, err := strconv.ParseUint(mxhd, 0, 8)
		if err == nil && i > 0 {
			return int64(i)
		}
	}
	return 1
}

func (s *ssdpServer) readSTS(req *http.Request) []string {
	st := req.Header.Get("ST")
	if st == "ssdp:all" {
		return s.ntList()
	}
	for _, nt := range s.ntList() {
		if st == nt {
			return []string{st}
		}
	}

	return nil
}

func (s *ssdpServer) ipnetContains(src net.IP) net.IP {
	for _, in := range s.addrs {
		if in.Contains(src) {
			return in.IP
		}
	}
	return nil
}
func (s *ssdpServer) readRequestRoutine(buf []byte, src *net.UDPAddr) {

	for _, ipnet := range s.DenyIps {
		if ipnet.Contains(src.IP) {
			s.notifyError(NewIPDenyError(src.IP))
			return
		}
	}
	var allowed = s.AllowIps == nil
	for _, ipnet := range s.AllowIps {
		if ipnet.Contains(src.IP) {
			allowed = true
			break
		}
	}
	if !allowed {
		s.notifyError(NewIPNotAllowError(src.IP))
		return
	}

	io := bufio.NewReader(bytes.NewReader(buf))
	req, err := http.ReadRequest(io)
	if err != nil {
		s.notifyError(&RequestError{src, err})
		return
	}
	if !s.checkRequest(req) {
		return
	}
	mx := s.readMX(req)
	sts := s.readSTS(req)

	// s.log.Info("receive", zap.String("from", src.String()), zap.String("body", string(buf)))

	ip := s.ipnetContains(src.IP)
	if ip == nil {
		return
	}
	// 单播响应
	for _, st := range sts {
		resp := s.makeResponse(ip, st)
		s.send(resp, src, time.Duration(rand.Int63n(mx)))
	}
}

func (s *ssdpServer) makeUSN(nt string) string {
	if s.UUID == nt {
		return nt
	}
	return s.UUID + "::" + nt
}

func (s *ssdpServer) ntList() []string {
	list := make([]string, 2)
	list[0] = "upnp:rootdevice"
	list[1] = s.UUID

	list = append(list, s.Devices...)
	list = append(list, s.Services...)

	return list
}

func (s *ssdpServer) send(buf []byte, ip *net.UDPAddr, delay time.Duration) {
	go func() {
		if delay > 0 {
			select {
			case <-time.After(delay):
				if s.conn == nil {
					// 拦截多个channel同时触发的情况
					return
				}
			case <-s.signal:
				return
			case <-s.ctx.Done():
				s.Close()
				return
			}
		}
		num, err := s.conn.WriteToUDP(buf, ip)
		if err != nil {
			s.notifyError(err)
		}
		if num != len(buf) {
			s.notifyError(fmt.Errorf("write error %d:%d", num, len(buf)))
		}
	}()
}

func appendHeaders(buf *bytes.Buffer, hd any) {
	switch hd := hd.(type) {
	case map[string]string:
		for k, v := range hd {
			fmt.Fprintf(buf, "%s: %s\r\n", k, v)
		}
	case string:
		fmt.Fprint(buf, hd)
	}
}

func (s *ssdpServer) makeResponse(ip net.IP, st string) []byte {
	head := map[string]string{
		"CACHE-CONTROL": fmt.Sprintf("max-age=%d", 5*s.NotifyInterval/2),
		"EXT":           "",
		"LOCATION":      s.Location(ip),
		"SERVER":        s.ServerDesc,
		"ST":            st,
		"USN":           s.makeUSN(st),
	}
	buf := &bytes.Buffer{}
	appendHeaders(buf, "HTTP/1.1 200 OK\r\n")
	appendHeaders(buf, head)
	return buf.Bytes()
}

func (s *ssdpServer) makeNotify(nt, nts string, extHeads map[string]string) []byte {
	head := map[string]string{
		"HOST":   MulticastAddrPort,
		"NT":     nt,
		"NTS":    nts,
		"SERVER": s.ServerDesc,
		"USN":    s.makeUSN(nt),
	}

	buf := &bytes.Buffer{}
	appendHeaders(buf, "NOTIFY * HTTP/1.1\r\n")
	appendHeaders(buf, head)
	appendHeaders(buf, extHeads)
	appendHeaders(buf, "\r\n")
	return buf.Bytes()
}

func (s *ssdpServer) sendByeBye() {
	for _, nt := range s.ntList() {
		buf := s.makeNotify(nt, "ssdp:byebye", nil)
		s.send(buf, s.listenAddr, 0)
	}
}

func (s *ssdpServer) sendAlive(extHeads map[string]string) {
	for _, nt := range s.ntList() {
		buf := s.makeNotify(nt, "ssdp:alive", extHeads)
		s.send(buf, s.listenAddr, time.Duration(rand.Int63n(100*int64(time.Millisecond))))
	}
}

func (s *ssdpServer) notifyError(err error) {
	if s.ErrChan == nil || len(s.ErrChan) == cap(s.ErrChan) {
		return
	}

	s.ErrChan <- &SSDPError{err}
}
