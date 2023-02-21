package connectionmanager1

import "encoding/xml"

type ArgInConnectionComplete struct {
	XMLName      xml.Name `xml:"urn:schemas-upnp-org:service:ConnectionManager:1 ConnectionComplete"`
	ConnectionID int32    `soap:"A_ARG_TYPE_ConnectionID"`
}
type ArgOutConnectionComplete struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:ConnectionManager:1 ConnectionCompleteResponse"`
}
type ArgInGetCurrentConnectionIDs struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:ConnectionManager:1 GetCurrentConnectionIDs"`
}
type ArgOutGetCurrentConnectionIDs struct {
	XMLName       xml.Name `xml:"urn:schemas-upnp-org:service:ConnectionManager:1 GetCurrentConnectionIDsResponse"`
	ConnectionIDs string   `soap:"CurrentConnectionIDs,sendevent"`
}
type ArgInGetCurrentConnectionInfo struct {
	XMLName      xml.Name `xml:"urn:schemas-upnp-org:service:ConnectionManager:1 GetCurrentConnectionInfo"`
	ConnectionID int32    `soap:"A_ARG_TYPE_ConnectionID"`
}
type ArgOutGetCurrentConnectionInfo struct {
	XMLName               xml.Name `xml:"urn:schemas-upnp-org:service:ConnectionManager:1 GetCurrentConnectionInfoResponse"`
	AVTransportID         int32    `soap:"A_ARG_TYPE_AVTransportID"`
	Direction             string   `soap:"A_ARG_TYPE_Direction" allowed:"Input,Output"`
	PeerConnectionID      int32    `soap:"A_ARG_TYPE_ConnectionID"`
	PeerConnectionManager string   `soap:"A_ARG_TYPE_ConnectionManager"`
	ProtocolInfo          string   `soap:"A_ARG_TYPE_ProtocolInfo"`
	RcsID                 int32    `soap:"A_ARG_TYPE_RcsID"`
	Status                string   `soap:"A_ARG_TYPE_ConnectionStatus" allowed:"OK,ContentFormatMismatch,InsufficientBandwidth,UnreliableChannel,Unknown"`
}
type ArgInGetProtocolInfo struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:ConnectionManager:1 GetProtocolInfo"`
}
type ArgOutGetProtocolInfo struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:ConnectionManager:1 GetProtocolInfoResponse"`
	Sink    string   `soap:"SinkProtocolInfo,sendevent"`
	Source  string   `soap:"SourceProtocolInfo,sendevent"`
}
type ArgInPrepareForConnection struct {
	XMLName               xml.Name `xml:"urn:schemas-upnp-org:service:ConnectionManager:1 PrepareForConnection"`
	Direction             string   `soap:"A_ARG_TYPE_Direction" allowed:"Input,Output"`
	PeerConnectionID      int32    `soap:"A_ARG_TYPE_ConnectionID"`
	PeerConnectionManager string   `soap:"A_ARG_TYPE_ConnectionManager"`
	RemoteProtocolInfo    string   `soap:"A_ARG_TYPE_ProtocolInfo"`
}
type ArgOutPrepareForConnection struct {
	XMLName       xml.Name `xml:"urn:schemas-upnp-org:service:ConnectionManager:1 PrepareForConnectionResponse"`
	AVTransportID int32    `soap:"A_ARG_TYPE_AVTransportID"`
	ConnectionID  int32    `soap:"A_ARG_TYPE_ConnectionID"`
	RcsID         int32    `soap:"A_ARG_TYPE_RcsID"`
}
