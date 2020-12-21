package client

import (
	"encoding/xml"
)

// Response struct for binding oai-pmh response
type Response struct {
	XMLName        xml.Name    `xml:"OAI-PMH"`
	Text           string      `xml:",chardata"`
	Xmlns          string      `xml:"xmlns,attr"`
	Xsi            string      `xml:"xsi,attr"`
	SchemaLocation string      `xml:"schemaLocation,attr"`
	ResponseDate   string      `xml:"responseDate"`
	Request        RequestFrom `xml:"request"`
	Error          OAIError    `xml:"error"`
	ListRecords    ListRecord  `xml:"ListRecords"`
}

// OAIError struct
type OAIError struct {
	Code    string `xml:"code,attr"`
	Message string `xml:",chardata"`
}

// RequestFrom request information
type RequestFrom struct {
	Text           string `xml:",chardata"`
	Verb           string `xml:"verb,attr"`
	MetadataPrefix string `xml:"metadataPrefix,attr"`
}

// ListRecord struct
type ListRecord struct {
	Text            string          `xml:",chardata"`
	Record          []Record        `xml:"record"`
	ResumptionToken ResumptionToken `xml:"resumptionToken"`
}

// Record struct
type Record struct {
	Text     string `xml:",chardata"`
	Header   Header `xml:"header"`
	Metadata struct {
		Text string `xml:",chardata"`
		Dc   Dc     `xml:"dc"`
	} `xml:"metadata"`
}

// Header struct
type Header struct {
	Text       string   `xml:",chardata"`
	Status     string   `xml:"status,attr"`
	Identifier string   `xml:"identifier"`
	Datestamp  string   `xml:"datestamp"`
	SetSpec    []string `xml:"setSpec"`
}

// Dc struct
type Dc struct {
	Title       []string `xml:"title"`
	Creator     []string `xml:"creator"`
	Subject     []string `xml:"subject"`
	Description []string `xml:"description"`
	Publisher   []string `xml:"publisher"`
	Contributor []string `xml:"contributor"`
	Date        []string `xml:"date"`
	Type        []string `xml:"type"`
	Identifier  []string `xml:"identifier"`
	Language    []string `xml:"language"`
	Rights      []string `xml:"rights"`
	Format      []string `xml:"format"`
	Source      []string `xml:"source"`
	Relation    []string `xml:"relation"`
	Coverage    []string `xml:"coverage"`
}

// ResumptionToken struct
type ResumptionToken struct {
	Text             string `xml:",chardata"`
	ExpirationDate   string `xml:"expirationDate,attr"`
	CompleteListSize string `xml:"completeListSize,attr"`
	Cursor           string `xml:"cursor,attr"`
}
