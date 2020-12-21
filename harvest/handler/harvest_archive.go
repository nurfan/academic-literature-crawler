package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nurfan/academic-literature-crawler/harvest/client"
)

// HarvestArchive initiate object
type HarvestArchive struct {
	client *client.Client
}

// Handle : handle request for this action
func (ha *HarvestArchive) Handle(c echo.Context) (err error) {

	conn := ha.client.NewClient()
	conn.SetRepository("EPRINTS")
	res, _ := conn.GetOAI()

	log.Println(res)
	return c.JSON(http.StatusCreated, "Tested Create Article PASS")
}

// NewHarvestArchive setup initiate object
func NewHarvestArchive() *HarvestArchive {
	return &HarvestArchive{}
}
