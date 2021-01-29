package repomd

import (
	"fmt"
	"io"

	"rs3.io/go/rpm/repomd/xml"
)

type Metadata struct {
	XMLName  xml.Name  `xml:"http://linux.duke.edu/metadata/common metadata"`
	XMLNSRPM string    `xml:"xmlns:rpm,attr"`
	Count    int       `xml:"packages,attr"`
	Packages []Package `xml:"package"`
}

func ReadMetadata(r io.Reader) (*Metadata, error) {
	var md Metadata
	if err := xml.NewDecoder(r).Decode(&md); err != nil {
		return nil, err
	}
	md.XMLNSRPM = "http://linux.duke.edu/metadata/rpm"
	return &md, nil
}

func (md *Metadata) WriteTo(w io.Writer) error {
	return md.WriteToIndent(w, "", "")
}

func (md *Metadata) WriteToIndent(
	w io.Writer,
	prefix string,
	indent string,
) error {
	encoder := xml.NewEncoder(w)
	encoder.DeepEmpty(true)
	encoder.Indent(prefix, indent)
	w.Write([]byte(xml.Header))
	return encoder.EncodeNS(md, map[string]string{
		"http://linux.duke.edu/metadata/rpm": "rpm",
	})
}

type Package struct {
	XMLName     xml.Name `xml:"package"`
	Type        string   `xml:"type,attr"`
	Name        string   `xml:"name"`
	Arch        string   `xml:"arch"`
	Version     Version  `xml:"version,selfclose"`
	Checksum    Checksum `xml:"checksum"`
	Summary     RawBytes `xml:"summary"`
	Description RawBytes `xml:"description"`
	Packager    string   `xml:"packager"`
	URL         string   `xml:"url"`
	Time        Time     `xml:"time,selfclose"`
	Size        Size     `xml:"size,selfclose"`
	Location    Location `xml:"location,selfclose"`
	Format      Format   `xml:"format"`
}

type Version struct {
	XMLName xml.Name `xml:"version"`
	EVR
}

type EVR struct {
	Epoch   *int   `xml:"epoch,attr,omitempty"`
	Version string `xml:"ver,attr,omitempty"`
	Release string `xml:"rel,attr,omitempty"`
}

type RawBytes struct {
	Data string `xml:",innerxml"`
}

type Checksum struct {
	XMLName xml.Name     `xml:"checksum"`
	Type    ChecksumType `xml:"type,attr"`
	Pkgid   string       `xml:"pkgid,attr"`
	Value   string       `xml:",innerxml"`
}

type ChecksumType int

const (
	SHA256 ChecksumType = iota
)

func (ct ChecksumType) MarshalText() (text []byte, err error) {
	if ct == SHA256 {
		return []byte("sha256"), nil
	}
	return nil, fmt.Errorf("unknown checksum type: %d", ct)
}

func (ct *ChecksumType) UnmarshalText(text []byte) error {
	switch string(text) {
	case "sha256":
		*ct = SHA256
	default:
		return fmt.Errorf("unknown checksum type: %s", string(text))
	}
	return nil
}

type Time struct {
	XMLName xml.Name `xml:"time"`
	File    int      `xml:"file,attr"`
	Build   int      `xml:"build,attr"`
}

type Size struct {
	XMLName   xml.Name `xml:"size"`
	Package   int      `xml:"package,attr"`
	Installed int      `xml:"installed,attr"`
	Archive   int      `xml:"archive,attr"`
}

type Location struct {
	XMLName xml.Name `xml:"location"`
	HRef    string   `xml:"href,attr"`
}
