package soap

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

const (
	EncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
	EnvelopeNS    = "http://schemas.xmlsoap.org/soap/envelope/"
)

type ServiceURN struct {
	Auth    string
	Type    string
	Version uint64
}

func (me ServiceURN) String() string {
	return fmt.Sprintf("urn:%s:service:%s:%d", me.Auth, me.Type, me.Version)
}

func ParseServiceURN(s string) (ret ServiceURN, err error) {
	segs := strings.Split(s, ":")

	if len(segs) != 5 || segs[0] != "urn" || segs[2] != "service" {
		err = fmt.Errorf("unknown URN %s", s)
		return
	}
	ret.Auth = segs[1]
	ret.Type = segs[3]
	ret.Version, err = strconv.ParseUint(segs[4], 0, 32)
	return
}

type SoapAction struct {
	ServiceURN
	Action string
}

func ParseSOAPAction(s string) (ret *SoapAction, err error) {
	if len(s) < 3 {
		err = fmt.Errorf("soap action too short: %s", s)
		return
	}
	if s[0] != '"' || s[len(s)-1] != '"' {
		err = fmt.Errorf("soap action invalid: %s", s)
		return
	}
	s = s[1 : len(s)-1]
	hashIndex := strings.LastIndex(s, "#")
	if hashIndex == -1 {
		err = fmt.Errorf("soap action not found: %s", s)
		return
	}
	ret = &SoapAction{}
	ret.Action = s[hashIndex+1:]
	ret.ServiceURN, err = ParseServiceURN(s[:hashIndex])
	return
}

type EnvelopeArg struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type EnvelopeAction struct {
	XMLName xml.Name
	Args    []EnvelopeArg
}

type EnvelopeBody struct {
	Action []byte `xml:",innerxml"`
}

type Envelope struct {
	XMLName       xml.Name     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          EnvelopeBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}
type EnvelopeResponse struct {
	XMLName       xml.Name     `xml:"s:Envelope"`
	XMLSpace      string       `xml:"xmlns:s,attr"`
	EncodingStyle string       `xml:"encodingStyle,attr"`
	Body          EnvelopeBody `xml:"s:Body"`
}

type UPnPErrorXML struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:control-1-0 UPnPError"`
	Code    uint     `xml:"errorCode"`
	Desc    string   `xml:"errorDescription"`
}

type FaultDetail struct {
	XMLName xml.Name `xml:"detail"`
	Data    interface{}
}

type Fault struct {
	XMLName     xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	FaultCode   string      `xml:"faultcode"`
	FaultString string      `xml:"faultstring"`
	Detail      FaultDetail `xml:"detail"`
}

func NewFault(s string, detail interface{}) *Fault {
	return &Fault{
		FaultCode:   EnvelopeNS + ":Client",
		FaultString: s,
		Detail: FaultDetail{
			Data: detail,
		},
	}
}
