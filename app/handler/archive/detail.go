package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nurfan/academic-literature-crawler/app/repo"
	m "github.com/nurfan/academic-literature-crawler/constants/model"
	"github.com/olivere/elastic/v7"
)

// DetailArchive initiate object
type DetailArchive struct {
	elastic *elastic.Client
	arcRepo repo.ArchiveElasticRepo
}

// Handle : handle request for this action
func (ha *DetailArchive) Handle(c echo.Context) (err error) {
	var result m.APIResponse
	ctx := c.Request().Context()

	archiveID := c.Param("ID")
	searchResult, err := ha.arcRepo.SearchByArchiveID(ctx, archiveID)

	var archive m.Archive
	if searchResult.TotalHits() > 0 {

		for _, hit := range searchResult.Hits.Hits {
			err := json.Unmarshal(hit.Source, &archive)
			if err == nil {
				log.Println(err)
			}
		}

		result.Code = http.StatusOK
		result.Message = http.StatusText(result.Code)
		result.Data = map[string]interface{}{
			"archive": archive,
		}

		return c.JSON(result.Code, result)
	}

	result.Code = http.StatusNotFound
	result.Message = http.StatusText(result.Code)

	return c.JSON(result.Code, result)
}

// NewDetailArchive setup initiate object
func NewDetailArchive(elasticConn *elastic.Client) *DetailArchive {
	return &DetailArchive{
		elastic: elasticConn,
		arcRepo: repo.NewArchiveIndex(elasticConn),
	}
}
