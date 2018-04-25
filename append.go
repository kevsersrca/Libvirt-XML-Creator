package bin

import (
	"fmt"
	"io/ioutil"
	"os"

	m "github.com/kevsersrca/libvirt-xml/models"
)

var LibvirtDomainName

var LibvirtDomainFullPath

var LibvirtDomainConfigurationPath = "/etc/libvirt"

var LibvirtDomainType = "qemu"


type LibvirtXml struct {
	LibvirtDomainName
	LibvirtDomainFullPath
	LibvirtDomainConfigurationPath
	LibvirtDomainType
}

func New(domainName string) *LibvirtXml{
	libvirt := new(LibvirtXml)
	libvirt.LibvirtDomainName = domainName
	libvirt.LibvirtDomainFullPath = returnFullPath(libvirt)
	return libvirt
}


func (l *LibvirtXml) UnmarshallDomain(filename string) *m.Domain {

}

func returnFullPath(libvirt *LibvirtXml) string{
	return fmt.Sprintf("%s/%s/%s.xml", libvirt.LibvirtDomainConfigurationPath,libvirt.LibvirtDomainType, libvirt.LibvirtDomainName)
}
