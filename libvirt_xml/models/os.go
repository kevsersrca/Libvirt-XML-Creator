package models

/*

<os>
	<type arch='x86_64' machine='pc-i440fx-xenial'>hvm</type>
	<boot dev='hd'/>
</os>

TODO:
InitEnv
Loader
NVRam
ACPI
BootMenu
BIOS
SMBios
*/
type DomainOS struct {
	Type        DomainOSType       `xml:"type"`
	Init        string             `xml:"init,omitempty"`
	InitArgs    []string           `xml:"initarg"`
	InitDir     string             `xml:"initdir,omitempty"`
	InitUser    string             `xml:"inituser,omitempty"`
	InitGroup   string             `xml:"initgroup,omitempty"`
	Kernel      string             `xml:"kernel,omitempty"`
	Initrd      string             `xml:"initrd,omitempty"`
	Cmdline     string             `xml:"cmdline,omitempty"`
	DTB         string             `xml:"dtb,omitempty"`
	BootDevices []DomainBootDevice `xml:"boot"`
}

type DomainBootDevice struct {
	DomainBootDeviceName string `xml:"dev,attr"`
}

type DomainOSType struct {
	Arch             string `xml:"arch,attr,omitempty"`
	Machine          string `xml:"machine,attr,omitempty"`
	DomainOsTypeName string `xml:",innerxml"`
}
