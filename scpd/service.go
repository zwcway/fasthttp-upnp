package scpd

import (
	"encoding/xml"
)

type SpecVersion struct {
	Major uint `xml:"major"`
	Minor uint `xml:"minor"`
}

type Service struct {
	XMLName     xml.Name `xml:"service"`
	ServiceType string   `xml:"serviceType"`
	ServiceId   string   `xml:"serviceId"`
	SCPDURL     string
	ControlURL  string `xml:"controlURL"`
	EventSubURL string `xml:"eventSubURL"`
}

type Device struct {
	DeviceType       string `xml:"deviceType"`
	FriendlyName     string `xml:"friendlyName"`
	Manufacturer     string `xml:"manufacturer"`
	ManufacturerURL  string `xml:"manufacturerURL,omitempty"`
	ModelName        string `xml:"modelName"`
	ModelDescription string `xml:"modelDescription,omitempty"`
	ModelNumber      uint   `xml:"modelNumber,omitempty"`
	ModelURL         string `xml:"modelURL,omitempty"`
	SerialNumber     uint   `xml:"serialNumber,omitempty"`
	UDN              string
	ServiceList      []Service `xml:"serviceList>service"`
}

type DeviceDesc struct {
	SpecVersion SpecVersion `xml:"specVersion"`
	Device      Device      `xml:"device"`
}

type Argument struct {
	Name            string `xml:"name"`
	Direction       string `xml:"direction"`
	RelatedStateVar string `xml:"relatedStateVariable"`
}

type Action struct {
	Name      string      `xml:"name"`
	Arguments []*Argument `xml:"argumentList>argument"`
}

type SCPD struct {
	XMLName           xml.Name    `xml:"urn:schemas-upnp-org:service-1-0 scpd"`
	SpecVersion       SpecVersion `xml:"specVersion"`
	ActionList        []Action    `xml:"actionList>action"`
	ServiceStateTable []*Variable `xml:"serviceStateTable>stateVariable"`
}

type Variable struct {
	SendEvents    string      `xml:"sendEvents,attr"`
	Name          string      `xml:"name"`
	DataType      string      `xml:"dataType"`
	Default       string      `xml:"defaultValue,omitempty"`
	AllowedValues *[]string   `xml:"allowedValueList>allowedValue,omitempty"`
	AllowedRange  *AllowRange `xml:"allowedValueRange,omitempty"`
}

type AllowRange struct {
	Min  int `xml:"minimum"`
	Max  int `xml:"maximum"`
	Step int `xml:"step,omitempty"`
}
