package upnp

const (
	AuthName = "schemas-upnp-org"

	DeviceType_MediaServer   = "MediaServer"
	DeviceType_MediaRenderer = "MediaRenderer"

	ServiceName_AVTransport = "AVTransport"
	ServiceName_ConnectionManager = "ConnectionManager"
	ServiceName_RenderingControl = "RenderingControl"
)

const (
	DirIn  string = "in"
	DirOut string = "out"

	BoolYes string = "yes"
	BoolNo  string = "no"

	DataTypeInt32  = "i4"
	DataTypeUint32 = "ui4"
	DataTypeStr    = "string"
	DataTypeInt16  = "i2"
	DataTypeUInt16 = "ui2"
	DataTypeBool   = "boolean"
)

const NTEvent = "upnp:event"

const ResponseContentTypeXML = `text/xml; charset="utf-8"`
