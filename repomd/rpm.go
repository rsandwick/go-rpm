package repomd

import (
	"rs3.io/go/rpm/repomd/xml"
)

type Format struct {
	XMLName   xml.Name `xml:"format"`
	License   string   `xml:"http://linux.duke.edu/metadata/rpm license"`
	Vendor    string   `xml:"http://linux.duke.edu/metadata/rpm vendor,selfclose"`
	Group     string   `xml:"http://linux.duke.edu/metadata/rpm group"`
	BuildHost string   `xml:"http://linux.duke.edu/metadata/rpm buildhost"`
	SourceRPM string   `xml:"http://linux.duke.edu/metadata/rpm sourcerpm"`

	HeaderRange RPMHeaderRange `xml:"http://linux.duke.edu/metadata/rpm header-range,selfclose"`

	Provides  RPMProvides  `xml:"http://linux.duke.edu/metadata/rpm provides,omitempty"`
	Requires  RPMRequires  `xml:"http://linux.duke.edu/metadata/rpm requires,omitempty"`
	Conflicts RPMConflicts `xml:"http://linux.duke.edu/metadata/rpm conflicts,omitempty"`
	Obsoletes RPMObsoletes `xml:"http://linux.duke.edu/metadata/rpm obsoletes,omitempty"`

	Files []File `xml:"file"`
}

type RPMHeaderRange struct {
	XMLName xml.Name `xml:"http://linux.duke.edu/metadata/rpm header-range,selfclose"`
	Start   int      `xml:"start,attr"`
	End     int      `xml:"end,attr"`
}

type RPMProvides struct {
	XMLName xml.Name   `xml:"http://linux.duke.edu/metadata/rpm provides"`
	Entries []RPMEntry `xml:"http://linux.duke.edu/metadata/rpm entry,omitempty,selfclose"`
}

type RPMRequires struct {
	XMLName xml.Name   `xml:"http://linux.duke.edu/metadata/rpm requires"`
	Entries []RPMEntry `xml:"http://linux.duke.edu/metadata/rpm entry,omitempty,selfclose"`
}

type RPMConflicts struct {
	XMLName xml.Name   `xml:"http://linux.duke.edu/metadata/rpm conflicts"`
	Entries []RPMEntry `xml:"http://linux.duke.edu/metadata/rpm entry,omitempty,selfclose"`
}

type RPMObsoletes struct {
	XMLName xml.Name   `xml:"http://linux.duke.edu/metadata/rpm obsoletes"`
	Entries []RPMEntry `xml:"http://linux.duke.edu/metadata/rpm entry,omitempty,selfclose"`
}

type RPMEntry struct {
	XMLName xml.Name `xml:"http://linux.duke.edu/metadata/rpm entry"`
	Name    string   `xml:"name,attr"`
	Flags   string   `xml:"flags,attr,omitempty"`
	EVR
	Pre int `xml:"pre,attr,omitempty"`
}

type File struct {
	XMLName xml.Name `xml:"file"`
	Type    string   `xml:"type,attr,omitempty"`
	Name    string   `xml:",innerxml"`
}
