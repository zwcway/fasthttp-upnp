package avtransport1

import "encoding/xml"

type ArgInGetCurrentTransportActions struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetCurrentTransportActions"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetCurrentTransportActions struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetCurrentTransportActionsResponse"`
	Actions string   `soap:"CurrentTransportActions"`
}
type ArgInGetDeviceCapabilities struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetDeviceCapabilities"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetDeviceCapabilities struct {
	XMLName         xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetDeviceCapabilitiesResponse"`
	PlayMedia       string   `soap:"PossiblePlaybackStorageMedia"`
	RecMedia        string   `soap:"PossibleRecordStorageMedia"`
	RecQualityModes string   `soap:"PossibleRecordQualityModes"`
}
type ArgInGetMediaInfo struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetMediaInfo"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetMediaInfo struct {
	XMLName            xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetMediaInfoResponse"`
	CurrentURI         string   `soap:"AVTransportURI"`
	CurrentURIMetaData string   `soap:"AVTransportURIMetaData"`
	MediaDuration      string   `soap:"CurrentMediaDuration"`
	NextURI            string   `soap:"NextAVTransportURI"`
	NextURIMetaData    string   `soap:"NextAVTransportURIMetaData"`
	NrTracks           uint32   `soap:"NumberOfTracks" range:"0,0,1"`
	PlayMedium         string   `soap:"PlaybackStorageMedium" allowed:"UNKNOWN,DV,MINI-DV,VHS,W-VHS,S-VHS,D-VHS,VHSC,VIDEO8,HI8,CD-ROM,CD-DA,CD-R,CD-RW,VIDEO-CD,SACD,MD-AUDIO,MD-PICTURE,DVD-ROM,DVD-VIDEO,DVD-R,DVD+RW,DVD-RW,DVD-RAM,DVD-AUDIO,DAT,LD,HDD,MICRO-MV,NETWORK,NONE,NOT_IMPLEMENTED"`
	RecordMedium       string   `soap:"RecordStorageMedium" allowed:"UNKNOWN,DV,MINI-DV,VHS,W-VHS,S-VHS,D-VHS,VHSC,VIDEO8,HI8,CD-ROM,CD-DA,CD-R,CD-RW,VIDEO-CD,SACD,MD-AUDIO,MD-PICTURE,DVD-ROM,DVD-VIDEO,DVD-R,DVD+RW,DVD-RW,DVD-RAM,DVD-AUDIO,DAT,LD,HDD,MICRO-MV,NETWORK,NONE,NOT_IMPLEMENTED"`
	WriteStatus        string   `soap:"RecordMediumWriteStatus" allowed:"WRITABLE,PROTECTED,NOT_WRITABLE,UNKNOWN,NOT_IMPLEMENTED"`
}
type ArgInGetPositionInfo struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetPositionInfo"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetPositionInfo struct {
	XMLName       xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetPositionInfoResponse"`
	AbsCount      int32    `soap:"AbsoluteCounterPosition"`
	AbsTime       string   `soap:"AbsoluteTimePosition"`
	RelCount      int32    `soap:"RelativeCounterPosition"`
	RelTime       string   `soap:"RelativeTimePosition"`
	Track         uint32   `soap:"CurrentTrack" range:"0,0,1"`
	TrackDuration string   `soap:"CurrentTrackDuration"`
	TrackMetaData string   `soap:"CurrentTrackMetaData"`
	TrackURI      string   `soap:"CurrentTrackURI"`
}
type ArgInGetTransportInfo struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetTransportInfo"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetTransportInfo struct {
	XMLName                xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetTransportInfoResponse"`
	CurrentSpeed           string   `soap:"TransportPlaySpeed" allowed:"1"`
	CurrentTransportState  string   `soap:"TransportState" allowed:"STOPPED,PAUSED_PLAYBACK,PAUSED_RECORDING,PLAYING,RECORDING,TRANSITIONING,NO_MEDIA_PRESENT"`
	CurrentTransportStatus string   `soap:"TransportStatus" allowed:"OK,ERROR_OCCURRED"`
}
type ArgInGetTransportSettings struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetTransportSettings"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutGetTransportSettings struct {
	XMLName        xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetTransportSettingsResponse"`
	PlayMode       string   `soap:"CurrentPlayMode NORMAL" allowed:"NORMAL,SHUFFLE,REPEAT_ONE,REPEAT_ALL,RANDOM,DIRECT_1,INTRO"`
	RecQualityMode string   `soap:"CurrentRecordQualityMode" allowed:"0:EP,1:LP,2:SP,0:BASIC,1:MEDIUM,2:HIGH,NOT_IMPLEMENTED"`
}
type ArgInNext struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Next"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutNext struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 NextResponse"`
}
type ArgInPause struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Pause"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutPause struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 PauseResponse"`
}
type ArgInPlay struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Play"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
	Speed      string   `soap:"TransportPlaySpeed" allowed:"1"`
}
type ArgOutPlay struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 PlayResponse"`
}
type ArgInPrevious struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Previous"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutPrevious struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 PreviousResponse"`
}
type ArgInRecord struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Record"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutRecord struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 RecordResponse"`
}
type ArgInSeek struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Seek"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
	Target     string   `soap:"A_ARG_TYPE_SeekTarget"`
	Unit       string   `soap:"A_ARG_TYPE_SeekMode" allowed:"ABS_TIME,REL_TIME,ABS_COUNT,REL_COUNT,TRACK_NR,CHANNEL_FREQ,TAPE-INDEX,FRAME"`
}
type ArgOutSeek struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 SeekResponse"`
}
type ArgInSetAVTransportURI struct {
	XMLName            xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetAVTransportURI"`
	CurrentURI         string   `soap:"AVTransportURI"`
	CurrentURIMetaData string   `soap:"AVTransportURIMetaData"`
	InstanceID         uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutSetAVTransportURI struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetAVTransportURIResponse"`
}
type ArgInSetNextAVTransportURI struct {
	XMLName         xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetNextAVTransportURI"`
	InstanceID      uint32   `soap:"A_ARG_TYPE_InstanceID"`
	NextURI         string   `soap:"NextAVTransportURI"`
	NextURIMetaData string   `soap:"NextAVTransportURIMetaData"`
}
type ArgOutSetNextAVTransportURI struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetNextAVTransportURIResponse"`
}
type ArgInSetPlayMode struct {
	XMLName     xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetPlayMode"`
	InstanceID  uint32   `soap:"A_ARG_TYPE_InstanceID"`
	NewPlayMode string   `soap:"CurrentPlayMode NORMAL" allowed:"NORMAL,SHUFFLE,REPEAT_ONE,REPEAT_ALL,RANDOM,DIRECT_1,INTRO"`
}
type ArgOutSetPlayMode struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetPlayModeResponse"`
}
type ArgInSetRecordQualityMode struct {
	XMLName              xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetRecordQualityMode"`
	InstanceID           uint32   `soap:"A_ARG_TYPE_InstanceID"`
	NewRecordQualityMode string   `soap:"CurrentRecordQualityMode" allowed:"0:EP,1:LP,2:SP,0:BASIC,1:MEDIUM,2:HIGH,NOT_IMPLEMENTED"`
}
type ArgOutSetRecordQualityMode struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetRecordQualityModeResponse"`
}
type ArgInStop struct {
	XMLName    xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Stop"`
	InstanceID uint32   `soap:"A_ARG_TYPE_InstanceID"`
}
type ArgOutStop struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 StopResponse"`
}
