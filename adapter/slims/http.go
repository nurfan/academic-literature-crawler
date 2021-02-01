package slims

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	m "github.com/nurfan/academic-literature-crawler/constants/model"
	"github.com/parnurzeal/gorequest"
)

// Slims struct object
type Slims struct {
	Host       string
	HTTPClient *http.Client
}

// GetBookInfo get oai file
func (s *Slims) GetBookInfo(idBook string) (*m.SlimsDetailBookResponse, error) {

	result, err := s.sendRequest(s.getFullURLBookInfo(idBook))

	if err != nil {
		return nil, err
	}

	return result, nil
}

// getFullURLBookInfo represents the slims book detail info
func (s *Slims) getFullURLBookInfo(idBook string) string {
	array := []string{}

	add := func(name, value string) {
		if value != "" {
			array = append(array, name+"="+value)
		}
	}

	add("p", "show_detail")
	add("inXML", "true")
	add("id", idBook)

	URL := strings.Join([]string{s.Host, "/index.php", "?", strings.Join(array, "&")}, "")

	return URL
}

// sendRequest get oai file from repository
func (s *Slims) sendRequest(url string) (*m.SlimsDetailBookResponse, error) {

	log.Println("Slims HOST : ", url)
	resp, _, err := gorequest.New().Get(url).End()

	if err != nil {
		log.Println("ERR sendRequest : ", err)
		return nil, err[0]
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	result := m.SlimsDetailBookResponse{}
	xml.Unmarshal(bodyBytes, &result)

	return &result, nil
}

// NewClient initiate oai-pmh connection
func NewClient() *Slims {
	return &Slims{
		Host: os.Getenv("SLIMS_HOST"),
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
