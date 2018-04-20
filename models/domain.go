package models

import "encoding/xml"

/*
TODO
BlockIOTune
MemoryTune
MemoryBacking
IOThreads
IOThreadIDs
CPUTune
NUMATune
Resource
SysInfo
IDMap
Perf
SecLabel
QEMUCommandline
LXCNamespace
VMWareDataCenterPath
KeyWrap
*/
type Domain struct {
	XMLName        xml.Name            `xml:"domain"`
	Type           string              `xml:"type,attr,omitempty"`
	ID             *int                `xml:"id,attr"`
	Name           string              `xml:"name,omitempty"`
	UUID           string              `xml:"uuid,omitempty"`
	Title          string              `xml:"title"`
	Description    string              `xml:"description"`
	Metadata       DomainMetadata      `xml:"metadata"`
	MaximumMemory  DomainMaxMemory     `xml:"maxMemory"`
	Memory         DomainMemory        `xml:"memory"`
	CurrentMemory  DomainCurrentMemory `xml:"currentMemory"`
	VCPU           DomainVCPU          `xml:"vcpu"`
	VCPUs          DomainVCPUs         `xml:"vcpus"`
	Bootloader     string              `xml:"bootloader,omitempty"`
	BootloaderArgs string              `xml:"bootloader_args,omitempty"`
	OS             DomainOS            `xml:"os"`
	Features       DomainFeatureList   `xml:"features"`
	CPU            DomainCPU           `xml:"cpu"`
	Clock          DomainClock         `xml:"clock"`
	OnPoweroff     string              `xml:"on_poweroff,omitempty"`
	OnReboot       string              `xml:"on_reboot,omitempty"`
	OnCrash        string              `xml:"on_crash,omitempty"`
	PM             DomainPM            `xml:"pm"`
	Devices        DomainDevices       `xml:"devices"`
}

/*
Domain Meta Data Struct
*/
type DomainMetadata struct {
	DomainMetaDataInner string `xml:",innerxml"`
}
