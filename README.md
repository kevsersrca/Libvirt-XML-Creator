# Libvirt Xml Operations Tool

This package manipulating libvirt XML documents. Define libvirt xml files with structs and used independently with libvirt.

Objects in the libvirt API are configured using XML documents to allow for ease of extension in future releases. Each XML document has an associated Relax-NG schema that can be used to validate documents prior to usage. For all formats : https://libvirt.org/format.html


# TODO
* [ ]  Models Todos completion
* [ ]  Xml Marshall
* [ ]  Xml Unmarshall
* [ ]  New Xml File Creation
* [ ]  Append Xml Structures
* [ ]  Delete Xml Structures
* [ ]  Cli 
* [ ]   Validating


# Configuration list to change 
<details>
 <summary>Domains</summary>

* [x] Element and attribute overview 
* [x] General metadata
* [ ] Operating system booting
* [ ] BIOS bootloader
* [ ] Host bootloader
* [ ] Direct kernel boot
* [ ] Container boot
* [ ] SMBIOS System Information
* [x] CPU Allocation
* [ ] IOThreads Allocation
* [ ] CPU Tuning
* [x] Memory Allocation
* [ ] Memory Backing
* [ ] Memory Tuning
* [ ] NUMA Node Tuning
* [ ] Block I/O Tuning
* [ ] Resource partitioning
* [x] CPU model and topology
* [x] Events configuration
* [x] Power Management
* [x] Hypervisor features
* [x] Time keeping
* [ ] Performance monitoring events
* [x] Devices
* [ ] Hard drives, floppy disks, CDROMs
* [ ] Filesystems
* [ ] Device Addresses
* [ ] Virtio-related options
* [x] Controllers
* [ ] Device leases
* [ ] Host device assignment
* [ ] USB / PCI / SCSI devices
* [ ] Block / character devices
* [ ] Redirected devices
* [ ] Smartcard devices
* [x] Network interfaces
* [x] Virtual network
* [x] Bridge to LAN
* [x] Userspace SLIRP stack
* [x] Generic ethernet connection
* [x] Direct attachment to physical interface
* [x] PCI Passthrough
* [x] Multicast tunnel
* [x] TCP tunnel
* [x] UDP unicast tunnel
* [x] Setting the NIC model
* [x] Setting NIC driver-specific options
* [x] Setting network backend-specific options
* [x] Overriding the target element
* [x] Specifying boot order
* [x] Interface ROM BIOS configuration
* [x] Setting up a network backend in a driver domain
* [x] Quality of service
* [x] Setting VLAN tag (on supported network types only)
* [x] Modifying virtual link state
* [x] MTU configuration
* [x] Coalesce settings
* [x] IP configuration
* [x] vhost-user interface
* [x] Traffic filtering with NWFilter
* [ ] Input devices
* [ ] Hub devices
* [x] Graphical framebuffers
* [ ] Video devices
* [ ] Consoles, serial, parallel & channel devices
* [ ] Guest interface
* [ ] Parallel port
* [ ] Serial port
* [ ] Console
* [ ] Relationship between serial ports and consoles
* [ ] Channel
* [ ] Host interface
* [ ] Domain logfile
* [ ] Device logfile
* [ ] Virtual console
* [ ] Null device
* [ ] Pseudo TTY
* [ ] Host device proxy
* [ ] Named pipe
* [ ] TCP client/server
* [ ] UDP network console
* [ ] UNIX domain socket client/server
* [ ] Spice channel
* [ ] Nmdm device
* [ ] Sound devices
* [ ] Watchdog device
* [x] Memory balloon device
* [ ] Random number generator device
* [ ] TPM device
* [ ] NVRAM device
* [ ] panic device
* [ ] Shared memory device
* [ ] Memory devices
* [ ] IOMMU devices
* [ ] Security label</details>

<details>
 <summary>Networks</summary>

</details>
<details>
 <summary>Network filtering</summary>

</details>
<details>
 <summary>Storage</summary>

</details>
<details>
 <summary>Storage encryption</summary>

</details>
<details>
 <summary>Capabilities</summary>

</details>
<details>
 <summary>Domain capabilities</summary>

</details>
<details>
 <summary>Node devices</summary>

</details>
<details>
 <summary>Secrets</summary>

</details>
<details>
 <summary>Snapshots</summary>
</details>
