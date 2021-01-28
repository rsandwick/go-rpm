package repomd

import (
	"fmt"

	"rs3.io/go/rpm/repomd/xml"
)

type Comps struct {
	XMLName xml.Name `xml:"comps"`
	Groups  []Group  `xml:"group"`
}

type Group struct {
	XMLName     xml.Name     `xml:"group"`
	Id          string       `xml:"id"`
	Name        string       `xml:"name"`
	Description string       `xml:"description"`
	Default     bool         `xml:"default"`
	UserVisible bool         `xml:"uservisible"`
	PackageList []PackageReq `xml:"packagelist>packagereq"`
}

type PackageReq struct {
	XMLName xml.Name `xml:"packagereq"`
	Type    ReqType  `xml:"type,attr"`
	Name    string   `xml:",chardata"`
}

type ReqType int

const (
	Mandatory ReqType = iota
	Optional
)

func (rt ReqType) MarshalText() (text []byte, err error) {
	switch rt {
	case Mandatory:
		text = []byte("mandatory")
	case Optional:
		text = []byte("optional")
	default:
		err = fmt.Errorf("unknown ReqType value: %d", rt)
	}
	return
}

func (rt *ReqType) UnmarshalText(text []byte) error {
	switch string(text) {
	case "mandatory":
		*rt = Mandatory
	case "optional":
		*rt = Optional
	default:
		return fmt.Errorf("unknown ReqType value: %d", rt)
	}
	return nil
}
