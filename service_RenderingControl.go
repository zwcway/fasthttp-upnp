package upnp

import (
	"encoding/xml"
	"github.com/zwcway/fasthttp-upnp/scpd"
)

var RC_A_ARG_TYPE_Channel = scpd.Variable{
	Name:          "A_ARG_TYPE_Channel",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"Master", "LF", "RF", "CF", "LFE", "LS", "RS", "LFC", "RFC", "SD", "SL", "SR ", "T", "B"},
}
var RC_A_ARG_TYPE_InstanceID = scpd.Variable{
	Name:       "A_ARG_TYPE_InstanceID",
	DataType:   DataTypeUint32,
	SendEvents: BoolNo,
}
var RC_A_ARG_TYPE_PresetName = scpd.Variable{
	Name:          "A_ARG_TYPE_PresetName",
	DataType:      DataTypeStr,
	SendEvents:    BoolNo,
	AllowedValues: &[]string{"FactoryDefaults", "InstallationDefaults"},
}
var RC_BlueVideoBlackLevel = scpd.Variable{
	Name:         "BlueVideoBlackLevel",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_BlueVideoGain = scpd.Variable{
	Name:         "BlueVideoGain",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_Brightness = scpd.Variable{
	Name:         "Brightness",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_ColorTemperature = scpd.Variable{
	Name:         "ColorTemperature",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_Contrast = scpd.Variable{
	Name:         "Contrast",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_GreenVideoBlackLevel = scpd.Variable{
	Name:         "GreenVideoBlackLevel",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_GreenVideoGain = scpd.Variable{
	Name:         "GreenVideoGain",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_HorizontalKeystone = scpd.Variable{
	Name:         "HorizontalKeystone",
	DataType:     DataTypeInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_LastChange = scpd.Variable{
	Name:       "LastChange",
	DataType:   DataTypeStr,
	SendEvents: BoolYes,
}
var RC_Loudness = scpd.Variable{
	Name:       "Loudness",
	DataType:   DataTypeBool,
	SendEvents: BoolNo,
}
var RC_Mute = scpd.Variable{
	Name:       "Mute",
	DataType:   DataTypeBool,
	SendEvents: BoolNo,
}
var RC_PresetNameList = scpd.Variable{
	Name:       "PresetNameList",
	DataType:   DataTypeStr,
	SendEvents: BoolNo,
}
var RC_RedVideoBlackLevel = scpd.Variable{
	Name:         "RedVideoBlackLevel",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_RedVideoGain = scpd.Variable{
	Name:         "RedVideoGain",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_Sharpness = scpd.Variable{
	Name:         "Sharpness",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_VerticalKeystone = scpd.Variable{
	Name:         "VerticalKeystone",
	DataType:     DataTypeInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_Volume = scpd.Variable{
	Name:         "Volume",
	DataType:     DataTypeUInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 1},
}
var RC_VolumeDB = scpd.Variable{
	Name:         "VolumeDB",
	DataType:     DataTypeInt16,
	SendEvents:   BoolNo,
	AllowedRange: &scpd.AllowRange{0, 0, 0},
}

type RCArgIn_GetBlueVideoBlackLevel struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetBlueVideoBlackLevel"`
	InstanceID uint32
}
type RCArgOut_GetBlueVideoBlackLevel struct {
	XMLName                    xml.Name `xml:"u:GetBlueVideoBlackLevelResponse"`
	XMLPrefix                  string   `xml:"xmlns:u,attr"`
	CurrentBlueVideoBlackLevel uint16
}

var RC_GetBlueVideoBlackLevel = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetBlueVideoBlackLevel{},
	ArgOut:  &RCArgOut_GetBlueVideoBlackLevel{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentBlueVideoBlackLevel", DirOut, &RC_BlueVideoBlackLevel},
	},
}

type RCArgIn_GetBlueVideoGain struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetBlueVideoGain"`
	InstanceID uint32
}
type RCArgOut_GetBlueVideoGain struct {
	XMLName              xml.Name `xml:"u:GetBlueVideoGainResponse"`
	XMLPrefix            string   `xml:"xmlns:u,attr"`
	CurrentBlueVideoGain uint16
}

var RC_GetBlueVideoGain = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetBlueVideoGain{},
	ArgOut:  &RCArgOut_GetBlueVideoGain{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentBlueVideoGain", DirOut, &RC_BlueVideoGain},
	},
}

type RCArgIn_GetBrightness struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetBrightness"`
	InstanceID uint32
}
type RCArgOut_GetBrightness struct {
	XMLName           xml.Name `xml:"u:GetBrightnessResponse"`
	XMLPrefix         string   `xml:"xmlns:u,attr"`
	CurrentBrightness uint16
}

var RC_GetBrightness = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetBrightness{},
	ArgOut:  &RCArgOut_GetBrightness{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentBrightness", DirOut, &RC_Brightness},
	},
}

type RCArgIn_GetColorTemperature struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetColorTemperature"`
	InstanceID uint32
}
type RCArgOut_GetColorTemperature struct {
	XMLName                 xml.Name `xml:"u:GetColorTemperatureResponse"`
	XMLPrefix               string   `xml:"xmlns:u,attr"`
	CurrentColorTemperature uint16
}

var RC_GetColorTemperature = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetColorTemperature{},
	ArgOut:  &RCArgOut_GetColorTemperature{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentColorTemperature", DirOut, &RC_ColorTemperature},
	},
}

type RCArgIn_GetContrast struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetContrast"`
	InstanceID uint32
}
type RCArgOut_GetContrast struct {
	XMLName         xml.Name `xml:"u:GetContrastResponse"`
	XMLPrefix       string   `xml:"xmlns:u,attr"`
	CurrentContrast uint16
}

var RC_GetContrast = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetContrast{},
	ArgOut:  &RCArgOut_GetContrast{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentContrast", DirOut, &RC_Contrast},
	},
}

type RCArgIn_GetGreenVideoBlackLevel struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetGreenVideoBlackLevel"`
	InstanceID uint32
}
type RCArgOut_GetGreenVideoBlackLevel struct {
	XMLName                     xml.Name `xml:"u:GetGreenVideoBlackLevelResponse"`
	XMLPrefix                   string   `xml:"xmlns:u,attr"`
	CurrentGreenVideoBlackLevel uint16
}

var RC_GetGreenVideoBlackLevel = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetGreenVideoBlackLevel{},
	ArgOut:  &RCArgOut_GetGreenVideoBlackLevel{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentGreenVideoBlackLevel", DirOut, &RC_GreenVideoBlackLevel},
	},
}

type RCArgIn_GetGreenVideoGain struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetGreenVideoGain"`
	InstanceID uint32
}
type RCArgOut_GetGreenVideoGain struct {
	XMLName               xml.Name `xml:"u:GetGreenVideoGainResponse"`
	XMLPrefix             string   `xml:"xmlns:u,attr"`
	CurrentGreenVideoGain uint16
}

var RC_GetGreenVideoGain = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetGreenVideoGain{},
	ArgOut:  &RCArgOut_GetGreenVideoGain{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentGreenVideoGain", DirOut, &RC_GreenVideoGain},
	},
}

type RCArgIn_GetHorizontalKeystone struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetHorizontalKeystone"`
	InstanceID uint32
}
type RCArgOut_GetHorizontalKeystone struct {
	XMLName                   xml.Name `xml:"u:GetHorizontalKeystoneResponse"`
	XMLPrefix                 string   `xml:"xmlns:u,attr"`
	CurrentHorizontalKeystone int16
}

var RC_GetHorizontalKeystone = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetHorizontalKeystone{},
	ArgOut:  &RCArgOut_GetHorizontalKeystone{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentHorizontalKeystone", DirOut, &RC_HorizontalKeystone},
	},
}

type RCArgIn_GetLoudness struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetLoudness"`
	InstanceID uint32
	Channel    string
}
type RCArgOut_GetLoudness struct {
	XMLName         xml.Name `xml:"u:GetLoudnessResponse"`
	XMLPrefix       string   `xml:"xmlns:u,attr"`
	CurrentLoudness bool
}

var RC_GetLoudness = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetLoudness{},
	ArgOut:  &RCArgOut_GetLoudness{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"Channel", DirIn, &RC_A_ARG_TYPE_Channel},
		{"CurrentLoudness", DirOut, &RC_Loudness},
	},
}

type RCArgIn_GetMute struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetMute"`
	InstanceID uint32
	Channel    string
}
type RCArgOut_GetMute struct {
	XMLName     xml.Name `xml:"u:GetMuteResponse"`
	XMLPrefix   string   `xml:"xmlns:u,attr"`
	CurrentMute bool
}

var RC_GetMute = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetMute{},
	ArgOut:  &RCArgOut_GetMute{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"Channel", DirIn, &RC_A_ARG_TYPE_Channel},
		{"CurrentMute", DirOut, &RC_Mute},
	},
}

type RCArgIn_GetRedVideoBlackLevel struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetRedVideoBlackLevel"`
	InstanceID uint32
}
type RCArgOut_GetRedVideoBlackLevel struct {
	XMLName                   xml.Name `xml:"u:GetRedVideoBlackLevelResponse"`
	XMLPrefix                 string   `xml:"xmlns:u,attr"`
	CurrentRedVideoBlackLevel uint16
}

var RC_GetRedVideoBlackLevel = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetRedVideoBlackLevel{},
	ArgOut:  &RCArgOut_GetRedVideoBlackLevel{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentRedVideoBlackLevel", DirOut, &RC_RedVideoBlackLevel},
	},
}

type RCArgIn_GetRedVideoGain struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetRedVideoGain"`
	InstanceID uint32
}
type RCArgOut_GetRedVideoGain struct {
	XMLName             xml.Name `xml:"u:GetRedVideoGainResponse"`
	XMLPrefix           string   `xml:"xmlns:u,attr"`
	CurrentRedVideoGain uint16
}

var RC_GetRedVideoGain = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetRedVideoGain{},
	ArgOut:  &RCArgOut_GetRedVideoGain{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentRedVideoGain", DirOut, &RC_RedVideoGain},
	},
}

type RCArgIn_GetSharpness struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetSharpness"`
	InstanceID uint32
}
type RCArgOut_GetSharpness struct {
	XMLName          xml.Name `xml:"u:GetSharpnessResponse"`
	XMLPrefix        string   `xml:"xmlns:u,attr"`
	CurrentSharpness uint16
}

var RC_GetSharpness = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetSharpness{},
	ArgOut:  &RCArgOut_GetSharpness{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentSharpness", DirOut, &RC_Sharpness},
	},
}

type RCArgIn_GetVerticalKeystone struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVerticalKeystone"`
	InstanceID uint32
}
type RCArgOut_GetVerticalKeystone struct {
	XMLName                 xml.Name `xml:"u:GetVerticalKeystoneResponse"`
	XMLPrefix               string   `xml:"xmlns:u,attr"`
	CurrentVerticalKeystone int16
}

var RC_GetVerticalKeystone = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetVerticalKeystone{},
	ArgOut:  &RCArgOut_GetVerticalKeystone{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentVerticalKeystone", DirOut, &RC_VerticalKeystone},
	},
}

type RCArgIn_GetVolume struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVolume"`
	InstanceID uint32
	Channel    string
}
type RCArgOut_GetVolume struct {
	XMLName       xml.Name `xml:"u:GetVolumeResponse"`
	XMLPrefix     string   `xml:"xmlns:u,attr"`
	CurrentVolume uint16
}

var RC_GetVolume = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetVolume{},
	ArgOut:  &RCArgOut_GetVolume{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"Channel", DirIn, &RC_A_ARG_TYPE_Channel},
		{"CurrentVolume", DirOut, &RC_Volume},
	},
}

type RCArgIn_GetVolumeDB struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVolumeDB"`
	InstanceID uint32
	Channel    string
}
type RCArgOut_GetVolumeDB struct {
	XMLName       xml.Name `xml:"u:GetVolumeDBResponse"`
	XMLPrefix     string   `xml:"xmlns:u,attr"`
	CurrentVolume int16
}

var RC_GetVolumeDB = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetVolumeDB{},
	ArgOut:  &RCArgOut_GetVolumeDB{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"Channel", DirIn, &RC_A_ARG_TYPE_Channel},
		{"CurrentVolume", DirOut, &RC_VolumeDB},
	},
}

type RCArgIn_GetVolumeDBRange struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVolumeDBRange"`
	InstanceID uint32
	Channel    string
}
type RCArgOut_GetVolumeDBRange struct {
	XMLName   xml.Name `xml:"u:GetVolumeDBRangeResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
	MinValue  int16
	MaxValue  int16
}

var RC_GetVolumeDBRange = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_GetVolumeDBRange{},
	ArgOut:  &RCArgOut_GetVolumeDBRange{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"Channel", DirIn, &RC_A_ARG_TYPE_Channel},
		{"MinValue", DirOut, &RC_VolumeDB},
		{"MaxValue", DirOut, &RC_VolumeDB},
	},
}

type RCArgIn_ListPresets struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 ListPresets"`
	InstanceID uint32
}
type RCArgOut_ListPresets struct {
	XMLName               xml.Name `xml:"u:ListPresetsResponse"`
	XMLPrefix             string   `xml:"xmlns:u,attr"`
	CurrentPresetNameList string
}

var RC_ListPresets = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_ListPresets{},
	ArgOut:  &RCArgOut_ListPresets{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"CurrentPresetNameList", DirOut, &RC_PresetNameList},
	},
}

type RCArgIn_SelectPreset struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SelectPreset"`
	InstanceID uint32
	PresetName string
}
type RCArgOut_SelectPreset struct {
	XMLName   xml.Name `xml:"u:SelectPresetResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SelectPreset = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SelectPreset{},
	ArgOut:  &RCArgOut_SelectPreset{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"PresetName", DirIn, &RC_A_ARG_TYPE_PresetName},
	},
}

type RCArgIn_SetBlueVideoBlackLevel struct {
	XMLName                    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetBlueVideoBlackLevel"`
	InstanceID                 uint32
	DesiredBlueVideoBlackLevel uint16
}
type RCArgOut_SetBlueVideoBlackLevel struct {
	XMLName   xml.Name `xml:"u:SetBlueVideoBlackLevelResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetBlueVideoBlackLevel = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetBlueVideoBlackLevel{},
	ArgOut:  &RCArgOut_SetBlueVideoBlackLevel{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredBlueVideoBlackLevel", DirIn, &RC_BlueVideoBlackLevel},
	},
}

type RCArgIn_SetBlueVideoGain struct {
	XMLName              xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetBlueVideoGain"`
	InstanceID           uint32
	DesiredBlueVideoGain uint16
}
type RCArgOut_SetBlueVideoGain struct {
	XMLName   xml.Name `xml:"u:SetBlueVideoGainResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetBlueVideoGain = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetBlueVideoGain{},
	ArgOut:  &RCArgOut_SetBlueVideoGain{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredBlueVideoGain", DirIn, &RC_BlueVideoGain},
	},
}

type RCArgIn_SetBrightness struct {
	XMLName           xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetBrightness"`
	InstanceID        uint32
	DesiredBrightness uint16
}
type RCArgOut_SetBrightness struct {
	XMLName   xml.Name `xml:"u:SetBrightnessResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetBrightness = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetBrightness{},
	ArgOut:  &RCArgOut_SetBrightness{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredBrightness", DirIn, &RC_Brightness},
	},
}

type RCArgIn_SetColorTemperature struct {
	XMLName                 xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetColorTemperature"`
	InstanceID              uint32
	DesiredColorTemperature uint16
}
type RCArgOut_SetColorTemperature struct {
	XMLName   xml.Name `xml:"u:SetColorTemperatureResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetColorTemperature = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetColorTemperature{},
	ArgOut:  &RCArgOut_SetColorTemperature{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredColorTemperature", DirIn, &RC_ColorTemperature},
	},
}

type RCArgIn_SetContrast struct {
	XMLName         xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetContrast"`
	InstanceID      uint32
	DesiredContrast uint16
}
type RCArgOut_SetContrast struct {
	XMLName   xml.Name `xml:"u:SetContrastResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetContrast = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetContrast{},
	ArgOut:  &RCArgOut_SetContrast{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredContrast", DirIn, &RC_Contrast},
	},
}

type RCArgIn_SetGreenVideoBlackLevel struct {
	XMLName                     xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetGreenVideoBlackLevel"`
	InstanceID                  uint32
	DesiredGreenVideoBlackLevel uint16
}
type RCArgOut_SetGreenVideoBlackLevel struct {
	XMLName   xml.Name `xml:"u:SetGreenVideoBlackLevelResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetGreenVideoBlackLevel = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetGreenVideoBlackLevel{},
	ArgOut:  &RCArgOut_SetGreenVideoBlackLevel{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredGreenVideoBlackLevel", DirIn, &RC_GreenVideoBlackLevel},
	},
}

type RCArgIn_SetGreenVideoGain struct {
	XMLName               xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetGreenVideoGain"`
	InstanceID            uint32
	DesiredGreenVideoGain uint16
}
type RCArgOut_SetGreenVideoGain struct {
	XMLName   xml.Name `xml:"u:SetGreenVideoGainResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetGreenVideoGain = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetGreenVideoGain{},
	ArgOut:  &RCArgOut_SetGreenVideoGain{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredGreenVideoGain", DirIn, &RC_GreenVideoGain},
	},
}

type RCArgIn_SetHorizontalKeystone struct {
	XMLName                   xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetHorizontalKeystone"`
	InstanceID                uint32
	DesiredHorizontalKeystone int16
}
type RCArgOut_SetHorizontalKeystone struct {
	XMLName   xml.Name `xml:"u:SetHorizontalKeystoneResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetHorizontalKeystone = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetHorizontalKeystone{},
	ArgOut:  &RCArgOut_SetHorizontalKeystone{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredHorizontalKeystone", DirIn, &RC_HorizontalKeystone},
	},
}

type RCArgIn_SetLoudness struct {
	XMLName         xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetLoudness"`
	InstanceID      uint32
	Channel         string
	DesiredLoudness bool
}
type RCArgOut_SetLoudness struct {
	XMLName   xml.Name `xml:"u:SetLoudnessResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetLoudness = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetLoudness{},
	ArgOut:  &RCArgOut_SetLoudness{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"Channel", DirIn, &RC_A_ARG_TYPE_Channel},
		{"DesiredLoudness", DirIn, &RC_Loudness},
	},
}

type RCArgIn_SetMute struct {
	XMLName     xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetMute"`
	InstanceID  uint32
	Channel     string
	DesiredMute bool
}
type RCArgOut_SetMute struct {
	XMLName   xml.Name `xml:"u:SetMuteResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetMute = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetMute{},
	ArgOut:  &RCArgOut_SetMute{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"Channel", DirIn, &RC_A_ARG_TYPE_Channel},
		{"DesiredMute", DirIn, &RC_Mute},
	},
}

type RCArgIn_SetRedVideoBlackLevel struct {
	XMLName                   xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetRedVideoBlackLevel"`
	InstanceID                uint32
	DesiredRedVideoBlackLevel uint16
}
type RCArgOut_SetRedVideoBlackLevel struct {
	XMLName   xml.Name `xml:"u:SetRedVideoBlackLevelResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetRedVideoBlackLevel = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetRedVideoBlackLevel{},
	ArgOut:  &RCArgOut_SetRedVideoBlackLevel{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredRedVideoBlackLevel", DirIn, &RC_RedVideoBlackLevel},
	},
}

type RCArgIn_SetRedVideoGain struct {
	XMLName             xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetRedVideoGain"`
	InstanceID          uint32
	DesiredRedVideoGain uint16
}
type RCArgOut_SetRedVideoGain struct {
	XMLName   xml.Name `xml:"u:SetRedVideoGainResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetRedVideoGain = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetRedVideoGain{},
	ArgOut:  &RCArgOut_SetRedVideoGain{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredRedVideoGain", DirIn, &RC_RedVideoGain},
	},
}

type RCArgIn_SetSharpness struct {
	XMLName          xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetSharpness"`
	InstanceID       uint32
	DesiredSharpness uint16
}
type RCArgOut_SetSharpness struct {
	XMLName   xml.Name `xml:"u:SetSharpnessResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetSharpness = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetSharpness{},
	ArgOut:  &RCArgOut_SetSharpness{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredSharpness", DirIn, &RC_Sharpness},
	},
}

type RCArgIn_SetVerticalKeystone struct {
	XMLName                 xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetVerticalKeystone"`
	InstanceID              uint32
	DesiredVerticalKeystone int16
}
type RCArgOut_SetVerticalKeystone struct {
	XMLName   xml.Name `xml:"u:SetVerticalKeystoneResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetVerticalKeystone = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetVerticalKeystone{},
	ArgOut:  &RCArgOut_SetVerticalKeystone{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"DesiredVerticalKeystone", DirIn, &RC_VerticalKeystone},
	},
}

type RCArgIn_SetVolume struct {
	XMLName       xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetVolume"`
	InstanceID    uint32
	Channel       string
	DesiredVolume uint16
}
type RCArgOut_SetVolume struct {
	XMLName   xml.Name `xml:"u:SetVolumeResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetVolume = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetVolume{},
	ArgOut:  &RCArgOut_SetVolume{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"Channel", DirIn, &RC_A_ARG_TYPE_Channel},
		{"DesiredVolume", DirIn, &RC_Volume},
	},
}

type RCArgIn_SetVolumeDB struct {
	XMLName       xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetVolumeDB"`
	InstanceID    uint32
	Channel       string
	DesiredVolume int16
}
type RCArgOut_SetVolumeDB struct {
	XMLName   xml.Name `xml:"u:SetVolumeDBResponse"`
	XMLPrefix string   `xml:"xmlns:u,attr"`
}

var RC_SetVolumeDB = Action{
	Handler: nil,
	ArgIn:   &RCArgIn_SetVolumeDB{},
	ArgOut:  &RCArgOut_SetVolumeDB{XMLPrefix: ServiceNS(ServiceName_RenderingControl, 1)},
	arguments: []Argument{
		{"InstanceID", DirIn, &RC_A_ARG_TYPE_InstanceID},
		{"Channel", DirIn, &RC_A_ARG_TYPE_Channel},
		{"DesiredVolume", DirIn, &RC_VolumeDB},
	},
}
var RenderingControlV1 = ActionMap{
	"GetBlueVideoBlackLevel":  &RC_GetBlueVideoBlackLevel,
	"GetBlueVideoGain":        &RC_GetBlueVideoGain,
	"GetBrightness":           &RC_GetBrightness,
	"GetColorTemperature":     &RC_GetColorTemperature,
	"GetContrast":             &RC_GetContrast,
	"GetGreenVideoBlackLevel": &RC_GetGreenVideoBlackLevel,
	"GetGreenVideoGain":       &RC_GetGreenVideoGain,
	"GetHorizontalKeystone":   &RC_GetHorizontalKeystone,
	"GetLoudness":             &RC_GetLoudness,
	"GetMute":                 &RC_GetMute,
	"GetRedVideoBlackLevel":   &RC_GetRedVideoBlackLevel,
	"GetRedVideoGain":         &RC_GetRedVideoGain,
	"GetSharpness":            &RC_GetSharpness,
	"GetVerticalKeystone":     &RC_GetVerticalKeystone,
	"GetVolume":               &RC_GetVolume,
	"GetVolumeDB":             &RC_GetVolumeDB,
	"GetVolumeDBRange":        &RC_GetVolumeDBRange,
	"ListPresets":             &RC_ListPresets,
	"SelectPreset":            &RC_SelectPreset,
	"SetBlueVideoBlackLevel":  &RC_SetBlueVideoBlackLevel,
	"SetBlueVideoGain":        &RC_SetBlueVideoGain,
	"SetBrightness":           &RC_SetBrightness,
	"SetColorTemperature":     &RC_SetColorTemperature,
	"SetContrast":             &RC_SetContrast,
	"SetGreenVideoBlackLevel": &RC_SetGreenVideoBlackLevel,
	"SetGreenVideoGain":       &RC_SetGreenVideoGain,
	"SetHorizontalKeystone":   &RC_SetHorizontalKeystone,
	"SetLoudness":             &RC_SetLoudness,
	"SetMute":                 &RC_SetMute,
	"SetRedVideoBlackLevel":   &RC_SetRedVideoBlackLevel,
	"SetRedVideoGain":         &RC_SetRedVideoGain,
	"SetSharpness":            &RC_SetSharpness,
	"SetVerticalKeystone":     &RC_SetVerticalKeystone,
	"SetVolume":               &RC_SetVolume,
	"SetVolumeDB":             &RC_SetVolumeDB,
}
