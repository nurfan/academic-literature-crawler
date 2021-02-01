package oaipmh

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	m "github.com/nurfan/academic-literature-crawler/constants/model"
	s "github.com/nurfan/academic-literature-crawler/constants/state"
	"github.com/nurfan/academic-literature-crawler/lib/errors"
	"github.com/parnurzeal/gorequest"
)

// OAI struct
type OAI struct {
	Request    Request
	HTTPClient *http.Client
	e          *errors.Error
}

// Request param for harvest oai-pmh
type Request struct {
	BaseURL         string
	Set             string
	MetadataPrefix  string
	Verb            string
	Identifier      string
	ResumptionToken string
	From            string
	Until           string
}

// SetRepository set repository HOST
// value : "OJS","SLIMS","EPRINTS"
func (c *OAI) SetRepository(repo string) {
	switch strings.ToUpper(repo) {
	case s.OJS:
		c.Request.BaseURL = os.Getenv("OJS_HOST")
	case s.SLIMS:
		c.Request.BaseURL = os.Getenv("SLIMS_HOST_OAI")
	case s.EPRINTs:
		c.Request.BaseURL = os.Getenv("EPRINTS_HOST")
	default:
		return
	}
}

// SetDateRange set date range query param
func (c *OAI) SetDateRange(from, until string) {
	if from != "" {
		c.Request.From = from
	}

	if until != "" {
		c.Request.Until = until
	}
}

// GetOAI get oai file
func (c *OAI) GetOAI() (*m.OaiResponse, error) {
	var result m.OaiResponse

	resp, err := c.sendRequest()
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(resp)
	err = json.Unmarshal(bytes, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// getFullURL represents the OAI Request in a string format
func (c *OAI) getFullURL() string {
	array := []string{}

	add := func(name, value string) {
		if value != "" {
			array = append(array, name+"="+value)
		}
	}

	add("verb", c.Request.Verb)
	add("set", c.Request.Set)
	add("metadataPrefix", c.Request.MetadataPrefix)
	add("resumptionToken", c.Request.ResumptionToken)
	add("identifier", c.Request.Identifier)
	add("from", c.Request.From)
	add("until", c.Request.Until)

	URL := strings.Join([]string{c.Request.BaseURL, "?", strings.Join(array, "&")}, "")

	return URL
}

// sendRequest get oai file from repository
func (c *OAI) sendRequest() (*m.OaiResponse, error) {
	var result m.OaiResponse
	log.Println("OAI-PMH HOST : ", c.getFullURL())
	resp, _, err := gorequest.New().Get(c.getFullURL()).End()

	if err != nil {
		msg := fmt.Sprint("ERR sendRequest : ", err)
		return nil, c.e.ProcessingError(msg)
	}

	bodyBytes, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		log.Println(err)
	}

	xml.Unmarshal(bodyBytes, &result)

	return &result, nil
}

// NewClient initiate oai-pmh connection
func NewClient() *OAI {
	return &OAI{
		Request: Request{
			Verb:           "ListRecords",
			MetadataPrefix: "oai_dc",
		},
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
