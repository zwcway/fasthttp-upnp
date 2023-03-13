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
	"runtime"
	"strconv"
	"time"

	"golang.org/x/net/ipv4"
)

const (
	MulticastAddrPort string = "239.255.255.250:1900"
)

type multiConn struct {
	ifi  *net.Interface
	conn *net.UDPConn
	addr []*net.IPNet
}

type SSDPServer interface {
	Init(ifaces []*net.Interface) error
	ListenAndServe() error
	Close()
}
type ssdpServer struct {
	ctx        context.Context
	signal     chan int
	listenAddr *net.UDPAddr
	multi      []*multiConn

	ErrorHandler ErrorHandler
	InfoHandler  InfoHandler

	uuids      []string
	ServerDesc string

	Services []string

	AllowIps []*net.IPNet
	DenyIps  []*net.IPNet

	Location func(uuid string, ip net.IP) string

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

func (s *ssdpServer) Init(ifaces []*net.Interface) error {
	if s.signal == nil {
		return fmt.Errorf("signal chan nil")
	}
	if s.NotifyInterval == 0 {
		return fmt.Errorf("NotifyInterval zero")
	}
	if s.Location == nil {
		return fmt.Errorf("location function nil")
	}
	if s.InterfaceAddrsFilter == nil {
		s.InterfaceAddrsFilter = func(iface *net.Interface, ip net.IP) bool { return true }
	}
	if len(s.uuids) == 0 {
		return fmt.Errorf("uuid can not empty")
	}

	s.ServerDesc = fmt.Sprintf("%s/%s %s", runtime.GOOS, runtime.Version(), s.ServerDesc)

	s.multi = make([]*multiConn, 0)
	if len(ifaces) == 0 {
		ifis, err := net.Interfaces()
		if err != nil {
			return err
		}
		for i := range ifis {
			ifi := ifis[i]
			ifaces = append(ifaces, &ifi)
		}
	}

	for _, ifi := range ifaces {
		addrs := s.interfaceAddrs(ifi)
		if len(addrs) == 0 {
			continue
		}
		mc := &multiConn{ifi, nil, addrs}
		s.multi = append(s.multi, mc)
	}

	return nil
}

func NewSSDPServer(ctx context.Context, ifaces []*net.Interface, uuids []string) (*ssdpServer, error) {
	s := &ssdpServer{
		ctx:    ctx,
		signal: make(chan int, 1),

		uuids: uuids,

		Location: func(uuid string, ip net.IP) string {
			return ip.String()
		},
		NotifyInterval: 30,
	}
	err := s.Init(ifaces)

	return s, err
}

func (s *ssdpServer) Close() {
	if s.signal == nil {
		return
	}
	s.signal <- 1 // 一切阻塞终将退出

	close(s.signal)
	s.signal = nil

	s.sendByeBye()
	for _, mc := range s.multi {
		if mc.conn != nil {
			mc.conn.Close()
		}
	}
	s.multi = nil
}

func (s *ssdpServer) ListenAndServe() (err error) {
	if s.signal == nil || s.multi == nil {
		return fmt.Errorf("please init ssdp server first")
	}
	s.listenAddr, err = net.ResolveUDPAddr("udp4", MulticastAddrPort)
	if err != nil {
		panic(err)
	}
	sucCount := 0
	for _, mc := range s.multi {
		mc.conn, err = net.ListenMulticastUDP("udp", mc.ifi, s.listenAddr)
		if err != nil {
			mc.conn = nil
			s.notifyError(&InterfaceError{mc.ifi, err})
			continue
		}
		pack := ipv4.NewPacketConn(mc.conn)
		err = pack.SetMulticastTTL(2)
		if err != nil {
			mc.conn.Close()
			mc.conn = nil
			s.notifyError(&InterfaceError{mc.ifi, err})
			continue
		}

		s.notifyInfo("SSDP: listen on " + mc.ifi.Name)

		go s.readUdpRoutine(mc)

		sucCount++
	}

	if sucCount > 0 {
		go s.multicast()
	} else {
		return fmt.Errorf("start failed")
	}

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
		extHeads := map[string]string{
			"CACHE-CONTROL": fmt.Sprintf("max-age=%d", 5*s.NotifyInterval/2),
			"LOCATION":      "",
		}
		for _, uuid := range s.uuids {
			for _, mc := range s.multi {
				if mc.conn == nil {
					continue
				}
				for _, addr := range mc.addr {
					extHeads["LOCATION"] = s.Location(uuid, addr.IP)
					s.sendAlive(mc, uuid, extHeads)
				}
			}
		}
	}
}

func (s *ssdpServer) readUdpRoutine(mc *multiConn) {
	bs := int(math.Max(65535, float64(mc.ifi.MTU)))
	if bs <= 0 {
		bs = 65535
	}

	buf := make([]byte, bs)
	for {
		num, src, err := mc.conn.ReadFromUDP(buf)

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
	return 2
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

func (s *ssdpServer) ipnetContains(src net.IP) (*multiConn, net.IP) {
	for _, mc := range s.multi {
		if mc.conn == nil {
			continue
		}

		for _, in := range mc.addr {
			if in.Contains(src) {
				return mc, in.IP
			}
		}
	}
	return nil, nil
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

	s.notifyInfo("SSDP: request from " + src.String())

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

	mc, ip := s.ipnetContains(src.IP)
	if ip == nil {
		return
	}
	// 单播响应
	for _, uuid := range s.uuids {
		for _, st := range sts {
			resp := s.makeResponse(uuid, ip, st)
			s.send(mc, resp, src, time.Duration(rand.Int63n(mx)))
		}
	}
}

func (s *ssdpServer) makeUSN(uuid, nt string) string {
	if uuid == nt {
		return nt
	}
	return "uuid:" + uuid + "::" + nt
}

func (s *ssdpServer) ntList() []string {
	list := make([]string, 1)
	list[0] = "upnp:rootdevice"

	list = append(list, s.uuids...)
	list = append(list, s.Services...)

	return list
}

func (s *ssdpServer) send(mc *multiConn, buf []byte, ip *net.UDPAddr, delay time.Duration) {
	go func() {
		if delay > 0 {
			select {
			case <-time.After(delay):
				if mc.conn == nil {
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
		num, err := mc.conn.WriteToUDP(buf, ip)
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

func (s *ssdpServer) makeResponse(uuid string, ip net.IP, st string) []byte {
	head := map[string]string{
		"CACHE-CONTROL": fmt.Sprintf("max-age=%d", 5*s.NotifyInterval/2),
		"EXT":           "",
		"LOCATION":      s.Location(uuid, ip),
		"SERVER":        s.ServerDesc,
		"ST":            st,
		"USN":           s.makeUSN(uuid, st),
	}
	buf := &bytes.Buffer{}
	appendHeaders(buf, "HTTP/1.1 200 OK\r\n")
	appendHeaders(buf, head)
	return buf.Bytes()
}

func (s *ssdpServer) makeNotify(uuid, nt, nts string, extHeads map[string]string) []byte {
	head := map[string]string{
		"HOST":   MulticastAddrPort,
		"NT":     nt,
		"NTS":    nts,
		"SERVER": s.ServerDesc,
		"USN":    s.makeUSN(uuid, nt),
	}

	buf := &bytes.Buffer{}
	appendHeaders(buf, "NOTIFY * HTTP/1.1\r\n")
	appendHeaders(buf, head)
	appendHeaders(buf, extHeads)
	appendHeaders(buf, "\r\n")
	return buf.Bytes()
}

func (s *ssdpServer) sendByeBye() {
	for _, uuid := range s.uuids {
		for _, mc := range s.multi {
			if mc.conn == nil {
				continue
			}
			s.notifyInfo("SSDP: byebye")
			for _, nt := range s.ntList() {
				buf := s.makeNotify(uuid, nt, "ssdp:byebye", nil)
				mc.conn.WriteToUDP(buf, s.listenAddr)
			}
		}
	}
}

func (s *ssdpServer) sendAlive(mc *multiConn, uuid string, extHeads map[string]string) {
	for _, nt := range s.ntList() {
		buf := s.makeNotify(uuid, nt, "ssdp:alive", extHeads)
		s.send(mc, buf, s.listenAddr, time.Duration(rand.Int63n(500*int64(time.Millisecond))))
	}
}

func (s *ssdpServer) notifyError(err error) {
	if s.ErrorHandler == nil {
		return
	}

	s.ErrorHandler(&SSDPError{err})
}
func (s *ssdpServer) notifyInfo(err string) {
	if s.InfoHandler == nil {
		return
	}
	s.InfoHandler(err)
}
