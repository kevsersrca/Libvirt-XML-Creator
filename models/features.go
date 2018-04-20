package models

import "encoding/xml"

/*
Example:
<features>
  <pae/>
  <acpi/>
  <apic/>
  <hap/>
  <privnet/>
  <hyperv>
    <relaxed state='on'/>
    <vapic state='on'/>
    <spinlocks state='on' retries='4096'/>
    <vpindex state='on'/>
    <runtime state='on'/>
    <synic state='on'/>
    <reset state='on'/>
    <vendor_id state='on' value='KVM Hv'/>
  </hyperv>
  <kvm>
    <hidden state='on'/>
  </kvm>
  <pvspinlock state='on'/>
  <gic version='2'/>
  <ioapic driver='qemu'/>
  <hpt resizing='required'/>
</features>


TODO:
GIC
IOAPIC
HPT
*/

type DomainFeatureList struct {
	PAE        DomainFeature       `xml:"pae"`
	ACPI       DomainFeature       `xml:"acpi"`
	APIC       DomainFeature       `xml:"apic"`
	HAP        DomainFeature       `xml:"hap"`
	PrivNet    DomainFeature       `xml:"privnet"`
	PVSpinlock DomainFeatureState  `xml:"pvspinlock"`
	HyperV     DomainFeatureHyperV `xml:"hyperv"`
	KVM        DomainFeatureKVM    `xml:"kvm"`
}

/*
Domain Feature HyperV Struct

Example:
<hyperv>
    <relaxed state='on'/>
    <vapic state='on'/>
    <spinlocks state='on' retries='4096'/>
    <vpindex state='on'/>
    <runtime state='on'/>
    <synic state='on'/>
    <reset state='on'/>
    <vendor_id state='on' value='KVM Hv'/>
  </hyperv>
*/
type DomainFeatureHyperV struct {
	DomainFeature
	Relaxed   DomainFeatureState           `xml:"relaxed"`
	VAPIC     DomainFeatureState           `xml:"vapic"`
	Spinlocks DomainFeatureHyperVSpinlocks `xml:"spinlocks"`
	VPIndex   DomainFeatureState           `xml:"vpindex"`
	Runtime   DomainFeatureState           `xml:"runtime"`
	Synic     DomainFeatureState           `xml:"synic"`
	STimer    DomainFeatureState           `xml:"stimer"`
	Reset     DomainFeatureState           `xml:"reset"`
	VendorId  DomainFeatureHyperVVendorId  `xml:"vendor_id"`
}

/*
Domain Feature KVM Struct

Example:
<kvm>
    <hidden state='on'/>
</kvm>

*/
type DomainFeatureKVM struct {
	Hidden DomainFeatureState `xml:"hidden"`
}

/*
Domain Feature State Attribute Struct

Example:
<relaxed state='on'/>

State string value is  on
*/
type DomainFeatureState struct {
	State string `xml:"state,attr,omitempty"`
}

/*
Domain Feature HyperV Spin Locks Struct

Example:
<spinlocks state='on' retries='4096'/>

Spinlocks Attributes return
*/

type DomainFeatureHyperVSpinlocks struct {
	DomainFeatureState
	Retries uint `xml:"retries,attr,omitempty"`
}

/*
Domain Feature HyperV Vendor Id Struct

Example:
<vendor_id state='on' value='KVM Hv'/>

VendorId Attributes return . There are state=on and value= KVM Hv
*/
type DomainFeatureHyperVVendorId struct {
	DomainFeatureState
	Value string `xml:"value,attr,omitempty"`
}

/*
Domain Feature Name Struct

Example:
<pae/>

DomainFeatureName is pae
*/
type DomainFeature struct {
	DomainFeatureName xml.Name `xml:","`
}
