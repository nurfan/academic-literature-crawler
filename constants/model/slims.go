package model

import "encoding/xml"

//SlimsDetailBookResponse struct for binding detail book response
type SlimsDetailBookResponse struct {
	XMLName        xml.Name `xml:"modsCollection"`
	Text           string   `xml:",chardata"`
	Xlink          string   `xml:"xlink,attr"`
	Xsi            string   `xml:"xsi,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Slims          string   `xml:"slims,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Mods           struct {
		Text      string `xml:",chardata"`
		Version   string `xml:"version,attr"`
		ID        string `xml:"ID,attr"`
		TitleInfo struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
		} `xml:"titleInfo"`
		Name struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			Authority string `xml:"authority,attr"`
			NamePart  string `xml:"namePart"`
			Role      struct {
				Text     string `xml:",chardata"`
				RoleTerm struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"roleTerm"`
			} `xml:"role"`
		} `xml:"name"`
		TypeOfResource struct {
			Text       string `xml:",chardata"`
			Manuscript string `xml:"manuscript,attr"`
			Collection string `xml:"collection,attr"`
		} `xml:"typeOfResource"`
		Genre struct {
			Text      string `xml:",chardata"`
			Authority string `xml:"authority,attr"`
		} `xml:"genre"`
		OriginInfo struct {
			Text  string `xml:",chardata"`
			Place struct {
				Text      string `xml:",chardata"`
				PlaceTerm struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"placeTerm"`
			} `xml:"place"`
			Publisher  string `xml:"publisher"`
			DateIssued string `xml:"dateIssued"`
			Issuance   string `xml:"issuance"`
			Edition    string `xml:"edition"`
		} `xml:"originInfo"`
		Language struct {
			Text         string `xml:",chardata"`
			LanguageTerm []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"languageTerm"`
		} `xml:"language"`
		PhysicalDescription struct {
			Text string `xml:",chardata"`
			Form struct {
				Text      string `xml:",chardata"`
				Authority string `xml:"authority,attr"`
			} `xml:"form"`
			Extent string `xml:"extent"`
		} `xml:"physicalDescription"`
		Note    string `xml:"note"`
		Subject struct {
			Text      string `xml:",chardata"`
			Authority string `xml:"authority,attr"`
			Topic     string `xml:"topic"`
		} `xml:"subject"`
		Classification string `xml:"classification"`
		Identifier     struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"identifier"`
		Location struct {
			Text             string `xml:",chardata"`
			PhysicalLocation string `xml:"physicalLocation"`
			ShelfLocator     string `xml:"shelfLocator"`
			HoldingSimple    struct {
				Text            string `xml:",chardata"`
				CopyInformation []struct {
					Text                    string `xml:",chardata"`
					NumerationAndChronology struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"numerationAndChronology"`
					Sublocation  string `xml:"sublocation"`
					ShelfLocator string `xml:"shelfLocator"`
				} `xml:"copyInformation"`
			} `xml:"holdingSimple"`
		} `xml:"location"`
		Image      string `xml:"image"`
		RecordInfo struct {
			Text               string `xml:",chardata"`
			RecordIdentifier   string `xml:"recordIdentifier"`
			RecordCreationDate struct {
				Text     string `xml:",chardata"`
				Encoding string `xml:"encoding,attr"`
			} `xml:"recordCreationDate"`
			RecordChangeDate struct {
				Text     string `xml:",chardata"`
				Encoding string `xml:"encoding,attr"`
			} `xml:"recordChangeDate"`
			RecordOrigin string `xml:"recordOrigin"`
		} `xml:"recordInfo"`
	} `xml:"mods"`
}

// SlimsBookInformation for
type SlimsBookInformation struct {
	Title               string    `json:"title"`
	Cover               string    `json:"cover"`
	Author              Authority `json:"author"`
	PublishDate         string    `json:"publish_date"`
	Publisher           string    `json:"publisher"`
	Edition             string    `json:"edition"`
	PhysicalDescription string    `json:"physical_description"`
	Subject             string    `json:"subject"`
	Classification      string    `json:"classification"`
	Locations           Location  `json:"location"`
}

// Authority for binding author info
type Authority struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Role string `json:"role"`
}

// Location is a struct
type Location struct {
	PhysicalLocation string            `json:"physical_location"`
	ShelfLocator     string            `json:"shelf_locator"`
	CopyInformations []CopyInformation `json:"copy_informations"`
}

// CopyInformation is a struct
type CopyInformation struct {
	Numeration   string `json:"numeration"`
	Sublocation  string `json:"sublocation"`
	ShelfLocator string `json:"shelflocation"`
}
