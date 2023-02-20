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
	src *net.UDPAddr
	err error
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("request error from %s '%s'", e.src.String(), e.err.Error())
}

func IsIPDenyError(err error) bool {
	if IsSSDPError(err) {
		return true
	}
	_, ok := err.(*IPDenyError)
	return ok
}
func IsRequestError(err error) bool {
	if IsSSDPError(err) {
		return true
	}
	_, ok := err.(*RequestError)
	return ok
}
func IsSSDPError(err error) bool {
	_, ok := err.(*SSDPError)
	return ok
}
