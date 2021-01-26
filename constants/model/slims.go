package model

//SlimsDetailBookResponse struct for binding detail book response
type SlimsDetailBookResponse struct {
	ModsCollection struct {
		XmlnsXlink        string `json:"-xmlns:xlink"`
		XmlnsXsi          string `json:"-xmlns:xsi"`
		Xmlns             string `json:"-xmlns"`
		XmlnsSlims        string `json:"-xmlns:slims"`
		XsiSchemaLocation string `json:"-xsi:schemaLocation"`
		Mods              struct {
			Version   string `json:"-version"`
			ID        string `json:"-ID"`
			TitleInfo struct {
				Title    string `json:"title"`
				SubTitle string `json:"subTitle"`
			} `json:"titleInfo"`
			Name struct {
				Type     string `json:"-type"`
				NamePart string `json:"namePart"`
				Role     struct {
					RoleTerm struct {
						Type string `json:"-type"`
						Text string `json:"#text"`
					} `json:"roleTerm"`
				} `json:"role"`
			} `json:"name"`
			TypeOfResource struct {
				Manuscript string `json:"-manuscript"`
				Collection string `json:"-collection"`
				Text       string `json:"#text"`
			} `json:"typeOfResource"`
			Genre struct {
				Authority string `json:"-authority"`
				Text      string `json:"#text"`
			} `json:"genre"`
			OriginInfo struct {
				Place struct {
					PlaceTerm struct {
						Type string `json:"-type"`
						Text string `json:"#text"`
					} `json:"placeTerm"`
				} `json:"place"`
				Publisher  string `json:"publisher"`
				DateIssued string `json:"dateIssued"`
				Issuance   string `json:"issuance"`
			} `json:"originInfo"`
			Language struct {
				LanguageTerm []struct {
					Type string `json:"-type"`
					Text string `json:"#text"`
				} `json:"languageTerm"`
			} `json:"language"`
			PhysicalDescription struct {
				Form struct {
					Authority string `json:"-authority"`
					Text      string `json:"#text"`
				} `json:"form"`
				Extent string `json:"extent"`
			} `json:"physicalDescription"`
			Note    string `json:"note"`
			Subject struct {
				Topic string `json:"topic"`
			} `json:"subject"`
			Classification string `json:"classification"`
			Identifier     struct {
				Type string `json:"-type"`
			} `json:"identifier"`
			Location struct {
				PhysicalLocation string `json:"physicalLocation"`
				ShelfLocator     string `json:"shelfLocator"`
				HoldingSimple    struct {
					CopyInformation []struct {
						NumerationAndChronology struct {
							Type string `json:"-type"`
							Text string `json:"#text"`
						} `json:"numerationAndChronology"`
						Sublocation  string `json:"sublocation"`
						ShelfLocator string `json:"shelfLocator,omitempty"`
					} `json:"copyInformation"`
				} `json:"holdingSimple"`
			} `json:"location"`
			SlimsImage string `json:"slims:image"`
			RecordInfo struct {
				RecordIdentifier   string `json:"recordIdentifier"`
				RecordCreationDate struct {
					Encoding string `json:"-encoding"`
					Text     string `json:"#text"`
				} `json:"recordCreationDate"`
				RecordChangeDate struct {
					Encoding string `json:"-encoding"`
					Text     string `json:"#text"`
				} `json:"recordChangeDate"`
				RecordOrigin string `json:"recordOrigin"`
			} `json:"recordInfo"`
		} `json:"mods"`
	} `json:"modsCollection"`
}
