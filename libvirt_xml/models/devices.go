package models

/*
TODO:
DomainLease
DomainFilesystem
DomainSmartcard
DomainParallel
DomainChannel
DomainInput
DomainTPM
DomainSound
DomainHostdev
DomainRedirDev
DomainRedirFilter
DomainHub
DomainWatchdog
RNGs
NVRAM
Panics
Shmems
Memorydevs
IOMMU
*/

type DomainDevices struct {
	Emulator    string             `xml:"emulator,omitempty"`
	Disks       []DomainDisk       `xml:"disk"`
	Controllers []DomainController `xml:"controller"`
	Interfaces  []DomainInterface  `xml:"interface"`
	Serial      []DomainSerial     `xml:"serial"`
	Consoles    []DomainConsole    `xml:"console"`
	Graphics    []DomainGraphic    `xml:"graphics"`
	Videos      []DomainVideo      `xml:"video"`
	MemBalloon  DomainMemBalloon   `xml:"memballoon"`
}
