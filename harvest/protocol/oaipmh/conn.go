package oaipmh

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

// Conn struct
type Conn struct {
	Request    Request
	HTTPClient *http.Client
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

// GetFullURL represents the OAI Request in a string format
func (c *Conn) GetFullURL() string {
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

// SendRequest get oai file from repository
func (c *Conn) SendRequest() (*Response, error) {
	var result Response
	log.Println("OAI-PMH HOST : ", c.GetFullURL())
	resp, _, err := gorequest.New().Get(c.GetFullURL()).End()

	if err != nil {
		log.Println("ERR sendRequest : ", err)
		return nil, err[0]
	}

	bodyBytes, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		log.Fatal(errr)
	}

	xml.Unmarshal(bodyBytes, &result)
	return &result, nil
}

// NewConnection initiate oai-pmh connection
func NewConnection(req Request) *Conn {
	return &Conn{
		Request: req,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
