package models

import "encoding/xml"

/*
TODO:
Driver
PCI
USB
VirtIOSerial
Alias
*/
type DomainController struct {
	XMLName xml.Name      `xml:"controller"`
	Type    string        `xml:"type,attr"`
	Index   *uint         `xml:"index,attr"`
	Model   string        `xml:"model,attr,omitempty"`
	Address DomainAddress `xml:"address"`
}

/*
Domain Adress Struct

TODO:
PCI
Drive
VirtioSerial
CCID
USB
SpaprVIO
VirtioS390
CCW
VirtioMMIO
ISA
DIMM
*/
type DomainAddress struct {
}
