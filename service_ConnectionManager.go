package upnp

import (
	"encoding/xml"

	"github.com/zwcway/fasthttp-upnp/scpd"
	"github.com/zwcway/fasthttp-upnp/soap"
)

var CM_A_ARG_TYPE_AVTransportID = scpd.Variable{
	Name:       "A_ARG_TYPE_AVTransportID",
	DataType:   DataTypeInt32,
	SendEvents: BoolNo,
}
var CM_A_ARG_TYPE_ConnectionID = scpd.Variable{
	Name:       "A_ARG_TYPE_ConnectionID",
	DataType:   DataTypeInt32,
	SendEvents: BoolNo,
}
var CM_A_ARG_TYPE_ConnectionManager = scpd.Variable{
	Name:       "A_ARG_TYPE_ConnectionManager",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var CM_A_ARG_TYPE_ConnectionStatus = scpd.Variable{
	Name:          "A_ARG_TYPE_ConnectionStatus",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"OK", "ContentFormatMismatch", "InsufficientBandwidth", "UnreliableChannel", "Unknown"},
}
var CM_A_ARG_TYPE_Direction = scpd.Variable{
	Name:          "A_ARG_TYPE_Direction",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"Input", "Output"},
}
var CM_A_ARG_TYPE_ProtocolInfo = scpd.Variable{
	Name:       "A_ARG_TYPE_ProtocolInfo",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var CM_A_ARG_TYPE_RcsID = scpd.Variable{
	Name:       "A_ARG_TYPE_RcsID",
	DataType:   DataTypeInt32,
	SendEvents: BoolNo,
}
var CM_CurrentConnectionIDs = scpd.Variable{
	Name:       "CurrentConnectionIDs",
	DataType:   DataTypeStr,
	SendEvents: BoolYes,
}
var CM_SinkProtocolInfo = scpd.Variable{
	Name:       "SinkProtocolInfo",
	DataType:   DataTypeStr,
	SendEvents: BoolYes,
}
var CM_SourceProtocolInfo = scpd.Variable{
	Name:       "SourceProtocolInfo",
	DataType:   DataTypeStr,
	SendEvents: BoolYes,
}

type CMArgIn_ConnectionComplete struct {
	XMLName      xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v ConnectionComplete"`
	ConnectionID int32
}
type CMArgOut_ConnectionComplete struct {
	XMLName   xml.Name `xml:"u:ConnectionCompleteResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var CM_ConnectionComplete = Action{
	Handler: nil,
	ArgIn:   CMArgIn_ConnectionComplete{},
	ArgOut:  CMArgOut_ConnectionComplete{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"ConnectionID", DirIn, &CM_A_ARG_TYPE_ConnectionID},
	},
}

type CMArgIn_GetCurrentConnectionIDs struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v GetCurrentConnectionIDs"`
}
type CMArgOut_GetCurrentConnectionIDs struct {
	XMLName       xml.Name `xml:"u:GetCurrentConnectionIDsResponse"`
	XMLPrefix     string   `xml:"xmlns:u,attr"`
	ConnectionIDs string
}

var CM_GetCurrentConnectionIDs = Action{
	Handler: nil,
	ArgIn:   CMArgIn_GetCurrentConnectionIDs{},
	ArgOut:  CMArgOut_GetCurrentConnectionIDs{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"ConnectionIDs", DirOut, &CM_CurrentConnectionIDs},
	},
}

type CMArgIn_GetCurrentConnectionInfo struct {
	XMLName      xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v GetCurrentConnectionInfo"`
	ConnectionID int32
}
type CMArgOut_GetCurrentConnectionInfo struct {
	XMLName               xml.Name `xml:"u:GetCurrentConnectionInfoResponse"`
	XMLPrefix             string   `xml:"xmlns:u,attr"`
	RcsID                 int32
	AVTransportID         int32
	ProtocolInfo          string
	PeerConnectionManager string
	PeerConnectionID      int32
	Direction             string
	Status                string
}

var CM_GetCurrentConnectionInfo = Action{
	Handler: nil,
	ArgIn:   CMArgIn_GetCurrentConnectionInfo{},
	ArgOut:  CMArgOut_GetCurrentConnectionInfo{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"ConnectionID", DirIn, &CM_A_ARG_TYPE_ConnectionID},
		{"RcsID", DirOut, &CM_A_ARG_TYPE_RcsID},
		{"AVTransportID", DirOut, &CM_A_ARG_TYPE_AVTransportID},
		{"ProtocolInfo", DirOut, &CM_A_ARG_TYPE_ProtocolInfo},
		{"PeerConnectionManager", DirOut, &CM_A_ARG_TYPE_ConnectionManager},
		{"PeerConnectionID", DirOut, &CM_A_ARG_TYPE_ConnectionID},
		{"Direction", DirOut, &CM_A_ARG_TYPE_Direction},
		{"Status", DirOut, &CM_A_ARG_TYPE_ConnectionStatus},
	},
}

type CMArgIn_GetProtocolInfo struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v GetProtocolInfo"`
}
type CMArgOut_GetProtocolInfo struct {
	XMLName   xml.Name `xml:"u:GetProtocolInfoResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
	Source    string
	Sink      string
}

var CM_GetProtocolInfo = Action{
	Handler: nil,
	ArgIn:   CMArgIn_GetProtocolInfo{},
	ArgOut:  CMArgOut_GetProtocolInfo{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"Source", DirOut, &CM_SourceProtocolInfo},
		{"Sink", DirOut, &CM_SinkProtocolInfo},
	},
}

type CMArgIn_PrepareForConnection struct {
	XMLName               xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v PrepareForConnection"`
	RemoteProtocolInfo    string
	PeerConnectionManager string
	PeerConnectionID      int32
	Direction             string
}
type CMArgOut_PrepareForConnection struct {
	XMLName       xml.Name `xml:"u:PrepareForConnectionResponse"`
	XMLPrefix     string   `xml:"xmlns:u,attr"`
	ConnectionID  int32
	AVTransportID int32
	RcsID         int32
}

var CM_PrepareForConnection = Action{
	Handler: nil,
	ArgIn:   CMArgIn_PrepareForConnection{},
	ArgOut:  CMArgOut_PrepareForConnection{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"RemoteProtocolInfo", DirIn, &CM_A_ARG_TYPE_ProtocolInfo},
		{"PeerConnectionManager", DirIn, &CM_A_ARG_TYPE_ConnectionManager},
		{"PeerConnectionID", DirIn, &CM_A_ARG_TYPE_ConnectionID},
		{"Direction", DirIn, &CM_A_ARG_TYPE_Direction},
		{"ConnectionID", DirOut, &CM_A_ARG_TYPE_ConnectionID},
		{"AVTransportID", DirOut, &CM_A_ARG_TYPE_AVTransportID},
		{"RcsID", DirOut, &CM_A_ARG_TYPE_RcsID},
	},
}
var ConnectionManagerV1 = ActionMap{
	"ConnectionComplete":       &CM_ConnectionComplete,
	"GetCurrentConnectionIDs":  &CM_GetCurrentConnectionIDs,
	"GetCurrentConnectionInfo": &CM_GetCurrentConnectionInfo,
	"GetProtocolInfo":          &CM_GetProtocolInfo,
	"PrepareForConnection":     &CM_PrepareForConnection,
}
