package models

import "encoding/xml"

/*
Domain Mem Balloon Struct
TODO:
Alias
Address
*/

type DomainMemBalloon struct {
	XMLName     xml.Name                `xml:"memballoon"`
	Model       string                  `xml:"model,attr"`
	AutoDeflate string                  `xml:"autodeflate,attr,omitempty"`
	Driver      *DomainMemBalloonDriver `xml:"driver"`
	Stats       *DomainMemBalloonStats  `xml:"stats"`
}

type DomainMemBalloonDriver struct {
	IOMMU string `xml:"iommu,attr,omitempty"`
	ATS   string `xml:"ats,attr,omitempty"`
}
type DomainMemBalloonStats struct {
	Period uint `xml:"period,attr"`
}

/*
Domain Memory Struct
TODO:

*/
type DomainMemory struct {
	Value    uint   `xml:",chardata"`
	Unit     string `xml:"unit,attr,omitempty"`
	DumpCore string `xml:"dumpCore,attr,omitempty"`
}

/*
Domain Current Memory Struct
TODO:

*/
type DomainCurrentMemory struct {
	Value uint   `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
}

/*
Domain domain Max Memory Struct
TODO:

*/
type DomainMaxMemory struct {
	Value uint   `xml:",chardata"`
	Unit  string `xml:"unit,attr,omitempty"`
	Slots uint   `xml:"slots,attr,omitempty"`
}
