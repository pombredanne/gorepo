package main

import (
	"encoding/xml"
)

type MetadataType struct {
	XMLName     xml.Name              `xml:"metadata"`
	Xmnls1      string                `xml:"xmlns:,attr"`
	Xmlns2      string                `xml:"xmlns:rpm,attr"`
	NumPackages int                   `xml:"packages,attr"`
	Packages    []MetadataPackageType `xml:"package"`
}

type MetadataPackageType struct {
	XMLName     xml.Name                   `xml:"package"`
	Type        string                     `xml:"type,attr"`
	Name        string                     `xml:"name"`
	Arch        string                     `xml:"arch"`
	Version     MetadataPackageVersionType `xml:"version"`
	Checksum    MetadataChecksumType       `xml:"checksum"`
	Summary     string                     `xml:"summary"`
	Description string                     `xml:"description"`
	Packager    string                     `xml:"packager"`
	Url         string                     `xml:"url"`
	Time        MetadataTimeType           `xml:"time"`
	Size        MetadataSizeType           `xml:"size"`
	Location    MetadataLocationType       `xml:"location"`
}

type MetadataPackageVersionType struct {
	Epoch string `xml:"epoch,attr"`
	Ver   string `xml:"ver,attr"`
	Rel   string `xml:"rel,attr"`
}

type MetadataChecksumType struct {
	Type  string `xml:"type,attr"`
	Pkgid string `xml:"pkgid,attr"`
	Value string `xml:",chardata"`
}

type MetadataTimeType struct {
	File  string `xml:"file,attr"`
	Build string `xml:"build,attr"`
}

type MetadataSizeType struct {
	Package   string `xml:"package,attr"`
	Installed string `xml:"installed,attr"`
	Archive   string `xml:"archive,attr"`
}

type MetadataLocationType struct {
	Href   string             `xml:"href,attr"`
	Format MetadataFormatType `xml:"format"`
}

type MetadataFormatType struct {
	Vendor      string                  `xml:"rpm:vendor"`
	Group       string                  `xml:"rpm:group"`
	Buildhost   string                  `xml:"rpm:buildhost"`
	HeaderRange MetadataHeaderRangeType `xml:"rpm:header-range"`
	Provides    []MetadataEntryType     `xml:"provides"`
	Requires    []MetadataEntryType     `xml:"requires"`
	Obsoletes   []MetadataEntryType     `xml:"obsoletes"`
	File        []string                `xml:"file"`
}

type MetadataHeaderRangeType struct {
	Start string `xml:"start,attr"`
	End   string `xml:"end,attr"`
}

type MetadataEntryType struct {
	Entry xml.Name `xml:"entry"`
	Name  string   `xml:"name,attr"`
	Flags string   `xml:"flags,attr"`
	Epoch string   `xml:"epoch,attr"`
	Ver   string   `xml:"ver,attr"`
}
