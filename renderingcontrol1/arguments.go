package renderingcontrol1

import "encoding/xml"

type ArgInGetBlueVideoBlackLevel struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetBlueVideoBlackLevel"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetBlueVideoBlackLevel struct {
	XMLName                    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetBlueVideoBlackLevelResponse"`
	CurrentBlueVideoBlackLevel uint16   `soap:"BlueVideoBlackLevel" range:"0,0,1"`
}
type ArgInGetBlueVideoGain struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetBlueVideoGain"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetBlueVideoGain struct {
	XMLName              xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetBlueVideoGainResponse"`
	CurrentBlueVideoGain uint16   `soap:"BlueVideoGain" range:"0,0,1"`
}
type ArgInGetBrightness struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetBrightness"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetBrightness struct {
	XMLName           xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetBrightnessResponse"`
	CurrentBrightness uint16   `soap:"Brightness" range:"0,0,1"`
}
type ArgInGetColorTemperature struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetColorTemperature"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetColorTemperature struct {
	XMLName                 xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetColorTemperatureResponse"`
	CurrentColorTemperature uint16   `soap:"ColorTemperature" range:"0,0,1"`
}
type ArgInGetContrast struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetContrast"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetContrast struct {
	XMLName         xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetContrastResponse"`
	CurrentContrast uint16   `soap:"Contrast" range:"0,0,1"`
}
type ArgInGetGreenVideoBlackLevel struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetGreenVideoBlackLevel"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetGreenVideoBlackLevel struct {
	XMLName                     xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetGreenVideoBlackLevelResponse"`
	CurrentGreenVideoBlackLevel uint16   `soap:"GreenVideoBlackLevel" range:"0,0,1"`
}
type ArgInGetGreenVideoGain struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetGreenVideoGain"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetGreenVideoGain struct {
	XMLName               xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetGreenVideoGainResponse"`
	CurrentGreenVideoGain uint16   `soap:"GreenVideoGain" range:"0,0,1"`
}
type ArgInGetHorizontalKeystone struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetHorizontalKeystone"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetHorizontalKeystone struct {
	XMLName                   xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetHorizontalKeystoneResponse"`
	CurrentHorizontalKeystone int16    `soap:"HorizontalKeystone" range:"0,0,1"`
}
type ArgInGetLoudness struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetLoudness"`
	Channel    string   `soap:"A_ARG_TYPE_Channel" allowed:"Master,LF,RF,CF,LFE,LS,RS,LFC,RFC,SD,SL,SR ,T,B"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetLoudness struct {
	XMLName         xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetLoudnessResponse"`
	CurrentLoudness bool     `soap:"Loudness"`
}
type ArgInGetMute struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetMute"`
	Channel    string   `soap:"A_ARG_TYPE_Channel" allowed:"Master,LF,RF,CF,LFE,LS,RS,LFC,RFC,SD,SL,SR ,T,B"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetMute struct {
	XMLName     xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetMuteResponse"`
	CurrentMute bool     `soap:"Mute"`
}
type ArgInGetRedVideoBlackLevel struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetRedVideoBlackLevel"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetRedVideoBlackLevel struct {
	XMLName                   xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetRedVideoBlackLevelResponse"`
	CurrentRedVideoBlackLevel uint16   `soap:"RedVideoBlackLevel" range:"0,0,1"`
}
type ArgInGetRedVideoGain struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetRedVideoGain"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetRedVideoGain struct {
	XMLName             xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetRedVideoGainResponse"`
	CurrentRedVideoGain uint16   `soap:"RedVideoGain" range:"0,0,1"`
}
type ArgInGetSharpness struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetSharpness"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetSharpness struct {
	XMLName          xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetSharpnessResponse"`
	CurrentSharpness uint16   `soap:"Sharpness" range:"0,0,1"`
}
type ArgInGetVerticalKeystone struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVerticalKeystone"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetVerticalKeystone struct {
	XMLName                 xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVerticalKeystoneResponse"`
	CurrentVerticalKeystone int16    `soap:"VerticalKeystone" range:"0,0,1"`
}
type ArgInGetVolume struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVolume"`
	Channel    string   `soap:"A_ARG_TYPE_Channel" allowed:"Master,LF,RF,CF,LFE,LS,RS,LFC,RFC,SD,SL,SR ,T,B"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetVolume struct {
	XMLName       xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVolumeResponse"`
	CurrentVolume uint16   `soap:"Volume" range:"0,0,1"`
}
type ArgInGetVolumeDB struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVolumeDB"`
	Channel    string   `soap:"A_ARG_TYPE_Channel" allowed:"Master,LF,RF,CF,LFE,LS,RS,LFC,RFC,SD,SL,SR ,T,B"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetVolumeDB struct {
	XMLName       xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVolumeDBResponse"`
	CurrentVolume int16    `soap:"VolumeDB" range:"0,0,1"`
}
type ArgInGetVolumeDBRange struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVolumeDBRange"`
	Channel    string   `soap:"A_ARG_TYPE_Channel" allowed:"Master,LF,RF,CF,LFE,LS,RS,LFC,RFC,SD,SL,SR ,T,B"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetVolumeDBRange struct {
	XMLName  xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 GetVolumeDBRangeResponse"`
	MaxValue int16    `soap:"VolumeDB" range:"0,0,1"`
	MinValue int16    `soap:"VolumeDB" range:"0,0,1"`
}
type ArgInListPresets struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 ListPresets"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutListPresets struct {
	XMLName               xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 ListPresetsResponse"`
	CurrentPresetNameList string   `soap:"PresetNameList"`
}
type ArgInSelectPreset struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SelectPreset"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
	PresetName string   `soap:"A_ARG_TYPE_PresetName" allowed:"FactoryDefaults,InstallationDefaults"`
}
type ArgOutSelectPreset struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SelectPresetResponse"`
}
type ArgInSetBlueVideoBlackLevel struct {
	XMLName                    xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetBlueVideoBlackLevel"`
	DesiredBlueVideoBlackLevel uint16   `soap:"BlueVideoBlackLevel" range:"0,0,1"`
	InstanceID                 uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetBlueVideoBlackLevel struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetBlueVideoBlackLevelResponse"`
}
type ArgInSetBlueVideoGain struct {
	XMLName              xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetBlueVideoGain"`
	DesiredBlueVideoGain uint16   `soap:"BlueVideoGain" range:"0,0,1"`
	InstanceID           uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetBlueVideoGain struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetBlueVideoGainResponse"`
}
type ArgInSetBrightness struct {
	XMLName           xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetBrightness"`
	DesiredBrightness uint16   `soap:"Brightness" range:"0,0,1"`
	InstanceID        uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetBrightness struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetBrightnessResponse"`
}
type ArgInSetColorTemperature struct {
	XMLName                 xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetColorTemperature"`
	DesiredColorTemperature uint16   `soap:"ColorTemperature" range:"0,0,1"`
	InstanceID              uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetColorTemperature struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetColorTemperatureResponse"`
}
type ArgInSetContrast struct {
	XMLName         xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetContrast"`
	DesiredContrast uint16   `soap:"Contrast" range:"0,0,1"`
	InstanceID      uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetContrast struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetContrastResponse"`
}
type ArgInSetGreenVideoBlackLevel struct {
	XMLName                     xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetGreenVideoBlackLevel"`
	DesiredGreenVideoBlackLevel uint16   `soap:"GreenVideoBlackLevel" range:"0,0,1"`
	InstanceID                  uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetGreenVideoBlackLevel struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetGreenVideoBlackLevelResponse"`
}
type ArgInSetGreenVideoGain struct {
	XMLName               xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetGreenVideoGain"`
	DesiredGreenVideoGain uint16   `soap:"GreenVideoGain" range:"0,0,1"`
	InstanceID            uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetGreenVideoGain struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetGreenVideoGainResponse"`
}
type ArgInSetHorizontalKeystone struct {
	XMLName                   xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetHorizontalKeystone"`
	DesiredHorizontalKeystone int16    `soap:"HorizontalKeystone" range:"0,0,1"`
	InstanceID                uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetHorizontalKeystone struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetHorizontalKeystoneResponse"`
}
type ArgInSetLoudness struct {
	XMLName         xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetLoudness"`
	Channel         string   `soap:"A_ARG_TYPE_Channel" allowed:"Master,LF,RF,CF,LFE,LS,RS,LFC,RFC,SD,SL,SR ,T,B"`
	DesiredLoudness bool     `soap:"Loudness"`
	InstanceID      uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetLoudness struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetLoudnessResponse"`
}
type ArgInSetMute struct {
	XMLName     xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetMute"`
	Channel     string   `soap:"A_ARG_TYPE_Channel" allowed:"Master,LF,RF,CF,LFE,LS,RS,LFC,RFC,SD,SL,SR ,T,B"`
	DesiredMute bool     `soap:"Mute"`
	InstanceID  uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetMute struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetMuteResponse"`
}
type ArgInSetRedVideoBlackLevel struct {
	XMLName                   xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetRedVideoBlackLevel"`
	DesiredRedVideoBlackLevel uint16   `soap:"RedVideoBlackLevel" range:"0,0,1"`
	InstanceID                uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetRedVideoBlackLevel struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetRedVideoBlackLevelResponse"`
}
type ArgInSetRedVideoGain struct {
	XMLName             xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetRedVideoGain"`
	DesiredRedVideoGain uint16   `soap:"RedVideoGain" range:"0,0,1"`
	InstanceID          uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetRedVideoGain struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetRedVideoGainResponse"`
}
type ArgInSetSharpness struct {
	XMLName          xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetSharpness"`
	DesiredSharpness uint16   `soap:"Sharpness" range:"0,0,1"`
	InstanceID       uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetSharpness struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetSharpnessResponse"`
}
type ArgInSetVerticalKeystone struct {
	XMLName                 xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetVerticalKeystone"`
	DesiredVerticalKeystone int16    `soap:"VerticalKeystone" range:"0,0,1"`
	InstanceID              uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetVerticalKeystone struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetVerticalKeystoneResponse"`
}
type ArgInSetVolume struct {
	XMLName       xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetVolume"`
	Channel       string   `soap:"A_ARG_TYPE_Channel" allowed:"Master,LF,RF,CF,LFE,LS,RS,LFC,RFC,SD,SL,SR ,T,B"`
	DesiredVolume uint16   `soap:"Volume" range:"0,0,1"`
	InstanceID    uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetVolume struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetVolumeResponse"`
}
type ArgInSetVolumeDB struct {
	XMLName       xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetVolumeDB"`
	Channel       string   `soap:"A_ARG_TYPE_Channel" allowed:"Master,LF,RF,CF,LFE,LS,RS,LFC,RFC,SD,SL,SR ,T,B"`
	DesiredVolume int16    `soap:"VolumeDB" range:"0,0,1"`
	InstanceID    uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetVolumeDB struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:RenderingControl:1 SetVolumeDBResponse"`
}
