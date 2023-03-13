package ssdp

import (
	"fmt"
	"net"
)

type IPDenyError struct {
	IP       net.IP
	FromDeny bool
}

func (e *IPDenyError) Error() string {
	if e.FromDeny {
		return fmt.Sprintf("ip '%s' denied from deny list", e.IP.String())
	} else {
		return fmt.Sprintf("ip '%s' denied from allow list", e.IP.String())
	}
}

func NewIPDenyError(ip net.IP) error {
	return &IPDenyError{ip, true}
}
func NewIPNotAllowError(ip net.IP) error {
	return &IPDenyError{ip, false}
}

type SSDPError struct {
	Err error
}

func (e *SSDPError) Error() string {
	return fmt.Sprintf("SSDP: %s", e.Err.Error())
}

type RequestError struct {
	Src *net.UDPAddr
	Err error
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("request error from %s '%s'", e.Src.String(), e.Err.Error())
}

type InterfaceError struct {
	Iface *net.Interface
	Err   error
}

func (e *InterfaceError) Error() (s string) {
	return fmt.Sprintf("the interface '%s' error : %s\n", e.Iface.Name, e.Err)
}

func IsIPDenyError(err error) bool {
	_, ok := err.(*IPDenyError)
	return ok
}
func IsRequestError(err error) bool {
	_, ok := err.(*RequestError)
	return ok
}
func IsSSDPError(err error) bool {
	_, ok := err.(*SSDPError)
	return ok
}
func IsInterfaceError(err error) bool {
	_, ok := err.(*InterfaceError)
	return ok
}

type ErrorHandler func(error)
type InfoHandler func(string)
