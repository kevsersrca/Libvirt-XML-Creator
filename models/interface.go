package models

import "encoding/xml"

type DomainInterfaceMAC struct {
	Address string `xml:"address,attr"`
}

type DomainInterfaceModel struct {
	Type string `xml:"type,attr"`
}

type DomainInterfaceSource struct {
	User      *DomainInterfaceSourceUser     `xml:"-"`
	Ethernet  *DomainInterfaceSourceEthernet `xml:"-"`
	VHostUser *DomainChardevSource           `xml:"-"`
	Server    *DomainInterfaceSourceServer   `xml:"-"`
	Client    *DomainInterfaceSourceClient   `xml:"-"`
	MCast     *DomainInterfaceSourceMCast    `xml:"-"`
	Network   *DomainInterfaceSourceNetwork  `xml:"-"`
	Bridge    *DomainInterfaceSourceBridge   `xml:"-"`
	Internal  *DomainInterfaceSourceInternal `xml:"-"`
	Direct    *DomainInterfaceSourceDirect   `xml:"-"`
	Hostdev   *DomainInterfaceSourceHostdev  `xml:"-"`
	UDP       *DomainInterfaceSourceUDP      `xml:"-"`
}

type DomainInterfaceSourceUser struct {
}

type DomainInterfaceSourceEthernet struct {
	IP    []DomainInterfaceIP    `xml:"ip"`
	Route []DomainInterfaceRoute `xml:"route"`
}

type DomainInterfaceSourceServer struct {
	Address string                      `xml:"address,attr,omitempty"`
	Port    uint                        `xml:"port,attr,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local"`
}

type DomainInterfaceSourceClient struct {
	Address string                      `xml:"address,attr,omitempty"`
	Port    uint                        `xml:"port,attr,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local"`
}

type DomainInterfaceSourceMCast struct {
	Address string                      `xml:"address,attr,omitempty"`
	Port    uint                        `xml:"port,attr,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local"`
}

type DomainInterfaceSourceNetwork struct {
	Network   string `xml:"network,attr,omitempty"`
	PortGroup string `xml:"portgroup,attr,omitempty"`
}

type DomainInterfaceSourceBridge struct {
	Bridge string `xml:"bridge,attr"`
}

type DomainInterfaceSourceInternal struct {
	Name string `xml:"name,attr,omitempty"`
}

type DomainInterfaceSourceDirect struct {
	Dev  string `xml:"dev,attr,omitempty"`
	Mode string `xml:"mode,attr,omitempty"`
}

type DomainInterfaceSourceHostdev struct {
	PCI *DomainHostdevSubsysPCISource `xml:"-"`
	USB *DomainHostdevSubsysUSBSource `xml:"-"`
}

type DomainInterfaceSourceUDP struct {
	Address string                      `xml:"address,attr,omitempty"`
	Port    uint                        `xml:"port,attr,omitempty"`
	Local   *DomainInterfaceSourceLocal `xml:"local"`
}

type DomainInterfaceSourceLocal struct {
	Address string `xml:"address,attr,omitempty"`
	Port    uint   `xml:"port,attr,omitempty"`
}

type DomainInterfaceTarget struct {
	Dev string `xml:"dev,attr"`
}

type DomainInterfaceLink struct {
	State string `xml:"state,attr"`
}

type DomainDeviceBoot struct {
	Order    uint   `xml:"order,attr"`
	LoadParm string `xml:"loadparm,attr,omitempty"`
}

type DomainInterfaceScript struct {
	Path string `xml:"path,attr"`
}

type DomainInterfaceDriver struct {
	Name        string                      `xml:"name,attr,omitempty"`
	TXMode      string                      `xml:"txmode,attr,omitempty"`
	IOEventFD   string                      `xml:"ioeventfd,attr,omitempty"`
	EventIDX    string                      `xml:"event_idx,attr,omitempty"`
	Queues      uint                        `xml:"queues,attr,omitempty"`
	RXQueueSize uint                        `xml:"rx_queue_size,attr,omitempty"`
	TXQueueSize uint                        `xml:"tx_queue_size,attr,omitempty"`
	IOMMU       string                      `xml:"iommu,attr,omitempty"`
	ATS         string                      `xml:"ats,attr,omitempty"`
	Host        *DomainInterfaceDriverHost  `xml:"host"`
	Guest       *DomainInterfaceDriverGuest `xml:"guest"`
}

type DomainInterfaceDriverHost struct {
	CSum     string `xml:"csum,attr,omitempty"`
	GSO      string `xml:"gso,attr,omitempty"`
	TSO4     string `xml:"tso4,attr,omitempty"`
	TSO6     string `xml:"tso6,attr,omitempty"`
	ECN      string `xml:"ecn,attr,omitempty"`
	UFO      string `xml:"ufo,attr,omitempty"`
	MrgRXBuf string `xml:"mrg_rxbuf,attr,omitempty"`
}

type DomainInterfaceDriverGuest struct {
	CSum string `xml:"csum,attr,omitempty"`
	TSO4 string `xml:"tso4,attr,omitempty"`
	TSO6 string `xml:"tso6,attr,omitempty"`
	ECN  string `xml:"ecn,attr,omitempty"`
	UFO  string `xml:"ufo,attr,omitempty"`
}

type DomainInterfaceVirtualPort struct {
	Params *DomainInterfaceVirtualPortParams `xml:"parameters"`
}

type DomainInterfaceVirtualPortParams struct {
	Any          *DomainInterfaceVirtualPortParamsAny          `xml:"-"`
	VEPA8021QBG  *DomainInterfaceVirtualPortParamsVEPA8021QBG  `xml:"-"`
	VNTag8011QBH *DomainInterfaceVirtualPortParamsVNTag8021QBH `xml:"-"`
	OpenVSwitch  *DomainInterfaceVirtualPortParamsOpenVSwitch  `xml:"-"`
	MidoNet      *DomainInterfaceVirtualPortParamsMidoNet      `xml:"-"`
}

type DomainInterfaceVirtualPortParamsAny struct {
	ManagerID     *uint  `xml:"managerid,attr"`
	TypeID        *uint  `xml:"typeid,attr"`
	TypeIDVersion *uint  `xml:"typeidversion,attr"`
	InstanceID    string `xml:"instanceid,attr,omitempty"`
	ProfileID     string `xml:"profileid,attr,omitempty"`
	InterfaceID   string `xml:"interfaceid,attr,omitempty"`
}

type DomainInterfaceVirtualPortParamsVEPA8021QBG struct {
	ManagerID     *uint  `xml:"managerid,attr"`
	TypeID        *uint  `xml:"typeid,attr"`
	TypeIDVersion *uint  `xml:"typeidversion,attr"`
	InstanceID    string `xml:"instanceid,attr,omitempty"`
}

type DomainInterfaceVirtualPortParamsVNTag8021QBH struct {
	ProfileID string `xml:"profileid,attr,omitempty"`
}

type DomainInterfaceVirtualPortParamsOpenVSwitch struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty"`
	ProfileID   string `xml:"profileid,attr,omitempty"`
}

type DomainInterfaceVirtualPortParamsMidoNet struct {
	InterfaceID string `xml:"interfaceid,attr,omitempty"`
}

type DomainInterfaceBandwidthParams struct {
	Average *int `xml:"average,attr"`
	Peak    *int `xml:"peak,attr"`
	Burst   *int `xml:"burst,attr"`
	Floor   *int `xml:"floor,attr"`
}

type DomainInterfaceBandwidth struct {
	Inbound  *DomainInterfaceBandwidthParams `xml:"inbound"`
	Outbound *DomainInterfaceBandwidthParams `xml:"outbound"`
}

type DomainInterfaceVLan struct {
	Trunk string                   `xml:"trunk,attr,omitempty"`
	Tags  []DomainInterfaceVLanTag `xml:"tag"`
}

type DomainInterfaceVLanTag struct {
	ID         uint   `xml:"id,attr"`
	NativeMode string `xml:"nativeMode,attr,omitempty"`
}

type DomainInterfaceGuest struct {
	Dev    string `xml:"dev,attr,omitempty"`
	Actual string `xml:"actual,attr,omitempty"`
}

type DomainInterfaceFilterRef struct {
	Filter     string                       `xml:"filter,attr"`
	Parameters []DomainInterfaceFilterParam `xml:"parameter"`
}

type DomainInterfaceFilterParam struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type DomainInterfaceBackend struct {
	Tap   string `xml:"tap,attr,omitempty"`
	VHost string `xml:"vhost,attr,omitempty"`
}

type DomainInterfaceTune struct {
	SndBuf uint `xml:"sndbuf"`
}

type DomainInterfaceMTU struct {
	Size uint `xml:"size,attr"`
}

type DomainInterfaceCoalesce struct {
	RX *DomainInterfaceCoalesceRX `xml:"rx"`
}

type DomainInterfaceCoalesceRX struct {
	Frames *DomainInterfaceCoalesceRXFrames `xml:"frames"`
}

type DomainInterfaceCoalesceRXFrames struct {
	Max *uint `xml:"max,attr"`
}

type DomainROM struct {
	Bar     string `xml:"bar,attr,omitempty"`
	File    string `xml:"file,attr,omitempty"`
	Enabled string `xml:"enabled,attr,omitempty"`
}

type DomainInterfaceIP struct {
	Address string `xml:"address,attr"`
	Family  string `xml:"family,attr,omitempty"`
	Prefix  uint   `xml:"prefix,attr,omitempty"`
	Peer    string `xml:"peer,attr,omitempty"`
}

type DomainInterfaceRoute struct {
	Family  string `xml:"family,attr,omitempty"`
	Address string `xml:"address,attr"`
	Netmask string `xml:"netmask,attr,omitempty"`
	Prefix  uint   `xml:"prefix,attr,omitempty"`
	Gateway string `xml:"gateway,attr"`
	Metric  uint   `xml:"metric,attr,omitempty"`
}

/*
Domain interface Struct
*/
type DomainInterface struct {
	XMLName             xml.Name                   `xml:"interface"`
	Managed             string                     `xml:"managed,attr,omitempty"`
	TrustGuestRXFilters string                     `xml:"trustGuestRxFilters,attr,omitempty"`
	MAC                 DomainInterfaceMAC         `xml:"mac"`
	Source              DomainInterfaceSource      `xml:"source"`
	Boot                DomainDeviceBoot           `xml:"boot"`
	VLan                DomainInterfaceVLan        `xml:"vlan"`
	VirtualPort         DomainInterfaceVirtualPort `xml:"virtualport"`
	IP                  []DomainInterfaceIP        `xml:"ip"`
	Route               []DomainInterfaceRoute     `xml:"route"`
	Script              DomainInterfaceScript      `xml:"script"`
	Target              DomainInterfaceTarget      `xml:"target"`
	Guest               DomainInterfaceGuest       `xml:"guest"`
	Model               DomainInterfaceModel       `xml:"model"`
	Driver              DomainInterfaceDriver      `xml:"driver"`
	Backend             DomainInterfaceBackend     `xml:"backend"`
	FilterRef           DomainInterfaceFilterRef   `xml:"filterref"`
	Tune                DomainInterfaceTune        `xml:"tune"`
	Link                DomainInterfaceLink        `xml:"link"`
	MTU                 DomainInterfaceMTU         `xml:"mtu"`
	Bandwidth           DomainInterfaceBandwidth   `xml:"bandwidth"`
	Coalesce            DomainInterfaceCoalesce    `xml:"coalesce"`
	ROM                 DomainROM                  `xml:"rom"`
	Alias               DomainAlias                `xml:"alias"`
	Address             DomainAddress              `xml:"address"`
}
