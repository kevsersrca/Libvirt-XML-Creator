package models

import "encoding/xml"

/*
Domain CPU  Struct

*/
type DomainCPUModel struct {
	Fallback string `xml:"fallback,attr,omitempty"`
	Value    string `xml:",chardata"`
	VendorID string `xml:"vendor_id,attr,omitempty"`
}

/*
Domain CPU  TopologyStruct

*/
type DomainCPUTopology struct {
	Sockets int `xml:"sockets,attr,omitempty"`
	Cores   int `xml:"cores,attr,omitempty"`
	Threads int `xml:"threads,attr,omitempty"`
}

/*
Domain CPU  Feature Struct

*/
type DomainCPUFeature struct {
	Policy string `xml:"policy,attr,omitempty"`
	Name   string `xml:"name,attr,omitempty"`
}

/*
Domain CPU Cache Struct

*/
type DomainCPUCache struct {
	Level uint   `xml:"level,attr,omitempty"`
	Mode  string `xml:"mode,attr"`
}

/*
Domain CPU  Numa Struct

*/
type DomainNuma struct {
	Cell []DomainCell `xml:"cell"`
}

/*
Domain CPU  Cell Struct

*/
type DomainCell struct {
	ID        *uint               `xml:"id,attr"`
	CPUs      string              `xml:"cpus,attr"`
	Memory    string              `xml:"memory,attr"`
	Unit      string              `xml:"unit,attr,omitempty"`
	MemAccess string              `xml:"memAccess,attr,omitempty"`
	Distances DomainCellDistances `xml:"distances"`
}

/*
Domain CPU  Distance Struct

*/
type DomainCellDistances struct {
	Siblings []DomainCellSibling `xml:"sibling"`
}

/*
Domain CPU  Cell Sibling Struct

*/
type DomainCellSibling struct {
	ID    uint `xml:"id,attr"`
	Value uint `xml:"value,attr"`
}

/*
Domain VCPU  Struct

*/
type DomainVCPU struct {
	Placement string `xml:"placement,attr,omitempty"`
	CPUSet    string `xml:"cpuset,attr,omitempty"`
	Current   string `xml:"current,attr,omitempty"`
	Value     int    `xml:",chardata"`
}

/*
Domain VCPU  Struct

*/
type DomainVCPUsVCPU struct {
	Id           *uint  `xml:"id,attr"`
	Enabled      string `xml:"enabled,attr,omitempty"`
	Hotpluggable string `xml:"hotpluggable,attr,omitempty"`
	Order        *uint  `xml:"order,attr"`
}

/*
Domain VCPUS  Struct

*/
type DomainVCPUs struct {
	VCPU []DomainVCPUsVCPU `xml:"vcpu"`
}

/*
Domain CPU  Struct

*/
type DomainCPU struct {
	XMLName  xml.Name           `xml:"cpu"`
	Match    string             `xml:"match,attr,omitempty"`
	Mode     string             `xml:"mode,attr,omitempty"`
	Check    string             `xml:"check,attr,omitempty"`
	Model    DomainCPUModel     `xml:"model"`
	Vendor   string             `xml:"vendor,omitempty"`
	Topology DomainCPUTopology  `xml:"topology"`
	Cache    DomainCPUCache     `xml:"cache"`
	Features []DomainCPUFeature `xml:"feature"`
	Numa     *DomainNuma        `xml:"numa"`
}
