package client

import (
	"log"
	"os"
	"strings"

	"github.com/nurfan/academic-literature-crawler/harvest/protocol/oaipmh"
	m "github.com/nurfan/academic-literature-crawler/model"
)

// Client  slims object
type Client struct {
	Request oaipmh.Request
}

// SetRepository set repository HOST
// value : "OJS","SLIMS","EPRINTS"
func (c *Client) SetRepository(repo string) {
	switch strings.ToUpper(repo) {
	case m.OJS:
		c.Request.BaseURL = os.Getenv("OJS_HOST")
	case m.SLIMS:
		c.Request.BaseURL = os.Getenv("SLIMS_HOST")
	case m.EPRINTs:
		c.Request.BaseURL = os.Getenv("EPRINTS_HOST")
	default:
		return
	}
}

// SetDateRange set date range query param
func (c *Client) SetDateRange(from, until string) {
	if from != "" {
		c.Request.From = from
	}

	if until != "" {
		c.Request.Until = until
	}
}

// GetOAI get oai file
func (c *Client) GetOAI() (*Response, error) {
	var res Response
	conn := oaipmh.NewConnection(c.Request)
	result, err := conn.SendRequest()

	log.Fatal("Client.GetOAI : ", result)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// NewClient initiate slims object
func (c *Client) NewClient() *Client {
	return &Client{
		Request: oaipmh.Request{
			Verb:           "ListRecords",
			MetadataPrefix: "oai_dc",
		},
	}
}
