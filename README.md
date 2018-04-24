# Libvirt Xml Operations Tool

This package manipulating libvirt XML documents. Define libvirt xml files with structs and used independently with libvirt.

<img src="https://raw.githubusercontent.com/kevsersrca/libvirt-xml/master/gopher.png" width="400" height="400" />


Objects in the libvirt API are configured using XML documents to allow for ease of extension in future releases. Each XML document has an associated Relax-NG schema that can be used to validate documents prior to usage. For all formats : https://libvirt.org/format.html

# Data Flow Diagram
![Gopher](https://raw.githubusercontent.com/kevsersrca/libvirt-xml/master/dataflowdiagram.png)


# TODO
* [ ]  Models Todos completion
* [ ]  Xml Marshall
* [ ]  Xml Unmarshall
* [ ]  New Xml File Creation
* [ ]  Append Xml Structures
* [ ]  Delete Xml Structures
* [ ]  Cli 
* []   Validating


# Configuration list to change 
<details>
 <summary>Domains</summary>

* [ ] Element and attribute overview
* [ ] General metadata
* [ ] Operating system booting
* [ ] BIOS bootloader
* [ ] Host bootloader
* [ ] Direct kernel boot
* [ ] Container boot
* [ ] SMBIOS System Information
* [ ] CPU Allocation
* [ ] IOThreads Allocation
* [ ] CPU Tuning
* [ ] Memory Allocation
* [ ] Memory Backing
* [ ] Memory Tuning
* [ ] NUMA Node Tuning
* [ ] Block I/O Tuning
* [ ] Resource partitioning
* [ ] CPU model and topology
* [ ] Events configuration
* [ ] Power Management
* [ ] Hypervisor features
* [ ] Time keeping
* [ ] Performance monitoring events
* [ ] Devices
* [ ] Hard drives, floppy disks, CDROMs
* [ ] Filesystems
* [ ] Device Addresses
* [ ] Virtio-related options
* [ ] Controllers
* [ ] Device leases
* [ ] Host device assignment
* [ ] USB / PCI / SCSI devices
* [ ] Block / character devices
* [ ] Redirected devices
* [ ] Smartcard devices
* [ ] Network interfaces
* [ ] Virtual network
* [ ] Bridge to LAN
* [ ] Userspace SLIRP stack
* [ ] Generic ethernet connection
* [ ] Direct attachment to physical interface
* [ ] PCI Passthrough
* [ ] Multicast tunnel
* [ ] TCP tunnel
* [ ] UDP unicast tunnel
* [ ] Setting the NIC model
* [ ] Setting NIC driver-specific options
* [ ] Setting network backend-specific options
* [ ] Overriding the target element
* [ ] Specifying boot order
* [ ] Interface ROM BIOS configuration
* [ ] Setting up a network backend in a driver domain
* [ ] Quality of service
* [ ] Setting VLAN tag (on supported network types only)
* [ ] Modifying virtual link state
* [ ] MTU configuration
* [ ] Coalesce settings
* [ ] IP configuration
* [ ] vhost-user interface
* [ ] Traffic filtering with NWFilter
* [ ] Input devices
* [ ] Hub devices
* [ ] Graphical framebuffers
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
* [ ] Memory balloon device
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
