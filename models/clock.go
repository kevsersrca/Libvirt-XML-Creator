package models

/*
Domain Clock  Struct

*/
type DomainClock struct {
	Offset     string        `xml:"offset,attr,omitempty"`
	Basis      string        `xml:"basis,attr,omitempty"`
	Adjustment string        `xml:"adjustment,attr,omitempty"`
	TimeZone   string        `xml:"timezone,attr,omitempty"`
	Timer      []DomainTimer `xml:"timer"`
}

/*
Domain Timer  Struct

*/
type DomainTimer struct {
	Name       string             `xml:"name,attr"`
	Track      string             `xml:"track,attr,omitempty"`
	TickPolicy string             `xml:"tickpolicy,attr,omitempty"`
	CatchUp    DomainTimerCatchUp `xml:"catchup"`
	Frequency  uint32             `xml:"frequency,attr,omitempty"`
	Mode       string             `xml:"mode,attr,omitempty"`
	Present    string             `xml:"present,attr,omitempty"`
}

/*
Domain Timer Catchup  Struct

*/
type DomainTimerCatchUp struct {
	Threshold uint `xml:"threshold,attr,omitempty"`
	Slew      uint `xml:"slew,attr,omitempty"`
	Limit     uint `xml:"limit,attr,omitempty"`
}

/*
Domain PM  Struct

*/
type DomainPM struct {
	SuspendToMem  DomainPMPolicy `xml:"suspend-to-mem"`
	SuspendToDisk DomainPMPolicy `xml:"suspend-to-disk"`
}

/*
Domain PM Policy  Struct

*/
type DomainPMPolicy struct {
	Enabled string `xml:"enabled,attr"`
}
