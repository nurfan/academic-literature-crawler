package handler

import (
	"encoding/json"
	"fmt"
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
	// create context
	ctx := c.Request().Context()

	archiveID := c.Param("ID")
	log.Println(archiveID)

	// Search with a term query
	termQuery := elastic.NewMultiMatchQuery(archiveID, "archive_id").Type("phrase_prefix")
	searchResult, err := ha.elastic.Search().
		Index("archives").
		Query(termQuery).
		Pretty(true).
		Do(ctx)

	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			panic(fmt.Sprintf("Document not found: %v", err))
		case elastic.IsTimeout(err):
			panic(fmt.Sprintf("Timeout retrieving document: %v", err))
		case elastic.IsConnErr(err):
			panic(fmt.Sprintf("Connection problem: %v", err))
		default:
			// Some other kind of error
			panic(err)
		}
	}

	var result m.APIResponse
	var t m.Archive

	if searchResult.TotalHits() > 0 {

		for _, hit := range searchResult.Hits.Hits {
			err := json.Unmarshal(hit.Source, &t)
			if err == nil {
				log.Println(err)
			}
		}

		result.Code = http.StatusOK
		result.Message = http.StatusText(result.Code)
		result.Data = map[string]interface{}{
			"archive": t,
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
