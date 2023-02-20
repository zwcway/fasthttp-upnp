package upnp

import (
	"crypto/md5"
	"fmt"
	"io"
	"net"
	"net/netip"
)

func MakeUUID(unique string) string {
	md := md5.New()

	if _, err := io.WriteString(md, unique); err != nil {
		panic(fmt.Errorf("make uuid failed"))
	}

	buf := md.Sum(nil)

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[:4], buf[4:6], buf[6:8], buf[8:10], buf[10:16])
}

var InterfaceAddrsFilter = func(iface *net.Interface, ip net.IP) bool {
	return !ip.IsLoopback() && !ip.IsMulticast() && !ip.IsInterfaceLocalMulticast() && !ip.IsLinkLocalMulticast() && !ip.IsLinkLocalUnicast()
}

func getListenAddress(iface *net.Interface, port uint16) (address string, err error) {
	if iface != nil {
		var addrs []net.Addr
		addrs, err = iface.Addrs()
		if err != nil {
			return
		}
		for _, addr := range addrs {
			if ip, ok := addr.(*net.IPNet); ok && InterfaceAddrsFilter(iface, ip.IP) {
				a, _ := netip.AddrFromSlice(ip.IP)
				return netip.AddrPortFrom(a, port).String(), nil
			} else if ip, ok := addr.(*net.IPAddr); ok && InterfaceAddrsFilter(iface, ip.IP) {
				a, _ := netip.AddrFromSlice(ip.IP)
				return netip.AddrPortFrom(a, port).String(), nil
			}
		}
	}

	return fmt.Sprintf(":%d", port), nil
}

func ServiceNS(service string, ver int) string {
	return fmt.Sprintf("urn:%s:service:%s:%d", AuthName, service, ver)
}
