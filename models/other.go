package models

/*
Domain Video Struct
TODO:

*/
type DomainVideo struct {
}

/*
Domain Graphic Struct
TODO:

*/
type DomainGraphic struct {
	Type     string `xml:"type,attr"`
	Port     string `xml:"port,attr"`
	Autoport string `xml:"autoport,attr"`
	Address  string `xml:"listen,attr"`
	Password string `xml:"passwd,attr"`
}

/*
Domain Console  Struct
TODO:

*/
type DomainConsole struct {
	Type string `xml:"type,attr"`
}

/*
Domain Serial  Struct
TODO:

*/
type DomainSerial struct {
	Type string `xml:"type,attr"`
}
