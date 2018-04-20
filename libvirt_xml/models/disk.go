package models

import "encoding/xml"

/*
Domain Disk Struct

TODO:
Auth
BackingStore
Geometry
BlockIO
Mirror
IOTune
ReadOnly
Shareable
Transient
Serial
WWN
Vendor
Product
Encryption
Boot
Alias
*/
type DomainDisk struct {
	XMLName  xml.Name          `xml:"disk"`
	Device   string            `xml:"device,attr,omitempty"`
	RawIO    string            `xml:"rawio,attr,omitempty"`
	SGIO     string            `xml:"sgio,attr,omitempty"`
	Snapshot string            `xml:"snapshot,attr,omitempty"`
	Driver   DomainDiskDriver  `xml:"driver"`
	Source   *DomainDiskSource `xml:"source"`
	Target   *DomainDiskTarget `xml:"target"`
	Address  *DomainAddress    `xml:"address"`
}

/*
Domain Disk Driver Struct
TODO:

*/
type DomainDiskDriver struct {
}

/*
Domain Disk Target Struct
TODO:
*/
type DomainDiskTarget struct {
}

/*
Domain Disk Source Struct
TODO:
*/
type DomainDiskSource struct {
}
