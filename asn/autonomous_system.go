package asn

import "encoding/xml"

type AutonomousSystem struct {
	Asn  string   `xml:"ASN"`
	Ipv4 []string `xml:"IPv4"`
	Ipv6 []string `xml:"IPv6"`
}

func NewAs(asn string) *AutonomousSystem {
	return &AutonomousSystem{Asn: asn}
}

func (a *AutonomousSystem) ToXml() ([]byte, error) {
	return xml.Marshal(a)
}
