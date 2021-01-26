package model

// Archive struct
type Archive struct {
	ArchiveID     string `json:"archive_id"`
	OaiIdentifier string `json:"oai_identifier"`
	Platform      string `json:"platform"`
	Title         string `json:"title"`
	Creator       string `json:"creator"`
	Subject       string `json:"subject"`
	Description   string `json:"description"`
	Publisher     string `json:"publisher"`
	Contributor   string `json:"contributor"`
	Date          string `json:"date"`
	Type          string `json:"type"`
	Identifier    string `json:"identifier"`
	Language      string `json:"language"`
	Rights        string `json:"rights"`
	Format        string `json:"format"`
	Source        string `json:"source"`
	Relation      string `json:"relation"`
	Coverage      string `json:"coverage"`
}
