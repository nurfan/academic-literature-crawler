package model

import "net/http"

// APIResponse struct
type APIResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Errors  string                 `json:"errors,omitempty"`
}

// SetErrorResponse builder for error response
func (r *APIResponse) SetErrorResponse(code int, message string) {
	r.Code = code
	r.Message = http.StatusText(code)
	r.Errors = message
}

// SetSuccessResponse builder for error response
func (r *APIResponse) SetSuccessResponse(code int, result map[string]interface{}) {
	r.Code = code
	r.Message = http.StatusText(code)
	r.Data = result
}

// ListArchive struct
type ListArchive struct {
	ArchiveID string `json:"archive_id"`
	Platform  string `json:"platform"`
	Title     string `json:"title"`
	Creator   string `json:"creator"`
	Subject   string `json:"subject"`
	Publisher string `json:"publisher"`
	Link      string `json:"_link"`
}

// DetailEprintsResponse struct
type DetailEprintsResponse struct {
	ArchiveID          string     `json:"archive_id"`
	OaiIdentifier      string     `json:"oai_identifier"`
	Platform           string     `json:"platform"`
	Title              string     `json:"title"`
	Creator            string     `json:"creator"`
	Subject            string     `json:"subject"`
	Description        string     `json:"description"`
	Publisher          string     `json:"publisher"`
	Contributor        string     `json:"contributor"`
	Date               string     `json:"date"`
	Type               string     `json:"type"`
	DocumentIdentifier string     `json:"document_identifier"`
	Documents          []Document `json:"documents"`
	Rights             string     `json:"rights"`
}

// Document struct for binding indentifier to document
type Document struct {
	FileName string `json:"filename"`
	Language string `json:"language"`
	Format   string `json:"format"`
	URL      string `json:"url"`
}
