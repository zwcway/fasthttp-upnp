package upnp

import (
	"encoding/xml"

	"github.com/zwcway/fasthttp-upnp/scpd"
	"github.com/zwcway/fasthttp-upnp/soap"
)

var AVT_AVTransportURI = scpd.Variable{
	Name:       "AVTransportURI",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_AVTransportURIMetaData = scpd.Variable{
	Name:       "AVTransportURIMetaData",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_A_ARG_TYPE_InstanceID = scpd.Variable{
	Name:       "A_ARG_TYPE_InstanceID",
	DataType:   DataTypeUint32,
	SendEvents: BoolNo,
}
var AVT_A_ARG_TYPE_SeekMode = scpd.Variable{
	Name:          "A_ARG_TYPE_SeekMode",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"ABS_TIME", "REL_TIME", "ABS_COUNT", "REL_COUNT", "TRACK_NR", "CHANNEL_FREQ", "TAPE-INDEX", "FRAME"},
}
var AVT_A_ARG_TYPE_SeekTarget = scpd.Variable{
	Name:       "A_ARG_TYPE_SeekTarget",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_AbsoluteCounterPosition = scpd.Variable{
	Name:       "AbsoluteCounterPosition",
	DataType:   DataTypeInt32,
	SendEvents: BoolNo,
}
var AVT_AbsoluteTimePosition = scpd.Variable{
	Name:       "AbsoluteTimePosition",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_CurrentMediaDuration = scpd.Variable{
	Name:       "CurrentMediaDuration",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_CurrentPlayMode = scpd.Variable{
	Name:          "CurrentPlayMode",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	Default:       "NORMAL",
	AllowedValues: &[]string{"NORMAL", "SHUFFLE", "REPEAT_ONE", "REPEAT_ALL", "RANDOM", "DIRECT_1", "INTRO"},
}
var AVT_CurrentRecordQualityMode = scpd.Variable{
	Name:          "CurrentRecordQualityMode",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"0:EP", "1:LP", "2:SP", "0:BASIC", "1:MEDIUM", "2:HIGH", "NOT_IMPLEMENTED"},
}
var AVT_CurrentTrack = scpd.Variable{
	Name:         "CurrentTrack",
	DataType:     DataTypeUint32,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var AVT_CurrentTrackDuration = scpd.Variable{
	Name:       "CurrentTrackDuration",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_CurrentTrackMetaData = scpd.Variable{
	Name:       "CurrentTrackMetaData",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_CurrentTrackURI = scpd.Variable{
	Name:       "CurrentTrackURI",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_CurrentTransportActions = scpd.Variable{
	Name:       "CurrentTransportActions",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_LastChange = scpd.Variable{
	Name:       "LastChange",
	DataType:   DataTypeStr,
	SendEvents: BoolYes,
}
var AVT_NextAVTransportURI = scpd.Variable{
	Name:       "NextAVTransportURI",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_NextAVTransportURIMetaData = scpd.Variable{
	Name:       "NextAVTransportURIMetaData",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_NumberOfTracks = scpd.Variable{
	Name:         "NumberOfTracks",
	DataType:     DataTypeUint32,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 0},
}
var AVT_PlaybackStorageMedium = scpd.Variable{
	Name:          "PlaybackStorageMedium",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"UNKNOWN", "DV", "MINI-DV", "VHS", "W-VHS", "S-VHS", "D-VHS", "VHSC", "VIDEO8", "HI8", "CD-ROM", "CD-DA", "CD-R", "CD-RW", "VIDEO-CD", "SACD", "MD-AUDIO", "MD-PICTURE", "DVD-ROM", "DVD-VIDEO", "DVD-R", "DVD+RW", "DVD-RW", "DVD-RAM", "DVD-AUDIO", "DAT", "LD", "HDD", "MICRO-MV", "NETWORK", "NONE", "NOT_IMPLEMENTED"},
}
var AVT_PossiblePlaybackStorageMedia = scpd.Variable{
	Name:       "PossiblePlaybackStorageMedia",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_PossibleRecordQualityModes = scpd.Variable{
	Name:       "PossibleRecordQualityModes",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_PossibleRecordStorageMedia = scpd.Variable{
	Name:       "PossibleRecordStorageMedia",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_RecordMediumWriteStatus = scpd.Variable{
	Name:          "RecordMediumWriteStatus",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"WRITABLE", "PROTECTED", "NOT_WRITABLE", "UNKNOWN", "NOT_IMPLEMENTED"},
}
var AVT_RecordStorageMedium = scpd.Variable{
	Name:          "RecordStorageMedium",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"UNKNOWN", "DV", "MINI-DV", "VHS", "W-VHS", "S-VHS", "D-VHS", "VHSC", "VIDEO8", "HI8", "CD-ROM", "CD-DA", "CD-R", "CD-RW", "VIDEO-CD", "SACD", "MD-AUDIO", "MD-PICTURE", "DVD-ROM", "DVD-VIDEO", "DVD-R", "DVD+RW", "DVD-RW", "DVD-RAM", "DVD-AUDIO", "DAT", "LD", "HDD", "MICRO-MV", "NETWORK", "NONE", "NOT_IMPLEMENTED"},
}
var AVT_RelativeCounterPosition = scpd.Variable{
	Name:       "RelativeCounterPosition",
	DataType:   DataTypeInt32,
	SendEvents: BoolNo,
}
var AVT_RelativeTimePosition = scpd.Variable{
	Name:       "RelativeTimePosition",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var AVT_TransportPlaySpeed = scpd.Variable{
	Name:          "TransportPlaySpeed",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"1"},
}
var AVT_TransportState = scpd.Variable{
	Name:          "TransportState",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"STOPPED", "PAUSED_PLAYBACK", "PAUSED_RECORDING", "PLAYING", "RECORDING", "TRANSITIONING", "NO_MEDIA_PRESENT"},
}
var AVT_TransportStatus = scpd.Variable{
	Name:          "TransportStatus",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"OK", "ERROR_OCCURRED"},
}

type AVTArgIn_GetCurrentTransportActions struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v GetCurrentTransportActions"`
	InstanceID uint32
}
type AVTArgOut_GetCurrentTransportActions struct {
	XMLName   xml.Name `xml:"u:GetCurrentTransportActionsResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
	Actions   string
}

var AVT_GetCurrentTransportActions = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_GetCurrentTransportActions{},
	ArgOut:  AVTArgOut_GetCurrentTransportActions{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"Actions", DirOut, &AVT_CurrentTransportActions},
	},
}

type AVTArgIn_GetDeviceCapabilities struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v GetDeviceCapabilities"`
	InstanceID uint32
}
type AVTArgOut_GetDeviceCapabilities struct {
	XMLName         xml.Name `xml:"u:GetDeviceCapabilitiesResponse"`
	XMLPrefix       string   `xml:"xmlns:u,attr"`
	PlayMedia       string
	RecMedia        string
	RecQualityModes string
}

var AVT_GetDeviceCapabilities = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_GetDeviceCapabilities{},
	ArgOut:  AVTArgOut_GetDeviceCapabilities{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"PlayMedia", DirOut, &AVT_PossiblePlaybackStorageMedia},
		{"RecMedia", DirOut, &AVT_PossibleRecordStorageMedia},
		{"RecQualityModes", DirOut, &AVT_PossibleRecordQualityModes},
	},
}

type AVTArgIn_GetMediaInfo struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v GetMediaInfo"`
	InstanceID uint32
}
type AVTArgOut_GetMediaInfo struct {
	XMLName            xml.Name `xml:"u:GetMediaInfoResponse"`
	XMLPrefix          string   `xml:"xmlns:u,attr"`
	NrTracks           uint32
	MediaDuration      string
	CurrentURI         string
	CurrentURIMetaData string
	NextURI            string
	NextURIMetaData    string
	PlayMedium         string
	RecordMedium       string
	WriteStatus        string
}

var AVT_GetMediaInfo = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_GetMediaInfo{},
	ArgOut:  AVTArgOut_GetMediaInfo{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"NrTracks", DirOut, &AVT_NumberOfTracks},
		{"MediaDuration", DirOut, &AVT_CurrentMediaDuration},
		{"CurrentURI", DirOut, &AVT_AVTransportURI},
		{"CurrentURIMetaData", DirOut, &AVT_AVTransportURIMetaData},
		{"NextURI", DirOut, &AVT_NextAVTransportURI},
		{"NextURIMetaData", DirOut, &AVT_NextAVTransportURIMetaData},
		{"PlayMedium", DirOut, &AVT_PlaybackStorageMedium},
		{"RecordMedium", DirOut, &AVT_RecordStorageMedium},
		{"WriteStatus", DirOut, &AVT_RecordMediumWriteStatus},
	},
}

type AVTArgIn_GetPositionInfo struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v GetPositionInfo"`
	InstanceID uint32
}
type AVTArgOut_GetPositionInfo struct {
	XMLName       xml.Name `xml:"u:GetPositionInfoResponse"`
	XMLPrefix     string   `xml:"xmlns:u,attr"`
	Track         uint32
	TrackDuration string
	TrackMetaData string
	TrackURI      string
	RelTime       string
	AbsTime       string
	RelCount      int32
	AbsCount      int32
}

var AVT_GetPositionInfo = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_GetPositionInfo{},
	ArgOut:  AVTArgOut_GetPositionInfo{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"Track", DirOut, &AVT_CurrentTrack},
		{"TrackDuration", DirOut, &AVT_CurrentTrackDuration},
		{"TrackMetaData", DirOut, &AVT_CurrentTrackMetaData},
		{"TrackURI", DirOut, &AVT_CurrentTrackURI},
		{"RelTime", DirOut, &AVT_RelativeTimePosition},
		{"AbsTime", DirOut, &AVT_AbsoluteTimePosition},
		{"RelCount", DirOut, &AVT_RelativeCounterPosition},
		{"AbsCount", DirOut, &AVT_AbsoluteCounterPosition},
	},
}

type AVTArgIn_GetTransportInfo struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v GetTransportInfo"`
	InstanceID uint32
}
type AVTArgOut_GetTransportInfo struct {
	XMLName                xml.Name `xml:"u:GetTransportInfoResponse"`
	XMLPrefix              string   `xml:"xmlns:u,attr"`
	CurrentTransportState  string
	CurrentTransportStatus string
	CurrentSpeed           string
}

var AVT_GetTransportInfo = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_GetTransportInfo{},
	ArgOut:  AVTArgOut_GetTransportInfo{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"CurrentTransportState", DirOut, &AVT_TransportState},
		{"CurrentTransportStatus", DirOut, &AVT_TransportStatus},
		{"CurrentSpeed", DirOut, &AVT_TransportPlaySpeed},
	},
}

type AVTArgIn_GetTransportSettings struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v GetTransportSettings"`
	InstanceID uint32
}
type AVTArgOut_GetTransportSettings struct {
	XMLName        xml.Name `xml:"u:GetTransportSettingsResponse"`
	XMLPrefix      string   `xml:"xmlns:u,attr"`
	PlayMode       string
	RecQualityMode string
}

var AVT_GetTransportSettings = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_GetTransportSettings{},
	ArgOut:  AVTArgOut_GetTransportSettings{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"PlayMode", DirOut, &AVT_CurrentPlayMode},
		{"RecQualityMode", DirOut, &AVT_CurrentRecordQualityMode},
	},
}

type AVTArgIn_Next struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v Next"`
	InstanceID uint32
}
type AVTArgOut_Next struct {
	XMLName   xml.Name `xml:"u:NextResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_Next = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_Next{},
	ArgOut:  AVTArgOut_Next{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
	},
}

type AVTArgIn_Pause struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v Pause"`
	InstanceID uint32
}
type AVTArgOut_Pause struct {
	XMLName   xml.Name `xml:"u:PauseResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_Pause = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_Pause{},
	ArgOut:  AVTArgOut_Pause{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
	},
}

type AVTArgIn_Play struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v Play"`
	InstanceID uint32
	Speed      string
}
type AVTArgOut_Play struct {
	XMLName   xml.Name `xml:"u:PlayResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_Play = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_Play{},
	ArgOut:  AVTArgOut_Play{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"Speed", DirIn, &AVT_TransportPlaySpeed},
	},
}

type AVTArgIn_Previous struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v Previous"`
	InstanceID uint32
}
type AVTArgOut_Previous struct {
	XMLName   xml.Name `xml:"u:PreviousResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_Previous = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_Previous{},
	ArgOut:  AVTArgOut_Previous{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
	},
}

type AVTArgIn_Record struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v Record"`
	InstanceID uint32
}
type AVTArgOut_Record struct {
	XMLName   xml.Name `xml:"u:RecordResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_Record = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_Record{},
	ArgOut:  AVTArgOut_Record{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
	},
}

type AVTArgIn_Seek struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v Seek"`
	InstanceID uint32
	Unit       string
	Target     string
}
type AVTArgOut_Seek struct {
	XMLName   xml.Name `xml:"u:SeekResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_Seek = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_Seek{},
	ArgOut:  AVTArgOut_Seek{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"Unit", DirIn, &AVT_A_ARG_TYPE_SeekMode},
		{"Target", DirIn, &AVT_A_ARG_TYPE_SeekTarget},
	},
}

type AVTArgIn_SetAVTransportURI struct {
	XMLName            xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v SetAVTransportURI"`
	InstanceID         uint32
	CurrentURI         string
	CurrentURIMetaData string
}
type AVTArgOut_SetAVTransportURI struct {
	XMLName   xml.Name `xml:"u:SetAVTransportURIResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_SetAVTransportURI = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_SetAVTransportURI{},
	ArgOut:  AVTArgOut_SetAVTransportURI{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"CurrentURI", DirIn, &AVT_AVTransportURI},
		{"CurrentURIMetaData", DirIn, &AVT_AVTransportURIMetaData},
	},
}

type AVTArgIn_SetNextAVTransportURI struct {
	XMLName         xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v SetNextAVTransportURI"`
	InstanceID      uint32
	NextURI         string
	NextURIMetaData string
}
type AVTArgOut_SetNextAVTransportURI struct {
	XMLName   xml.Name `xml:"u:SetNextAVTransportURIResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_SetNextAVTransportURI = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_SetNextAVTransportURI{},
	ArgOut:  AVTArgOut_SetNextAVTransportURI{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"NextURI", DirIn, &AVT_NextAVTransportURI},
		{"NextURIMetaData", DirIn, &AVT_NextAVTransportURIMetaData},
	},
}

type AVTArgIn_SetPlayMode struct {
	XMLName     xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v SetPlayMode"`
	InstanceID  uint32
	NewPlayMode string
}
type AVTArgOut_SetPlayMode struct {
	XMLName   xml.Name `xml:"u:SetPlayModeResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_SetPlayMode = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_SetPlayMode{},
	ArgOut:  AVTArgOut_SetPlayMode{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"NewPlayMode", DirIn, &AVT_CurrentPlayMode},
	},
}

type AVTArgIn_SetRecordQualityMode struct {
	XMLName              xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v SetRecordQualityMode"`
	InstanceID           uint32
	NewRecordQualityMode string
}
type AVTArgOut_SetRecordQualityMode struct {
	XMLName   xml.Name `xml:"u:SetRecordQualityModeResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_SetRecordQualityMode = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_SetRecordQualityMode{},
	ArgOut:  AVTArgOut_SetRecordQualityMode{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
		{"NewRecordQualityMode", DirIn, &AVT_CurrentRecordQualityMode},
	},
}

type AVTArgIn_Stop struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:serviceType:v Stop"`
	InstanceID uint32
}
type AVTArgOut_Stop struct {
	XMLName   xml.Name `xml:"u:StopResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var AVT_Stop = Action{
	Handler: nil,
	ArgIn:   AVTArgIn_Stop{},
	ArgOut:  AVTArgOut_Stop{XMLPrefix: soap.ActionNS},
	arguments: []Argument{
		{"InstanceID", DirIn, &AVT_A_ARG_TYPE_InstanceID},
	},
}
var AVTransportV1 = ActionMap{
	"GetCurrentTransportActions": &AVT_GetCurrentTransportActions,
	"GetDeviceCapabilities":      &AVT_GetDeviceCapabilities,
	"GetMediaInfo":               &AVT_GetMediaInfo,
	"GetPositionInfo":            &AVT_GetPositionInfo,
	"GetTransportInfo":           &AVT_GetTransportInfo,
	"GetTransportSettings":       &AVT_GetTransportSettings,
	"Next":                       &AVT_Next,
	"Pause":                      &AVT_Pause,
	"Play":                       &AVT_Play,
	"Previous":                   &AVT_Previous,
	"Record":                     &AVT_Record,
	"Seek":                       &AVT_Seek,
	"SetAVTransportURI":          &AVT_SetAVTransportURI,
	"SetNextAVTransportURI":      &AVT_SetNextAVTransportURI,
	"SetPlayMode":                &AVT_SetPlayMode,
	"SetRecordQualityMode":       &AVT_SetRecordQualityMode,
	"Stop":                       &AVT_Stop,
}
