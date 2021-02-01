package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nurfan/academic-literature-crawler/app/repo"
	m "github.com/nurfan/academic-literature-crawler/constants/model"
	s "github.com/nurfan/academic-literature-crawler/constants/state"
	"github.com/olivere/elastic/v7"
	"github.com/parnurzeal/gorequest"
)

// Journal initiate object
type Journal struct {
	elastic *elastic.Client
	arcRepo repo.ArchiveElasticRepo
}

// Handle : handle request for this action
func (j *Journal) Handle(c echo.Context) (err error) {
	var result m.APIResponse
	ctx := c.Request().Context()

	archiveID := c.Param("ID")
	log.Println(archiveID)

	searchResult, err := j.arcRepo.SearchByArchiveID(ctx, archiveID)

	var archive m.Archive
	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			err := json.Unmarshal(hit.Source, &archive)
			if err == nil {
				log.Println(err)
			}
		}

		if archive.Platform == s.OJS {
			resp, body, errs := gorequest.New().Get(archive.Relation).End()

			if errs != nil {
				log.Println(errs)
			}

			return c.Blob(resp.StatusCode, "application/pdf", []byte(body))
		}
	}

	result.Code = http.StatusNotFound
	result.Message = http.StatusText(result.Code)

	return c.JSON(result.Code, result)
}

// NewJournal setup initiate object
func NewJournal(elasticConn *elastic.Client) *Journal {
	return &Journal{
		elastic: elasticConn,
		arcRepo: repo.NewArchiveIndex(elasticConn),
	}
}
